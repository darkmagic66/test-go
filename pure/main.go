package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type test struct {
	Name string `json:"name"`
	Test string `json:"test"`
}

var arr []test

// fix order parameter
func moviesHandler(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	fmt.Fprintf(w, "hello %s ", method)

	if method == "POST" {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(w, "error: %v", err)
			return
		}
		fmt.Println(string(body))
		obj := test{}
		// json -> obj
		err = json.Unmarshal(body, &obj)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "error: %v", err)
			return
		}
		arr = append(arr, obj)
		return
	}
	// string  struct
	// fmt.Fprintf(w, "arr %v",arr)
	obj_json, err := json.Marshal(arr)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	// stirng that look like json
	fmt.Fprintf(w, "arr_json %v", string(obj_json))

	// change contentype to json
	w.Header().Set("Conten-Type", "application/json; charset=utf-8")
	w.Write(obj_json)
}

func main() {
	http.HandleFunc("/root", moviesHandler)

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
