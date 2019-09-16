package c

import "fmt"

//go:generate ./genvers.sh

func F(s string) string {
	return fmt.Sprintf("C %s: %s", Version, s)
}

func G(s string) string {
	return fmt.Sprintf("D %s: %s", Version, s)
}
