package cast

import (
	"rosenzu/database/model"
	"rosenzu/gen/rosenzu"
)

// structのキャスト

// goaのLine型にキャスト
func CastedLine(line model.Line) *rosenzu.Line {
	casted_line := rosenzu.Line{
		Name:      &line.Name,
		Elements:  CastedElements(line.Elements),
		Relations: CastedRelations(line.Relations),
	}
	return &casted_line
}

// goaのOpoerationalpointCollection型にキャスト
func CastedOperationalpoints(ops []model.OperationalPoint) rosenzu.OperationalpointCollection {
	casted_ops := rosenzu.OperationalpointCollection{}
	for _, e := range ops {
		casted_ops = append(casted_ops, CastedOperationalpoint(e))
	}
	return casted_ops
}

// goaのOperationalpoint型にキャスト
func CastedOperationalpoint(op model.OperationalPoint) *rosenzu.Operationalpoint {
	casted_op := rosenzu.Operationalpoint{
		Name:      &op.Name,
		ElementID: &op.ElementID,
	}
	return &casted_op
}

// goaのRelationCollection型にキャスト
func CastedRelations(relations []model.Relation) rosenzu.RelationCollection {
	casted_relations := rosenzu.RelationCollection{}
	for _, e := range relations {
		casted_relations = append(casted_relations, CastedRelation(e))
	}
	return casted_relations
}

// goaのRelation型にキャスト
func CastedRelation(relation model.Relation) *rosenzu.Relation {
	casted_relation := rosenzu.Relation{
		ElementA: &relation.ElementA,
		ElementB: &relation.ElementB,
	}
	return &casted_relation
}

// goaのElementCollection型にキャスト
func CastedElements(elements []model.Element) rosenzu.ElementCollection {
	casted_elements := rosenzu.ElementCollection{}
	for _, e := range elements {
		casted_elements = append(casted_elements, CastedElement(e))
	}
	return casted_elements
}

// goaのElement型にキャスト
func CastedElement(element model.Element) *rosenzu.Element {
	casted_element := rosenzu.Element{
		ID: &element.ID,
	}
	return &casted_element
}

// goaのCoordinateCollection型にキャスト
func CastedCoordinates(coordinates []model.Coordinate) rosenzu.CoordinateCollection {
	casted_coordinates := rosenzu.CoordinateCollection{}
	for _, e := range coordinates {
		casted_coordinates = append(casted_coordinates, CastedCoordinate(e))
	}
	return casted_coordinates
}

// goaのCoordinate型にキャスト
func CastedCoordinate(coordinate model.Coordinate) *rosenzu.Coordinate {
	casted_coordinate := rosenzu.Coordinate{
		Latitude:  &coordinate.Latitude,
		Longitude: &coordinate.Longitude,
	}
	return &casted_coordinate
}
