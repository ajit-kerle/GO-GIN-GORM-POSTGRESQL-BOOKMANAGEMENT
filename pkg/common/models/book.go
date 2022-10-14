package models



type Book struct {
	Title       string `gorm:"not null" json:"title" `
	// AuthorI      string `json:"author"`
	Description string `gorm:"not null" json:"description"`

	// author id referencing from author models
	AuthorID    int64  `gorm:"not null" json:"author_id" binding:"required" `
	Author      *Author `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"author_name,omitempty"`

	// publisher id referencing from publisher models
	PublisherID    int64  `gorm:"not null" json:"publisher_id" binding:"required" `
	Publisher      *Publisher `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"publisher_name,omitempty"`
}


// postman data sending format
// {
// 	"title":"good vibe",
// 	"description":"best book ever",
// 	"author_id":4,
// 	"publisher_id":123456789
// }
