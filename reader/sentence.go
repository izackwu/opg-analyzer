package reader

import (
	"bufio"
	"github.com/keithnull/opg-analyzer/types"
	"io"
	"os"
	"strings"
)

func ReadSentences(reader io.Reader) ([]types.TokenList, error) {
	scanner := bufio.NewScanner(reader)
	sentences := make([]types.TokenList, 0)
	for scanner.Scan() {
		tokenStrings := strings.Fields(scanner.Text())
		if len(tokenStrings) == 0 { // skip empty lines
			continue
		}
		tl := make(types.TokenList, len(tokenStrings))
		for i, s := range tokenStrings {
			tl[i] = types.Token{
				Name:       s,
				IsTerminal: true,
			}
		}
		sentences = append(sentences, tl)
	}
	return sentences, nil
}

func ReadSentencesFromFile(filepath string) ([]types.TokenList, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ReadSentences(file)
}
