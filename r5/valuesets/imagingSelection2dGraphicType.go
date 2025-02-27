// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/imagingselection-2dgraphictype
*/type ImagingSelection2DGraphicType string

const (
	// A single location denoted by a single (x,y) pair.
	ImagingSelection2DGraphicTypePoint ImagingSelection2DGraphicType = "point"
	// A series of connected line segments with ordered vertices denoted by (x,y) triplets; the points need not be coplanar.
	ImagingSelection2DGraphicTypePolyline ImagingSelection2DGraphicType = "polyline"
	// An n-tuple list of (x,y) pair end points between which some form of implementation dependent curved lines are to be drawn. The rendered line shall pass through all the specified points.
	ImagingSelection2DGraphicTypeInterpolated ImagingSelection2DGraphicType = "interpolated"
	// Two points shall be present; the first point is to be interpreted as the center and the second point as a point on the circumference of a circle, some form of implementation dependent representation of which is to be drawn.
	ImagingSelection2DGraphicTypeCircle ImagingSelection2DGraphicType = "circle"
	// An ellipse defined by four (x,y) pairs, the first two pairs specifying the endpoints of the major axis and the second two pairs specifying the endpoints of the minor axis.
	ImagingSelection2DGraphicTypeEllipse ImagingSelection2DGraphicType = "ellipse"
)
