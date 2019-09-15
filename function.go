package c

import "fmt"

//go:generate ./genvers.sh

func F(s string) string {
	return fmt.Sprintf("C %s: %s", Version, s)
}
