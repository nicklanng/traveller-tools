package worlds

import (
	"fmt"
	"strings"
)

type World struct {
	Name       string
	HexX       int
	HexY       int
	Population int
	TravelZone TravelZone
	TradeCodes []TradeCode
}

func (w World) String() (s string) {
	s = w.Name + "\t"
	s += fmt.Sprintf("%02d", w.HexX) + fmt.Sprintf("%02d", w.HexY) + "\t"

	var codes []string
	for _, v := range w.TradeCodes {
		codes = append(codes, v.Code)
	}
	s += strings.Join(codes, " ") + "\t"

	return
}
