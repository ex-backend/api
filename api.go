package main

import (
	"fmt"
	"context"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo/bson"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/paper/{type}", paperHandler)
	http.Handle("/", r)
}

func paperHandler(w http.ResponseWriter, r *http.Requst) {
	vars := mux.Vars(r)
	db, err := getTable()
	if err != nil {
		w.WriteHeader(http.StatusError)
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
	filter := bson.D{{"type", mux.Vars(r)["type"]}}
	var result Paper
	err = db.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		w.WriteHeader(http.StatusError)
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
	resJSON, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusError)
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", string(resJSON))
}
