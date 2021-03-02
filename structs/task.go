package structs

import "time"

type Task struct {
	TaskId     int
	Time       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"ordered_at"`
	Project_Id int
	Title_Id   int
	Status_id  int
}
