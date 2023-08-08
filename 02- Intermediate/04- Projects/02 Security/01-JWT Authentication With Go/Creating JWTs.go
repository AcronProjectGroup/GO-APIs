import (
	"github.com/golang-jwt/jwt"
	"<your-module>/service"
	"log"
	"time"
  )
  
  // user login validation should occur here
  
  userClaims := service.UserClaims{
   Id:    "01H4EKGQSY5637MQP395283JR8",
   First: "Leeroy",
   Last:  "Jenkins",
   StandardClaims: jwt.StandardClaims{
	IssuedAt:  time.Now().Unix(),
	ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
   },
  }
  
  signedAccessToken, err := service.NewAccessToken(userClaims)
  if err != nil {
	log.Fatal("error creating access token")
  }
  
  refreshClaims := jwt.StandardClaims{
	IssuedAt:  time.Now().Unix(),
	ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
  }
  
  signedRefreshToken, err := service.NewRefreshToken(refreshClaims)
  if err != nil {
	log.Fatal("error creating refresh token")
  }
  
  // do something with access, and refresh tokens.
  // i.e. return them in an HTTP response for a successful login