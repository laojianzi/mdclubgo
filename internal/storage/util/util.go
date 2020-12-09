package util

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/disintegration/imaging"
)

func GetThumbLocation(location, size string) string {
	locationArr := strings.Split(location, ".")
	locationArr[len(locationArr)-2] = fmt.Sprintf("%s_%s", locationArr[len(locationArr)-2], size)

	return strings.Join(locationArr, ".")
}

func CROP(reader io.Reader, location string, thumbs map[string][2]int,
	callback func(string, io.Reader) error) error {
	img, err := imaging.Decode(reader)
	if err != nil {
		return fmt.Errorf("imaging decode: %w", err)
	}

	originalWidth := img.Bounds().Dx()
	originalHeight := img.Bounds().Dy()
	for size, thumb := range thumbs {
		width := thumb[0]
		height := thumb[1]
		newImage := imaging.Clone(img)
		if height == 0 {
			height = originalHeight / originalWidth * width
		}

		newImageWidth := width
		if originalWidth/originalHeight < width/height {
			newImageWidth = originalWidth
			newImage = imaging.CropCenter(newImage, newImageWidth, originalWidth/(width/height))
		} else if originalWidth/originalHeight > width/height {
			newImageWidth = originalHeight * (width / height)
			newImage = imaging.CropCenter(newImage, newImageWidth, originalHeight)
		}

		if width < newImageWidth {
			newImage = imaging.Resize(newImage, width, height, imaging.NearestNeighbor)
		}

		buf := bytes.NewBuffer([]byte{})
		f, err := imaging.FormatFromFilename(location)
		if err != nil {
			return fmt.Errorf("imaging format from file name: %w", err)
		}

		err = imaging.Encode(buf, newImage, f)
		if err != nil {
			return fmt.Errorf("imaging encode: %w", err)
		}

		cropLocation := GetThumbLocation(location, size)
		if err = callback(cropLocation, buf); err != nil {
			return fmt.Errorf("callback: %w", err)
		}
	}

	return nil
}
