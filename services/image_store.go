package service

import (
	"bytes"
	"fmt"
	"os"
	"sync"

	"github.com/jasonrowsell/speccy/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ImageStore interface {
	SaveImage(id string, format string, imageData bytes.Buffer) error
}

type DiskImageStore struct {
	mutex       sync.RWMutex
	imageFolder string
	images      map[string]*ImageInfo
}

type ImageInfo struct {
	Id     string
	Format string
	Path   string
}

func NewDiskImageStore(imageFolder string) *DiskImageStore {
	return &DiskImageStore{
		imageFolder: imageFolder,
		images:      make(map[string]*ImageInfo),
	}
}

func (store *DiskImageStore) SaveImage(id string, format string, imageData bytes.Buffer) error {
	if err := validateImageId(store, id); err != nil {
		return err
	}

	imagePath := fmt.Sprintf("%s/%s%s", store.imageFolder, id, format)

	file, err := os.Create(imagePath)
	if err != nil {
		return status.Errorf(codes.Internal, "cannot create image file: %v", err)
	}

	_, err = imageData.WriteTo(file)
	if err != nil {
		return status.Errorf(codes.Internal, "cannot write image data to file: %v", err)
	}

	store.mutex.Lock()
	defer store.mutex.Unlock()

	store.images[id] = &ImageInfo{
		Id:     id,
		Format: format,
		Path:   imagePath,
	}

	return nil
}

func validateImageId(store *DiskImageStore, id string) error {
	if err := utils.ValidateLaptopId(id); err != nil {
		return err
	}

	if _, ok := store.images[id]; ok {
		return status.Errorf(codes.AlreadyExists, "image with ID: %s already exists", id)
	}
	return nil
}
