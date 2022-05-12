package core

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

type TestApi struct{}

func (a TestApi) Foo(in string) (int, error) {
	res, err := strconv.Atoi(in)
	if err != nil {
		return 0, fmt.Errorf("Could not convert.")
	}
	return res, nil
}

func TestGetSignature(t *testing.T) {
	expected := Signature{}

	master := TestApi{}
	method := reflect.ValueOf(master).Method(0)

	var in string
	var succ int

	expected.Params = append(expected.Params, reflect.TypeOf(in))
	expected.Outputs = append(expected.Outputs, reflect.TypeOf(succ))
	expected.Outputs = append(expected.Outputs, reflect.TypeOf((*error)(nil)).Elem())

	res, err := GetSignature(&method)
	if err != nil {
		t.Errorf("Function failed to run.")
	}

	for i, p := range res.Params {
		if expected.Params[i] != p {
			t.Errorf("Signature parsing did not work as expected. Want: %v, Got: %v", expected.Outputs[i], p)
		}
	}

	for i, o := range res.Outputs {
		if expected.Outputs[i] != o {
			t.Errorf("Signature parsing did not work as expected. Want: %v, Got: %v", expected.Outputs[i], o)
		}
	}
}
