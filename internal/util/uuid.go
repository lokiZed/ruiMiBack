package util

import (
	"github.com/google/uuid"
	"strings"
)

func GetUuid() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
