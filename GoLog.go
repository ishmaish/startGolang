package main

import "log"

/* The main purpose of logging is to get a trace of what's happening in the program,
where it's happening, and when it's happening.
Logs can be providing code tracing, profiling, and analytics*/

func main() {
	// Println writes to the standard logger.
	log.Println("main started")

	// Fatalln is Println() followed by a call to os.Exit(1)
	log.Fatalln("fatal message")

	// Panicln is Println() followed by a call to panic()
	log.Panicln("panic message")
}

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}
