package main
import (
	"path/filepath"
	"log"
	"os"
	"io"
	"flag"
)

func main() {
	var srcPath string
	var dstPath string

	flag.StringVar(&srcPath, "src", "", "Path to source files")
	flag.StringVar(&dstPath, "dst", "", "Path to copy the matching files to")

	flag.Parse()

	if srcPath == "" {
		log.Fatalf("You need to set the -src flag")
	}

	if dstPath == "" {
		log.Fatalf("You need to set the -dst flag")
	}

	newImages, err := filepath.Glob(srcPath + "/*.png")
	if err != nil {
		log.Fatalf("Error getting new images: %s", err)
	}

	existingImages, err := filepath.Glob(dstPath + "/*.png")
	if err != nil {
		log.Fatalf("Error getting existing images: %s", err)
	}

	log.Printf("newImages: %d / existingImages: %d", len(newImages), len(existingImages))

	imageMap := make(map[string]string)

	for _, newImage := range newImages {
		_, fileName := filepath.Split(newImage)
		imageMap[fileName] = newImage
	}

	for _, existingImage := range existingImages {
		_, existingFileName := filepath.Split(existingImage)
		newImage, exists := imageMap[existingFileName]

		if !exists {
			log.Printf("No new image exists for %s!", existingFileName)
		} else {
			err := cp(existingImage, newImage)

			if err != nil {
				log.Fatalf("Unable to copy file %s: %s", existingFileName, err)
			}
		}
	}
}

func cp(dst, src string) error {
	s, err := os.Open(src)
	if err != nil {
		return err
	}
	// no need to check errors on read only file, we already got everything
	// we need from the filesystem, so nothing can go wrong now.
	defer s.Close()
	d, err := os.Create(dst)
	if err != nil {
		return err
	}
	if _, err := io.Copy(d, s); err != nil {
		d.Close()
		return err
	}
	return d.Close()
}
