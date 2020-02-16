package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// AcceptedState is the int const value of the final state
const AcceptedState = -1

// InitialState is the int const value of the initial state
const InitialState = 1

// NullByte is the rune const value of null (C-style end of string)
const NullByte = '\x00'

// InitalStateSymbolGrammar is the rune const value of the inital state when we process a .rg file
const InitalStateSymbolGrammar = 'S'

//Lambda is the rune const value of the lambda symbol
const Lambda = '\\'

// Key type of the transition matrix map
type Key struct {
	state  int
	symbol rune
}

func main() {
	file, str := os.Args[1], strings.Replace(os.Args[2]+string(NullByte), "\"", "", -1) //Append null byte and remove quotes
	//transitionMatrix, err := readTransitionMatrixFromCsvFile(file)
	transitionMatrix, err := readTransitionMatrixFromGrammarFile(file)
	if err != nil {
		fmt.Println("Error reading file:: \n", err)
		os.Exit(-1)
	}

	currentState := InitialState
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

func readTransitionMatrixFromGrammarFile(file string) (transitionMatrix map[Key]int, err error) {
	grammarFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer grammarFile.Close()

	transitionMatrix = make(map[Key]int)

	scanner := bufio.NewScanner(grammarFile)
	for scanner.Scan() {
		line := scanner.Text()
		var currentState, nextState int
		var symbol rune
		//Format: M->tN or M->\
		for index, runeValue := range line {
			switch index {
			case 0: // Processing leftmost terminal symbol
				if runeValue == InitalStateSymbolGrammar {
					currentState = InitialState
				} else {
					currentState = int(runeValue)
				}
			case 3: //Processing leftmost non-terminal symbol
				if runeValue == Lambda { // Accept state
					symbol = NullByte
					nextState = AcceptedState
				} else {
					symbol = runeValue
				}
			case 4: //Process rightmost terminal symbol
				if runeValue == InitalStateSymbolGrammar {
					nextState = InitialState
				} else {
					nextState = int(runeValue)
				}
			default:
			}
		}
		transitionMatrix[Key{currentState, symbol}] = nextState
	}

	return transitionMatrix, nil
}

func readTransitionMatrixFromCsvFile(file string) (transitionMatrix map[Key]int, err error) {
	recordFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer recordFile.Close()
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
			token = strings.Replace(token, "EOS", string(NullByte), -1) //Replace EOS with "/00"
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
	return transitionMatrix, nil
}
