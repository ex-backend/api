package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/paper/{type}", paperHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func paperHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	db, err := getTable()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
	filter := bson.D{{"type", vars["type"]}}
	var result Paper
	err = db.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
	resJSON, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", string(resJSON))
}
