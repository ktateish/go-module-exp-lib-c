package c

import (
	"log"
	"strings"
	"testing"
)

func TestF(t *testing.T) {
	s := "foo"
	x := F(s)
	log.Println(x)
	if !strings.HasPrefix(x, "C v") {
		t.Errorf("errornous prefix")
	}
	if !strings.HasSuffix(x, s) {
		t.Errorf("errornous suffix")
	}
}

func TestG(t *testing.T) {
	s := "foo"
	x := G(s)
	log.Println(x)
	if !strings.HasPrefix(x, "D v") {
		t.Errorf("errornous prefix")
	}
	if !strings.HasSuffix(x, s) {
		t.Errorf("errornous suffix")
	}
}
