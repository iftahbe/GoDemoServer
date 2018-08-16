package Server

import (
	"net/http"
	"github.com/ravendb/ravendb-go-client"
	"github.com/GoDemoServer/Entity"
)

func (s *Server) Store(w http.ResponseWriter, r *http.Request) {

	session, err := s.DocumentStoreHolder.Store.OpenSession()
	if err != nil {
		ErrorResponse(w, 500, err.Error())
		return
	}
	defer session.Close()

	c := Entity.Company{
		Name:       "Hibernating Rhinos",
		ExternalId: "HR",
		Phone:      "+972 4 622 7811",
		Fax:        "+972 153 4 622 7811"}

	session.Store(c)

	if err = session.SaveChanges(); err != nil {
		ErrorResponse(w, 500, err.Error())
		return
	}

	JsonResponse(w, 201, map[string]string{"Result": "Success"})
}

func (s *Server) Edit(w http.ResponseWriter, r *http.Request) {
	session, err := s.DocumentStoreHolder.Store.OpenSession()
	if err != nil {
		ErrorResponse(w, 500, err.Error())
		return
	}
	defer session.Close()

	// pluralize() is not fully implemented, this needs to change to companies/1-A
	obj, err := session.Load(ravendb.GetTypeOf(&Entity.Company{}), "companys/1-A")
	if err != nil {
		ErrorResponse(w, 500, err.Error())
		return
	}

	doc, ok := obj.(*Entity.Company)
	if !ok{
		ErrorResponse(w, 500, "Type assertion failed")
		return
	}
	doc.Name = "Hibernating Elephants"

	err3 := session.SaveChanges()
	if err3 != nil {
		ErrorResponse(w, 500, err3.Error())
		return
	}

	JsonResponse(w, 201, map[string]string{"Result": "Success"})
}

func (s *Server) Delete(w http.ResponseWriter, r *http.Request) {
	session, err := s.DocumentStoreHolder.Store.OpenSession()
	if err != nil {
		ErrorResponse(w, 500, err.Error())
		return
	}
	defer session.Close()

	// pluralize() is not fully implemented, this needs to change to companies/1-A
	obj, err := session.Load(ravendb.GetTypeOf(&Entity.Company{}), "companys/1-A")
	if err != nil {
		ErrorResponse(w, 500, err.Error())
		return
	}

	err = session.DeleteEntity(obj)
	if err != nil {
		ErrorResponse(w, 500, err.Error())
		return
	}

	err3 := session.SaveChanges()
	if err3 != nil {
		ErrorResponse(w, 500, err3.Error())
		return
	}

	JsonResponse(w, 201, map[string]string{"Result": "Success"})
}

