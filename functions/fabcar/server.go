package main

import (
	"fmt"
	"log"
	"strings"
	"net/http"
	"encoding/json"
	"fabcar/functions"
)

func main() {
	//queryallcars := func(w http.ResponseWriter, _ *http.Request) {
	//	
	//	js, err := json.Marshal(functions.QueryAll())
	//	
	//	if err != nil {
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//	cleanjs := strings.Replace(string(js), "\\", "", -1)
	//	w.Write([]byte(strings.Replace(cleanjs, "}},", "}}\n", -1)))
	//}

	querycar := func(w http.ResponseWriter, r *http.Request) {
		
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
	   	}
	   	carID := r.PostFormValue("carID")
		js, err := json.Marshal(functions.GetMachine(carID))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		cleanjs := strings.Replace(string(js), "\\", "", -1)
		
		if cleanjs == "" {
			w.Write([]byte("Car does not exist"))
			return
		}

		w.Write([]byte(strings.Replace(cleanjs, "}},", "}}\n", -1)))
	}

	createcar := func(w http.ResponseWriter, r *http.Request) {
		
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
	   	}
		
		carid := r.PostFormValue("carid")
		make := r.PostFormValue("make")
		model := r.PostFormValue("model")
		colour := r.PostFormValue("colour")
		owner := r.PostFormValue("owner")

		js, err := json.Marshal(functions.NewMachine(carid, make, model, colour, owner))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		w.Write(js)
	}

	//changeowner := func(w http.ResponseWriter, r *http.Request) {
	//	
	//	err := r.ParseForm()
	//	if err != nil {
	//		fmt.Println(err)
	// 	}
	//	
	//	carid := r.PostFormValue("carid")
	//	owner := r.PostFormValue("owner")
	//
	//	js, err := json.Marshal(functions.ChangeOwner(carid, owner))
	//
	//	if err != nil {
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//	
	//	w.Write(js)
	//}

	http.Handle("/", http.FileServer(http.Dir("./static")))
	//http.HandleFunc("/queryallcars", queryallcars)
	http.HandleFunc("/querycar", querycar)
	http.HandleFunc("/createcar", createcar)
	//http.HandleFunc("/changeowner", changeowner)

	log.Fatal(http.ListenAndServe(":8080", nil))
}