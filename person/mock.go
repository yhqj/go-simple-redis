package person

import "fmt"

func Mock() map[string]Person {
	ret := make(map[string]Person)
	for i := 100000000; i < 100000100; i++ {
		ret[fmt.Sprintf("person-%d", i)] = Person{
			Name:     fmt.Sprintf("person-%d", i),
			Age:      i,
			height:   i,
			weight:   i,
			location: fmt.Sprintf("location-%d", i),
			married:  false,
		}
	}
	return ret
}
