package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// RateLimit Rate limit configuration for a single domain.
type RateLimit struct {
	ID uint `gorm:"primarykey"`

	Name        string
	Domain      string                `gorm:"unique"`
	Descriptors []RateLimitDescriptor `gorm:"foreignKey:RateLimitId"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// TODO: support sub descriptor
// RateLimitDescriptor Rate limit configuration descriptor.
type RateLimitDescriptor struct {
	ID uint `gorm:"primarykey"`

	// rate limit id
	RateLimitId uint

	Key   string `gorm:"uniqueIndex:idx_key_value"`
	Value string `gorm:"uniqueIndex:idx_key_value;default:null"`

	Policy RateLimitPolicy `gorm:"type:json"`

	ShadowMode bool `gorm:"default:false"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// RateLimitPolicy a policy of RateLimitDescriptor
type RateLimitPolicy struct {
	Unit int32 `json:"unit"`
	// The time unit: SECOND, MINUTE, HOUR, DAY
	RequestsPerUnit int32 `json:"requests_per_unit"`
}

func (p *RateLimitPolicy) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, p)
	}
	return nil
}

func (p RateLimitPolicy) Value() (driver.Value, error) {
	return json.Marshal(p)
}
