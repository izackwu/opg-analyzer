package analyzer

import (
	"fmt"
	"github.com/keithnull/opg-analyzer/types"
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
	for inputIdx < len(sentence) {
		iteration += 1
		fmt.Printf("Iteration %2d: %v %v ", iteration, analyzeStack, sentence[inputIdx:])
		stackTopIdx := len(analyzeStack) - 1
		for ; !analyzeStack[stackTopIdx].IsTerminal; stackTopIdx-- {
		}
		inputToken, stackTopToken := sentence[inputIdx], analyzeStack[stackTopIdx]
		relation, ok := opTable.Relations[types.TokenPair{
			Left:  stackTopToken,
			Right: inputToken,
		}]
		if !ok {
			return fmt.Errorf("invalid sentence: %v", sentence)
		}
		switch relation {
		case types.Lower, types.Equal:
			fmt.Println("Shift")
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
			return fmt.Errorf("invalid terminal relation between %v and %v", stackTopToken, inputToken)
		}
	}
	if len(analyzeStack) != 3 || !(analyzeStack[0] == endToken &&
		!analyzeStack[1].IsTerminal && analyzeStack[2] == endToken) {
		return fmt.Errorf("invalid sentence")
	}
	fmt.Println("Accept")
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
