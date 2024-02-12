package usecase

import (
	"github.com/cezarcruz/gym-address/internal/domain/model"
	"github.com/cezarcruz/gym-address/internal/infra/integration"
	"github.com/cezarcruz/gym-address/internal/infra/storage"
)

func GetAndCacheAddress(zipcode string) *model.Address {

	a := storage.GetAddressBy(zipcode)

	if a == nil {
		a = integration.GetAddressBy(zipcode)
		storage.SaveAddress(a)
	}

	return a

}
