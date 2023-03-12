package model

import (
	"time"

	"gorm.io/gorm"
)

// 路線
type Line struct {
	ID                int                `gorm:"primaryKey"`
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
	ID          int `gorm:"primaryKey"`
	Coordinates []Coordinate
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Coordinate struct {
	ID        int `gorm:"primaryKey"`
	ElementID int
	Longitude float32 // 緯度
	Latitude  float32 // 経度
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// 繋がり
type Relation struct {
	ID        int `gorm:"primaryKey"`
	ElementA  int
	ElementB  int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// バス停
type OperationalPoint struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	ElementID int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
