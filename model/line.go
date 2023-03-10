package model

import (
	"time"

	"gorm.io/gorm"
)

// 路線
type Line struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique"`
	Elements  []Element
	Relations []Relation
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// 要素
type Element struct {
	ID                uint `gorm:"primaryKey"`
	LineID            uint
	Longitude         float32 // 緯度
	Latitude          float32 // 経度
	OperationalPoints []OpelationalPoint
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

// 繋がり
type Relation struct {
	ID        uint `gorm:"primaryKey"`
	LineID    uint
	ElementA  uint
	ElementB  uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// バス停
type OpelationalPoint struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	ElementID uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
