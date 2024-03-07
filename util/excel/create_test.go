package excel

import (
	"context"
	"testing"
)

func TestCreate(t *testing.T) {
	cli := NewClient()
	err := cli.Create(context.Background(), "./test.xlsx", [][]string{
		{"Tom", "hh", "B"},
		{"B", "C", "HH"},
	})
	if err != nil {
		t.Fatal(err)
	}
}
