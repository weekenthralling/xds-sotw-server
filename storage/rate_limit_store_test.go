package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xds-sotw-server/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGetAllRateLimits(t *testing.T) {

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}

	db.AutoMigrate(&models.RateLimit{}, &models.RateLimitDescriptor{})

	rateLimits := []models.RateLimit{
		{
			ID:     1,
			Name:   "Test RateLimit 1",
			Domain: "example1.com",
			Descriptors: []models.RateLimitDescriptor{
				{
					ID:  1,
					Key: "foo",
					Policy: models.RateLimitPolicy{
						Unit:            2,
						RequestsPerUnit: 5,
					},
				},
			},
		},
		{
			ID:     2,
			Name:   "Test RateLimit 2",
			Domain: "example2.com",
			Descriptors: []models.RateLimitDescriptor{
				{
					ID:  2,
					Key: "foo",
				},
			},
		},
	}
	db.Create(rateLimits)

	store := NewRateLimitStore(db)
	results, err := store.ListRateLimit()

	assert.NoError(t, err, "List rate limit should not return an error")
	assert.Len(t, results, 2)
}
