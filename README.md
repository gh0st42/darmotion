# darmotion

A simple gui frontend based on giu and golang to automate the generation of movement traces using [bonnmotion](https://www.bonnmotion.net).

Features:

- Example values for all parameter values
- Batch creation of mobility traces
- Movement model preview using [the-ONE](https://github.com/akeranen/the-one)

Supported movement models at the moment:

- Random Waypoint
- SMOOTH

## Installation

Just run `go build` in the repo root.

To use `darmotion` make sure that *bonnmotion* is callable as `bonnmotion` somewhere in your `PATH`, e.g., `ln -s /home/user01/src/bonnmotion-src/bin/bm /usr/local/bin/bonnmotion`. 

Furthermore, for quick previews make sure that `one.sh` of *the-ONE* DTN simulator is callable from anywhere and in your `PATH`. By default this is not possible, a patched startup script can be found in this [fork](https://github.com/gh0st42/the-one).

