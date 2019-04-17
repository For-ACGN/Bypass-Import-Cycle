package a

import (
	"fmt"

	"Bypass-Import-Cycle/bridge/context"
	"Bypass-Import-Cycle/bridge/parameter/a"
	"Bypass-Import-Cycle/bridge/parameter/b"
)

type A struct {
	ctx *context.Context
}

func New_A(ctx *context.Context) *A {
	return &A{ctx: ctx}
}

func (this *A) Func_1(config *a.Config) {
	fmt.Println("package a: call A.Func_1", config.Str)
	config_b := &b.Config{Str: "hello package b! form package a"}
	this.ctx.B.Func_1(config_b)
}

func (this *A) Func_2(config *a.Config) {
	fmt.Println("package a receive:", config.Str)
}
