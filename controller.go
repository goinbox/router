package router

import (
	"github.com/goinbox/pcontext"
)

type Controller interface {
	Name() string
}

type Action[T pcontext.Context] interface {
	Name() string

	Before(ctx T)
	Run(ctx T)
	After(ctx T)
	Destruct(ctx T)
}

type BaseAction[T pcontext.Context] struct {
}

func (b *BaseAction[T]) Before(ctx T) {
}

func (b *BaseAction[T]) After(ctx T) {
}

func (b *BaseAction[T]) Destruct(ctx T) {
}
