package entity

import "time"

type Article struct {
	Id      int       `gorm:"id"`
	Author  string    `gorm:"author"`
	Title   string    `gorm:"title"`
	Body    string    `gorm:"body"`
	Created time.Time `gorm:"created"`
}

func (Article) TableName() string {
	return "article"
  }