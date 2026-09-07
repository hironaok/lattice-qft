import sys
import matplotlib.pyplot as plot
import numpy as np

from pathlib import Path

data = np.loadtxt(sys.argv[1], dtype = float)
X = data[1:, 1]
Y = data[1:, 3]
E = data[1:, 5]

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
    r'The density-density correlator $\left<\rho_u(x)\rho_u(0)\right>$ (disconnected part)',
    fontsize = 15,
    pad = 15
)
plot.xlabel(
    '$x$',
    fontsize = 15,
    labelpad = 5
)
plot.ylabel(
    r'$\left<\mathrm{tr}\,\gamma_0 S_u(x, x) \, \mathrm{tr}\, \gamma_0 S_u(0, 0)\right>_\mathrm{gauge}$',
    fontsize = 15,
    labelpad = 5
)

# plot.grid()

figname = Path(sys.argv[1]).stem
plot.savefig(f'/home/hironao/lattice-qft/pictures/{figname}.png', dpi = 200)
