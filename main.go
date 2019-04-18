package main

import (
	"fmt"

	"Bypass-Import-Cycle/a"
	"Bypass-Import-Cycle/b"
	p_a "Bypass-Import-Cycle/bridge/a"
	p_b "Bypass-Import-Cycle/bridge/b"
	"Bypass-Import-Cycle/bridge/context"
)

func main() {
	c := &context.Context{}
	oa := a.New_A(c)
	ob := b.New_B(c)
	c.A = oa
	c.B = ob
	oa.Func_1(&p_a.Config{Str: "aaa"})
	fmt.Println("---------------")
	ob.Func_2(&p_b.Config{Str: "bbb"})
}
