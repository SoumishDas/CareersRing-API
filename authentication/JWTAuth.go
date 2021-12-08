package authentication

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//jwt service
type JWTService interface {
	GenerateToken(email string, isUser bool,userId uint64) Tokens
	ValidateToken(token string) (*jwt.Token, error)
}
type    authCustomClaims struct {
	Name string `json:"name"`
	User bool   `json:"user"`
	
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}
type Tokens struct {
	AccessToken string
	RefreshToken    string
}

//auth-jwt
func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issure:    "CareersRing",
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtServices) GenerateToken(email string, isUser bool,userId uint64) Tokens {
	claims := &authCustomClaims{
		email,
		isUser,
		
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
			Subject: strconv.Itoa(int(userId)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		println(err)
	}
	
	rtClaims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 120).Unix(),
		Issuer:    service.issure,
		IssuedAt:  time.Now().Unix(),
		Subject: strconv.Itoa(int(userId)),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	rt, err := refreshToken.SignedString([]byte(service.secretKey))
	if err != nil {
		println(err)
	}



	return Tokens{AccessToken: t,RefreshToken: rt}
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})

}