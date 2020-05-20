package printer

import (
	"fmt"
	"github.com/keithnull/opg-analyzer/types"
	"strings"
)

func PrintOPTable(opTable *types.OPTable) error {
	fmt.Println("The OP table is:")
	fmt.Println("Terminals:", opTable.Terminals)
	fmt.Println("Relations:")
	maxLength, extraSpace := 0, 3
	for _, t := range opTable.Terminals {
		if maxLength < len(t.Name) {
			maxLength = len(t.Name)
		}
	}
	// print table header
	fmt.Print(strings.Repeat(" ", maxLength+extraSpace))
	for _, t := range opTable.Terminals {
		fmt.Print(t.Name, strings.Repeat(" ", maxLength+extraSpace-len(t.Name)))
	}
	fmt.Print("\n")
	// print each row
	for _, t := range opTable.Terminals {
		fmt.Print(t.Name, strings.Repeat(" ", maxLength+extraSpace-len(t.Name)))
		for _, t2 := range opTable.Terminals {
			var s string
			if p, ok := opTable.Relations[types.TokenPair{t, t2}]; !ok {
				s = ""
			} else {
				s = p.String()
			}
			fmt.Print(s, strings.Repeat(" ", maxLength+extraSpace-len(s)))
		}
		fmt.Print("\n")
	}
	return nil
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
