package printer

import (
	"fmt"
	"github.com/keithnull/opg-analyzer/types"
	"io"
	"os"
	"strings"
)

func PrintOPTable(opTable *types.OPTable, writer io.Writer) (err error) {
	_, err = fmt.Fprintln(writer, "The OP table is:")
	_, err = fmt.Fprintln(writer, "Terminals:", opTable.Terminals)
	_, err = fmt.Fprintln(writer, "Relations:")
	maxLength, extraSpace := 0, 3
	for _, t := range opTable.Terminals {
		if maxLength < len(t.Name) {
			maxLength = len(t.Name)
		}
	}
	// print table header
	_, err = fmt.Fprint(writer, strings.Repeat(" ", maxLength+extraSpace))
	for _, t := range opTable.Terminals {
		_, err = fmt.Fprint(writer, t.Name, strings.Repeat(" ", maxLength+extraSpace-len(t.Name)))
	}
	_, err = fmt.Fprint(writer, "\n")
	// print each row
	for _, t := range opTable.Terminals {
		_, err = fmt.Fprint(writer, t.Name, strings.Repeat(" ", maxLength+extraSpace-len(t.Name)))
		for _, t2 := range opTable.Terminals {
			var s string
			if p, ok := opTable.Relations[types.TokenPair{t, t2}]; !ok {
				s = ""
			} else {
				s = p.String()
			}
			_, err = fmt.Fprint(writer, s, strings.Repeat(" ", maxLength+extraSpace-len(s)))
		}
		_, err = fmt.Fprint(writer, "\n")
	}
	return
}

func PrintGrammar(grammar *types.Grammar, writer io.Writer) (err error) {
	_, err = fmt.Fprintln(writer, "The grammar is:")
	_, err = fmt.Fprintln(writer, "Terminals:", grammar.Terminals)
	_, err = fmt.Fprintln(writer, "Non-terminals:", grammar.NonTerminals)
	_, err = fmt.Fprintln(writer, "Productions:")
	for left, productions := range grammar.Productions {
		_, err = fmt.Fprintln(writer, left, "->", productions)
	}
	return
}

// create a file and then write OP table to it
func PrintOPTableToFile(opTable *types.OPTable, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	return PrintOPTable(opTable, file)
}
