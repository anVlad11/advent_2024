package utils

import (
	"github.com/anVlad11/advent_2024/pkg/data"
	"io"
	"strconv"
	"strings"
)

func MustParseInt64(in string) int64 {
	val, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		panic(err)
	}

	return val
}

func MustParseInt(in string) int {
	val, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}

	return val
}

func MustParseFloat64(in string) float64 {
	val, err := strconv.ParseFloat(in, 64)
	if err != nil {
		panic(err)
	}

	return val
}

func FormatInt64(in int64) string {
	return strconv.FormatInt(in, 10)
}

func ConvertSlice[T any, K any](in []T, converter func(T) K) []K {
	out := make([]K, len(in))

	for i := range in {
		out[i] = converter(in[i])
	}

	return out
}

func GetInput(path string) ([]string, error) {
	inputs, err := data.Inputs.Open(path)
	if err != nil {
		return nil, err
	}

	input, err := io.ReadAll(inputs)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(input), "\n"), nil
}
