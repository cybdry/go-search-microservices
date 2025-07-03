package data

var data = []Kitten{
	{
		ID:     "1",
		Name:   "Felix",
		Weight: 12.3,
	},
	{
		ID:     "2",
		Name:   "Fat Freddy's Cat",
		Weight: 20.0,
	},
	{
		ID:     "3",
		Name:   "Gartfield",
		Weight: 35.0,
	},
}

type MemoryStore struct{}

func (m *MemoryStore) Search(name string) []Kitten {
	var kitten []Kitten

	for _, k := range data {
		if k.Name == name {
			kitten = append(kitten, k)
		}
	}

	return kitten
}
