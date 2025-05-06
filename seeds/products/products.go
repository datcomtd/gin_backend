package products

import (
	"datcomtd/backend/initializers"
	"datcomtd/backend/models"
    "github.com/lib/pq"
)

var products = []models.Product{
    {
        Title:       "Team T-Shirt",
        Description: "Official team merchandise",
        Category:    "merchandise",
        Price:       29.99,
        Count:       50,
        Photos:      pq.StringArray{"mug1.jpg", "mug2.jpg"},
        CreatedBy:   "admin",
        LastUpdatedBy: "admin",
    },
    {
        Title:       "Coffee Mug",
        Description: "Team branded coffee mug",
        Category:    "merchandise",
        Price:       14.99,
        Count:       100,
        Photos:      pq.StringArray{"mug1.jpg", "mug2.jpg"},
        CreatedBy:   "john_doe",
        LastUpdatedBy: "john_doe",
    },
}

func LoadProducts() {
    for _, product := range products {
        result := initializers.DB.Model(&models.Product{}).Create(&product)
        if result.Error != nil {
            panic(result.Error)
        }
    }
}