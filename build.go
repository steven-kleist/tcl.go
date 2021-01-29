// !build
package main

import "github.com/pellared/taskflow"

func taskBuild() taskflow.Task {
	return taskflow.Task{
		Name:        "build",
		Description: "Builds the project",
		Command:     taskflow.Exec("go", "build", "-o", "./build/tcl.exe", "./cmd/tcl"),
	}
}

func main() {
	flow := taskflow.New()
	flow.MustRegister(taskBuild())
	flow.Main()
}
