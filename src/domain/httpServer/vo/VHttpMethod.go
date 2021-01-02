package vo

import (
	"errors"
	"github.com/wxnacy/wgo/arrays"
	"net/http"
)

type VHttpMethod struct {
	Method string `json:"method"`
}

func NewVHttpMethod(method string) *VHttpMethod {
	arr := []string{
		http.MethodGet,
		http.MethodHead,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodConnect,
		http.MethodOptions,
		http.MethodTrace,
	}
	if arrays.ContainsString(arr, method) == -1 {
		panic(errors.New("methods not allowed: " + method))
	}
	return &VHttpMethod{Method: method}
}
