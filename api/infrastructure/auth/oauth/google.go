package oauth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type ResponseToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

type ResponseUserInfo struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func RequestToken(code string) (string, error) {
	client := &http.Client{}
	url := fmt.Sprintf("https://oauth2.googleapis.com/token?client_id=%s&client_secret=%s&code=%s&grant_type=authorization_code&redirect_uri=%s",
		os.Getenv("NEXT_PUBLIC_GOOGLE_CLIENT_ID"),
		os.Getenv("GOOGLE_CLIENT_SECRET"),
		code,
		"http://localhost:3001/login")

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "curl/7.54.0")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		return "", fmt.Errorf("invalid request")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data ResponseToken
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	return data.AccessToken, nil
}

func RequestUserInfo(access_token string) (ResponseUserInfo, error) {
	var userInfo ResponseUserInfo
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v2/userinfo", nil)
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
