package Server

import (
	"net/http"
)

func (s *Server) Store(w http.ResponseWriter, r *http.Request) {

	session := s.DocumentStoreHolder.Store.OpenSession()
	defer session.Close()

	c := Company{
		Name:       "Hibernating Rhinos",
		ExternalId: "HR",
		Phone:      "+972 4 622 7811",
		Fax:        "+972 153 4 622 7811"}

	session.Store(c)

	err := session.SaveChanges()

	if err != nil {
		ErrorResponse(w, 500, err.Error())
		return
	}

	JsonResponse(w, 201, map[string]string{"Result": "Success"})
}

func (s *Server) Edit(w http.ResponseWriter, r *http.Request) {
	session := s.DocumentStoreHolder.Store.OpenSession()
	defer session.Close()

	var result *Company
	err := session.Load(&result, "companies/1")
	if err != nil {
		ErrorResponse(w, 500, err.Error())
		return
	}

	result.Name = "Hibernating Bears"

	err3 := session.SaveChanges()
	if err3 != nil {
		ErrorResponse(w, 500, err3.Error())
		return
	}

	JsonResponse(w, 201, map[string]string{"Result": "Success"})
}

func (s *Server) Delete(w http.ResponseWriter, r *http.Request) {
	session := s.DocumentStoreHolder.Store.OpenSession()
	defer session.Close()

	var result *Company

	err := session.Load(&result, "companies/1")
	if err != nil {
		ErrorResponse(w, 500, err.Error())
		return
	}

	err2 := session.DeleteEntity(result)
	if err2 != nil {
		ErrorResponse(w, 500, err2.Error())
		return
	}

	err3 := session.SaveChanges()
	if err3 != nil {
		ErrorResponse(w, 500, err3.Error())
		return
	}

	JsonResponse(w, 201, map[string]string{"Result": "Success"})
}