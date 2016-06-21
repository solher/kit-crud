package app

import (
	"net/http"

	"github.com/go-kit/kit/log"

	"github.com/pressly/chi"
	"github.com/spf13/viper"
)

func Builder(v *viper.Viper, l log.Logger) (http.Handler, error) {
	router := chi.NewRouter()
	return router, nil
}
