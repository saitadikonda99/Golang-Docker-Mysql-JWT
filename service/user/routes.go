package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/saitadikonda99/Golang-Docker-Mysql-JWT/types"
	"github.com/saitadikonda99/Golang-Docker-Mysql-JWT/utils"
)


type Handler struct {
	store types.UserStore
}

 
func NewHandler() *Handler {
	return &Handler{

	}
}

func (c *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", c.handleLogin).Methods("POST");
	router.HandleFunc("/register", c.handleRegister).Methods("POST");
}

func (c *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (c *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

	var payload types.UserPayLoad  

	if err := utils.ParseJSON(*r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	
	c.store.GetUserByMail(payload.Email)
}