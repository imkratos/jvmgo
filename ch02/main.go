package main

import "fmt"
import "strings"
import "github.com/imkratos/jvmgo/ch02/classpath"

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJvm(cmd)
	}

	a := "com/sun/javafx/tools/ant/Antlog.class"
	b := "com/sun/javafx/tools/ant/AntLog.class"
	fmt.Println(strings.EqualFold(a, b))
}

func startJvm(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n", cp, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	fmt.Printf("className: %s", className)
	classData, _, err := cp.ReadClass(className)

	if err != nil {
		fmt.Printf("Could not find or load main class %s \n", cmd.class)
		return
	}

	fmt.Printf("class data: %v\n", classData)

}
