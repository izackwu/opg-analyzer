package analyzer

import (
	"fmt"
	"github.com/keithnull/opg-analyzer/types"
	"strings"
)

func ParseSentence(opTable *types.OPTable, sentence types.TokenList) error {
	analyzeStack := make(types.TokenList, 1)
	inputIdx, iteration := 0, 0
	endToken := types.Token{
		Name:       "$",
		IsTerminal: true,
	}
	analyzeStack[0] = endToken
	sentence = append(sentence, endToken)
	formatWidth := len(sentence.String())
	if formatWidth < 5 {
		formatWidth = 5
	}
	fmt.Printf("Start to parse %v\n", sentence)
	fmt.Println(strings.Repeat("-", 20+formatWidth*2))
	defer fmt.Println(strings.Repeat("-", 20+formatWidth*2))
	fmt.Printf("%v %-*v %*v %v\n", "Iteration", formatWidth, "Stack", formatWidth, "Input",
		"Action")
	for inputIdx < len(sentence) {
		iteration += 1
		fmt.Printf("%-9d %-*v %*v ", iteration, formatWidth,
			analyzeStack, formatWidth, sentence[inputIdx:])
		stackTopIdx := len(analyzeStack) - 1
		for ; !analyzeStack[stackTopIdx].IsTerminal; stackTopIdx-- {
		}
		inputToken, stackTopToken := sentence[inputIdx], analyzeStack[stackTopIdx]
		relation, ok := opTable.Relations[types.TokenPair{
			Left:  stackTopToken,
			Right: inputToken,
		}]
		if !ok {
			fmt.Println("Error")
			return fmt.Errorf("invalid sentence: %v", sentence)
		}
		switch relation {
		case types.Lower, types.Equal:
			if len(analyzeStack) == 2 && inputIdx == len(sentence)-1 && !analyzeStack[1].
				IsTerminal {
				fmt.Println("Accept")
			} else {
				fmt.Println("Shift")
			}
			analyzeStack = append(analyzeStack, inputToken)
			inputIdx += 1
		case types.Higher:
			fmt.Println("Reduce")
			lastPop := inputToken
			leftIdx := stackTopIdx
			for ; leftIdx >= 0; leftIdx-- {
				if !analyzeStack[leftIdx].IsTerminal {
					continue
				}
				if opTable.Relations[types.TokenPair{
					Left:  analyzeStack[leftIdx],
					Right: lastPop,
				}] == types.Lower {
					break
				}
				lastPop = analyzeStack[leftIdx]
			}
			analyzeStack = append(analyzeStack[:leftIdx+1], types.Token{
				Name:       "X",
				IsTerminal: false,
			})
		default:
			fmt.Println("Error")
			return fmt.Errorf("invalid terminal relation between %v and %v", stackTopToken, inputToken)
		}
	}
	return nil
}

func ParseSentences(opTable *types.OPTable, sentences []types.TokenList) error {
	for _, sentence := range sentences {
		err := ParseSentence(opTable, sentence)
		if err != nil {
			return err
		}
	}
	return nil
}
