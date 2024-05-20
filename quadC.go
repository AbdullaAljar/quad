package piscine

import (
	"github.com/01-edu/z01"
)

func QuadC(x, y int) {
	if x >= 1 && y >= 1 {
		z01.PrintRune('A')
		for i := 1; i <= x-2; i++ {
			z01.PrintRune('B')
		}
		if x > 1 {
			z01.PrintRune('A')
		}

		if y > 1 {
			z01.PrintRune('\n')
			for i := 1; i <= y-2; i++ {
				z01.PrintRune('B')

				for j := 1; j <= x-2; j++ {
					z01.PrintRune(' ')
				}
				if x > 1 {
					z01.PrintRune('B')
				}
				z01.PrintRune('\n')
			}
		}

		if x >= 1 && y > 1 {
			z01.PrintRune('C')
		}

		if y > 1 {
			for i := 1; i <= x-2; i++ {
				z01.PrintRune('B')
			}
		}

		if x > 1 && y > 1 {
			z01.PrintRune('C')
		}
		z01.PrintRune('\n')

	}
}

