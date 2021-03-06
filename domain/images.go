package domain

import (
	"io"
	"os"
	"path/filepath"

	"strings"

	"github.com/maleck13/local/config"
	"github.com/pkg/errors"
)

// Uploader defines an interface for handling uploads
type Uploader interface {
	Upload(name string, r io.Reader) (string, error)
}

// UploadService is responsible for handling uploads
type UploadService struct {
	Config *config.Config
}

//TODO at somepoint will likely want to use aws s3 so we will need a factory to return the correct implementation of Uploader

// Upload implements Uploader
func (us UploadService) Upload(name string, r io.Reader) (string, error) {

	if us.Config.Files.Local.Enabled {
		dirPath := us.Config.Files.Local.DirPath
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			return "", errors.Wrap(err, "failed to get current dir")
		}
		fPath := filepath.Join(dir, dirPath, name)
		if err := os.MkdirAll(dirPath, 0655); err != nil {
			return "", errors.Wrap(err, "failed to create local dir ")
		}
		localF, err := os.OpenFile(fPath, os.O_CREATE|os.O_RDWR, 0655)
		if err != nil {
			return "", errors.Wrap(err, "failed to create local file "+fPath)
		}
		if _, err := io.Copy(localF, r); err != nil {
			return "", errors.Wrap(err, "failed to copy contents to file")
		}
		dirPath = strings.Replace(dirPath, "web/dist/", "", 1)
		return filepath.Join(dirPath, name), nil
	}
	return "", nil
}
