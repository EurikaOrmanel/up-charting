package utils

import (
	"EurikaOrmanel/up-charter/internal/cache"
	"EurikaOrmanel/up-charter/internal/schemas"
	"os"
	"strconv"
	"time"

	"EurikaOrmanel/up-charter/internal/models"
)

func GenerateAuthResponse(cacheConfig cache.CacheConfig, userAuth models.Admin) (schemas.AuthResponse, error) {
	refreshToken, err := GenerateRefreshToken()
	if err != nil {
		return schemas.AuthResponse{}, err
	}
	accessExpiresInMin, err := strconv.Atoi(os.Getenv("ACCESS_EXPIRES_IN_MIN"))
	if err != nil {
		accessExpiresInMin = 1
	}

	refreshExpiresInMin, err := strconv.Atoi(os.Getenv("REFRESH_EXPIRES_IN_MIN"))

	if err != nil {
		refreshExpiresInMin = 15
	}

	cacheConfig.DeleteTokenCache(userAuth.ID.String())
	cacheConfig.SetRefreshToken(userAuth.ID.String(), refreshToken)
	refreshExpiresAt := time.Now().Add(time.Minute * time.Duration(refreshExpiresInMin))
	accessToken, err := GenerateAccessToken(userAuth, accessExpiresInMin)
	respBody := schemas.AuthResponse{
		AccessToken:           &accessToken,
		RefreshToken:          &refreshToken,
		AccessExpiresAtEpoach: time.Now().Add(1 * time.Minute).Unix(),

		RefreshTokenExpiresAtEpoach: refreshExpiresAt.Unix(),
	}
	return respBody, nil
}
