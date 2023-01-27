package birdle

//

type Mark int8

const (
	Normal Mark = iota
	Correct
	Misplaced
)

//

type Guess struct {
	word  string
	marks []Mark
}

func (g Guess) Word() string    { return g.word }
func (g Guess) Mark(i int) Mark { return g.marks[i] }
