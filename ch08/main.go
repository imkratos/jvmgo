package main

import "fmt"
import (
	"github.com/imkratos/jvmgo/ch08/classpath"
	"github.com/imkratos/jvmgo/ch08/rtda/heap"
	"strings"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJvm(cmd)
	}
}

func startJvm(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	classLoader := heap.NewClassLoader(cp, cmd.verboseClassFlag)

	className := strings.Replace(cmd.class, ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()

	if mainMethod != nil {
		interpret(mainMethod, cmd.verboseClassFlag, cmd.args)
	} else {
		fmt.Printf("Main method not found in class %s \n", cmd.class)
	}
}
