package errors

type ValidationErrors map[string][]string

func (s ValidationErrors) Add(field string, message string) {
	if _, ok := s[field]; !ok {
		s[field] = make([]string, 0)
	}
	s[field] = append(s[field], message)
}
