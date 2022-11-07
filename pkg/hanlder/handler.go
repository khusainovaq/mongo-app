package hanlder

import (
	"mongo-l3/pkg/service"

	"github.com/gorilla/mux"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/sign-up", h.CreateUser).Methods("POST")
	r.HandleFunc("/sign-in", h.GetUser).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	{
		api.HandleFunc("/update", h.UpdateUser).Methods("PUT", "POST")
		api.HandleFunc("/delete", h.DeleteUser).Methods("DELETE")
	}

	return r
}
