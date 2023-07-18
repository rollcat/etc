package must

func Assert(err error) {
	if err != nil {
		panic(err)
	}
}

func Must[T any](x T, err error) T {
	if err != nil {
		panic(err)
	}
	return x
}

func Must2[T1, T2 any](x1 T1, x2 T2, err error) (T1, T2) {
	if err != nil {
		panic(err)
	}
	return x1, x2
}
