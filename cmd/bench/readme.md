# bench

Benchmark a command by running it a couple of times in a loop.

## Usage

    bench [-t RUNTIME] [-i ITERATIONS] CMD [args...]

Option `-t` chooses how long to keep re-running the command. The
default is one second; you can specify e.g. `-t 10s` to keep running
it for about 10 seconds.

Option `-i` chooses the minimum number of iterations (default: 3) that
will be performed (regardless of the value of `-t`).

You should probably always run this command twice and discard the
first result, as loading the benchmarked program into the memory for
the first time will likely have significant overhead.

Note that this program runs the command directly, without the usage of
a shell. If you'd like to benchmark evaluating a shell snippet, you
can do this instead:

    bench sh -c "echo 42"

(Note that the shell has a non-negligible startup overhead.)

## Example output

The system is Macmini9,1 (2020; M1), running macOS 13.4.1.

```
$ bench true
iters:  864
total:  997.387237ms
min:    1.008834ms
max:    6.688833ms
mean:   1.154383ms
stddev: 237.106µs
```

```
$ bench sleep 0.1
iters:  10
total:  1.099196583s
min:    108.403916ms
max:    111.793834ms
mean:   109.919658ms
stddev: 1.012818ms
```

```
$ bench sh -c "echo 42"
iters:  436
total:  999.19813ms
min:    2.016042ms
max:    7.057292ms
mean:   2.291738ms
stddev: 424.521µs
```
