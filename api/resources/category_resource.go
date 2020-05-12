package resources

import (
	"blog/comm"
)

type CategoryResource struct {
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	ImageURL    string     `json:"image_url"`
	CreatedAt   comm.XTime `json:"created_at"`
}
