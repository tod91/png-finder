package loader

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
)

type pngData struct {
	Image string
	Dir   string
}

type Loader struct {
	Root                  string
	SrcImages, DestImages []pngData
}

func (l *Loader) SetRootPath(newPath string) {
	l.Root = newPath
}

func (l *Loader) GetRootSrcPath() string {
	return l.Root
}

func (l *Loader) LoadSrcImages() {
	libRegEx, e := regexp.Compile("([^\\s]+(\\.(?i)(png))$)")
	if e != nil {
		log.Fatal(e)
	}

	e = filepath.Walk(l.Root, func(path string, info os.FileInfo, err error) error {
		if err == nil && libRegEx.MatchString(info.Name()) {
			tmp := pngData{
				Image: info.Name(),
				Dir:   path,
			}
			l.SrcImages = append(l.SrcImages, tmp)
		}
		return nil
	})
	if e != nil {
		log.Fatal(e)
	}
}

func (l *Loader) LoadDestImages() {
	libRegEx, e := regexp.Compile("([^\\s]+(\\.(?i)(png))$)")
	if e != nil {
		log.Fatal(e)
	}

	e = filepath.Walk(l.Root, func(path string, info os.FileInfo, err error) error {
		if err == nil && libRegEx.MatchString(info.Name()) {
			tmp := pngData{
				Image: info.Name(),
				Dir:   path,
			}
			l.DestImages = append(l.DestImages, tmp)
		}
		return nil
	})
	if e != nil {
		log.Fatal(e)
	}
}
