package set

type StringSet map[string]int

func NewStringSet(items ...string) StringSet {
	s := StringSet{}
	if len(items) > 0 {
		for _, item := range items {
			s.Add(item)
		}
	}
	return s
}

func (s StringSet) Has(item string) (exist bool) {
	if len(item) == 0 || s == nil {
		return false
	}

	_, exist = s[item]
	return
}

func (s StringSet) Add(item string) StringSet {
	if len(item) == 0 || s == nil {
		return s
	}

	s[item] = 1
	return s
}

func (s StringSet) Remove(item string) StringSet {
	if _, exist := s[item]; exist {
		delete(s, item)
	}
	return s
}

func (s StringSet) Len() int {
	if s == nil {
		return 0
	}
	return len(s)
}
