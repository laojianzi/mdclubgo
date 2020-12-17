package avatar

import (
	"crypto/md5" // nolint:gosec
	"fmt"
	"io"
	"path/filepath"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/laojianzi/mdclubgo/conf"
	"github.com/laojianzi/mdclubgo/internal/storage"
	"github.com/laojianzi/mdclubgo/internal/web"
)

// Brandable object logo e.g: user-avatar, user-cover, topic-cover
type Brandable interface {
	// e.g: user-avatar, user-cover, topic-cover
	BrandType() string

	// e.g:
	//
	// {
	//   'small': [width number, height number],
	//   'middle': [width number, height number],
	//   'large': [width number, height number]
	// }
	BrandSize() map[string][2]int

	// e.g: {'original': 'url', 'size string': 'url'}
	DefaultBrandUrls() map[string]string
}

// BrandPath get image relative path
func BrandPath(brand Brandable, userID uint, filename string) string {
	hash := fmt.Sprintf("%x", md5.Sum([]byte(strconv.Itoa(int(userID))))) // nolint:gosec
	path := filepath.Join(brand.BrandType(), hash[0:2], hash[2:4])

	return filepath.Join(fmt.Sprintf("/%s", path), filename)
}

// BrandUrls get image urls by brand
func BrandUrls(ctx echo.Context, brand Brandable, userID uint, filename string) map[string]string {
	var result map[string]string
	var needSuffix bool
	var defaultStaticPath string
	if filename == "" {
		needSuffix = true
		defaultStaticPath = "static"
		result = brand.DefaultBrandUrls()
	} else {
		path := BrandPath(brand, userID, filename)
		thumbs := brand.BrandSize()
		result = storage.Read(path, thumbs)
	}

	if conf.Storage.Type == storage.Local {
		if defaultStaticPath == "" {
			defaultStaticPath = "upload"
		}

		suffix := ".jpg"
		if web.SupportWebp(ctx) {
			suffix = ".webp"
		}

		for k := range result {
			result[k] = fmt.Sprintf("%s%s", web.StaticPath(ctx, defaultStaticPath), result[k])
			if needSuffix {
				result[k] = fmt.Sprintf("%s%s", result[k], suffix)
			}
		}
	}

	return result
}

// DeleteImage delete image file
func DeleteImage(brand Brandable, userID uint, filename string) error {
	path := BrandPath(brand, userID, filename)
	thumbs := brand.BrandSize()

	return storage.Delete(path, thumbs)
}

// UploadImage upload image and return a filename
func UploadImage(brand Brandable, userID uint, mediaType string, reader io.Reader) (string, error) {
	token := middleware.DefaultRequestIDConfig.Generator()
	suffix := "png"
	if mediaType != "image/png" {
		suffix = "jpg"
	}

	filename := fmt.Sprintf("%s.%s", token, suffix)
	path := BrandPath(brand, userID, filename)
	thumbs := brand.BrandSize()
	return filename, storage.Write(path, reader, thumbs)
}
