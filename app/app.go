package app

import (
	"net/http"

	"github.com/go-kit/kit/log"

	"github.com/solher/kit-crud/library"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
)

func Builder(v *viper.Viper, l log.Logger) (http.Handler, error) {
	ctx := context.Background()
	ls := library.NewService()
	return library.MakeHTTPHandler(ctx, ls), nil
}
