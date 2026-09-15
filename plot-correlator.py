import sys
import argparse
import matplotlib.pyplot as plot
import numpy as np

from pathlib import Path

parser = argparse.ArgumentParser()
parser.add_argument('datafile', type = str)
parser.add_argument('-l', '--logscale', action="store_true")
parser.add_argument('-t', '--title')
parser.add_argument('-k', '--kind', choices = ['connected', 'disconnected'])
parser.add_argument('-s', '--start', type = int, default = 0)
parser.add_argument('-e', '--end', type = int, default = 32)
args = parser.parse_args()

xlabel = r'Distance $x$'
ylabel = r''
title  = r'The density-density correlator $\left<\rho_u(x)\rho_u(0)\right>$ (' + args.title + ')'

if args.kind == 'connected':
    ylabel = r'$\left<\mathrm{tr}\,\gamma_5 \gamma_0 S_u(x, 0) \, \gamma_0 \gamma_5 S_u(x, 0)^\dagger\right>_\mathrm{gauge}$'

if args.kind == 'disconnected':
    ylabel = r'$\left<\mathrm{tr}\,\gamma_0 S_u(x, x) \, \mathrm{tr}\, \gamma_0 S_u(0, 0)\right>_\mathrm{gauge}$'

if args.logscale:
    ylabel = r'$\mathrm{log}$ ' + ylabel

data = np.loadtxt(args.datafile, dtype = float)
X = data[args.start:args.end, 1]
Y = data[args.start:args.end, 3]
E = data[args.start:args.end, 5]

plot.rcParams['text.usetex'] = True

plot.errorbar(
    X,
    Y,
    yerr = E,
    capsize = 5,
    fmt = 'o',
    markersize = 5,
    ecolor = 'black',
    markeredgecolor = "black",
    color = 'w'
)

plot.title(
    title,
    fontsize = 15,
    pad = 15
)

plot.xlabel(
    xlabel,
    fontsize = 15,
    labelpad = 5
)

plot.ylabel(
    ylabel,
    fontsize = 15,
    labelpad = 5
)

plot.ticklabel_format(
    axis="y",
    style="sci",
    scilimits=(0, 0)
)

plot.grid()

figname = Path(sys.argv[1]).stem
plot.savefig(f'/home/hironao/lattice-qft/pictures/{figname}.png', dpi = 200)
