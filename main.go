package main

import (
	"finder/checker"
	"finder/loader"
	"flag"
	"fmt"
)

func main() {
	var src string
	var dest string
	flag.StringVar(&src, "src", "", "the root directory of all the src images")

	flag.StringVar(&dest, "dest", "", "the root directory of all the destination images images")

	flag.Parse()

	var l loader.Loader
	switch {
	case src != "":
		l.SetRootPath(src)
		l.LoadSrcImages()
		fallthrough
	case dest != "":
		l.SetRootPath(dest)
		l.LoadDestImages()
	}

	for _, src := range l.SrcImages {
		for _, dest := range l.DestImages {
			if checker.Equal(src.Dir, dest.Dir) {
				fmt.Println(src.Image + " -> " + dest.Dir)
				break
			}
		}
	}
}
