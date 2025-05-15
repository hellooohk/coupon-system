package models

import "time"

type UsageType string
type DiscountType string

const (
	OneTime   UsageType = "one_time"
	MultiUse  UsageType = "multi_use"
	TimeBased UsageType = "time_based"

	Flat    DiscountType = "flat"
	Percent DiscountType = "percent"
)

type Category struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"uniqueIndex"`
}

type Medicine struct {
	ID         string `gorm:"primaryKey"` // e.g., "med_123"
	Name       string
	BatchNo    string
	ExpiryDate time.Time
	Quantity   int
	CategoryID uint
	Category   Category `gorm:"foreignKey:CategoryID"`
}

type Coupon struct {
	ID                   uint   `gorm:"primaryKey"`
	CouponCode           string `gorm:"uniqueIndex"`
	ExpiryDate           time.Time
	UsageType            UsageType
	MinOrderValue        float64
	ValidTimeWindowStart *time.Time
	ValidTimeWindowEnd   *time.Time
	TermsAndConditions   string
	DiscountType         DiscountType
	DiscountValue        float64
	MaxUsagePerUser      int
	CreatedAt            time.Time
	UpdatedAt            time.Time

	// Relations
	CouponMedicines  []CouponMedicine `gorm:"foreignKey:CouponID;constraint:OnDelete:CASCADE"`
	CouponCategories []CouponCategory `gorm:"foreignKey:CouponID;constraint:OnDelete:CASCADE"`
}

type CouponMedicine struct {
	ID         uint `gorm:"primaryKey"`
	CouponID   uint
	MedicineID string
	Medicine   Medicine `gorm:"foreignKey:MedicineID;references:ID"`
}

type CouponCategory struct {
	ID         uint `gorm:"primaryKey"`
	CouponID   uint
	CategoryID uint
	Category   Category `gorm:"foreignKey:CategoryID;references:ID"`
}
