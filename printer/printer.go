package printer

import (
	"fmt"
	"github.com/keithnull/opg-analyzer/types"
)

func PrintOPTable(optable *types.OPTable) error {
	return fmt.Errorf("PrintOPTable unimplemented")
}

func PrintGrammar(grammar *types.Grammar) error {
	fmt.Println("The grammar is:")
	fmt.Println("Terminals:", grammar.Terminals)
	fmt.Println("Non-terminals:", grammar.NonTerminals)
	fmt.Println("Productions:")
	for left, productions := range grammar.Productions {
		fmt.Println(left, "->", productions)
	}
	return nil
}
