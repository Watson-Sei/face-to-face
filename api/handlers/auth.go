package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/Watson-Sei/face-to-face/database"
	"github.com/Watson-Sei/face-to-face/models"
	"github.com/Watson-Sei/face-to-face/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RequestCode struct {
	Code string `json:"code"`
}

type ResponseToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

type ResponseTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type ResponseUserInfo struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Id    string `json:"id"`
}

type ResponseServiceToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// google oauth2.0 -> create account & update account -> generate token
func GetToken(c echo.Context) error {
	reqCode := new(RequestCode)
	if err := c.Bind(reqCode); err != nil {
		return err
	}

	client := &http.Client{}
	url := fmt.Sprintf("https://oauth2.googleapis.com/token?client_id=%s&client_secret=%s&code=%s&grant_type=authorization_code&redirect_uri=%s",
		os.Getenv("NEXT_PUBLIC_GOOGLE_CLIENT_ID"),
		os.Getenv("GOOGLE_CLIENT_SECRET"),
		reqCode.Code,
		"http://localhost:3001/login")

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "curl/7.54.0")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		return c.JSON(http.StatusBadRequest, "Bad request")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var data ResponseToken
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	userInfo, err := getUserInfo(data.AccessToken)
	if err != nil {
		return err
	}
	var user models.User
	if err := database.DB.Db.Where("email = ?", userInfo.Email).First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		if err := database.DB.Db.Create(&models.User{
			Name:  userInfo.Name,
			Email: userInfo.Email,
			Role:  "guest",
		}).Error; err != nil {
			return err
		}
		if err := database.DB.Db.Where("email = ?", userInfo.Email).First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	} else {
		if err := database.DB.Db.Model(&user).Updates(&models.User{
			Name:  userInfo.Name,
			Email: userInfo.Email,
		}).Error; err != nil {
			return err
		}
	}

	// Generate Token
	token, err := utils.GenerateJWT(strconv.Itoa(int(user.ID)), user.Role)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ResponseServiceToken{
		AccessToken: token,
	})
}

// 検証用ハンドラー
func Check(c echo.Context) error {
	claims, ok := c.Get("jwt").(utils.Claims)
	if !ok {
		return fmt.Errorf("invalid claims type")
	}

	return c.JSON(http.StatusOK, claims.UserID)
}

// get userinfo from google
func getUserInfo(access_token string) (ResponseUserInfo, error) {
	var userInfo ResponseUserInfo
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://www.googleapis.com/userinfo/v2/me", nil)
	if err != nil {
		return userInfo, err
	}

	req.Header.Set("Authorization", "Bearer "+access_token)

	resp, err := client.Do(req)
	if err != nil {
		return userInfo, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return userInfo, err
	}

	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		return userInfo, err
	}

	return userInfo, nil
}
