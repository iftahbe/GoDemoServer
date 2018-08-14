package Server

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"log"
)


type Server struct {
	router *mux.Router
	DocumentStoreHolder *DocumentStoreHolder
}

func NewServer(holder *DocumentStoreHolder) *Server {
	s := Server{
		router: mux.NewRouter(),
		DocumentStoreHolder: holder}
	return &s
}

func (s *Server) Init(){
	s.DocumentStoreHolder.Store.Initialize()
	s.DocumentStoreHolder.MediaStore.Initialize()

	s.router.HandleFunc("/basic/store", s.Store).Methods("GET")
	s.router.HandleFunc("/basic/edit", s.Edit).Methods("GET")
	s.router.HandleFunc("/basic/delete", s.Delete).Methods("GET")
}

func (s *Server) Start(){

	http.ListenAndServe(":80", s.router)

}

func ErrorResponse(w http.ResponseWriter, code int, message string) {
	JsonResponse(w, code, map[string]string{"error": message})
}

func JsonResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

type Company struct {
	Name		string `json:"Name"`
	ExternalId  string `json:"ExternalId"`
	Phone       string `json:"Phone"`
	Fax 		string `json:"Fax"`
}

