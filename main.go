package main

import (
	"fmt"
)

type coringas []string

func main() {
	names := coringas{"R>T", "A>L", "P>O", "O>R", "G>A", "T>U", "U>G"}
	cleanNames := clear(names)

	fmt.Println(cleanNames)

	fmt.Println(whatever(cleanNames))

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

func whatever(cards []string) (cardsOrdered []string){
	total := len(cards)
	tempCard := ""

	for n := 0; n < total; n++ {
		fmt.Printf("n = %d\n", n)
		for x:=0; x < total;{
			fmt.Printf("X1 = %d\n", x)
			for x = n+1 ; x < total; x++ {
				fmt.Printf("X2 = %d\n", x)
				if cards[x][1] == cards[n][0] {
					tempCard = cards[x]
					cards[x] = cards[n]
					cards[n] = tempCard
					break
				}
			} 
		}
	}

	return cards
}
