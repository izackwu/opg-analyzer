package analyzer

import (
	"fmt"
	"github.com/keithnull/opg-analyzer/types"
)

func generateVT(grammar *types.Grammar, reverse bool) (map[types.Token]types.TokenList, error) {
	VT := make(map[types.Token]types.TokenList)
	// initialize first sets for all non-terminals
	for _, nt := range grammar.NonTerminals {
		VT[nt] = make(types.TokenList, 0)
	}
	// if (P, Q) in containing relationships, then VT(P) contains VT(Q)
	containing := make(types.TokenPairList, 0)
	// roughly compute VT and containing relationships
	for left, productions := range grammar.Productions {
		for _, production := range productions {
			firstIdx, secondIdx := 0, 1 // for FirstVT
			if reverse {                // for LastVT
				firstIdx, secondIdx = len(production)-1, len(production)-2
			}
			if production[firstIdx].IsTerminal {
				//    P -> x ..., then x is in FirstVT(P)
				// or P -> ... x, then x is in LastVT(P)
				VT[left], _ = types.AppendUniqueTokenList(VT[left], production[firstIdx])
			} else {
				//    P -> Q ..., then FirstVT(P) contains FirstVT(Q)
				// or P -> ... Q, then LastVT(P) contains LastVT(Q)
				tokenPair := types.TokenPair{
					Left:  left,
					Right: production[0],
				}
				containing, _ = types.AppendUniqueTokenPairList(containing, tokenPair)
				//    P -> Q x ..., then x is in FirstVT(P)
				// or P -> ... x Q, then x is in LastVT(P)
				if len(production) >= 2 && production[secondIdx].IsTerminal {
					VT[left], _ = types.AppendUniqueTokenList(VT[left], production[secondIdx])
				}
			}
		}
	}
	// iteratively update VT until it converges
	return updateIteratively(VT, containing, 1000)
}

func updateIteratively(VT map[types.Token]types.TokenList, containing types.TokenPairList,
	maxIteration int) (map[types.Token]types.TokenList, error) {
	converge := false
	iteration := 0
	for !converge {
		iteration += 1
		if maxIteration > 0 && iteration > maxIteration {
			return nil, fmt.Errorf("maximum number of iterations exceeded")
		}
		converge = true
		for _, pair := range containing {
			for _, nt := range VT[pair.Right] {
				var changeMade bool
				VT[pair.Left], changeMade = types.AppendUniqueTokenList(VT[pair.Left],
					nt)
				converge = converge && !changeMade
			}
		}
	}
	fmt.Println("Converge after", iteration, "iteration(s)")
	return VT, nil
}

func generateFirstVT(grammar *types.Grammar) (map[types.Token]types.TokenList, error) {
	return generateVT(grammar, false)
}

func generateLastVT(grammar *types.Grammar) (map[types.Token]types.TokenList, error) {
	return generateVT(grammar, true)
}

func generateOPTable(grammar *types.Grammar,
	firstVT, lastVT map[types.Token]types.TokenList) (*types.OPTable, error) {
	return nil, fmt.Errorf("generateOPTable unimplemented")
}

func GenerateOPTable(grammar *types.Grammar) (*types.OPTable, error) {
	firstVT, err := generateFirstVT(grammar)
	if err != nil {
		return nil, err
	}
	fmt.Println("FirstVT:", firstVT)
	lastVT, err := generateLastVT(grammar)
	if err != nil {
		return nil, err
	}
	fmt.Println("LastVT:", lastVT)
	return generateOPTable(grammar, firstVT, lastVT)
}
