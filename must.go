package must

import "errors"

func Do(err error) {
	if err != nil {
		panic(err)
	}
}

func Be[T any](t T, ok bool) T {
	if !ok {
		panic(errors.New("not ok!"))
	}
	return t
}

func Return[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
