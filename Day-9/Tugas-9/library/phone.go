package library

import "strconv"

type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

type CetakData interface {
	Print() string
}

func (p Phone) Print() string {
	var colors string
	for _, color := range p.Colors {
		colors += color + ", "
	}
	return "Name : " + p.Name + "\nBrand : " + p.Brand + "\nyear : " + strconv.Itoa(p.Year) + "\ncolors : " + colors
}
