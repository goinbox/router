package router

import (
	"github.com/goinbox/pcontext"
)

type Controller interface {
	Name() string
}

type Action[T pcontext.Context] interface {
	Name() string

	Before(ctx T) error
	Run(ctx T) error
	After(ctx T) error
	Destruct(ctx T) error
}

type BaseAction[T pcontext.Context] struct {
}

func (b *BaseAction[T]) Before(ctx T) error {
	return nil
}

func (b *BaseAction[T]) After(ctx T) error {
	return nil
}

func (b *BaseAction[T]) Destruct(ctx T) error {
	return nil
}
