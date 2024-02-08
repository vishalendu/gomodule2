package gomodule2

import (
	"fmt"

	"github.com/vishalendu/gomodule1"
)

func Firstchild() {
	fmt.Println("Inside First Child - function exposed")
	gomodule1.Parentfunc()
}

func firstchild1() {
	fmt.Println("Inside First Child - function not expoded")
}
