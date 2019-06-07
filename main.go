package main

import (
	"flag"
	"log"
	"os"
	"time"
)

func main() {
	var lines = flag.Int("l", 60, "Number of lines per minute")
	var source = flag.String("t", "shakespeare", "The type of logs")
	var destination = flag.String("d", "program.log", "The file to log to")
	var remove = flag.Bool("r", false, "Remove previous file content")

	flag.Parse()

	var gen Generator
	var prefix string
	switch *source {
	case "shakespeare":
		gen = &Shakespeare{}
		prefix = ""
	default:
		log.Panicf("UNRECOGNIZED TYPE\n")
	}

	if *remove {
		os.Remove(*destination)
	}

	f, err := os.OpenFile(*destination, os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Panic(err)
	}

	defer f.Close()

	logger := log.New(f, prefix, 0644)

	var interval = time.Minute / time.Duration(*lines)

	run(gen, interval, logger)
}

func run(gen Generator, delay time.Duration, logger *log.Logger) {
	ticker := time.NewTicker(delay).C

	for {
		<-ticker
		logger.Println(gen.Next())
	}
}
