package router

import (
	"fmt"
)

type demoController struct {
}

func (c *demoController) Name() string {
	return "demo"
}

func (c *demoController) AddAction() *addAction {
	return new(addAction)
}

func (c *demoController) DelAction() *delAction {
	return new(delAction)
}

func (c *demoController) EditAction() *editAction {
	return new(editAction)
}

func (c *demoController) GetAction() *getAction {
	return new(getAction)
}

type demoAction interface {
	Name() string
	Run() error
}

type addAction struct {
}

func (a *addAction) Name() string {
	return "add"
}

func (a *addAction) Run() error {
	fmt.Println("run add")
	return nil
}

type delAction struct {
}

func (a *delAction) Name() string {
	return "del"
}

func (a *delAction) Run() error {
	fmt.Println("run del")
	return nil
}

type editAction struct {
}

func (a *editAction) Name() string {
	return "edit"
}

func (a *editAction) Run() error {
	fmt.Println("run edit")
	return nil
}

type getAction struct {
}

func (a *getAction) Name() string {
	return "get"
}

func (a *getAction) Run() error {
	fmt.Println("run get")
	return nil
}
