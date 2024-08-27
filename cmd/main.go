package main

import (
	"fmt"
	"github/xclamation/logger-project/logger"
	"io"
	"os"
	"time"
)

func main() {
	var startEnable bool = true
	var startOutput io.Writer = os.Stdout
	// Do not use the same names for variables and imported packages.
	// 1) Renaming the variable: log, appLog, ...
	// - when the variable name isn't critical for understanding
	// 2) Alias for the package: logpkg "github.com/username/projectname/logger"
	// - when the variable name is critical or the package name isn't used often
	// 3) Contextual Naming: requestLogger, errorLogger, ...
	// - when the context of the logger is important for understanding the code
	log := logger.NewLogger(logger.WithEnabled(startEnable),
		logger.WithOutput(startOutput))
	// We can use var log *logger.Logger = ..., but this is not a common approach.
	// Also we can create type Logger = logger.Logger outside main function
	// to more short name but NOT RECOMENDED.

	// Log custom text
	log.Log("Custom log message in console: Hello World!")

	// Mark the beginnging of an execution block
	log.Begin()

	// Simulate som processing
	time.Sleep(2 * time.Second)

	//Mark the end of an execution block
	log.End()

	//Disable logging
	log.Disable()

	// This message won't be logged since logging is disabled
	log.Log("This won't be logged.")

	// Change the output to a file
	file, err := os.Create("logfile.txt")
	if err != nil {
		fmt.Println("Error creating log file:", err)
		return
	}
	defer file.Close()

	// You should close files after any interactions with them that can cause an error.
	// func CopyFile(dstName, srcName string) (written int64, err error) {
	// 	src, err := os.Open(srcName)
	// 	if err != nil {
	// 		return
	// 	}
	// 	defer src.Close()

	// 	dst, err := os.Create(dstName)
	//  // If there will be an error and function will return nil then "defer src.Close" will be executed.
	// 	// That's why we do not put "defer src.Close" after src's opening and not in the end of surrounding function.
	// 	if err != nil {
	// 		return
	// 	}
	// 	defer dst.Close()

	// 	return io.Copy(dst, src)
	// }

	// 1) A deffered function runs after the surrounding function has finished,
	// right before the function actually returns.
	// Return of defer function is earlier than return of surrounding function

	// 2) A deferred function’s arguments are evaluated when the defer statement is evaluated.

	// IMPORTANT MENTION

	// func a() {
	// 	x := 1
	// 	defer fmt.Println(x) // `x` is evaluated immediately, so 1 is captured
	// 	x++
	// 	return
	// }

	// func main() {
	// 	a() // Prints 1
	// }

	// func a() {
	// 	x := 1
	// 	defer func() {
	// 		fmt.Println(x) // `x` is captured by reference, so the current value is used when this function executes
	// 	}()
	// 	x++
	// 	return
	// }

	// func main() {
	// 	a() // Prints 2
	// }

	// defer fmt.Println(x):
	// The value of x is captured immediately when the defer statement is executed, so whatever x is at that moment will be used when fmt.Println eventually runs.

	// defer func() { fmt.Println(x) }(): The value of x is not captured immediately.
	// Instead, the anonymous function holds a reference to x,
	// and the actual value used by fmt.Println is whatever x is at the time the deferred function is executed.

	// 2) If you defer multiple functions,
	// they are executed in a last-in-first-out (LIFO) order.

	// 3) 'defer' is commonly used to close recources like files or networl connections,
	// unlock mutexes, or recover from panic.
	// Deferred functions may read and assign to the returning function’s named return values.

	// Set new output and enable logging
	log.SetOutput(file)
	log.Enable()

	// Log to the file
	// The file will be rewritten
	log.Log("This will be logged to the file.")
	log.Begin()
	log.End()

	// Usage of Logf
	name := "Ramil"
	age := 22
	log.Logf("The name was %q, %d years old.", name, age)
}
