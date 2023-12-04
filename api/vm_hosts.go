package api

import (
	"github.com/maas/gomaasclient/entity"
)

// VMHosts is an interface for listing and creating VMHost objects
type VMHosts interface {
	Get() ([]entity.VMHost, error)
	Create(params *entity.VMHostParams) (*entity.VMHost, error)
}
