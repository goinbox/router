package router

import (
	"fmt"

	"github.com/goinbox/pcontext"
)

type demoController struct {
}

func (c *demoController) Name() string {
	return "demo"
}

func (c *demoController) AddAction() Action[pcontext.Context] {
	return new(addAction)
}

func (c *demoController) DelAction() Action[pcontext.Context] {
	return new(delAction)
}

func (c *demoController) EditAction() Action[pcontext.Context] {
	return new(editAction)
}

func (c *demoController) GetAction() Action[pcontext.Context] {
	return new(getAction)
}

type addAction struct {
	BaseAction[pcontext.Context]
}

func (a *addAction) Name() string {
	return "add"
}

func (a *addAction) Run(ctx pcontext.Context) {
	fmt.Println("run add")
}

type delAction struct {
	BaseAction[pcontext.Context]
}

func (a *delAction) Name() string {
	return "del"
}

func (a *delAction) Run(ctx pcontext.Context) {
	fmt.Println("run del")
}

type editAction struct {
	BaseAction[pcontext.Context]
}

func (a *editAction) Name() string {
	return "edit"
}

func (a *editAction) Run(ctx pcontext.Context) {
	fmt.Println("run edit")
}

type getAction struct {
	BaseAction[pcontext.Context]
}

func (a *getAction) Name() string {
	return "get"
}

func (a *getAction) Run(ctx pcontext.Context) {
	fmt.Println("run get")
}
