package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	Nx    = 48
	Ny    = 16
	Nz    = 16	
	Nt    = 4
	Npos  = Nx / 2 + 1

	T0    = 0.76
	T1    = 1.20
	T2    = 1.69
	T3    = 3.01
)

func main() {
	var Nconf int
	var temp, part string

	flag.IntVar(&Nconf, "n", 0, "number of configurations")
	flag.StringVar(&temp, "t", "", "temperature")
	flag.StringVar(&part, "p", "", "connected or disconnected")
	flag.Parse()
	
	correlator := make([][]complex128, Nconf)
	for iconf := 0; iconf < Nconf; iconf++ {
		data, err := Parse(fmt.Sprintf("/home/hironao/lattice-qft/data/%s/%s/%d.dat", temp, part, iconf))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		correlator[iconf] = data
        }
	
	pos := make([]complex128, Npos)
	for ipos := 0; ipos < Npos; ipos++ {
		pos[ipos] = complex(float64(ipos), 0)
	}

	ensemble := NewEnsemble(correlator)
	average  := ensemble.Average()
	error    := ensemble.Error()

	Display(RescaleLength(pos, T0), RescaleLength(pos, T1), RescaleLength(pos, T2), RescaleLength(pos, T3))
}

func Parse(filename string) ([]complex128, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data := make([]complex128, Nx)
	scanner := bufio.NewScanner(f)
	for i := 0; i < Nx; i++ {
		scanner.Scan()
		line := strings.Split(strings.TrimSpace(scanner.Text()), "   ")
		re, err := strconv.ParseFloat(strings.TrimSpace(line[1]), 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		im, err := strconv.ParseFloat(strings.TrimSpace(line[2]), 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		data[i] = complex(re, im)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return data, nil
}

func Display(records ...[]complex128) {
        for ipos := 0; ipos < Npos; ipos++ {
                fmt.Printf("%4d", ipos)
                for _, r := range records {
                        fmt.Printf("  %20.12e  %20.12e", real(r[ipos]), imag(r[ipos]))
                }
                fmt.Printf("\n")
        }
}

func RescaleLength(record []complex128, temp float64) []complex128 {
        values := make([]complex128, Npos)
        for ipos := 0; ipos < Npos; ipos++ {
                values[ipos] = record[ipos] / complex(Nt * temp, 0)
        }
        return values
}

func RescaleCorrelator(record []complex128, temp float64) []complex128 {
        values := make([]complex128, Npos)
        for ipos := 0; ipos < Npos; ipos++ {
                values[ipos] = record[ipos] * complex(math.Pow(Nt * temp, 6), 0)
        }
        return values
}

func LogScale(record []complex128) []complex128 {
        values := make([]complex128, Npos)
        for ipos := 0; ipos < Npos; ipos++ {
                values[ipos] = complex(math.Log(real(record[ipos])), math.Log(imag(record[ipos])))
        }
        return values
}

type Ensemble [][]complex128
	
func NewEnsemble(raw [][]complex128) Ensemble {
	var ensemble Ensemble
	for iconf := 0; iconf < Nconf; iconf++ {
		lhs := make([]complex128, Npos)
		rhs := make([]complex128, Npos)
		for ipos := 0; ipos < Npos; ipos++ {
			lhs[ipos] = (raw[iconf])[ipos]
			rhs[ipos] = (raw[iconf])[(Nx - ipos) % Nx]
		}
		ensemble = append(ensemble, lhs)
		ensemble = append(ensemble, rhs)
	}
	return ensemble
}

func (e Ensemble) Average() []complex128 {
	average := make([]complex128, Npos)
	for ipos := 0; ipos < Npos; ipos++ {
		for i := 0; i < len(e); i++ {
			average[ipos] += e[i][ipos] / complex(float64(len(e)), 0)
		}
	}
	return average
}

func (e Ensemble) Error() []complex128 {
        average := e.Average()

        re    := make([]float64, Npos)
        im    := make([]float64, Npos)
	error := make([]complex128, Npos)
        for ipos := 0; ipos < Npos; ipos++ {
                for i := 0; i < len(e); i++ {
                        re[ipos] += math.Pow(real(e[i][ipos] - average[ipos]), 2) / float64(len(e))
                        im[ipos] += math.Pow(imag(e[i][ipos] - average[ipos]), 2) / float64(len(e))
                }
		error[ipos] = complex(math.Sqrt(re[ipos] / float64(len(e) - 1)), math.Sqrt(im[ipos] / float64(len(e) - 1)))
        }
        return error	
}
