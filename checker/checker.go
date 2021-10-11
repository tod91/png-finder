package checker

import (
	"github.com/vitali-fedulov/images/v2"
)

func Equal(src, dest string) bool {

	// Open photos.
	imgA, err := images.Open(src)
	if err != nil {
		panic(err)
	}
	imgB, err := images.Open(dest)
	if err != nil {
		panic(err)
	}

	// Calculate hashes and image sizes.
	hashA, imgSizeA := images.Hash(imgA)
	hashB, imgSizeB := images.Hash(imgB)

	// Image comparison.
	if images.Similar(hashA, hashB, imgSizeA, imgSizeB) {
		return true
	}
	return false
}
