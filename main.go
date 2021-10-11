package main

import (
	"finder/checker"
	"finder/loader"
	"flag"
	"fmt"
	"sync"
)

func main() {
	var src string
	var dest string
	flag.StringVar(&src, "src", "", "the root directory of all the src images")

	flag.StringVar(&dest, "dest", "", "the root directory of all the destination images images")

	flag.Parse()

	var l loader.Loader
	var wg1 sync.WaitGroup
	switch {
	case src != "":
		l.SetRootSrcPath(src)
		wg1.Add(1)
		go func() {
			defer wg1.Done()
			l.LoadSrcImages()
		}()
		fallthrough
	case dest != "":
		l.SetRootDestPath(dest)
		wg1.Add(1)
		go func() {
			defer wg1.Done()
			l.LoadDestImages()
		}()
	}
	wg1.Wait()

	for _, src := range l.SrcImages {
		for _, dest := range l.DestImages {
			if checker.Equal(src.Dir, dest.Dir) {
				fmt.Println(src.Image + " -> " + dest.Dir)
				break
			}
		}
	}
}
