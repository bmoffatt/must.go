package must

func Do(err error) {
	if err != nil {
		panic(err)
	}
}

func Return[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
