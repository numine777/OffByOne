package core

import (
	"fmt"
	"reflect"
)

type Signature struct {
	Params  []string
	Outputs []string
}

func GetSignature(f interface{}) (Signature, error) {
    t := reflect.TypeOf(f)
	if t.String() != "reflect.Method" {
        fmt.Printf("Received %v\n", t)
		return Signature{}, fmt.Errorf("<not a function>")
	}

	signature := Signature{}
	for i := 0; i < t.NumIn(); i++ {
		paramsSlice := signature.Params
		paramsSlice = append(paramsSlice, t.In(i).String())
		signature.Params = paramsSlice
	}
	if numOut := t.NumOut(); numOut > 0 {
		for i := 0; i < t.NumOut(); i++ {
			outputsSlice := signature.Outputs
			outputsSlice = append(outputsSlice, t.In(i).String())
			signature.Outputs = outputsSlice

		}
	}

	return signature, nil
}
