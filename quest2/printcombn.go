package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n == 1 {
		for i := '0'; i <= '9'; i++ {
			z01.PrintRune(i)
			if i != '9' {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	if n == 2 {
		for i := '0'; i <= '9'; i++ {
			for j := i + 1; j <= '9'; j++ {
				z01.PrintRune(i)
				z01.PrintRune(j)
				if i != '8' || j != '9' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}

	if n == 3 {
		for i := '0'; i <= '9'; i++ {
			for j := i + 1; j <= '9'; j++ {
				for k := j + 1; k <= '9'; k++ {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
					if i != '7' || j != '8' || k != '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}

	if n == 4 {
		for i := '0'; i <= '9'; i++ {
			for j := i + 1; j <= '9'; j++ {
				for k := j + 1; k <= '9'; k++ {
					for l := k + 1; l <= '9'; l++ {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(k)
						z01.PrintRune(l)
						if i != '6' || j != '7' || k != '8' || l != '9' {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
	if n == 5 {
		for i := '0'; i <= '9'; i++ {
			for j := i + 1; j <= '9'; j++ {
				for k := j + 1; k <= '9'; k++ {
					for l := k + 1; l <= '9'; l++ {
						for m := l + 1; m <= '9'; m++ {
							z01.PrintRune(i)
							z01.PrintRune(j)
							z01.PrintRune(k)
							z01.PrintRune(l)
							z01.PrintRune(m)
							if i != '5' || j != '6' || k != '7' || l != '8' || m != '9' {
								z01.PrintRune(',')
								z01.PrintRune(' ')
							}
						}
					}
				}
			}
		}
	}
	if n == 6 {
		for i := '0'; i <= '9'; i++ {
			for j := i + 1; j <= '9'; j++ {
				for k := j + 1; k <= '9'; k++ {
					for l := k + 1; l <= '9'; l++ {
						for m := l + 1; m <= '9'; m++ {
							for n := m + 1; n <= '9'; n++ {
								z01.PrintRune(i)
								z01.PrintRune(j)
								z01.PrintRune(k)
								z01.PrintRune(l)
								z01.PrintRune(m)
								z01.PrintRune(n)
								if i != '4' || j != '5' || k != '6' || l != '7' || m != '8' || n != '9' {
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
	if n == 7 {
		for i := '0'; i <= '9'; i++ {
			for j := i + 1; j <= '9'; j++ {
				for k := j + 1; k <= '9'; k++ {
					for l := k + 1; l <= '9'; l++ {
						for m := l + 1; m <= '9'; m++ {
							for n := m + 1; n <= '9'; n++ {
								for o := n + 1; o <= '9'; o++ {
									z01.PrintRune(i)
									z01.PrintRune(j)
									z01.PrintRune(k)
									z01.PrintRune(l)
									z01.PrintRune(m)
									z01.PrintRune(n)
									z01.PrintRune(o)
									if i != '3' || j != '4' || k != '5' || l != '6' || m != '7' || n != '8' || o != '9' {
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
	if n == 8 {
		for i := '0'; i <= '9'; i++ {
			for j := i + 1; j <= '9'; j++ {
				for k := j + 1; k <= '9'; k++ {
					for l := k + 1; l <= '9'; l++ {
						for m := l + 1; m <= '9'; m++ {
							for n := m + 1; n <= '9'; n++ {
								for o := n + 1; o <= '9'; o++ {
									for p := o + 1; p <= '9'; p++ {
										z01.PrintRune(i)
										z01.PrintRune(j)
										z01.PrintRune(k)
										z01.PrintRune(l)
										z01.PrintRune(m)
										z01.PrintRune(n)
										z01.PrintRune(o)
										z01.PrintRune(p)
										if i != '2' || j != '3' || k != '4' || l != '5' || m != '6' || n != '7' || o != '8' || p != '9' {
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
	if n == 9 {
		for i := '0'; i <= '9'; i++ {
			for j := i + 1; j <= '9'; j++ {
				for k := j + 1; k <= '9'; k++ {
					for l := k + 1; l <= '9'; l++ {
						for m := l + 1; m <= '9'; m++ {
							for n := m + 1; n <= '9'; n++ {
								for o := n + 1; o <= '9'; o++ {
									for p := o + 1; p <= '9'; p++ {
										for q := p + 1; q <= '9'; q++ {
											z01.PrintRune(i)
											z01.PrintRune(j)
											z01.PrintRune(k)
											z01.PrintRune(l)
											z01.PrintRune(m)
											z01.PrintRune(n)
											z01.PrintRune(o)
											z01.PrintRune(p)
											z01.PrintRune(q)
											if i != '1' || j != '2' || k != '3' || l != '4' || m != '5' || n != '6' || o != '7' || p != '8' || q != '9' {
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
	}
	z01.PrintRune('\n')
}
