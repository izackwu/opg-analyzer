package main

import (
	"fmt"
	"github.com/keithnull/opg-analyzer/analyzer"
	"github.com/keithnull/opg-analyzer/printer"
	"github.com/keithnull/opg-analyzer/reader"
	"os"
)

func main() {
	// read grammar from file
	grammar, err := reader.ReadFromFile("example_grammar.txt")
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// print the grammar for sanity check
	_ = printer.PrintGrammar(grammar)
	// analyze the grammar to generate OP (operator precedence) table
	optable, err := analyzer.GenerateOPTable(grammar)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// output the OP table
	_ = printer.PrintOPTable(optable)
}
