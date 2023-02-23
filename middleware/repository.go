package middleware

import (
	"context"
	"log"
	"net/http"
	"github.com/Go-Run-React/Catalog-Microservice-Backend/data/models"
	u "github.com/Go-Run-React/Catalog-Microservice-Backend/utils"
)

var l *log.Logger

func init() {
	l = u.Mlog
}

func Repository(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//fetching repository
		repo := models.GetRepository()

		ctx := context.WithValue(r.Context(), "repo", repo)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
  }