package entity

import (
	"time"
)

type Recipe struct {
	ID       uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	UserID   uint   `json:"userid"`
	Title    string `json:"title" gorm:"type:varchar(255);not null"`
	Duration string `json:"duration" gorm:"type:varchar(255);not null"`
	Servings int    `json:"servings"`
	imageUrl string `json:"imageUrl" gorm:"type:varchar(255);not null"`
	Comments []Comment
}

type User struct {
	ID       uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	FullName string `json:"fullname" gorm:"type:varchar(255);not null"`
	Recipes  []Recipe
}

type Comment struct {
	ID        uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	UserID    uint   `json:"userid"`
	UserName  string `json:"username" gorm:"type:varchar(255);not null`
	RecipeID  uint   `json:"recipeid" gorm:"type:varchar(255);not null`
	Message   string `json:"message" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
}
