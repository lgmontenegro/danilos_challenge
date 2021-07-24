package main

import (
	"fmt"
)

type coringas []string

func main() {
	names := coringas{"R>T", "A>L", "P>O", "O>R", "G>A", "T>U", "U>G"}
	cleanNames := clear(names)

	fmt.Println(cleanNames)

	name := ""

	for _, first := range cleanNames {
		f := first[0]
		found := false
		for _, last := range cleanNames {
			l := last[1]
			if f == l {
				found = true
				name = name + string(l)
				break
			}
		}
		if !found {
			name = name + string(f)
		}
	}

	fmt.Println(name)

}

func clean(a []byte) (b []byte) {
	return []byte{a[0], a[2]}
}

func clear(a []string) (b []string) {
	for _, c := range a {
		b = append(b, string(clean([]byte(c))))
	}
	return b
}
