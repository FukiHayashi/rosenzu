package database

import (
	"encoding/csv"
	"log"
	"os"
	"rosenzu/database/model"
	"strconv"
)

func DbSeed() {
	var line model.Line
	if err := Db.First(&line).Error; err != nil {
		// 各テーブルを生成
		initLines()
		initElements()
		initRelations()
		initOperationalPoints()
		// 各アソシエーションテーブルを生成
		initLineElements()
		initLineRelations()
		initLineOperationalPoints()
	}
}

// Line-OperationalPointsのアソシエーションを生成
func initLineOperationalPoints() {
	var lines []model.Line
	Db.Find(&lines)
	rows := loadCSV("./database/csv/operationalpoints.csv")
	for _, row := range rows {
		if row[0] != "" {
			var op model.OperationalPoint
			if err := Db.Where("name = ?", row[1]).First(&op).Error; err == nil {
				for _, line := range lines {
					for j := 2; j < len(row); j++ {
						if row[j] != "" {
							line_id, _ := strconv.ParseUint(row[j], 10, 32)
							if uint(line_id) == line.ID {
								Db.Model(&line).Association("OperationalPoints").Append(&model.OperationalPoint{ID: op.ID})
							}
						}
					}
				}
			}
		}
	}
	Db.Save(&lines)
}

// Line-Relationsのアソシエーションを生成
func initLineRelations() {
	var lines []model.Line
	Db.Find(&lines)
	rows := loadCSV("./database/csv/relations.csv")
	for i, row := range rows {
		if row[0] != "" {
			relation_id := i + 1 //, _ := strconv.Atoi(row[0])
			var relation model.Relation
			if err := Db.Where("id = ?", relation_id).First(&relation).Error; err == nil {
				for _, line := range lines {
					for j := 2; j < len(row); j++ {
						if row[j] != "" {
							line_id, _ := strconv.ParseUint(row[j], 10, 32)
							if uint(line_id) == line.ID {
								Db.Model(&line).Association("Relations").Append(&relation)
							}
						}
					}
				}
			}
		}
	}
	Db.Save(&lines)
}

// Line-Elementsのアソシエーションを生成
func initLineElements() {
	var lines []model.Line
	Db.Find(&lines)
	rows := loadCSV("./database/csv/elements.csv")
	for _, row := range rows {
		if row[0] != "" {
			element_id, _ := strconv.Atoi(row[0])
			var element model.Element
			if err := Db.Where("id = ?", element_id).First(&element).Error; err == nil {
				for _, line := range lines {
					for j := 4; j < len(row); j++ {
						if row[j] != "" {
							line_id, _ := strconv.ParseUint(row[j], 10, 32)
							if uint(line_id) == line.ID {
								Db.Model(&line).Association("Elements").Append(&element)
							}
						}
					}
				}
			}
		}
	}
	Db.Save(&lines)
}

// OperationalPointのシードを読み込み
func initOperationalPoints() {
	ops := []model.OperationalPoint{}
	rows := loadCSV("./database/csv/operationalpoints.csv")
	for _, row := range rows {
		if row[0] != "" {
			element_id, _ := strconv.ParseUint(row[0], 10, 32)
			op := model.OperationalPoint{
				Name:      row[1],
				ElementID: uint(element_id),
			}
			ops = append(ops, op)
		}
	}
	Db.Save(&ops)
}

// Relationのシードを読み込み
func initRelations() {
	relations := []model.Relation{}
	rows := loadCSV("./database/csv/relations.csv")
	for _, row := range rows {
		if row[0] != "" {
			a, _ := strconv.ParseUint(row[0], 10, 32)
			b, _ := strconv.ParseUint(row[1], 10, 32)
			relation := model.Relation{
				ElementA: uint(a),
				ElementB: uint(b),
			}
			relations = append(relations, relation)
		}
	}
	Db.Save(&relations)
}

// Elementのシードを読み込み
func initElements() {
	elements := []model.Element{}
	rows := loadCSV("./database/csv/elements.csv")
	for _, row := range rows {
		if row[0] != "" {
			if row[4] != "" {
				elements = append(elements, newElement(row[0]))
			}
			// 座標を追加
			coordinate := newCoordinate(row[1], row[2])
			elements[len(elements)-1].Coordinates = append(elements[len(elements)-1].Coordinates, coordinate)
		}
	}
	Db.Save(&elements)
}

// Elementを生成
func newElement(istr string) model.Element {
	id, _ := strconv.ParseUint(istr, 10, 32)
	element := model.Element{
		ID: uint(id),
	}
	return element
}

// Coordinateを生成
func newCoordinate(long string, lat string) model.Coordinate {
	longitude, _ := strconv.ParseFloat(long, 32)
	latitude, _ := strconv.ParseFloat(lat, 32)
	coordinate := model.Coordinate{
		Longitude: float32(longitude),
		Latitude:  float32(latitude),
	}
	return coordinate
}

// Lineのシードを読み込み
func initLines() {
	lines := []model.Line{}
	rows := loadCSV("./database/csv/lines.csv")
	for _, row := range rows {
		if row[0] != "" {
			id, _ := strconv.ParseUint(row[0], 10, 32)
			line := model.Line{
				ID:   uint(id),
				Name: row[1],
			}
			lines = append(lines, line)
		}
	}
	Db.Save(&lines)
}

// csv読み込み
func loadCSV(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return rows
}
