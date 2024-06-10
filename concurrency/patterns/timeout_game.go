package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

type problem struct {
	q, a string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readCSVFile(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

func parseProblems(records [][]string) []problem {
	problems := make([]problem, len(records))
	for i := 0; i < len(records); i++ {
		problems[i] = problem{q: records[i][0], a: records[i][1]}
	}
	return problems
}

func main() {
	records, err := readCSVFile("qa.csv")
	if err != nil {
		log.Fatalln(err)
	}
	problems := parseProblems((records))

	correct, incorrect := 0, 0
	answerCh := make(chan string)
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s ", i+1, p.q)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-time.NewTimer(time.Duration(5) * time.Second).C:
			fmt.Printf("\nCorrect %d, Incorrect %d, Total %d\n", correct, incorrect, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			} else {
				incorrect++
			}
		}
	}
	fmt.Printf("\nCorrect %d, Incorrect %d, Total %d\n", correct, incorrect, len(problems))
}
