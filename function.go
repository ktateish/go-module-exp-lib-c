package c

import "fmt"

//go:generate ./genvers.sh

func F(s string) string {
	return fmt.Sprintf("C %s: bar %s", Version, s)
}

func G(s string) string {
	return fmt.Sprintf("D %s: buz %s", Version, s)
}
