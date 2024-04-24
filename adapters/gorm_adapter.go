package adapters

import (
	"hexagonal_go/core"

	"gorm.io/gorm"
)

// Secondary Adapter
type GormOrderRepository struct {
	db *gorm.DB
}

func NewGormOrderRepository(db *gorm.DB) core.OrderRepository {
	return &GormOrderRepository{db: db}
}

// Save implements core.OrderRepository.
func (g *GormOrderRepository) Save(order core.Order) error {
	if result := g.db.Create(&order); result.Error != nil {
		return result.Error
	}
	return nil
}
