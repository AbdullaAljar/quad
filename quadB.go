package piscine

import (
	"github.com/01-edu/z01"
)

func QuadB(x, y int) {
	if x >= 1 && y >= 1 {
		z01.PrintRune('/')
		for i := 1; i <= x-2; i++ {
			z01.PrintRune('*')
		}
		if x > 1 {
			z01.PrintRune('\\')
		}

		if y > 1 {
			z01.PrintRune('\n')
			for i := 1; i <= y-2; i++ {
				z01.PrintRune('*')

				for j := 1; j <= x-2; j++ {
					z01.PrintRune(' ')
				}
				if x > 1 {
					z01.PrintRune('*')
				}
				z01.PrintRune('\n')
			}
		}

		if x >= 1 && y > 1 {
			z01.PrintRune('\\')
		}

		if y > 1 {
			for i := 1; i <= x-2; i++ {
				z01.PrintRune('*')
			}
		}

		if x > 1 && y > 1 {
			z01.PrintRune('/')
		}
		z01.PrintRune('\n')

	}
}
