package Server

import "github.com/ravendb/ravendb-go-client"

type DocumentStoreHolder struct {
	Store *ravendb.DocumentStore
	MediaStore *ravendb.DocumentStore
}