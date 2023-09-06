package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"umbrella.github.com/advanced_go/advanced_6/key-value-database/db"
)

var mdb *db.Driver

func init() {
	mdb, _ = db.New("db")
}

const listenPort = ":3010"

func requestHandler(rw http.ResponseWriter, req *http.Request) {
	urlPart := strings.Split(req.URL.Path, "/")
	var err error
	var document string
	var resource string

	if len(urlPart) == 4 {
		document = urlPart[2]
		resource = urlPart[3]
	}

	if len(urlPart) == 3 {
		document = urlPart[2]
	}

	switch req.Method {
	case http.MethodGet:
		var v *json.RawMessage
		err = mdb.Read(document, resource, &v)
		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
		} else {
			rw.Write(*v)
		}
	case http.MethodDelete:
		err = mdb.Delete(document, resource)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusNoContent)
			return
		}
		rw.WriteHeader(http.StatusOK)
		return
	case http.MethodPut:
		v, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusNoContent)
			return
		}
		err = mdb.Write(document, resource, v)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusNoContent)
			return
		}
		rw.WriteHeader(http.StatusOK)
		return
	default:
		rw.WriteHeader(http.StatusExpectationFailed)
		return
	}
}

func main() {
	http.HandleFunc("/db/", requestHandler)
	go func() {
		sigchan := make(chan os.Signal, 10)
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan
		mdb.Close()
		os.Exit(0)
	}()

	http.ListenAndServe(listenPort, nil)
}
