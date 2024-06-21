package models

import "gorm.io/gorm"

// RateLimit Rate limit configuration for a single domain.
type RateLimit struct {
	gorm.Model
	Name        string
	Domain      string
	Descriptors []RateLimitDescriptor `gorm:"foreignKey:RateLimitId"`
}

// RateLimitDescriptor Rate limit configuration descriptor.
type RateLimitDescriptor struct {
	gorm.Model
	RateLimitId uint
	Key         string
	Value       string `gorm:"default:null"`
	// The time unit: SECOND, MINUTE, HOUR, DAY
	Unit            int32 `gorm:"default:0"`
	RequestsPerUnit int32 `gorm:"default:null"`
	ShadowMode      bool
	Descriptors     []RateLimitDescriptor `gorm:"foreignKey:ParentDescriptorID"`

	ParentDescriptorID *uint
}
