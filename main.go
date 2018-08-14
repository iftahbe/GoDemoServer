package main

import (
	"github.com/ravendb/ravendb-go-client"
	"github.com/GoDemoServer/Server"
)

func main() {
	holder := Server.DocumentStoreHolder{
		Store: ravendb.NewDocumentStoreWithUrlAndDatabase("http://localhost:8080","Northwind"),
		MediaStore: ravendb.NewDocumentStoreWithUrlAndDatabase("http://localhost:8080","Media")}

	defer holder.Store.Close()
	defer holder.MediaStore.Close()

	server := Server.NewServer(&holder)
	server.Init()
	server.Start()

}

