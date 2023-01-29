package repository

import (
	"github.com/Jack-Music-Streaming/server/src/models"
	"gorm.io/gorm"
)

type scopeRepository struct {
	DB *gorm.DB
}

func NewScopeRepository(db *gorm.DB) models.ScopeRepository {
	return &scopeRepository{
		DB: db,
	}
}

func (s *scopeRepository) FindScope(roleID interface{}, planID interface{}) (*models.Scope, error) {
	var scope models.Scope

	err := s.DB.Model(&scope).Where("role_id = ? AND plan_id = ?", roleID, planID).First(&scope).Error

	if err != nil {
		return &scope, err
	}

	return &scope, nil
}
