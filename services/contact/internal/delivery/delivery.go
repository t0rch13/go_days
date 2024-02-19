package delivery

import (
	"architecture_go/services/contact/internal/useCase"
)

type ContactDeliveryImpl struct {
	useCase useCase.ContactUseCase
}

func NewContactDelivery(useCase useCase.ContactUseCase) ContactDelivery {
	return &ContactDeliveryImpl{
		useCase: useCase,
	}
}

type ContactDelivery interface{}
