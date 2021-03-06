package utils

import "github.com/AnuchitPrasertsang/roshambo/decide"

func MapUserKey(key string) string {
	m := map[string]string{
		"w": decide.Scissors,
		"b": decide.Rock,
		"p": decide.Paper,
	}
	return m[key]
}
