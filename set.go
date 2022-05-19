package main

type void struct{}

type set map[string]void

func (s set) add(item string) set {
	s[item] = void{}
	return s
}

func (s set) has(item string) bool {
	_, ok := s[item]
	return ok
}

func head(d set) string {
	for item := range d {
		return item
	}
	return ""
}
