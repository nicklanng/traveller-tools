package characters

var (
	RaceHuman = &Race{
		Name: "Human",
	}
)

type Race struct {
	Name   string
	Traits []*RaceTrait
}

func (r *Race) String() string {
	return r.Name
}
