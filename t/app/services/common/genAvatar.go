package common

import (
	"github.com/google/uuid"
	"github.com/o1egl/govatar"
	"service-api/src/core/helpers"
)

func GenAvatar(key string) string {
	dir := helpers.PathInstance.Resource("avatar/member/" + uuid.NewString() + ".jpg")
	err := govatar.GenerateFileForUsername(govatar.MALE, key, dir)
	if err != nil {
		return ""
	}

	return dir
}
