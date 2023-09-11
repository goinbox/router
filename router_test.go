package router

import (
	"testing"
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
		action := vs[0].Interface().(demoAction)
		t.Log("action", action.Name())
		err := action.Run()
		t.Log(err)
	}
}
