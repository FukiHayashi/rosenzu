package repository

import (
	"log"
	"rosenzu/database"
	"rosenzu/database/model"

	"gorm.io/gorm/clause"
)

// 路線名の検索
func FindLine(name string) model.Line {
	var line model.Line
	if err := database.Db.Preload(clause.Associations).Preload("Elements.Coordinates").Where("name = ?", name).First(&line).Error; err != nil {
		log.Fatal("Not Found Line")
	}
	return line
}
