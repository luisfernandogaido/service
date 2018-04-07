package modelo

import (
	"strings"
	"fmt"
)

func ft (txt string) string {
	termos := strings.Split(txt, " ")
	uteis := make([]string, 0)
	for _, t := range termos {
		if t != "" {
			uteis = append(uteis, fmt.Sprintf("+%v*", t))
		}
	}
	return strings.Join(uteis, " ")
}