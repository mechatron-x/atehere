package port

import "github.com/mechatron-x/atehere/internal/billing/domain/aggregate"

type (
	PostOrderRepository interface {
		Save(postOrder *aggregate.PostOrder) error
	}
)
