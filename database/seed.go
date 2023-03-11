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
		initLines()
		initElements()
		initRelations()
		initOperationalPoints()
	}
}

func initOperationalPoints() {
	ops := []model.OperationalPoint{}
	rows := loadCSV("./csv/elements.csv")
	for _, row := range rows {
		if row[0] != "" {
			element_id, _ := strconv.Atoi(row[0])
			op := model.OperationalPoint{
				Name:      row[1],
				ElementID: element_id,
			}
			ops = append(ops, op)
		}
	}
	Db.Save(ops)
}

func initRelations() {
	relations := []model.Relation{}
	rows := loadCSV("./csv/elements.csv")
	for _, row := range rows {
		if row[0] != "" {
			a, _ := strconv.Atoi(row[0])
			b, _ := strconv.Atoi(row[1])
			relation := model.Relation{
				ElementA: a,
				ElementB: b,
			}
			relations = append(relations, relation)
		}
	}
	Db.Save(relations)
}

func initElements() {
	elements := []model.Element{}
	rows := loadCSV("./csv/elements.csv")
	for i, row := range rows {
		if row[0] != "" {
			if i == 0 {
				elements = append(elements, newElement(row[0]))
			}
			// 座標を追加
			coordinate := newCoordinate(row[1], row[2])
			elements[len(elements)-1].Coordinates = append(elements[len(elements)-1].Coordinates, coordinate)

			if row[4] != "" {
				elements = append(elements, newElement(row[0]))
			}
		}
	}
	Db.Save(elements)
}

func newElement(istr string) model.Element {
	id, _ := strconv.Atoi(istr)
	element := model.Element{
		ID: id,
	}
	return element
}

func newCoordinate(long string, lat string) model.Coordinate {
	longitude, _ := strconv.ParseFloat(long, 32)
	latitude, _ := strconv.ParseFloat(lat, 32)
	coordinate := model.Coordinate{
		Longitude: float32(longitude),
		Latitude:  float32(latitude),
	}
	return coordinate
}

func initLines() {
	lines := []model.Line{}
	rows := loadCSV("./csv/lines.csv")
	for _, row := range rows {
		if row[0] != "" {
			id, _ := strconv.Atoi(row[0])
			line := model.Line{
				ID:   id,
				Name: row[1],
			}
			lines = append(lines, line)
		}
	}
	Db.Save(lines)
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
