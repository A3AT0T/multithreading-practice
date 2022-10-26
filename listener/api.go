package listener

import (
	"context"
)

type Listener[T any] interface {
	AddChan(<-chan T)
	Listen(context.Context) <-chan T
}

type GenericListener[T any] struct {
	chArr []<-chan T
	out   chan T
}

func NewListener[T any]() GenericListener[T] {
	return GenericListener[T]{
		chArr: make([]<-chan T, 0, 4),
		out:   make(chan T, 2),
	}
}

func (l *GenericListener[T]) AddChan(ch <-chan T) {
	l.chArr = append(l.chArr, ch)
}

func (l *GenericListener[T]) Listen(ctx context.Context) <-chan T {
	for _, ch := range l.chArr {
		go func(ch <-chan T) {
			for {
				select {
				case dataItem := <-ch:
					l.out <- dataItem
				case <-ctx.Done():
					return
				}
			}
		}(ch)
	}
	return l.out
}
