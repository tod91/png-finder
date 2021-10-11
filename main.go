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
	var wg sync.WaitGroup
	switch {
	case src != "":
		l.SetRootSrcPath(src)
		wg.Add(1)
		func() {
			go l.LoadSrcImages()
			wg.Done()
		}()
		fallthrough
	case dest != "":
		l.SetRootDestPath(dest)
		wg.Add(1)
		func() {
			go l.LoadDestImages()
			wg.Done()
		}()
		wg.Wait()
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
