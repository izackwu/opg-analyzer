package types

func (tl TokenList) Contains(t Token) bool {
	for _, x := range tl {
		if x.Name == t.Name {
			return true
		}
	}
	return false
}
