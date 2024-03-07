package sets

type StringSets map[string]struct{}

func (s StringSets) Add(v string) {
	s[v] = struct{}{}
}

func (s StringSets) Contains(v string) bool {
	_, ok := s[v]
	return ok
}

func (s StringSets) ToSlice() []string {
	slice := make([]string, 0, len(s))
	for key, _ := range s {
		slice = append(slice, key)
	}
	return slice
}
