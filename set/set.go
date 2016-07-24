package set

type Set map[interface{}]bool

func New() Set {
	return make(Set)
}

func (s Set) Size() int {
	return len(s)
}

func (s Set) Add(e interface{}) {
	s[e] = true
}

func (s Set) Contains(e interface{}) bool {
	_, err := s[e]
	return err
}
