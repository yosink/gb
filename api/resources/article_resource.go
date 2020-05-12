package resources

import (
	"blog/comm"
)

type ArticleResource struct {
	ID              uint              `json:"id"`
	Slug            string            `json:"slug"`
	Title           string            `json:"title"`
	SubTitle        string            `json:"sub_title"`
	Content         string            `json:"content"`
	PageImage       string            `json:"page_image"`
	MetaDescription string            `json:"meta_description"`
	Recommend       *uint8            `json:"recommend"`
	Sort            int               `json:"sort"`
	ViewCount       int               `json:"view_count"`
	CreatedAt       comm.XTime        `json:"created_at"`
	Category        *CategoryResource `json:"category"`
	User            *UserResource     `json:"user"`
}
