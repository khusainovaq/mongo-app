package hanlder

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"mongo-l3/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(HandlerResp(http.StatusBadRequest, err.Error()))
	}

	var u entity.User
	json.Unmarshal(body, &u)

	if err = h.services.CreateUser(u); err != nil {
		json.NewEncoder(w).Encode(HandlerResp(http.StatusInternalServerError, err.Error()))
		return
	}

	json.NewEncoder(w).Encode(HandlerResp(http.StatusOK, "Created"))
	return
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var u interface{}

	body, _ := ioutil.ReadAll(r.Body)
	bson.UnmarshalExtJSON(body, true, &u)

	u, err := h.services.GetUser(u)
	if err != nil {
		json.NewEncoder(w).Encode(HandlerResp(http.StatusInternalServerError, err.Error()))
		return
	}

	json.NewEncoder(w).Encode(HandlerResp(http.StatusOK, u))
	return
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(HandlerResp(http.StatusBadRequest, err.Error()))
	}

	var u entity.User
	json.Unmarshal(body, &u)

	if err = h.services.EditUser(u.Username, u); err != nil {
		json.NewEncoder(w).Encode(HandlerResp(http.StatusInternalServerError, err.Error()))
		return
	}

	json.NewEncoder(w).Encode(HandlerResp(http.StatusOK, "Updated"))
	return
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var u interface{}

	body, _ := ioutil.ReadAll(r.Body)
	bson.UnmarshalExtJSON(body, true, &u)

	if err := h.services.DeleteUser(u); err != nil {
		json.NewEncoder(w).Encode(HandlerResp(http.StatusInternalServerError, err.Error()))
		return
	}

	json.NewEncoder(w).Encode(HandlerResp(http.StatusOK, "Deleted"))
	return
}

func HandlerResp(code int, message interface{}) *map[string]interface{} {
	return &map[string]interface{}{
		"code":    code,
		"message": message,
	}
}
