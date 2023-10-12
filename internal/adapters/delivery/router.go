package delivery

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (d *DeliveryHTTP) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", d.Index).Methods(http.MethodGet)
	r.HandleFunc("/usr", d.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/usr_update", d.UpdateUser).Methods(http.MethodPost)
	return r
}
