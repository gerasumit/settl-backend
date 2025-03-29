package services

import (
    "context"
    "errors"
    "os"
    "time"

    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "google.golang.org/api/idtoken"
)

type GoogleLoginRequest struct {
    IDToken string `json:"id_token" binding:"required"`
}

type Claims struct {
    Email string `json:"email"`
    jwt.StandardClaims
}

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func HandleGoogleLogin(request GoogleLoginRequest) (string, error) {
    payload, err := idtoken.Validate(context.Background(), request.IDToken, os.Getenv("GOOGLE_CLIENT_ID"))
    if err != nil {
        return "", errors.New("invalid Google token")
    }

    email, ok := payload.Claims["email"].(string)
    if !ok {
        return "", errors.New("failed to parse email from token")
    }

    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        Email: email,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedToken, err := token.SignedString(jwtKey)
    if err != nil {
        return "", errors.New("failed to generate JWT")
    }

    return signedToken, nil
}

func HandleLogout(c *gin.Context) error {
    // In a stateless JWT setup, logout is typically handled client-side by deleting the token.
    return nil
}