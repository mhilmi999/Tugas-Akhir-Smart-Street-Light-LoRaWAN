package jwt

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/app/config"
)

func ValidationToken(tokenEncoded string) (*jwt.Token, error){
	conf, _ := config.Init()
	token, err := jwt.Parse(tokenEncoded, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid token")
		}

		return []byte(conf.App.Secret_key), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}

func SpawnToken(userID int) (string, error) {
	conf, _ := config.Init()
	SECRET_KEY := []byte(conf.App.Secret_key)
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}