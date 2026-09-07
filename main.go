package main

import (
	"bufio"	
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	Nconf = 860
	Nx    = 64
	Ny    = 16
	Nz    = 16	
	Nt    = 4
	Npos  = Nx / 2 + 1

	T0    = 0.70
	T1    = 1.20
	T2    = 3.01
)

func main() {
	charge := make([][]complex128, Nconf)
	for iconf := 0; iconf < Nconf; iconf++ {
		charge[iconf] = make([]complex128, Nx)
		file, err := os.Open(fmt.Sprintf("/home/hironao/lattice-qft/per-configuration/experiment/%d.dat", iconf))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
                for i:= 0; i < Nx; i++ {
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
			(charge[iconf])[i] = complex(re, im)
                }
                if err := scanner.Err(); err != nil {
                        fmt.Fprintf(os.Stderr, "%v\n", err)
                        os.Exit(1)
                }		
	}

	correlator := make([][]complex128, Nconf)
	for iconf := 0; iconf < Nconf; iconf++ {
		correlator[iconf] = make([]complex128, Nx)
		for i := 0; i < Nx; i++ {
			(correlator[iconf])[i] = (charge[iconf])[i] * (charge[iconf])[0]
		}
	}

	ensemble := NewEnsemble(correlator)
	average  := ensemble.Average()
	error    := ensemble.Error()

	pos := make([]complex128, Npos)
	for ipos := 0; ipos < Npos; ipos++ {
		pos[ipos] = complex(float64(ipos), 0)
	}
	
	Display(RescaleLength(pos), RescaleCorrelator(average), RescaleCorrelator(error))
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

func RescaleLength(record []complex128) []complex128 {
        values := make([]complex128, Npos)
        for ipos := 0; ipos < Npos; ipos++ {
                values[ipos] = record[ipos] / (Nt * T1)
        }
        return values
}

func RescaleCorrelator(record []complex128) []complex128 {
        values := make([]complex128, Npos)
        for ipos := 0; ipos < Npos; ipos++ {
                values[ipos] = record[ipos] * complex(math.Pow(Nt * T1, 6), 0)
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
