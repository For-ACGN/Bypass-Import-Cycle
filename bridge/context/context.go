package context

import (
	"Bypass-Import-Cycle/bridge/parameter/a"
	"Bypass-Import-Cycle/bridge/parameter/b"
)

type Context struct {
	A A
	B B
}

type A interface {
	Func_1(config *a.Config)
	Func_2(config *a.Config)
}

type B interface {
	Func_1(config *b.Config)
	Func_2(config *b.Config)
}
