import argparse
import numpy as np
import matplotlib.pyplot as plot

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('filenames', nargs = '+', type = str)
    args = parser.parse_args()

    fig, ax = plot.subplots()

    for i, fn in enumerate(args.filenames):
        data = np.loadtxt(fn, dtype = float)
        X = np.arange(3000)
        Y = data
        ax.scatter(X[2000:3000:50], Y[2000:3000:50])

    fig.savefig('plot.png', dpi = 200)
    
if __name__ == '__main__':
    main()
