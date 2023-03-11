package repository

import (
	"log"
	"rosenzu/database"
	"rosenzu/database/model"
)

func FindLine(name string) model.Line {
	var line model.Line
	if err := database.Db.Where("line_name = ?", name).First(&line).Error; err != nil {
		log.Fatal("Not Found Line")
	}
	return line
}
