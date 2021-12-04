package test

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	useInput = flag.Bool("use_input", false, "Indicate that the example should be loaded")
)
 
func Read[T any](toT func(string) (T, error)) ([]T, error) {
	filename := "example.txt"
	if *useInput {
		filename = "input.txt"
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("os.Readfile:%v; %w", filename, err)
	}

	parts := strings.Split(string(content), "\n")
	typedResult := []T{}
	for _, part := range parts {
		res, err := toT(part)
		if err != nil {
			return nil, fmt.Errorf("toT failed; %w", err)
		}

		typedResult = append(typedResult, res)
	}
	return typedResult, nil
}

func Dec2Int(txt string) (int, error) {
	tmp, err := strconv.ParseInt(txt, 10, 32)
	if err != nil {
		return -1, err
	}
	return int(tmp), nil
}
