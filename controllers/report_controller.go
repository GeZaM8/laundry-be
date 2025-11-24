package controllers

import (
	"net/http"
	"time"

	"github.com/GeZaM8/laundry-be/config"
	"github.com/GeZaM8/laundry-be/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReportController struct{
	DB *gorm.DB
}

func NewReportController(db *gorm.DB) *ReportController {
	return &ReportController{DB: db}
}

func (rc *ReportController) RevenueAll(c *gin.Context) {

	var total float64

	err := rc.DB.
		Table("orders").
		Select("SUM(total_price)").
		Scan(&total).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Revenue Report",
		Data: gin.H{
			"total_revenue": total,
		},
	})
}

func (rc *ReportController) Revenue(c *gin.Context) {
	start := c.Query("start")
	end := c.Query("end")

	if start == "" || end == "" {
		c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: "start dan end wajib diisi (yyyy-mm-dd)",
		})
		return
	}

	var total float64

	err := rc.DB.
		Table("orders").
		Select("SUM(total_price)").
		Where("DATE(created_at) BETWEEN ? AND ?", start, end).
		Scan(&total).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Revenue Report",
		Data: gin.H{
			"start":        start,
			"end":          end,
			"total_revenue": total,
		},
	})
}

func (rc *ReportController) Daily(c *gin.Context) {
	date := c.Query("date")

	if date == "" {
		date = time.Now().Format("2006-01-02")
	}

	var orders []model.Order

	err := rc.DB.
		Preload("Customer").
		Preload("Items.Category").
		Where("DATE(created_at) = ?", date).
		Find(&orders).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Daily Orders",
		Data: gin.H{
			"date":   date,
			"orders": orders,
		},
	})
}
func (rc *ReportController) ItemsAll(c *gin.Context) {
	type ItemReport struct {
		CategoryName string  `json:"category_name"`
		TotalQty     int     `json:"total_qty"`
		TotalWeight  float64 `json:"total_weight"`
		TotalSales   float64 `json:"total_sales"`
	}

	var report []ItemReport

	err := rc.DB.Raw(`
		SELECT 
			categories.name AS category_name,
			SUM(order_items.qty) AS total_qty,
			SUM(order_items.weight_kg) AS total_weight,
			SUM(order_items.price) AS total_sales
		FROM order_items
		JOIN orders ON order_items.order_id = orders.id
		JOIN categories ON order_items.category_id = categories.id
		GROUP BY categories.name
		ORDER BY categories.name ASC
	`).Scan(&report).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Items Report",
		Data: gin.H{
			"items":  report,
		},
	})
}

func (rc *ReportController) Items(c *gin.Context) {
	start := c.Query("start")
	end := c.Query("end")

	if start == "" || end == "" {
		c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: "start dan end wajib diisi (yyyy-mm-dd)",
		})
		return
	}

	type ItemReport struct {
		CategoryName string  `json:"category_name"`
		TotalQty     int     `json:"total_qty"`
		TotalWeight  float64 `json:"total_weight"`
		TotalSales   float64 `json:"total_sales"`
	}

	var report []ItemReport

	err := rc.DB.Raw(`
		SELECT 
			categories.name AS category_name,
			SUM(order_items.qty) AS total_qty,
			SUM(order_items.weight_kg) AS total_weight,
			SUM(order_items.price) AS total_sales
		FROM order_items
		JOIN orders ON order_items.order_id = orders.id
		JOIN categories ON order_items.category_id = categories.id
		WHERE DATE(orders.created_at) BETWEEN ? AND ?
		GROUP BY categories.name
		ORDER BY categories.name ASC
	`, start, end).Scan(&report).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Items Report",
		Data: gin.H{
			"start":  start,
			"end":    end,
			"items":  report,
		},
	})
}

type MonthlyChart struct {
	Month int     `json:"month"`
	Total float64 `json:"total"`
	Count int     `json:"count"`
}

func (ReportController) ChartMonthly(c *gin.Context) {
	year := c.Query("year")

	// Kalau tahun tidak diberikan → pakai tahun berjalan
	if year == "" {
		year = time.Now().Format("2006")
	}


	var raw []MonthlyChart

	err := config.DB.Raw(`
		SELECT 
			MONTH(created_at) AS month,
			SUM(total_price) AS total,
			COUNT(*) AS count
		FROM orders
		WHERE YEAR(created_at) = ?
		GROUP BY MONTH(created_at)
		ORDER BY MONTH(created_at)
	`, year).Scan(&raw).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	// Lengkapi 12 bulan
	final := fillMissingMonths(raw)

	c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Monthly Chart",
		Data:    final,
	})
}
type MonthlyFinal struct {
	MonthName string  `json:"month"`
	Total     float64 `json:"total"`
	Count     int     `json:"count"`
}

func fillMissingMonths(raw []MonthlyChart) []MonthlyFinal {
	monthNames := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

	// mapping data ke map
	m := map[int]MonthlyChart{}
	for _, r := range raw {
		m[r.Month] = r
	}

	// hasil final
	final := []MonthlyFinal{}

	for i := 1; i <= 12; i++ {
		if val, ok := m[i]; ok {
			final = append(final, MonthlyFinal{
				MonthName: monthNames[i-1],
				Total:     val.Total,
				Count:     val.Count,
			})
		} else {
			final = append(final, MonthlyFinal{
				MonthName: monthNames[i-1],
				Total:     0,
				Count:     0,
			})
		}
	}

	return final
}
