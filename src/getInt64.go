package main

import (
	"fmt"
	"strconv"
)

func getInt64(txt string) (int64, error) {

	a, err := strconv.ParseInt(txt, 10, 64)

	if err != nil {
		return 0,
			fmt.Errorf(
				"must be valid 64-bit signed integer in string: '%s' error: %w",
				txt,
				err,
			)
	}

	return a, nil
}
