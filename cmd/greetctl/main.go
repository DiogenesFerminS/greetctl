package main

import (
	"flag"
	"fmt"

	"github.com/DiogenesFerminS/greetctl/internal/logx"
	gr "github.com/DiogenesFerminS/greetctl/pkg/greet"
	"github.com/fatih/color"
)

func main() {
	name := flag.String("name", "mundo", "name to greet")
	flag.Parse()
	msg := gr.Hello(*name)
	fmt.Println(msg)
	fmt.Println("Proyecto Funcionando DESDE CMD!!!")

	color.Green("Hello %s :)", *name)
	color.Red("Hello %s :)", *name)

	logx.Info("TEST PACKAGE INTERNAL")
}
