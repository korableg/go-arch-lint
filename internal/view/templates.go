package view

import (
	"fmt"

	"github.com/fe3dback/go-arch-lint/internal/models"
)

var Templates = map[string]string{
	tpl(models.Version{}): Version,
	tpl(models.Check{}):   Check,
}

func tpl(model interface{}) string {
	return fmt.Sprintf("%T", model)
}
