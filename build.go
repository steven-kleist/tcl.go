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

func taskRun() taskflow.Task {
	return taskflow.Task{
		Name:        "run",
		Description: "Executes the build result",
		Command:     taskflow.Exec("./build/tcl.exe", "./test/test.tcl"),
	}
}

func main() {
	flow := taskflow.New()

	build := flow.MustRegister(taskBuild())
	run := flow.MustRegister(taskRun())

	flow.MustRegister(taskflow.Task{
		Name:        "all",
		Description: "build pipeline",
		Dependencies: taskflow.Deps{
			build,
			run,
		},
	})
	flow.Main()
}
