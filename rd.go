package rijksdriehoek

import "math"

type coefficients struct {
	p  float64
	q  float64
	pq float64
}

const (
	x0 = 155000
	y0 = 463000

	phi0 = 52.15517440
	lam0 = 5.38720621
)

var k = []coefficients{
	{p: 0, q: 1, pq: 3235.65389},
	{p: 2, q: 0, pq: -32.58297},
	{p: 0, q: 2, pq: -0.24750},
	{p: 2, q: 1, pq: -0.84978},
	{p: 0, q: 3, pq: -0.06550},
	{p: 2, q: 2, pq: -0.01709},
	{p: 1, q: 0, pq: -0.00738},
	{p: 4, q: 0, pq: 0.00530},
	{p: 2, q: 3, pq: -0.00039},
	{p: 4, q: 1, pq: 0.00033},
	{p: 1, q: 1, pq: -0.00012},
}

var l = []coefficients{
	{p: 1, q: 0, pq: 5260.52916},
	{p: 1, q: 1, pq: 105.94684},
	{p: 1, q: 2, pq: 2.45656},
	{p: 3, q: 0, pq: -0.81885},
	{p: 1, q: 3, pq: 0.05594},
	{p: 3, q: 1, pq: -0.05607},
	{p: 0, q: 1, pq: 0.01199},
	{p: 3, q: 2, pq: -0.00256},
	{p: 1, q: 4, pq: 0.00128},
	{p: 0, q: 2, pq: 0.00022},
	{p: 2, q: 0, pq: -0.00022},
	{p: 5, q: 0, pq: 0.00026},
}

// RDtoWGS84 converts a RD coordinate to WGS84 longitude and latitude.
func RDtoWGS84(x, y float64) (long, lat float64) {
	dx := math.Pow10(-5) * (x - x0)
	dy := math.Pow10(-5) * (y - y0)

	long = phi0
	lat = lam0

	for _, coef := range k {
		long += coef.pq * math.Pow(dx, coef.p) * math.Pow(dy, coef.q) / 3600
	}

	for _, coef := range l {
		lat += coef.pq * math.Pow(dx, coef.p) * math.Pow(dy, coef.q) / 3600
	}

	return long, lat
}

var r = []coefficients{
	{p: 0, q: 1, pq: 190094.945},
	{p: 1, q: 1, pq: -11832.228},
	{p: 2, q: 1, pq: -114.221},
	{p: 0, q: 3, pq: -32.391},
	{p: 1, q: 0, pq: -0.705},
	{p: 3, q: 1, pq: -2.340},
	{p: 1, q: 3, pq: -0.608},
	{p: 0, q: 2, pq: -0.008},
	{p: 2, q: 3, pq: 0.148},
}

var s = []coefficients{
	{p: 1, q: 0, pq: 309056.544},
	{p: 0, q: 2, pq: 3638.893},
	{p: 2, q: 0, pq: 73.077},
	{p: 1, q: 2, pq: -157.984},
	{p: 3, q: 0, pq: 59.788},
	{p: 0, q: 1, pq: 0.433},
	{p: 2, q: 2, pq: -6.439},
	{p: 1, q: 1, pq: -0.032},
	{p: 0, q: 4, pq: 0.092},
	{p: 1, q: 4, pq: -0.054},
}

// WGS84toRD converts a WGS84 longitude and latitude to RD coordinates.
func WGS84toRD(long, lat float64) (x, y float64) {
	dphi := 0.36 * (long - phi0)
	dlam := 0.36 * (lat - lam0)

	x = x0
	y = y0

	for _, coef := range r {
		x += coef.pq * math.Pow(dphi, coef.p) * math.Pow(dlam, coef.q)
	}

	for _, coef := range s {
		y += coef.pq * math.Pow(dphi, coef.p) * math.Pow(dlam, coef.q)
	}

	return x, y
}
