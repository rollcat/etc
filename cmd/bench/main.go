package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
	"unsafe"

	"github.com/aclements/go-moremath/stats"
	"github.com/rollcat/getopt"
)

func cloneCmd(cmd *exec.Cmd) *exec.Cmd {
	return &exec.Cmd{
		Path: cmd.Path,
		Args: cmd.Args,
	}
}

func runCmd(cmd *exec.Cmd) time.Duration {
	start := time.Now()
	err := cmd.Run()
	if err != nil {
		fmt.Printf("error:   %v\n", err)
		os.Exit(1)
	}
	end := time.Now()
	return end.Sub(start)
}

func benchmark(cmd *exec.Cmd, runtime time.Duration, iterations int) (results []time.Duration) {
	start := time.Now()
	for {
		results = append(results, runCmd(cloneCmd(cmd)))
		totalElapsed := time.Now().Sub(start)
		if len(results) >= iterations && totalElapsed >= runtime {
			break
		}
	}
	return
}

type num interface{ ~float64 | ~int64 }

func sum[T num](xs []T) (o T) {
	for _, x := range xs {
		o += x
	}
	return
}

func floats[T num](xs []T) []float64 {
	out := make([]float64, len(xs))
	for i, x := range xs {
		out[i] = float64(x)
	}
	return out
}

func min[T num](xs []T) T {
	if len(xs) == 0 {
		panic("len(xs) == 0")
	}
	m := xs[0]
	for _, x := range xs {
		if x < m {
			m = x
		}
	}
	return m
}

func max[T num](xs []T) T {
	if len(xs) == 0 {
		panic("len(xs) == 0")
	}
	m := xs[0]
	for _, x := range xs {
		if x > m {
			m = x
		}
	}
	return m
}

func usage() {
	println("Usage: bench [-t RUNTIME] [-i ITERATIONS] CMD [args...]")
}

func main() {
	args, opts, err := getopt.GetOpt(
		os.Args[1:],
		"ht:i:",
		nil,
	)
	if err != nil || len(args) == 0 {
		usage()
		os.Exit(1)
	}

	var runtime = 1 * time.Second
	var iterations int = 3
	for _, opt := range opts {
		switch opt.Opt() {
		case "-h":
			usage()
			os.Exit(0)
		case "-t":
			runtime, err = time.ParseDuration(opt.Arg())
			if err != nil {
				fmt.Printf("%v\n", err)
				usage()
				os.Exit(1)
			}
		case "-i":
			var i int64
			i, err = strconv.ParseInt(
				opt.Arg(),
				10,
				int(unsafe.Sizeof(iterations)),
			)
			iterations = int(i)
			if err != nil {
				fmt.Printf("%v\n", err)
				usage()
				os.Exit(1)
			}
		default:
			panic("unexpected argument")
		}
	}

	cmd := exec.Command(args[0], args[1:]...)
	results := benchmark(cmd, runtime, iterations)
	resultsF := floats(results)
	fmt.Printf("iters:  %v\n", len(results))
	fmt.Printf("total:  %v\n", sum(results))
	fmt.Printf("min:    %v\n", min(results))
	fmt.Printf("max:    %v\n", max(results))
	fmt.Printf("mean:   %v\n", time.Duration(stats.Mean(resultsF)))
	fmt.Printf("stddev: %v\n", time.Duration(stats.StdDev(resultsF)))
}
