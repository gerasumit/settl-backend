package services

import (
	"context"
	"errors"
	"fmt"
	"os"

	"settl-backend/config"
	"settl-backend/token"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/idtoken"
)

type GoogleLoginRequest struct {
	IDToken string `json:"id_token" binding:"required"`
	Email   string `json:"email" binding:required`
}

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func HandleGoogleLogin(config config.Config, request GoogleLoginRequest) (string, error) {
	tokenMaker, err := token.GenerateJWTMaker(config.TokenSymmetricKey)

	if err != nil {
		return "", fmt.Errorf("Cannot create token maker")
	}

	payload, err := idtoken.Validate(context.Background(), request.IDToken, config.GoogleClientID)

	if err != nil {
		return "", errors.New("invalid Google token")
	}

	email, ok := payload.Claims["email"].(string)
	if !ok {
		return "", errors.New("failed to parse email from token")
	}

	signedToken, err := tokenMaker.CreateToken(email, config.AccessTokenDuration)

	if err != nil {
		return "", errors.New("failed to generate JWT")
	}

	return signedToken, nil
}

func HandleLogout(c *gin.Context) error {
	// In a stateless JWT setup, logout is typically handled client-side by deleting the token.
	return nil
}
