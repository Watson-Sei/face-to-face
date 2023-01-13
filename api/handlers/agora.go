package handlers

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/AgoraIO-Community/go-tokenbuilder/rtctokenbuilder"
	"github.com/AgoraIO-Community/go-tokenbuilder/rtmtokenbuilder"
	"github.com/labstack/echo/v4"
)

var appID, appCertificate string

type AgoraEnv struct {
	AppID          string `json:"appID"`
	AppCertificate string `json:"appCertificate"`
}

func getEnv() (AgoraEnv, error) {
	appIDEnv, appIDExists := os.LookupEnv("AGORA_APP_ID")
	appCertEnv, appCertExists := os.LookupEnv("AGORA_APP_CERTIFICATE")

	if !appIDExists || !appCertExists {
		log.Fatal("FATAL ERROR: ENV not properly configured, check APP_ID and APP_CERTIFICATE")
		return AgoraEnv{}, fmt.Errorf("no enviroment variables")
	} else {
		return AgoraEnv{AppID: appIDEnv, AppCertificate: appCertEnv}, nil
	}
}

func GetRtcToken(c echo.Context) error {
	agoraEnv, err := getEnv()
	if err != nil {
		return c.JSON(400, err)
	}
	appID = agoraEnv.AppID
	appCertificate = agoraEnv.AppCertificate
	channelName, tokentype, uidStr, role, expireTimestamp, err := parseRtcParams(c)
	if err != nil {
		return c.JSON(400, err)
	}

	rtcToken, tokenErr := generateRtcToken(channelName, uidStr, tokentype, role, expireTimestamp)
	if tokenErr != nil {
		log.Print(tokenErr)
		return c.JSON(400, tokenErr)
	}
	log.Println("RTC Token generated")
	return c.JSON(200, map[string]string{
		"rtcToken": rtcToken,
	})
}

func GetRtmToken(c echo.Context) error {
	log.Printf("rtm token¥n")
	uidStr, expireTimestamp, err := parseRtmParams(c)

	if err != nil {
		return c.JSON(400, err)
	}

	rtmToken, tokenErr := rtmtokenbuilder.BuildToken(appID, appCertificate, uidStr, rtmtokenbuilder.RoleRtmUser, expireTimestamp)
	if tokenErr != nil {
		log.Print(err)
		return c.JSON(400, tokenErr)
	}
	log.Println("RTM Token generated")
	return c.JSON(200, map[string]string{
		"rtmToken": rtmToken,
	})
}

func GetBothTokens(c echo.Context) error {
	log.Printf("dual token¥n")
	channelName, tokentype, uidStr, role, expireTimestamp, rtcParamErr := parseRtcParams(c)

	if rtcParamErr != nil {
		return c.JSON(400, rtcParamErr)
	}

	// generate the rtcToken
	rtcToken, err := generateRtcToken(channelName, uidStr, tokentype, role, expireTimestamp)
	if err != nil {
		return c.JSON(400, err)
	}
	// generate the rtmToken
	rtmToken, err := rtmtokenbuilder.BuildToken(appID, appCertificate, uidStr, rtmtokenbuilder.RoleRtmUser, expireTimestamp)
	if err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, map[string]string{
		"rtcToken": rtcToken,
		"rtmToken": rtmToken,
	})
}

func parseRtcParams(c echo.Context) (channelName, tokentype, uidStr string, role rtctokenbuilder.Role, expireTimestamp uint32, err error) {
	// get param values
	channelName = c.Param("channelName")
	roleStr := c.Param("role")
	tokentype = c.Param("tokentype")
	uidStr = c.Param("uid")
	expireTime := "3600"

	if roleStr == "publisher" {
		role = rtctokenbuilder.RolePublisher
	} else {
		role = rtctokenbuilder.RoleSubscriber
	}

	expireTime64, parseErr := strconv.ParseUint(expireTime, 10, 64)
	if parseErr != nil {
		err = fmt.Errorf("failed to parse expireTime: %s, causing error: %s", expireTime, parseErr)
	}

	// set timestamps
	expireTimeInSeconds := uint32(expireTime64)
	currentTimestamp := uint32(time.Now().UTC().Unix())
	expireTimestamp = currentTimestamp + expireTimeInSeconds

	return channelName, tokentype, uidStr, role, expireTimestamp, err
}

func parseRtmParams(c echo.Context) (uidStr string, expireTimestamp uint32, err error) {
	// get param values
	uidStr = c.Param("uid")
	expireTime := "3600"

	expireTime64, parseErr := strconv.ParseUint(expireTime, 10, 64)
	if parseErr != nil {
		err = fmt.Errorf("failed to parse expireTime: %s, causing error: %s", expireTime, parseErr)
	}

	// set timestamps
	expireTimeInSeconds := uint32(expireTime64)
	currentTimestamp := uint32(time.Now().UTC().Unix())
	expireTimestamp = currentTimestamp + expireTimeInSeconds

	return uidStr, expireTimestamp, err
}

func generateRtcToken(channelName, uidStr, tokentype string, role rtctokenbuilder.Role, expireTimestamp uint32) (rtcToken string, err error) {
	if tokentype == "userAccount" {
		log.Printf("Building Token with userAccount: %s¥n", uidStr)
		rtcToken, err = rtctokenbuilder.BuildTokenWithUserAccount(appID, appCertificate, channelName, uidStr, role, expireTimestamp)
		return rtcToken, err
	}
	if tokentype == "uid" {
		uid64, parseErr := strconv.ParseUint(uidStr, 10, 64)
		if parseErr != nil {
			err = fmt.Errorf("failed to parse uidStr: %s, to uint causing error: %s", uidStr, parseErr)
			return "", err
		}

		uid := uint32(uid64)
		log.Printf("Building Token with uid: %d¥n", uid)
		rtcToken, err = rtctokenbuilder.BuildTokenWithUID(appID, appCertificate, channelName, uid, role, expireTimestamp)
		return rtcToken, err
	}

	err = fmt.Errorf("failed to generate RTC token for Unknown Tokentype: %s", tokentype)
	log.Println(err)
	return "", err
}
