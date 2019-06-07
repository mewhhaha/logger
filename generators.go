package main

type Generator interface {
	Next() string
}

type Shakespeare struct {
	i int
}

func (s *Shakespeare) Next() string {
	line := ShakespeareWiki[s.i]
	s.i = (s.i + 1) % len(ShakespeareWiki)
	return line
}
