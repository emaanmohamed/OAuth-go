package auth

import (
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"log"
	"os"
)

const (
	key    = "randomStr"
	MaxAge = 86400 * 30
	IsProd = false
)

func NewAuth() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	cookieStore := sessions.NewCookieStore([]byte(key))
	cookieStore.MaxAge(MaxAge)

	cookieStore.Options.Path = "/"
	cookieStore.Options.HttpOnly = true
	cookieStore.Options.Secure = IsProd

	gothic.Store = cookieStore
	goth.UseProviders(
		google.New(googleClientId, googleClientSecret, "http://127.0.0.1:3000/auth/google/callback"),
	)

}
