package b

import (
	"fmt"

	"Bypass-Import-Cycle/bridge/context"
	"Bypass-Import-Cycle/bridge/parameter/a"
	"Bypass-Import-Cycle/bridge/parameter/b"
)

type B struct {
	ctx *context.Context
}

func New_B(ctx *context.Context) *B {
	return &B{ctx: ctx}
}

func (this *B) Func_1(config *b.Config) {
	fmt.Println("package b receive:", config.Str)
}

func (this *B) Func_2(config *b.Config) {
	fmt.Println("package b: call B.Func_2", config.Str)
	config_a := &a.Config{Str: "hello package a! form package b"}
	this.ctx.A.Func_2(config_a)
}
