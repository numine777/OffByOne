package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"offByOne/core"
	"reflect"
)

func handlePost(handler *reflect.Value, request *http.Request, writer http.ResponseWriter) {
	if request.Body == nil {
		http.Error(writer, "Please send a request body", 400)
		return
	}
	sig, err := core.GetSignature(handler)
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(request.Body)

	parsed := reflect.New(sig.Params[0])
	elem := parsed.Interface()
	err = decoder.Decode(&elem)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", parsed.Elem())

	params := make([]reflect.Value, len(sig.Params))
	params[0] = parsed.Elem()

	body := (*handler).Call(params)
	fmt.Printf(body[0].String())
	json.NewEncoder(writer).Encode(body)
}

func superSmartPost(handler *reflect.Value) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		handlePost(handler, r, w)
	}
}
