package jwtAuth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"goiot/pkg/cache"
	"goiot/pkg/conf"
	"goiot/pkg/logger"
	"goiot/pkg/utils"
	"time"
)

type JWTConfig struct {
	SigningKey   string `mapstructure:"signedKey"`
	RTokenExpire string `mapstructure:"rTokenExpire"`
	ATokenExpire string `mapstructure:"aTokenExpire"`
	Issuer       string `mapstructure:"issuer"`
}

type TokenInfo struct {
	AccessToken      string
	AccessExpiresAt  int64
	RefreshToken     string
	RefreshExpiresAt int64
}

var jwtConfig = JWTConfig{
	SigningKey:   "goiot",
	RTokenExpire: "720h",
	ATokenExpire: "1h",
	Issuer:       "goiot",
}

func InitJWT() error {
	return conf.Conf.UnmarshalKey("jwt", &jwtConfig)
}

type UserClaims struct {
	UserId    string `json:"userId"`
	TokenType string `json:"tokenType"`
	jwt.RegisteredClaims
}

func (u *UserClaims) Valid() error {
	exist, err := cache.GetRedis(cache.PermissionDB).Exists(cache.BlackListKey + u.RegisteredClaims.ID).Result()
	if err != nil {
		return err
	}
	if exist > 0 {
		return errors.New("token is in the blacklist")
	}
	return nil
}

func GenToken(userID string) (*TokenInfo, error) {
	rtExpire, err := time.ParseDuration(jwtConfig.RTokenExpire)
	if err != nil {
		return nil, err
	}
	atExpire, err := time.ParseDuration(jwtConfig.ATokenExpire)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	ac := &UserClaims{
		UserId:    userID,
		TokenType: "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(atExpire)),
			Issuer:    jwtConfig.Issuer,
			ID:        utils.GetUUIDFull(),
		},
	}
	rc := &UserClaims{
		UserId:    userID,
		TokenType: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(rtExpire)),
			Issuer:    jwtConfig.Issuer,
			ID:        utils.GetUUIDFull(),
		},
	}

	secret := []byte(jwtConfig.SigningKey)
	aToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, ac).SignedString(secret)
	if err != nil {
		return nil, err
	}
	rToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, rc).SignedString(secret)
	if err != nil {
		return nil, err
	}

	return &TokenInfo{
		AccessToken:      aToken,
		AccessExpiresAt:  ac.ExpiresAt.Time.Unix(),
		RefreshToken:     rToken,
		RefreshExpiresAt: rc.ExpiresAt.Time.Unix(),
	}, nil
}

func ParseToken(tokenStr string) (*UserClaims, error) {
	uc := new(UserClaims)
	_, err := jwt.ParseWithClaims(tokenStr, uc, keyFunc)
	if err != nil {
		logger.Logger.Errorf("ParseWithClaims error: %v", err)
		return uc, err
	}
	return uc, nil
}

func RefreshToken(aToken, rToken string) (*TokenInfo, error) {
	_, err := jwt.Parse(rToken, keyFunc)
	if err != nil {
		return nil, err
	}

	uc := new(UserClaims)
	parse, err := jwt.ParseWithClaims(aToken, uc, keyFunc)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) || !parse.Valid {
			return GenToken(uc.UserId)
		}
		return nil, err
	}

	logger.Logger.Debugf("access token %s has not expired, but still refresh", aToken)
	return GenToken(uc.UserId)
}

func AddTokenToBlacklist(tokenId string, duration time.Duration) error {
	return cache.GetRedis(cache.PermissionDB).Set(cache.BlackListKey+tokenId, "invalid", duration).Err()
}

func keyFunc(token *jwt.Token) (any, error) {
	return []byte(jwtConfig.SigningKey), nil
}

func GetTokenRemainDuration(tokenStr string) (time.Duration, string, error) {
	uc := new(UserClaims)
	_, err := jwt.ParseWithClaims(tokenStr, uc, keyFunc)
	if err != nil {
		logger.Logger.Errorf("ParseWithClaims error: %v", err)
		return 0, "", err
	}

	now := time.Now()
	expiresAt := uc.RegisteredClaims.ExpiresAt
	remainDuration := expiresAt.Sub(now)

	if remainDuration <= 0 {
		logger.Logger.Warnf("oken has already expired")
		return 0, "", nil
	}

	return remainDuration, uc.RegisteredClaims.ID, nil
}
