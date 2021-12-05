package test

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var filename = "example.txt"

func Read[T any](toT func(string) (T, error)) ([]T, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("os.Readfile:%v; %w", filename, err)
	}

	if len(content) == 0 {
		return nil, fmt.Errorf("%s: no data", filename)
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

func String(txt string) (string, error) {
	return txt, nil
}
