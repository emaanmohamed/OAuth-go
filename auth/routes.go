package auth

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/markbates/goth/gothic"
	"golang.org/x/net/context"
	"net/http"
)

type Auth struct {
	// Fields of the auth type, if any
}

func (a *Auth) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/auth/{provider}", a.getAuthCallbackFunction)

	return r

}

func (s *Auth) getAuthCallbackFunction(w http.ResponseWriter, r *http.Request) {

	fmt.Println("asdasdasd")
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Println(user)
	http.Redirect(w, r, "http://127.0.0.1:5173", http.StatusFound)

}
