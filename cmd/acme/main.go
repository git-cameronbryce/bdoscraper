package main

import (
	"flag"

	"github.com/git-cameronbryce/internal"
)

func main() {
	var cname = flag.String("character", "Ohwaaa", "Lookup a character, return their stats.")
	flag.Parse()

	internal.Selector(cname)
	internal.FamilyStats(cname)
}
