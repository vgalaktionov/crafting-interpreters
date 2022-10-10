package scanner

type Scanner struct {
	source string
}

func NewScanner(source string) *Scanner {
	return &Scanner{source}
}

func (s *Scanner) ScanTokens() []Token {
	return []Token{}
}
