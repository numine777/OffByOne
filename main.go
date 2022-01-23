package main

import (
	"fmt"
	"log"
	"net/http"
	"offByOne/core"
	"reflect"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Off By One!")
	fmt.Println("Endpoint Hit: homePage")
}

func dumbPost(name string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to %v Off By One!", name)
		fmt.Println("Endpoint Hit: homePage")
	}
}

func handleRequests() {
	master := core.GetApi()
	methods := reflect.TypeOf(master)
	// http.HandleFunc("/", homePage)
	for i := 0; i < methods.NumMethod(); i++ {
		method := methods.Method(i)
		methodName := method.Name
		http.HandleFunc(fmt.Sprintf("/api/%v", methodName), superSmartPost(method))
	}
	log.Fatal(http.ListenAndServe(":6969", nil))
}

func main() {
	handleRequests()
}
