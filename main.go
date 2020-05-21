package main

import (
	"flag"
	"fmt"
	"github.com/keithnull/opg-analyzer/analyzer"
	"github.com/keithnull/opg-analyzer/printer"
	"github.com/keithnull/opg-analyzer/reader"
	"os"
)

var (
	grammarFile  string
	opTableFile  string
	sentenceFile string
)

func init() {
	flag.StringVar(&grammarFile, "grammar", "example_grammar.txt",
		" input: OPG file")
	flag.StringVar(&opTableFile, "table", "example_table.txt",
		"output: OP table file")
	flag.StringVar(&sentenceFile, "sentences", "",
		" input: sentences to parse in a file")
	flag.Parse()
}

func main() {
	// read grammar from file
	grammar, err := reader.ReadGrammarFromFile(grammarFile)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Failed to read and parse grammar form file:\n", err)
		os.Exit(1)
	}
	// print the grammar for sanity check
	_ = printer.PrintGrammar(grammar, os.Stdout)
	// analyze the grammar to generate OP (operator precedence) table
	opTable, err := analyzer.GenerateOPTable(grammar)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Failed to generate OP table:\n", err)
		os.Exit(1)
	}
	// output the OP table
	_ = printer.PrintOPTable(opTable, os.Stdout)
	err = printer.PrintOPTableToFile(opTable, opTableFile)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Failed to write OP table to file:\n", err)
		os.Exit(1)
	}
	if len(sentenceFile) != 0 { // there are sentences to parse
		sentences, err := reader.ReadSentencesFromFile(sentenceFile)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "Failed to load sentences from file:\n", err)
			os.Exit(1)
		}
		_ = analyzer.ParseSentences(opTable, sentences)
	}
}
