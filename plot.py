import argparse

import matplotlib.pyplot as plot
import numpy as np

EXAMPLE_FILE = "demo.dat"

MIN_DISTANCE = 0
MAX_DISTANCE = 32

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('filenames', nargs = '+', type = str, default = EXAMPLE_FILE)
    parser.add_argument('-s', '--start', type = int, default = MIN_DISTANCE)
    parser.add_argument('-e', '--end',   type = int, default = MAX_DISTANCE)
    args = parser.parse_args()

    plot.rcParams['text.usetex'] = True
    fig, ax = plot.subplots()

    for i, fn in enumerate(args.filenames):
        data = np.loadtxt(fn, dtype = float)
        X = data[args.start:args.end, 1]
        Y = data[args.start:args.end, 3]
        E = data[args.start:args.end, 5]

        ax.errorbar(
            X,
            Y,
            xerr = None,
            yerr = E,
            capsize = 5,
            fmt = 'o',
            markersize = 5
        )

    ax.grid()
    
    ax.ticklabel_format(
        axis="y",
        style="sci",
        scilimits=(0, 0)
    )
    
    ax.set_title(
        r"The density-density correlator $\left<\rho_u(x)\rho_u(0)\right>$ (connected)",
        fontsize = 15
    )
    
    ax.set_xlabel(
        r"Distance $x$",
        fontsize = 15
    )
    
    ax.set_ylabel(
        r"$\left<\mathrm{tr}\left\{\gamma_5\gamma_0 S_u(x, 0)\gamma_0\gamma_5S_u(x, 0)^\dagger\right\}\right>_\mathrm{gauge}$",
        fontsize = 15
    )
    fig.savefig('plot.png', dpi = 200)

if __name__ == '__main__':
    main()
