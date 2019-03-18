package vocabulary

type Set map[string]bool

func NewSet() Set {
return make(Set)
}

func (s *Set) Add(v string) {
(*s)[v] = true
}

func (s Set) Contains(v string) bool {
_, ok := s[v]
return ok
}
