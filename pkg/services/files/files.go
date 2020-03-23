package files

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/mushroomsir/mimetypes"
	"io"
	"log"
	"os"
	"path/filepath"
)

type FileService struct {
	storagePath string
}

type StoragePathType string

func NewFilesSvc(mediaPath StoragePathType) *FileService {
	if mediaPath == "" {
		panic(errors.New("media path can't be nil")) // <- be accurate
	}

	return &FileService{storagePath: string(mediaPath)}
}

func (receiver *FileService) Save(src io.Reader, contentType string) (name string, err error) {
	extensions := mimetypes.Extension(contentType)

	uuidV4 := uuid.New().String()

	name = fmt.Sprintf("%s%s", uuidV4, "." + extensions)
	path := filepath.Join(receiver.storagePath, name)

	dst, _ := os.Create(path)
	defer func() {
		errdefer := dst.Close()
		if errdefer != nil {
			log.Print(errdefer)
		}
	}()
	_, err = io.Copy(dst, src)
	if err != nil {
		return "", err
	}
	return name, nil
}
