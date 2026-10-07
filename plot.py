import argparse
import pathlib

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
    parser.add_argument('-t', '--temperature', type = str)
    args = parser.parse_args()

    plot.rcParams['text.usetex'] = True
    fig, ax = plot.subplots()
    axzoom = ax.inset_axes([0.3, 0.3, 0.6, 0.5])

    colors = ["red", "blue"]
    labels = ["connected", "disconnected"]
    for i, fn in enumerate(args.filenames):
        data = np.loadtxt(fn, dtype = float)
        
        X = data[args.start:args.end, 1]
        Y = data[args.start:args.end, 3]
        E = data[args.start:args.end, 5]

        ax.errorbar(
            X,
            Y,
            xerr       = None,
            yerr       = E,
            capsize    = 5,
            fmt        = 'o',
            markersize = 3,
            
            label = labels[i],
            color = colors[i],
            alpha = 0.5
        )

        axzoom.errorbar(
            X,
            Y,
            xerr       = None,
            yerr       = E,
            capsize    = 5,
            fmt        = 'o',
            markersize = 3,
            
            label = labels[i],
            color = colors[i],
            alpha = 0.5
        )
        

    axzoom.set_xlim(0.9, 1.75)
    
    axzoom.set_ylim(-1.05, 0.55)

    axzoom.grid()

    ax.indicate_inset_zoom(axzoom)

    ax.grid()
    
    ax.legend()
    
    ax.ticklabel_format(
        axis      = "y",
        style     = "sci",
        scilimits = (0, 0)
    )
    
    ax.set_title(
        r"The density-density correlator at $T = " + args.temperature + r"\,T_\mathrm{pc}$",
        fontsize = 15
    )
    
    ax.set_xlabel(
        r"Distance $x$",
        fontsize = 15
    )
    
    ax.set_ylabel(
        r"Connected/disconnected part of $\left<\rho_u(x)\rho_u(0)\right>$",
        fontsize = 15
    )
    fig.savefig('plot.png', dpi = 200)

if __name__ == '__main__':
    main()
