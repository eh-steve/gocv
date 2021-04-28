package gocv

/*
#include <stdlib.h>
#include "imgproc.h"
*/
import "C"
import (
	"image"
	"image/color"
	"unsafe"
)

// BorderType type of border.
type BorderType int

const (
	// BorderConstant border type
	BorderConstant BorderType = 0

	// BorderReplicate border type
	BorderReplicate BorderType = 1

	// BorderReflect border type
	BorderReflect BorderType = 2

	// BorderWrap border type
	BorderWrap BorderType = 3

	// BorderReflect101 border type
	BorderReflect101 BorderType = 4

	// BorderTransparent border type
	BorderTransparent BorderType = 5

	// BorderDefault border type
	BorderDefault = BorderReflect101

	// BorderIsolated border type
	BorderIsolated BorderType = 16
)

// GaussianBlur blurs an image Mat using a Gaussian filter.
// The function convolves the src Mat image into the dst Mat using
// the specified Gaussian kernel params.
//
// For further details, please see:
// http://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gaabe8c836e97159a9193fb0b11ac52cf1
//
func GaussianBlur(src Mat, dst *Mat, ksize image.Point, sigmaX float64,
	sigmaY float64, borderType BorderType) {
	pSize := C.struct_Size{
		width:  C.int(ksize.X),
		height: C.int(ksize.Y),
	}

	C.GaussianBlur(src.p, dst.p, pSize, C.double(sigmaX), C.double(sigmaY), C.int(borderType))
}

// ArrowedLine draws a arrow segment pointing from the first point
// to the second one.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga0a165a3ca093fd488ac709fdf10c05b2
//
func ArrowedLine(img *Mat, pt1 image.Point, pt2 image.Point, c color.RGBA, thickness int) {
	sp1 := C.struct_Point{
		x: C.int(pt1.X),
		y: C.int(pt1.Y),
	}

	sp2 := C.struct_Point{
		x: C.int(pt2.X),
		y: C.int(pt2.Y),
	}

	sColor := C.struct_Scalar{
		val1: C.double(c.B),
		val2: C.double(c.G),
		val3: C.double(c.R),
		val4: C.double(c.A),
	}

	C.ArrowedLine(img.p, sp1, sp2, sColor, C.int(thickness))
}

// Circle draws a circle.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#gaf10604b069374903dbd0f0488cb43670
//
func Circle(img *Mat, center image.Point, radius int, c color.RGBA, thickness int) {
	pc := C.struct_Point{
		x: C.int(center.X),
		y: C.int(center.Y),
	}

	sColor := C.struct_Scalar{
		val1: C.double(c.B),
		val2: C.double(c.G),
		val3: C.double(c.R),
		val4: C.double(c.A),
	}

	C.Circle(img.p, pc, C.int(radius), sColor, C.int(thickness))
}

// Ellipse draws a simple or thick elliptic arc or fills an ellipse sector.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga28b2267d35786f5f890ca167236cbc69
//
func Ellipse(img *Mat, center, axes image.Point, angle, startAngle, endAngle float64, c color.RGBA, thickness int) {
	pc := C.struct_Point{
		x: C.int(center.X),
		y: C.int(center.Y),
	}
	pa := C.struct_Point{
		x: C.int(axes.X),
		y: C.int(axes.Y),
	}

	sColor := C.struct_Scalar{
		val1: C.double(c.B),
		val2: C.double(c.G),
		val3: C.double(c.R),
		val4: C.double(c.A),
	}

	C.Ellipse(img.p, pc, pa, C.double(angle), C.double(startAngle), C.double(endAngle), sColor, C.int(thickness))
}

// Line draws a line segment connecting two points.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga7078a9fae8c7e7d13d24dac2520ae4a2
//
func Line(img *Mat, pt1 image.Point, pt2 image.Point, c color.RGBA, thickness int) {
	sp1 := C.struct_Point{
		x: C.int(pt1.X),
		y: C.int(pt1.Y),
	}

	sp2 := C.struct_Point{
		x: C.int(pt2.X),
		y: C.int(pt2.Y),
	}

	sColor := C.struct_Scalar{
		val1: C.double(c.B),
		val2: C.double(c.G),
		val3: C.double(c.R),
		val4: C.double(c.A),
	}

	C.Line(img.p, sp1, sp2, sColor, C.int(thickness))
}

// Rectangle draws a simple, thick, or filled up-right rectangle.
// It renders a rectangle with the desired characteristics to the target Mat image.
//
// For further details, please see:
// http://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga346ac30b5c74e9b5137576c9ee9e0e8c
//
func Rectangle(img *Mat, r image.Rectangle, c color.RGBA, thickness int) {
	cRect := C.struct_Rect{
		x:      C.int(r.Min.X),
		y:      C.int(r.Min.Y),
		width:  C.int(r.Size().X),
		height: C.int(r.Size().Y),
	}

	sColor := C.struct_Scalar{
		val1: C.double(c.B),
		val2: C.double(c.G),
		val3: C.double(c.R),
		val4: C.double(c.A),
	}

	C.Rectangle(img.p, cRect, sColor, C.int(thickness))
}

// FillPoly fills the area bounded by one or more polygons.
//
// For more information, see:
// https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#gaf30888828337aa4c6b56782b5dfbd4b7
func FillPoly(img *Mat, pts PointsVector, c color.RGBA) {
	sColor := C.struct_Scalar{
		val1: C.double(c.B),
		val2: C.double(c.G),
		val3: C.double(c.R),
		val4: C.double(c.A),
	}

	C.FillPoly(img.p, pts.p, sColor)
}

// Polylines draws several polygonal curves.
//
// For more information, see:
// https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga1ea127ffbbb7e0bfc4fd6fd2eb64263c
func Polylines(img *Mat, pts PointsVector, isClosed bool, c color.RGBA, thickness int) {
	sColor := C.struct_Scalar{
		val1: C.double(c.B),
		val2: C.double(c.G),
		val3: C.double(c.R),
		val4: C.double(c.A),
	}

	C.Polylines(img.p, pts.p, C.bool(isClosed), sColor, C.int(thickness))
}

// HersheyFont are the font libraries included in OpenCV.
// Only a subset of the available Hershey fonts are supported by OpenCV.
//
// For more information, see:
// http://sources.isc.org/utils/misc/hershey-font.txt
//
type HersheyFont int

const (
	// FontHersheySimplex is normal size sans-serif font.
	FontHersheySimplex HersheyFont = 0
	// FontHersheyPlain issmall size sans-serif font.
	FontHersheyPlain HersheyFont = 1
	// FontHersheyDuplex normal size sans-serif font
	// (more complex than FontHersheySIMPLEX).
	FontHersheyDuplex HersheyFont = 2
	// FontHersheyComplex i a normal size serif font.
	FontHersheyComplex HersheyFont = 3
	// FontHersheyTriplex is a normal size serif font
	// (more complex than FontHersheyCOMPLEX).
	FontHersheyTriplex HersheyFont = 4
	// FontHersheyComplexSmall is a smaller version of FontHersheyCOMPLEX.
	FontHersheyComplexSmall HersheyFont = 5
	// FontHersheyScriptSimplex is a hand-writing style font.
	FontHersheyScriptSimplex HersheyFont = 6
	// FontHersheyScriptComplex is a more complex variant of FontHersheyScriptSimplex.
	FontHersheyScriptComplex HersheyFont = 7
	// FontItalic is the flag for italic font.
	FontItalic HersheyFont = 16
)

// LineType are the line libraries included in OpenCV.
//
// For more information, see:
// https://vovkos.github.io/doxyrest-showcase/opencv/sphinx_rtd_theme/enum_cv_LineTypes.html
//
type LineType int

const (
	// Filled line
	Filled LineType = -1
	// Line4 4-connected line
	Line4 LineType = 4
	// Line8 8-connected line
	Line8 LineType = 8
	// LineAA antialiased line
	LineAA LineType = 16
)

// GetTextSize calculates the width and height of a text string.
// It returns an image.Point with the size required to draw text using
// a specific font face, scale, and thickness.
//
// For further details, please see:
// http://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga3d2abfcb995fd2db908c8288199dba82
//
func GetTextSize(text string, fontFace HersheyFont, fontScale float64, thickness int) image.Point {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	sz := C.GetTextSize(cText, C.int(fontFace), C.double(fontScale), C.int(thickness))
	return image.Pt(int(sz.width), int(sz.height))
}

// GetTextSizeWithBaseline calculates the width and height of a text string including the basline of the text.
// It returns an image.Point with the size required to draw text using
// a specific font face, scale, and thickness as well as its baseline.
//
// For further details, please see:
// http://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga3d2abfcb995fd2db908c8288199dba82
//
func GetTextSizeWithBaseline(text string, fontFace HersheyFont, fontScale float64, thickness int) (image.Point, int) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))
	cBaseline := C.int(0)

	sz := C.GetTextSizeWithBaseline(cText, C.int(fontFace), C.double(fontScale), C.int(thickness), &cBaseline)
	return image.Pt(int(sz.width), int(sz.height)), int(cBaseline)
}

// PutText draws a text string.
// It renders the specified text string into the img Mat at the location
// passed in the "org" param, using the desired font face, font scale,
// color, and line thinkness.
//
// For further details, please see:
// http://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga5126f47f883d730f633d74f07456c576
//
func PutText(img *Mat, text string, org image.Point, fontFace HersheyFont, fontScale float64, c color.RGBA, thickness int) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	pOrg := C.struct_Point{
		x: C.int(org.X),
		y: C.int(org.Y),
	}

	sColor := C.struct_Scalar{
		val1: C.double(c.B),
		val2: C.double(c.G),
		val3: C.double(c.R),
		val4: C.double(c.A),
	}

	C.PutText(img.p, cText, pOrg, C.int(fontFace), C.double(fontScale), sColor, C.int(thickness))
	return
}

// PutTextWithParams draws a text string.
// It renders the specified text string into the img Mat at the location
// passed in the "org" param, using the desired font face, font scale,
// color, and line thinkness.
//
// For further details, please see:
// http://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga5126f47f883d730f633d74f07456c576
//
func PutTextWithParams(img *Mat, text string, org image.Point, fontFace HersheyFont, fontScale float64, c color.RGBA, thickness int, lineType LineType, bottomLeftOrigin bool) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	pOrg := C.struct_Point{
		x: C.int(org.X),
		y: C.int(org.Y),
	}

	sColor := C.struct_Scalar{
		val1: C.double(c.B),
		val2: C.double(c.G),
		val3: C.double(c.R),
		val4: C.double(c.A),
	}

	C.PutTextWithParams(img.p, cText, pOrg, C.int(fontFace), C.double(fontScale), sColor, C.int(thickness), C.int(lineType), C.bool(bottomLeftOrigin))
	return
}

// InterpolationFlags are bit flags that control the interpolation algorithm
// that is used.
type InterpolationFlags int

const (
	// InterpolationNearestNeighbor is nearest neighbor. (fast but low quality)
	InterpolationNearestNeighbor InterpolationFlags = 0

	// InterpolationLinear is bilinear interpolation.
	InterpolationLinear InterpolationFlags = 1

	// InterpolationCubic is bicube interpolation.
	InterpolationCubic InterpolationFlags = 2

	// InterpolationArea uses pixel area relation. It is preferred for image
	// decimation as it gives moire-free results.
	InterpolationArea InterpolationFlags = 3

	// InterpolationLanczos4 is Lanczos interpolation over 8x8 neighborhood.
	InterpolationLanczos4 InterpolationFlags = 4

	// InterpolationDefault is an alias for InterpolationLinear.
	InterpolationDefault = InterpolationLinear

	// InterpolationMax indicates use maximum interpolation.
	InterpolationMax InterpolationFlags = 7
)

// CvtColor converts an image from one color space to another.
// It converts the src Mat image to the dst Mat using the
// code param containing the desired ColorConversionCode color space.
//
// For further details, please see:
// http://docs.opencv.org/master/d7/d1b/group__imgproc__misc.html#ga4e0972be5de079fed4e3a10e24ef5ef0
//
func CvtColor(src Mat, dst *Mat, code ColorConversionCode) {
	C.CvtColor(src.p, dst.p, C.int(code))
}

// Resize resizes an image.
// It resizes the image src down to or up to the specified size, storing the
// result in dst. Note that src and dst may be the same image. If you wish to
// scale by factor, an empty sz may be passed and non-zero fx and fy. Likewise,
// if you wish to scale to an explicit size, a non-empty sz may be passed with
// zero for both fx and fy.
//
// For further details, please see:
// https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga47a974309e9102f5f08231edc7e7529d
func Resize(src Mat, dst *Mat, sz image.Point, fx, fy float64, interp InterpolationFlags) {
	pSize := C.struct_Size{
		width:  C.int(sz.X),
		height: C.int(sz.Y),
	}

	C.Resize(src.p, dst.p, pSize, C.double(fx), C.double(fy), C.int(interp))
	return
}

// GetPerspectiveTransform returns 3x3 perspective transformation for the
// corresponding 4 point pairs as image.Point.
//
// For further details, please see:
// https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga8c1ae0e3589a9d77fffc962c49b22043
func GetPerspectiveTransform(src, dst PointVector) Mat {
	return newMat(C.GetPerspectiveTransform(src.p, dst.p))
}

// GetPerspectiveTransform2f returns 3x3 perspective transformation for the
// corresponding 4 point pairs as gocv.Point2f.
//
// For further details, please see:
// https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga8c1ae0e3589a9d77fffc962c49b22043
func GetPerspectiveTransform2f(src, dst Point2fVector) Mat {
	return newMat(C.GetPerspectiveTransform2f(src.p, dst.p))
}

// Apply phaseCorrelate.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/df3/group__imgproc__motion.html#ga552420a2ace9ef3fb053cd630fdb4952
//
func PhaseCorrelate(src1, src2, window Mat) (phaseShift Point2f, response float64) {
	var responseDouble C.double
	result := C.PhaseCorrelate(src1.p, src2.p, window.p, &responseDouble)

	return Point2f{
		X: float32(result.x),
		Y: float32(result.y),
	}, float64(responseDouble)
}
