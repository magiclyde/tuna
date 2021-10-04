/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2021/10/4 下午11:42
 * @refer: https://www.nexmo.com/legacy-blog/2020/09/07/using-jwt-for-authentication-in-a-golang-application-dr-2
 */

package util

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v7"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"strings"
	"time"
)

// data stored in client side
type TokenDetails struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AccessUuid   string `json:"access_uuid"`
	RefreshUuid  string `json:"refresh_uuid"`
	UserId       string `json:"user_id"`
	AtExpires    int64  `json:"at_expires"`
	RtExpires    int64  `json:"rt_expires"`
}

// data stored in server side
type AccessDetails struct {
	AccessUuid string `json:"access_uuid,omitempty"`
	UserId     string `json:"user_id,omitempty"`
}

type Jwt struct {
	AccessKey  string
	RefreshKey string
	Prefix     string
	Client     *redis.Client
}

func NewJwt(accessKey, refreshKey, prefix string, client *redis.Client) *Jwt {
	if len(prefix) <= 0 {
		prefix = "jwt:"
	}
	return &Jwt{
		AccessKey:  accessKey,
		RefreshKey: refreshKey,
		Prefix:     prefix,
		Client:     client,
	}
}

// CreateJwtToken create access N refresh token
func (j *Jwt) CreateJwtToken(userId string) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.UserId = userId

	td.AtExpires = time.Now().Add(time.Hour * 2).Unix()
	td.AccessUuid = StrBuilder(j.Prefix, userId)

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = StrBuilder(j.Prefix, uuid.New().String())

	var err error
	// Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userId
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(j.AccessKey))
	if err != nil {
		return nil, err
	}

	// Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userId
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(j.RefreshKey))
	if err != nil {
		return nil, err
	}
	return td, nil
}

// CreateAuth
func (j *Jwt) CreateAuth(td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) // converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	details, err := jsoniter.MarshalToString(AccessDetails{
		UserId: td.UserId,
	})
	if err != nil {
		return err
	}

	errAccess := j.Client.Set(td.AccessUuid, details, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}

	errRefresh := j.Client.Set(td.RefreshUuid, details, rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}

	return nil
}

// DeleteAuth
func (j *Jwt) DeleteAuth(givenUuid string) (int64, error) {
	deleted, err := j.Client.Del(givenUuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

// ExtractJwtToken extract access_token
func ExtractJwtToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// VerifyJwtToken verify access_token
func (j *Jwt) VerifyJwtToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractJwtToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.AccessKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// JwtTokenValid
func (j *Jwt) JwtTokenValid(r *http.Request) error {
	token, err := j.VerifyJwtToken(r)
	if err != nil {
		return fmt.Errorf("JwtTokenValid, verifyJwtToken err: %s", err.Error())
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return errors.New("JwtTokenValid, invalid token")
	}
	return nil
}

// ExtractJwtTokenMetadata extract metadata from access_token
func (j *Jwt) ExtractJwtTokenMetadata(r *http.Request) (string, error) {
	token, err := j.VerifyJwtToken(r)
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return "", errors.New("ExtractJwtTokenMetadata, invalid payload")
		}
		return accessUuid, nil
	}
	return "", err
}

// FetchAuth
func (j *Jwt) FetchAuth(givenUuid string) (*AccessDetails, error) {
	details, err := j.Client.Get(givenUuid).Result()
	if err != nil {
		return nil, err
	}

	var authD AccessDetails
	if err := jsoniter.UnmarshalFromString(details, &authD); err != nil {
		return nil, err
	}
	return &authD, nil
}
