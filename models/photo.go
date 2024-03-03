// models/photo.go
package models

import (
    "gorm.io/gorm"
    "time"
)

type Photo struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Title     string    `json:"title"`
    Caption   string    `json:"caption"`
    PhotoURL  string    `json:"photoUrl" binding:"required"`
    UserID    uint      `json:"userId"`
    User      User      `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}
