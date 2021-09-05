package main

import (
	"fmt"

	"github.com/dylandreimerink/go-rijksdriehoek"
)

func main() {
	rd := [][]float64{
		{121687, 487484}, // Amsterdam
		{92565, 437428},  // Rotterdam
		{176331, 317462}, // Maastricht
	}

	wgs := [][]float64{
		{52.37422, 4.89801}, // Amsterdam
		{51.92183, 4.47959}, // Rotterdam
		{50.84660, 5.69006}, // Maastricht
	}

	for _, xy := range rd {
		x, y := rijksdriehoek.RDtoWGS84(xy[0], xy[1])
		fmt.Printf("WGS - x: %f, y: %f\n", x, y)
	}

	for _, xy := range wgs {
		x, y := rijksdriehoek.WGS84toRD(xy[0], xy[1])
		fmt.Printf("RD - x: %f, y: %f\n", x, y)
	}
}
