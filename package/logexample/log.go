package logexample

import (
	"log"
	"os"
)

// LogPrint example
func LogPrint() {
	log.Print("Print")
	log.Printf("%s", "Printf log")
	log.Println("Println")
	log.Println("Println")
}

// LogFatal example
func LogFatal() {
	log.Fatal("Log Fatal")
	log.Fatalln("Log Fatalln")
	log.Fatalf("%s", "Log Fatalf")
}

// LogPanic example
func LogPanic() {
	log.Panic("Log Panic")
	log.Panicln("Log Panicln")
	log.Panicf("%s", "Log Panicln")
}

// LoggerTest example
func LoggerTest() {
	logger := log.New(os.Stdout, "LoggerTest", log.LstdFlags|log.Lshortfile)
	logger.Printf("%s", "Log New Test")
	logger.Fatalf("%s", "New Logger Fatalf")
	logger.Panicf("%s", "New Logger Panicf")
}
