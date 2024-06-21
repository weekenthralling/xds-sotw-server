package storage

import (
	"github.com/xds-sotw-server/models"
	"gorm.io/gorm"
)

type RateLimitStore struct {
	db *gorm.DB
}

func NewRateLimitStore(db *gorm.DB) *RateLimitStore {
	return &RateLimitStore{
		db: db,
	}
}

// ListRateLimit retrieves all RateLimit with its descriptors.
func (s *RateLimitStore) ListRateLimit() ([]models.RateLimit, error) {
	var rateLimits []models.RateLimit
	result := s.db.Preload("Descriptors").Preload("Descriptors.Descriptors").Find(&rateLimits)
	return rateLimits, result.Error
}

// GetRateLimit retrieves a RateLimit by ID with its descriptors.
func (s *RateLimitStore) GetRateLimit(id uint) (*models.RateLimit, error) {
	var rateLimit models.RateLimit
	if err := s.db.Preload("Descriptors").Preload("Descriptors.Descriptors").First(&rateLimit, id).Error; err != nil {
		return nil, err
	}
	return &rateLimit, nil
}

// GetRateLimitByDomain find rate limit by domain
func (s *RateLimitStore) GetRateLimitByDomain(domain string) (*models.RateLimit, error) {
	var rateLimit models.RateLimit
	if err := s.db.First(&rateLimit, "domain = ?", domain).Error; err != nil {
		return nil, err
	}
	return &rateLimit, nil
}

// SaveRateLimitDescriptor save rate limit descriptor
func (s *RateLimitStore) SaveRateLimitDescriptor(descriptor *models.RateLimitDescriptor) error {
	result := s.db.Save(descriptor)
	return result.Error
}
