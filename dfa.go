package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// AcceptedState is the numerical const value of the final state
const AcceptedState = -1

// Key type of the transition matrix map
type Key struct {
	state  int
	symbol rune
}

func main() {
	file, str := os.Args[1], strings.Replace(os.Args[2]+"\x00", "\"", "", -1) //Append null byte and remove quotes
	transitionMatrix, err := readTransitionMatrixFromFile(file)
	if err != nil {
		fmt.Println("Error reading file:: \n", err)
		os.Exit(-1)
	}

	currentState := 1
	for _, token := range str {
		nextState, exists := transitionMatrix[Key{currentState, token}]
		if !exists {
			break
		}
		currentState = nextState
	}
	var accepted string
	if currentState == AcceptedState {
		accepted = "yes"
	} else {
		accepted = "no"
	}
	fmt.Println("Accepted: ", accepted)
}

func readTransitionMatrixFromFile(file string) (transitionMatrix map[Key]int, err error) {
	recordFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(recordFile)
	header, err := reader.Read()
	if err != nil {
		return nil, err
	}
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	transitionMatrix = make(map[Key]int, len(records))
	for _, row := range records {
		currentState, err := strconv.Atoi(row[0])
		if err != nil {
			return nil, err
		}
		tokens := header[1:]
		for j, token := range tokens {
			token = strings.Replace(token, "EOS", "\x00", -1) //Replace EOS with "/00"
			nextStateStr := row[j+1]
			if nextStateStr != "" { //There's a state associated with this symbol
				nextState, err := strconv.Atoi(strings.Replace(nextStateStr, "accept", strconv.Itoa(AcceptedState), -1)) //Replace accept state with const
				if err != nil {
					return nil, err
				}
				transitionMatrix[Key{currentState, []rune(token)[0]}] = nextState
			}
		}
	}
	recordFile.Close()
	return transitionMatrix, nil
}
