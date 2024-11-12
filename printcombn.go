package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	switch {
	case n == 1:
		for i := '0'; i <= '9'; i++ {
			z01.PrintRune(i)
			if i != '9' {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	case n == 2:
		for i := '0'; i <= '8'; i++ {
			for j := i + 1; j <= '9'; j++ {
				z01.PrintRune(i)
				z01.PrintRune(j)
				if !(i == '8' && j == '9') {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	case n == 3:
		for i := '0'; i <= '7'; i++ {
			for j := i + 1; j <= '8'; j++ {
				for k := j + 1; k <= '9'; k++ {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)

					if !(i == '7' && j == '8' && k == '9') {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
		z01.PrintRune('\n')
	case n == 4:
		for i := '0'; i <= '6'; i++ {
			for j := i + 1; j <= '7'; j++ {
				for k := j + 1; k <= '8'; k++ {
					for l := k + 1; l <= '9'; l++ {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(k)
						z01.PrintRune(l)

						if !(i == '6' && j == '7' && k == '8' && l == '9') {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
		z01.PrintRune('\n')
	case n == 5:
		for i := '0'; i <= '5'; i++ {
			for j := i + 1; j <= '6'; j++ {
				for k := j + 1; k <= '7'; k++ {
					for l := k + 1; l <= '8'; l++ {
						for m := l + 1; m <= '9'; m++ {
							z01.PrintRune(i)
							z01.PrintRune(j)
							z01.PrintRune(k)
							z01.PrintRune(l)
							z01.PrintRune(m)

							if !(i == '5' && j == '6' && k == '7' && l == '8' && m == '9') {
								z01.PrintRune(',')
								z01.PrintRune(' ')
							}
						}
					}
				}
			}
		}
		z01.PrintRune('\n')
	case n == 6:
		for i := '0'; i <= '4'; i++ {
			for j := i + 1; j <= '5'; j++ {
				for k := j + 1; k <= '6'; k++ {
					for l := k + 1; l <= '7'; l++ {
						for m := l + 1; m <= '8'; m++ {
							for n := m + 1; n <= '9'; n++ {
								z01.PrintRune(i)
								z01.PrintRune(j)
								z01.PrintRune(k)
								z01.PrintRune(l)
								z01.PrintRune(m)
								z01.PrintRune(n)

								if !(i == '4' && j == '5' && k == '6' && l == '7' && m == '8' && n == '9') {
									z01.PrintRune(',')
									z01.PrintRune(' ')
								}
							}
						}
					}
				}
			}
		}
		z01.PrintRune('\n')
	case n == 7:
		for i := '0'; i <= '3'; i++ {
			for j := i + 1; j <= '4'; j++ {
				for k := j + 1; k <= '5'; k++ {
					for l := k + 1; l <= '6'; l++ {
						for m := l + 1; m <= '7'; m++ {
							for n := m + 1; n <= '8'; n++ {
								for p := n + 1; p <= '9'; p++ {
									z01.PrintRune(i)
									z01.PrintRune(j)
									z01.PrintRune(k)
									z01.PrintRune(l)
									z01.PrintRune(m)
									z01.PrintRune(n)
									z01.PrintRune(p)

									if !(i == '3' && j == '4' && k == '5' && l == '6' && m == '7' && n == '8' && p == '9') {
										z01.PrintRune(',')
										z01.PrintRune(' ')
									}
								}
							}
						}
					}
				}
			}
		}
		z01.PrintRune('\n')
	case n == 8:
		for i := '0'; i <= '2'; i++ {
			for j := i + 1; j <= '3'; j++ {
				for k := j + 1; k <= '4'; k++ {
					for l := k + 1; l <= '5'; l++ {
						for m := l + 1; m <= '6'; m++ {
							for n := m + 1; n <= '7'; n++ {
								for p := n + 1; p <= '8'; p++ {
									for q := p + 1; q <= '9'; q++ {
										z01.PrintRune(i)
										z01.PrintRune(j)
										z01.PrintRune(k)
										z01.PrintRune(l)
										z01.PrintRune(m)
										z01.PrintRune(n)
										z01.PrintRune(p)
										z01.PrintRune(q)

										if !(i == '2' && j == '3' && k == '4' && l == '5' && m == '6' && n == '7' && p == '8' && q == '9') {
											z01.PrintRune(',')
											z01.PrintRune(' ')
										}
									}
								}
							}
						}
					}
				}
			}
		}
		z01.PrintRune('\n')
	case n == 9:
		for i := '0'; i <= '2'; i++ {
			for j := i + 1; j <= '3'; j++ {
				for k := j + 1; k <= '4'; k++ {
					for l := k + 1; l <= '5'; l++ {
						for m := l + 1; m <= '6'; m++ {
							for n := m + 1; n <= '7'; n++ {
								for p := n + 1; p <= '8'; p++ {
									for q := p + 1; q <= '9'; q++ {
										for z := q + 1; z <= '9'; z++ {
											z01.PrintRune(i)
											z01.PrintRune(j)
											z01.PrintRune(k)
											z01.PrintRune(l)
											z01.PrintRune(m)
											z01.PrintRune(n)
											z01.PrintRune(p)
											z01.PrintRune(q)
											z01.PrintRune(z)
											if !(i == '1' && j == '2' && k == '3' && l == '4' && m == '5' && n == '6' && p == '7' && q == '8' && z == '9') {
												z01.PrintRune(',')
												z01.PrintRune(' ')
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
		z01.PrintRune('\n')
	default:
		return
	}
}
