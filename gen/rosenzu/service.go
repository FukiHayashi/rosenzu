// Code generated by goa v3.11.1, DO NOT EDIT.
//
// rosenzu service
//
// Command:
// $ goa gen rosenzu/design

package rosenzu

import (
	"context"
	rosenzuviews "rosenzu/gen/rosenzu/views"

	goa "goa.design/goa/v3/pkg"
)

// 路線データ
type Service interface {
	// Find implements find.
	Find(context.Context, *FindPayload) (res *Line, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "rosenzu"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"find"}

type Coordinate struct {
	// 緯度
	Longitude *float32
	// 経度
	Latitude *float32
}

type CoordinateCollection []*Coordinate

type Element struct {
	// id
	ID          *int
	Coordinates CoordinateCollection
}

type ElementCollection []*Element

// FindPayload is the payload type of the rosenzu service find method.
type FindPayload struct {
	// 路線名
	Name string
}

// Line is the result type of the rosenzu service find method.
type Line struct {
	// 路線名
	Name              *string
	Elements          ElementCollection
	Relations         RelationCollection
	OperationalPoints OperationalpointCollection
}

type Operationalpoint struct {
	// バス停名
	Name *string
	// elementID
	ElementID *int
}

type OperationalpointCollection []*Operationalpoint

type Relation struct {
	// elementA
	ElementA *int
	// elementB
	ElementB *int
}

type RelationCollection []*Relation

// MakeUnauthorized builds a goa.ServiceError from an error.
func MakeUnauthorized(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "Unauthorized", false, false, false)
}

// MakeBadRequest builds a goa.ServiceError from an error.
func MakeBadRequest(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "BadRequest", false, false, false)
}

// NewLine initializes result type Line from viewed result type Line.
func NewLine(vres *rosenzuviews.Line) *Line {
	return newLine(vres.Projected)
}

// NewViewedLine initializes viewed result type Line from result type Line
// using the given view.
func NewViewedLine(res *Line, view string) *rosenzuviews.Line {
	p := newLineView(res)
	return &rosenzuviews.Line{Projected: p, View: "default"}
}

// newLine converts projected type Line to service type Line.
func newLine(vres *rosenzuviews.LineView) *Line {
	res := &Line{
		Name: vres.Name,
	}
	if vres.Relations != nil {
		res.Relations = make([]*Relation, len(vres.Relations))
		for i, val := range vres.Relations {
			res.Relations[i] = transformRosenzuviewsRelationViewToRelation(val)
		}
	}
	if vres.Elements != nil {
		res.Elements = newElementCollection(vres.Elements)
	}
	if vres.OperationalPoints != nil {
		res.OperationalPoints = newOperationalpointCollection(vres.OperationalPoints)
	}
	return res
}

// newLineView projects result type Line to projected type LineView using the
// "default" view.
func newLineView(res *Line) *rosenzuviews.LineView {
	vres := &rosenzuviews.LineView{
		Name: res.Name,
	}
	if res.Relations != nil {
		vres.Relations = make([]*rosenzuviews.RelationView, len(res.Relations))
		for i, val := range res.Relations {
			vres.Relations[i] = transformRelationToRosenzuviewsRelationView(val)
		}
	}
	if res.Elements != nil {
		vres.Elements = newElementCollectionView(res.Elements)
	}
	if res.OperationalPoints != nil {
		vres.OperationalPoints = newOperationalpointCollectionView(res.OperationalPoints)
	}
	return vres
}

// newElementCollection converts projected type ElementCollection to service
// type ElementCollection.
func newElementCollection(vres rosenzuviews.ElementCollectionView) ElementCollection {
	res := make(ElementCollection, len(vres))
	for i, n := range vres {
		res[i] = newElement(n)
	}
	return res
}

// newElementCollectionView projects result type ElementCollection to projected
// type ElementCollectionView using the "default" view.
func newElementCollectionView(res ElementCollection) rosenzuviews.ElementCollectionView {
	vres := make(rosenzuviews.ElementCollectionView, len(res))
	for i, n := range res {
		vres[i] = newElementView(n)
	}
	return vres
}

// newElement converts projected type Element to service type Element.
func newElement(vres *rosenzuviews.ElementView) *Element {
	res := &Element{
		ID: vres.ID,
	}
	if vres.Coordinates != nil {
		res.Coordinates = newCoordinateCollection(vres.Coordinates)
	}
	return res
}

// newElementView projects result type Element to projected type ElementView
// using the "default" view.
func newElementView(res *Element) *rosenzuviews.ElementView {
	vres := &rosenzuviews.ElementView{
		ID: res.ID,
	}
	if res.Coordinates != nil {
		vres.Coordinates = newCoordinateCollectionView(res.Coordinates)
	}
	return vres
}

// newCoordinateCollectionDafault converts projected type CoordinateCollection
// to service type CoordinateCollection.
func newCoordinateCollectionDafault(vres rosenzuviews.CoordinateCollectionView) CoordinateCollection {
	res := make(CoordinateCollection, len(vres))
	for i, n := range vres {
		res[i] = newCoordinateDafault(n)
	}
	return res
}

// newCoordinateCollection converts projected type CoordinateCollection to
// service type CoordinateCollection.
func newCoordinateCollection(vres rosenzuviews.CoordinateCollectionView) CoordinateCollection {
	res := make(CoordinateCollection, len(vres))
	for i, n := range vres {
		res[i] = newCoordinate(n)
	}
	return res
}

// newCoordinateCollectionViewDafault projects result type CoordinateCollection
// to projected type CoordinateCollectionView using the "dafault" view.
func newCoordinateCollectionViewDafault(res CoordinateCollection) rosenzuviews.CoordinateCollectionView {
	vres := make(rosenzuviews.CoordinateCollectionView, len(res))
	for i, n := range res {
		vres[i] = newCoordinateViewDafault(n)
	}
	return vres
}

// newCoordinateCollectionView projects result type CoordinateCollection to
// projected type CoordinateCollectionView using the "default" view.
func newCoordinateCollectionView(res CoordinateCollection) rosenzuviews.CoordinateCollectionView {
	vres := make(rosenzuviews.CoordinateCollectionView, len(res))
	for i, n := range res {
		vres[i] = newCoordinateView(n)
	}
	return vres
}

// newCoordinateDafault converts projected type Coordinate to service type
// Coordinate.
func newCoordinateDafault(vres *rosenzuviews.CoordinateView) *Coordinate {
	res := &Coordinate{
		Longitude: vres.Longitude,
		Latitude:  vres.Latitude,
	}
	return res
}

// newCoordinate converts projected type Coordinate to service type Coordinate.
func newCoordinate(vres *rosenzuviews.CoordinateView) *Coordinate {
	res := &Coordinate{
		Longitude: vres.Longitude,
		Latitude:  vres.Latitude,
	}
	return res
}

// newCoordinateViewDafault projects result type Coordinate to projected type
// CoordinateView using the "dafault" view.
func newCoordinateViewDafault(res *Coordinate) *rosenzuviews.CoordinateView {
	vres := &rosenzuviews.CoordinateView{
		Longitude: res.Longitude,
		Latitude:  res.Latitude,
	}
	return vres
}

// newCoordinateView projects result type Coordinate to projected type
// CoordinateView using the "default" view.
func newCoordinateView(res *Coordinate) *rosenzuviews.CoordinateView {
	vres := &rosenzuviews.CoordinateView{
		Longitude: res.Longitude,
		Latitude:  res.Latitude,
	}
	return vres
}

// newRelationCollectionDafault converts projected type RelationCollection to
// service type RelationCollection.
func newRelationCollectionDafault(vres rosenzuviews.RelationCollectionView) RelationCollection {
	res := make(RelationCollection, len(vres))
	for i, n := range vres {
		res[i] = newRelationDafault(n)
	}
	return res
}

// newRelationCollection converts projected type RelationCollection to service
// type RelationCollection.
func newRelationCollection(vres rosenzuviews.RelationCollectionView) RelationCollection {
	res := make(RelationCollection, len(vres))
	for i, n := range vres {
		res[i] = newRelation(n)
	}
	return res
}

// newRelationCollectionViewDafault projects result type RelationCollection to
// projected type RelationCollectionView using the "dafault" view.
func newRelationCollectionViewDafault(res RelationCollection) rosenzuviews.RelationCollectionView {
	vres := make(rosenzuviews.RelationCollectionView, len(res))
	for i, n := range res {
		vres[i] = newRelationViewDafault(n)
	}
	return vres
}

// newRelationCollectionView projects result type RelationCollection to
// projected type RelationCollectionView using the "default" view.
func newRelationCollectionView(res RelationCollection) rosenzuviews.RelationCollectionView {
	vres := make(rosenzuviews.RelationCollectionView, len(res))
	for i, n := range res {
		vres[i] = newRelationView(n)
	}
	return vres
}

// newRelationDafault converts projected type Relation to service type Relation.
func newRelationDafault(vres *rosenzuviews.RelationView) *Relation {
	res := &Relation{
		ElementA: vres.ElementA,
		ElementB: vres.ElementB,
	}
	return res
}

// newRelation converts projected type Relation to service type Relation.
func newRelation(vres *rosenzuviews.RelationView) *Relation {
	res := &Relation{
		ElementA: vres.ElementA,
		ElementB: vres.ElementB,
	}
	return res
}

// newRelationViewDafault projects result type Relation to projected type
// RelationView using the "dafault" view.
func newRelationViewDafault(res *Relation) *rosenzuviews.RelationView {
	vres := &rosenzuviews.RelationView{
		ElementA: res.ElementA,
		ElementB: res.ElementB,
	}
	return vres
}

// newRelationView projects result type Relation to projected type RelationView
// using the "default" view.
func newRelationView(res *Relation) *rosenzuviews.RelationView {
	vres := &rosenzuviews.RelationView{
		ElementA: res.ElementA,
		ElementB: res.ElementB,
	}
	return vres
}

// newOperationalpointCollection converts projected type
// OperationalpointCollection to service type OperationalpointCollection.
func newOperationalpointCollection(vres rosenzuviews.OperationalpointCollectionView) OperationalpointCollection {
	res := make(OperationalpointCollection, len(vres))
	for i, n := range vres {
		res[i] = newOperationalpoint(n)
	}
	return res
}

// newOperationalpointCollectionView projects result type
// OperationalpointCollection to projected type OperationalpointCollectionView
// using the "default" view.
func newOperationalpointCollectionView(res OperationalpointCollection) rosenzuviews.OperationalpointCollectionView {
	vres := make(rosenzuviews.OperationalpointCollectionView, len(res))
	for i, n := range res {
		vres[i] = newOperationalpointView(n)
	}
	return vres
}

// newOperationalpoint converts projected type Operationalpoint to service type
// Operationalpoint.
func newOperationalpoint(vres *rosenzuviews.OperationalpointView) *Operationalpoint {
	res := &Operationalpoint{
		Name:      vres.Name,
		ElementID: vres.ElementID,
	}
	return res
}

// newOperationalpointView projects result type Operationalpoint to projected
// type OperationalpointView using the "default" view.
func newOperationalpointView(res *Operationalpoint) *rosenzuviews.OperationalpointView {
	vres := &rosenzuviews.OperationalpointView{
		Name:      res.Name,
		ElementID: res.ElementID,
	}
	return vres
}

// transformRosenzuviewsRelationViewToRelation builds a value of type *Relation
// from a value of type *rosenzuviews.RelationView.
func transformRosenzuviewsRelationViewToRelation(v *rosenzuviews.RelationView) *Relation {
	if v == nil {
		return nil
	}
	res := &Relation{
		ElementA: v.ElementA,
		ElementB: v.ElementB,
	}

	return res
}

// transformRelationToRosenzuviewsRelationView builds a value of type
// *rosenzuviews.RelationView from a value of type *Relation.
func transformRelationToRosenzuviewsRelationView(v *Relation) *rosenzuviews.RelationView {
	if v == nil {
		return nil
	}
	res := &rosenzuviews.RelationView{
		ElementA: v.ElementA,
		ElementB: v.ElementB,
	}

	return res
}
