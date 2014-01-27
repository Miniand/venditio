package web

import (
	"github.com/Miniand/venditio/cmd"
	"github.com/Miniand/venditio/config"
	"github.com/Miniand/venditio/core"
	"github.com/Miniand/venditio/inject"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	DEP_ROUTER          = "httpRouter"
	CONFIG_BIND_ADDRESS = "BIND_ADDRESS"
)

func Register(v *core.Venditio) {
	v.BindFactory(DEP_ROUTER, func(i inject.Injector) interface{} {
		return mux.NewRouter()
	})
	// Config
	c := v.MustGet(config.DEP_CONFIG).(*config.Config)
	c.Defaults[CONFIG_BIND_ADDRESS] = "127.0.0.1:8080"
	// Serve command
	cmd := v.MustGet(cmd.DEP_COMMANDER).(cmd.Commander)
	cmd.Register("serve", func(args []string) error {
		return Run(v)
	})
}

func Run(v *core.Venditio) error {
	http.Handle("/", v.MustGet(DEP_ROUTER).(*mux.Router))
	addr := v.MustGet(config.DEP_CONFIG).(*config.Config).Get(
		CONFIG_BIND_ADDRESS)
	log.Printf("Listening on %s\n", addr)
	return http.ListenAndServe(addr, nil)
}
