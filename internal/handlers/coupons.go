package handlers

import (
	"coupon-system/internal/models"
	"coupon-system/internal/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type CreateCouponRequest struct {
	CouponCode            string  `json:"coupon_code" binding:"required"`
	ExpiryDate            string  `json:"expiry_date" binding:"required"` // string format (ISO8601)
	UsageType             string  `json:"usage_type" binding:"required,oneof=one_time multi_use time_based"`
	ApplicableMedicineIDs string  `json:"applicable_medicine_ids"` // comma-separated medicine IDs
	ApplicableCategories  string  `json:"applicable_categories"`   // comma-separated category IDs (uint)
	MinOrderValue         float64 `json:"min_order_value"`
	ValidTimeWindowStart  *string `json:"valid_time_window_start"` // optional ISO8601 string
	ValidTimeWindowEnd    *string `json:"valid_time_window_end"`   // optional ISO8601 string
	TermsAndConditions    string  `json:"terms_and_conditions"`
	DiscountType          string  `json:"discount_type" binding:"required,oneof=flat percent"`
	DiscountValue         float64 `json:"discount_value" binding:"required"`
	MaxUsagePerUser       int     `json:"max_usage_per_user"`
}

func parseTimePtr(tstr *string) *time.Time {
	if tstr == nil || *tstr == "" {
		return nil
	}
	t, err := time.Parse(time.RFC3339, *tstr)
	if err != nil {
		return nil
	}
	return &t
}

func CreateCouponHandler(c *gin.Context) {
	var req CreateCouponRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse times
	expiryDate, err := time.Parse(time.RFC3339, req.ExpiryDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expiry_date format. Use RFC3339 format"})
		return
	}
	start := parseTimePtr(req.ValidTimeWindowStart)
	end := parseTimePtr(req.ValidTimeWindowEnd)

	coupon := models.Coupon{
		CouponCode:           req.CouponCode,
		ExpiryDate:           expiryDate,
		UsageType:            models.UsageType(req.UsageType),
		MinOrderValue:        req.MinOrderValue,
		ValidTimeWindowStart: start,
		ValidTimeWindowEnd:   end,
		TermsAndConditions:   req.TermsAndConditions,
		DiscountType:         models.DiscountType(req.DiscountType),
		DiscountValue:        req.DiscountValue,
		MaxUsagePerUser:      req.MaxUsagePerUser,
	}

	tx := repository.DB.Begin()
	if err := tx.Create(&coupon).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create coupon"})
		return
	}

	// Parse medicine IDs and link
	if req.ApplicableMedicineIDs != "" {
		medIDs := strings.Split(req.ApplicableMedicineIDs, ",")
		for _, medID := range medIDs {
			medID = strings.TrimSpace(medID)
			cm := models.CouponMedicine{
				CouponID:   coupon.ID,
				MedicineID: medID,
			}
			if err := tx.Create(&cm).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to link medicines"})
				return
			}
		}
	}

	// Parse category IDs and link
	if req.ApplicableCategories != "" {
		catIDs := strings.Split(req.ApplicableCategories, ",")
		for _, catID := range catIDs {
			catID = strings.TrimSpace(catID)
			idUint, err := strconv.ParseUint(catID, 10, 64)
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID: " + catID})
				return
			}
			cc := models.CouponCategory{
				CouponID:   coupon.ID,
				CategoryID: uint(idUint),
			}
			if err := tx.Create(&cc).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to link categories"})
				return
			}
		}
	}

	tx.Commit()
	c.JSON(http.StatusCreated, gin.H{"message": "Coupon created successfully", "coupon": coupon})
}
