package strings

type Interner struct {
	m map[string]string
}

func Intern(ss ...string) Interner {
	m := make(map[string]string, len(ss))
	for _, s := range ss {
		m[s] = s
	}
	return Interner{m}
}

func (i Interner) Add(s string) string {
	if s, ok := i.m[s]; ok {
		return s
	}
	i.m[s] = s
	return s
}

func (i Interner) Get(s string) string {
	if s, ok := i.m[s]; ok {
		return s
	}
	return s
}
