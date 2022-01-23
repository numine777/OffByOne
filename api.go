package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"offByOne/core"
	"reflect"
)

func handlePost(handler interface{}, request *http.Request, writer http.ResponseWriter) {
    if request.Body == nil {
        http.Error(writer, "Please send a request body", 400)
        return
    }
	decoder := json.NewDecoder(request.Body)

	var parsed interface{}
	err := decoder.Decode(&parsed)
	if err != nil {
		panic(err)
	}

	sig, err := core.GetSignature(handler)
	if err != nil {
		panic(err)
	}
	pMap := reflect.ValueOf(parsed)
	for i, key := range pMap.MapKeys() {
		val := pMap.MapIndex(key)
		if val.Type().String() != sig.Params[i] {
			err := fmt.Errorf("Input of type %v do not match the parameter of type %v", val.Type().String(), sig.Params[i])
			http.Error(writer, err.Error(), 500)
		}
	}
	body := reflect.ValueOf(handler).Call(pMap.MapKeys())
    json.NewEncoder(writer).Encode(body)
}

func superSmartPost (handler interface{}) func (w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        handlePost(handler, r, w)
    }
}
