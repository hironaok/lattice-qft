package main

import (
        "bufio"
        "fmt"
        "os"
        "strconv"
        "strings"
)

const (
        Nconf = 860
        Nx    = 16
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
                read, err := os.Open(fmt.Sprintf("/home/hironao/lattice-qft/per-configuration/experiment/charge/%d.dat", iconf))
                if err != nil {
                        fmt.Fprintf(os.Stderr, "%v\n", err)
                        os.Exit(1)
                }
                defer read.Close()

                scanner := bufio.NewScanner(read)
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

	for iconf := 0; iconf < Nconf; iconf++ {
		write, err := os.Create(fmt.Sprintf("/home/hironao/lattice-qft/per-configuration/experiment/correlator/%d.dat", iconf))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		for i := 0; i < Nx; i++ {
			fmt.Fprintf(write, "%4d", i)
			fmt.Fprintf(write, "  %20.12e  %20.12e", real((correlator[iconf])[i]), imag((correlator[iconf])[i]))
			fmt.Fprintf(write, "\n")
		}
	}
}
