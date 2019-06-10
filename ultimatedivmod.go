package piscine



func UltimateDivMod(a *int, b *int) {
	c := *b
	*b = *a % *b
	*a =  *a / c
}
