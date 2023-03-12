package model

import (
	"time"

	"gorm.io/gorm"
)

// 路線
type Line struct {
	ID                uint               `gorm:"primaryKey"`
	Name              string             `gorm:"unique"`
	Elements          []Element          `gorm:"many2many:line_elements"`
	Relations         []Relation         `gorm:"many2many:line_relations"`
	OperationalPoints []OperationalPoint `gorm:"many2many:line_operationalpoints"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

// 要素
type Element struct {
	ID          uint `gorm:"primaryKey"`
	Coordinates []Coordinate
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Coordinate struct {
	ID        uint `gorm:"primaryKey"`
	ElementID uint
	Longitude float32 // 緯度
	Latitude  float32 // 経度
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// 繋がり
type Relation struct {
	ID        uint `gorm:"primaryKey"`
	ElementA  uint
	ElementB  uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// バス停
type OperationalPoint struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	ElementID uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
