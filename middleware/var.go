package middleware
import (
	"github.com/satori/go.uuid"
)

func GenerateUuid() string {
	u4 := uuid.NewV4()
	return u4.String()
}