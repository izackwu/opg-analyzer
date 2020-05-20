package types

import "fmt"

func (tl TokenList) Contains(t Token) bool {
	for _, x := range tl {
		if x.Name == t.Name {
			return true
		}
	}
	return false
}

// Adds one token into the token list if it's not already in it
// return true if appending is performed
func AppendUniqueTokenList(tl TokenList, t Token) (TokenList, bool) {
	if tl.Contains(t) {
		return tl, false
	}
	return append(tl, t), true
}

func (tpl TokenPairList) Contains(tp TokenPair) bool {
	for _, x := range tpl {
		if x.Left.Name == tp.Left.Name && x.Right.Name == tp.Right.Name {
			return true
		}
	}
	return false
}

// Adds one token pair into the token pair list if it's not already in it
// return true if appending is performed
func AppendUniqueTokenPairList(tpl TokenPairList, tp TokenPair) (TokenPairList, bool) {
	if tpl.Contains(tp) {
		return tpl, false
	}
	return append(tpl, tp), true
}

func (opt *OPTable) InsertRelation(leftToken, rightToken Token, precedence Precedence) error {
	pair := TokenPair{
		Left:  leftToken,
		Right: rightToken,
	}
	if _, ok := opt.Relations[pair]; ok {
		return fmt.Errorf("precedence conflicts for %v and %v", leftToken, rightToken)
	}
	opt.Relations[pair] = precedence
	return nil
}

func (e Precedence) String() string {
	switch e {
	case Equal:
		return "="
	case Higher:
		return ">"
	case Lower:
		return "<"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}
