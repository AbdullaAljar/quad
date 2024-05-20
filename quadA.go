package piscine

import (
	"github.com/01-edu/z01"
)

func QuadA(x, y int) {
	if x >= 1 && y >= 1 {
		z01.PrintRune('o')

		for i := 1; i <= x-2; i++ {
			z01.PrintRune('-')
		}

		if x > 1 {
			z01.PrintRune('o')
		}
		z01.PrintRune('\n')

		for j := 1; j <= y-2; j++ {
			z01.PrintRune('|')

			for i := 1; i <= x-2; i++ {
				z01.PrintRune(' ')
			}

			if x > 1 {
				z01.PrintRune('|')
			}
			z01.PrintRune('\n')
		}

		if y >= 2 {
			z01.PrintRune('o')

			for i := 1; i <= x-2; i++ {
				z01.PrintRune('-')
			}
			if x > 1 {
				z01.PrintRune('o')
			}
			z01.PrintRune('\n')
		}
	}
}

