package strings

type Secret struct {
	s string
}

func SecretOf(s string) Secret {
	return Secret{s}
}

func (s Secret) String() string {
	return "***"
}

func (s Secret) Value() string {
	return s.s
}
