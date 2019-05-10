package sizefmt

import (
	// "math"
	"strconv"
)

const (
	kilo  = 1000
	mega  = kilo * kilo
	giga  = kilo * mega
	tera  = kilo * giga
	peta  = kilo * tera
	exa   = kilo * peta
	zetta = kilo * exa
	yotta = kilo * zetta
)
const (
	kibi = 1024
	mebi = kibi * kibi
	gibi = kibi * mebi
	tebi = kibi * gibi
	pebi = kibi * tebi
	exbi = kibi * pebi
	zebi = kibi * pebi
	yobi = kibi * zebi
)

const units = "KMGTPEZY"

const (
	IEC  = "iec"
	IECi = "iec-i"
	SI   = "si"
)

func FormatIEC(n float64, i bool) string {
	str := IEC
	if i {
		str = IECi
	}
	return Format(n, str)
}

func FormatSI(n float64) string {
	return Format(n, SI)
}

func Format(n float64, meth string) string {
	var sizes []float64
	switch meth {
	default:
	case SI:
		sizes = []float64{kilo, mega, giga, tera, peta, exa, zetta, yotta}
	case IEC, IECi:
		sizes = []float64{kibi, mebi, gibi, tebi, pebi, exbi, zebi, yobi}
	}

	var xs []byte
	if len(sizes) > 0 {
		xs = formatSize(n, sizes)
		if meth == IECi {
			xs = append(xs, 'i')
		}
	} else {
		xs = trimZeros(appendFloat(n, 0))
	}
	return string(xs)
}

func formatSize(n float64, sizes []float64) []byte {
	if n < sizes[0] {
		return appendFloat(n, 0)
	}
	var (
		div  float64
		unit byte
	)
	for i := 1; i < len(sizes); i++ {
		if n < sizes[i] {
			div, unit = sizes[i-1], units[i-1]
			break
		}
	}
	if div == 0 {
		div, unit = sizes[len(sizes)-1], units[len(units)-1]
	}
	return appendFloat(n/div, unit)
}

func appendFloat(n float64, unit byte) []byte {
	buf := make([]byte, 0, 16)
	buf = strconv.AppendFloat(buf, n, 'f', 2, 64)
	if unit != 0 {
		buf = append(trimZeros(buf), unit)
	}
	return buf
}

func trimZeros(buf []byte) []byte {
	i := len(buf) - 1
	for i >= 0 && buf[i] == '0' {
		i--
	}
	if buf[i] != '.' {
		i++
	}
	return buf[:i]
}
