package delivery

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (d *DeliveryHTTP) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", d.Index).Methods(http.MethodGet)
	r.HandleFunc("/usr", d.UsrCreate).Methods(http.MethodPost)
	r.HandleFunc("/usr_update", d.UsrUpdate).Methods(http.MethodPost)
	return r
}
