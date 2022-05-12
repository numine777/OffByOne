package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"github.com/numine777/OffByOne/core"
	"reflect"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Off By One!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests(db *gorm.DB) {
	master := core.GetApi(db)
	methods := reflect.TypeOf(master)
	for i := 0; i < methods.NumMethod(); i++ {
		method := reflect.ValueOf(master).Method(i)
		methodName := methods.Method(i).Name
		http.HandleFunc(fmt.Sprintf("/api/%v", methodName), superSmartPost(&method))
	}
    http.HandleFunc(fmt.Sprintf("/"), homePage)
	log.Fatal(http.ListenAndServe(":6969", nil))
}

func main() {
	db, err := gorm.Open(sqlite.Open("testdb.sqlite3"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	handleRequests(db)
}
