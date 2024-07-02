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
	result := s.db.Preload("Descriptors").Find(&rateLimits)
	return rateLimits, result.Error
}

// GetRateLimit retrieves a RateLimit by ID with its descriptors.
func (s *RateLimitStore) GetRateLimit(id uint) (*models.RateLimit, error) {
	var rateLimit models.RateLimit
	if err := s.db.Preload("Descriptors").First(&rateLimit, id).Error; err != nil {
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
	// find rate_limit descriptor with key and value
	query := s.db.Where("key = ?", descriptor.Key)
	if descriptor.Value != "" {
		query = query.Where("value = ?", descriptor.Value)
	}

	var descriptors []models.RateLimitDescriptor
	if err := query.Find(descriptors).Error; err != nil {
		return err
	}

	// find descriptor from descriptorList and update
	for _, instance := range descriptors {
		if instance.Policy.Unit == descriptor.Policy.Unit {
			// update if exist
			descriptor.ID = instance.ID
			return s.db.Save(descriptor).Error
		}
	}

	// create new one if not found
	return s.db.Create(descriptor).Error
}
