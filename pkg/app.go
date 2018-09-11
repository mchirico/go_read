package pkg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
	at     *int
}

func (a *App) Initilize() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", a.getRoot).Methods("GET")
	a.Router.HandleFunc("/products", a.getProducts).Methods("GET")
	// a.Router.HandleFunc("/product", a.createProduct).Methods("POST")

}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", addr), a.Router))
}

func (a *App) getRoot(w http.ResponseWriter, r *http.Request) {

	log.Printf("get Root")
	products, err := getRoot(a.DB, 0, 5)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, products)
}

func (a *App) getProducts(w http.ResponseWriter, r *http.Request) {

	products, err := getProducts(a.DB, 0, 5)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
