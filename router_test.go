package router

import (
	"testing"

	"github.com/goinbox/pcontext"
)

func TestRouter(t *testing.T) {
	r := NewRouter()
	r.MapRouteItems(new(demoController))

	pathList := []string{
		"demo/add",
		"demo/del",
		"demo/edit",
		"demo/get",
	}

	for _, path := range pathList {
		route := r.FindRoute(path)
		if route == nil {
			t.Fatal(path, "not find route")
		}
		t.Log("controller", route.C.Name())

		vs := route.NewActionFunc.Call(nil)
		action := vs[0].Interface().(Action[pcontext.Context])
		t.Log("action", action.Name())
		err := action.Run(nil)
		t.Log(err)
	}
}
