package resources

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xds-sotw-server/models"
	"github.com/xds-sotw-server/storage"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestListResources(t *testing.T) {

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}

	db.AutoMigrate(&models.RateLimit{}, &models.RateLimitDescriptor{})

	rateLimits := []models.RateLimit{
		{
			Model:  gorm.Model{ID: 1},
			Name:   "Test RateLimit 1",
			Domain: "example1.com",
			Descriptors: []models.RateLimitDescriptor{
				{
					Model: gorm.Model{ID: 1},
					Key:   "foo",
					Descriptors: []models.RateLimitDescriptor{
						{
							Model:           gorm.Model{ID: 3},
							Key:             "foo",
							Value:           "bar",
							Unit:            1,
							RequestsPerUnit: 10,
						},
					},
				},
			},
		},
		{
			Model:  gorm.Model{ID: 2},
			Name:   "Test RateLimit 2",
			Domain: "example2.com",
			Descriptors: []models.RateLimitDescriptor{
				{
					Model: gorm.Model{ID: 2},
					Key:   "foo",
					Descriptors: []models.RateLimitDescriptor{
						{
							Model:           gorm.Model{ID: 4},
							Key:             "foo",
							Value:           "bar",
							Unit:            2,
							RequestsPerUnit: 10,
						},
					},
				},
			},
		},
	}
	db.Create(rateLimits)

	store := storage.NewRateLimitStore(db)

	resources := NewResources(store)
	snap := resources.GenerateSnapshot()

	assert.NotEmpty(t, snap.Resources)
}
