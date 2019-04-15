package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type TestEvent struct {
	Time    time.Time // encodes as an RFC3339-format string
	Action  string
	Package string
	Test    string
	Elapsed float64 // seconds
	Output  string
}

func main() {
	resultsFileName := "smoke.json"
	f, err := os.Open(resultsFileName)
	if err != nil {
		fmt.Println("error opening file ", err)
		os.Exit(1)
	}
	defer f.Close()

	var run, skip, pass, fail int
	reader := bufio.NewReader(f)
	line, _, err := reader.ReadLine()
	for  err == nil {
		var event TestEvent
		err = json.Unmarshal(line, &event)
		if err != nil {
			fmt.Printf("Error %s on line [%s]\n", err, line)
			os.Exit(2)
		}

		switch event.Action {
		case "run": run++
		case "pass": pass++
		case "fail": fail++
		case "skip": skip++
		default:
		}
		if event.Action != "output" {
			fmt.Printf("Action: %s Test %s\n", event.Action, event.Test)
		}

		line, _, err = reader.ReadLine();
	}

	fmt.Printf("Run %d Pass %d Fail %d Skip %d\n", run, pass, fail, skip)
}
