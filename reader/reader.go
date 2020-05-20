package reader

import (
	"bufio"
	"fmt"
	"github.com/keithnull/opg-analyzer/types"
	"os"
	"strings"
)

func processLine(line string, grammar *types.Grammar, lineno int) error {
	fields := strings.Fields(line)
	if len(fields) == 0 { // skip empty lines
		return nil
	}
	if len(fields) <= 2 || fields[1] != "->" {
		return fmt.Errorf("incorrect grammar format in Line %d: %s", lineno, line)
	}
	leftToken := types.Token{
		Name:       fields[0],
		IsTerminal: false,
	}
	// add it to non-terminals
	if !grammar.NonTerminals.Contains(leftToken) {
		grammar.NonTerminals = append(grammar.NonTerminals, leftToken)
	}
	if _, ok := grammar.Productions[leftToken]; !ok {
		grammar.Productions[leftToken] = make([]types.Production, 0)
	}
	pos := 2
	for pos < len(fields) {
		production := make(types.Production, 0)
		for ; pos < len(fields) && fields[pos] != "|"; pos++ {
			token := types.Token{
				Name:       fields[pos],
				IsTerminal: true, // by default, treat it as a terminal and correct this later
			}
			production = append(production, token)
		}
		if len(production) == 0 {
			return fmt.Errorf("invalid production in Line %d: %s", lineno, line)
		}
		grammar.Productions[leftToken] = append(grammar.Productions[leftToken], production)
		pos += 1 // skip "|" or go outer of range
	}
	return nil
}

func correctTokenType(grammar *types.Grammar) error {
	// oh my god, these lines are so ugly!
	// but actually, I need to iterate through the grammar and have no other graceful ways
	for left, productions := range grammar.Productions {
		for i, production := range productions {
			for j, token := range production {
				tokenPointer := &grammar.Productions[left][i][j]
				if grammar.NonTerminals.Contains(token) {
					tokenPointer.IsTerminal = false
					// check whether it's a valid OPG
					if j != 0 && production[j-1].IsTerminal == false {
						return fmt.Errorf("invalid grammar: consecutive nonterminals in one"+
							" production: %v -> %v ... ", left, production[:j+1])
					}
				} else {
					grammar.Terminals = append(grammar.Terminals, *tokenPointer)
				}
			}
		}
	}
	return nil
}

func ReadFromFile(filepath string) (*types.Grammar, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lineno := 1
	grammar := &types.Grammar{
		Terminals:    make([]types.Token, 0),
		NonTerminals: make([]types.Token, 0),
		Productions:  make(map[types.Token][]types.Production, 0),
	}
	for scanner.Scan() {
		line := scanner.Text()
		if err := processLine(line, grammar, lineno); err != nil {
			return nil, err
		}
		lineno += 1
	}
	if err := correctTokenType(grammar); err != nil {
		return nil, err
	}
	return grammar, nil
}

func ReadFromText(text string) (*types.Grammar, error) {
	return nil, fmt.Errorf("ReadFromText unimplemented")
}
