package user

import (
	"context"
	"errors"
	"time"

	"github.com/fakriardian/Go-kelas.work/src/model"
	"github.com/fakriardian/Go-kelas.work/src/tracing"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.StandardClaims
}

func (ur *userRepo) CreateUserSession(ctx context.Context, userID string) (model.UserSession, error) {
	ctx, span := tracing.CreateSpan(ctx, "CreateUserSession")
	defer span.End()

	accessToken, err := ur.generateAccessToken(ctx, userID)
	if err != nil {
		return model.UserSession{}, err
	}

	return model.UserSession{
		JWTToken: accessToken,
	}, nil

}

func (ur *userRepo) generateAccessToken(ctx context.Context, userID string) (string, error) {
	_, span := tracing.CreateSpan(ctx, "generateAccessToken")
	defer span.End()

	accessTokenExp := time.Now().Add(ur.accessExp).Unix()

	accessClaim := Claims{
		jwt.StandardClaims{
			Subject:   userID,
			ExpiresAt: accessTokenExp,
		},
	}

	accessJwt := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), accessClaim)

	return accessJwt.SignedString(ur.signKey)
}

func (ur *userRepo) CheckSession(ctx context.Context, data model.UserSession) (userID string, err error) {
	_, span := tracing.CreateSpan(ctx, "CheckSession")
	defer span.End()

	accessToken, err := jwt.ParseWithClaims(data.JWTToken, Claims{}, func(t *jwt.Token) (interface{}, error) {
		return &ur.signKey.PublicKey, nil
	})

	if err != nil {
		return "", err
	}

	accessTokenClaims, ok := accessToken.Claims.(*Claims)
	if !ok {
		return "", errors.New("unauthorized")
	}

	if accessToken.Valid {
		return accessTokenClaims.Subject, nil
	}

	return "", errors.New("unauthorized")
}
