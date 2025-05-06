package documents

import (
	"datcomtd/backend/initializers"
	"datcomtd/backend/models"
)

var documents = []models.Document{
    {
        Title:          "Meeting Minutes Q1",
        Description:    "Quarterly team meeting notes",
        Category:       "meetings",
        CreatedBy:      "admin",
        LastUpdatedBy:  "admin",
        Filename:       "meeting_minutes_q1.pdf",
        Source:         "/storage/meeting_minutes_q1.pdf",
    },
    {
        Title:          "Project Guidelines",
        Description:    "Team project guidelines and requirements",
        Category:       "guidelines",
        CreatedBy:      "john_doe",
        LastUpdatedBy:  "john_doe",
        Filename:       "project_guidelines.pdf",
        Source:         "/storage/project_guidelines.pdf",
    },
}

func LoadDocuments() {
    for _, doc := range documents {
		result := initializers.DB.Model(&models.Document{}).Create(&doc)
        if result.Error != nil {
            panic(result.Error)
        }
    }
}