package gocv

/*
#include <stdlib.h>
#include "core.h"
*/
import "C"
import (
	"errors"
	"image"
	"reflect"
	"unsafe"
)

const (
	// MatChannels1 is a single channel Mat.
	MatChannels1 = 0

	// MatChannels2 is 2 channel Mat.
	MatChannels2 = 8

	// MatChannels3 is 3 channel Mat.
	MatChannels3 = 16

	// MatChannels4 is 4 channel Mat.
	MatChannels4 = 24
)

// MatType is the type for the various different kinds of Mat you can create.
type MatType int

const (
	// MatTypeCV8U is a Mat of 8-bit unsigned int
	MatTypeCV8U MatType = 0

	// MatTypeCV8S is a Mat of 8-bit signed int
	MatTypeCV8S MatType = 1

	// MatTypeCV16U is a Mat of 16-bit unsigned int
	MatTypeCV16U MatType = 2

	// MatTypeCV16S is a Mat of 16-bit signed int
	MatTypeCV16S MatType = 3

	// MatTypeCV16SC2 is a Mat of 16-bit signed int with 2 channels
	MatTypeCV16SC2 = MatTypeCV16S + MatChannels2

	// MatTypeCV32S is a Mat of 32-bit signed int
	MatTypeCV32S MatType = 4

	// MatTypeCV32F is a Mat of 32-bit float
	MatTypeCV32F MatType = 5

	// MatTypeCV64F is a Mat of 64-bit float
	MatTypeCV64F MatType = 6

	// MatTypeCV8UC1 is a Mat of 8-bit unsigned int with a single channel
	MatTypeCV8UC1 = MatTypeCV8U + MatChannels1

	// MatTypeCV8UC2 is a Mat of 8-bit unsigned int with 2 channels
	MatTypeCV8UC2 = MatTypeCV8U + MatChannels2

	// MatTypeCV8UC3 is a Mat of 8-bit unsigned int with 3 channels
	MatTypeCV8UC3 = MatTypeCV8U + MatChannels3

	// MatTypeCV8UC4 is a Mat of 8-bit unsigned int with 4 channels
	MatTypeCV8UC4 = MatTypeCV8U + MatChannels4

	// MatTypeCV8SC1 is a Mat of 8-bit signed int with a single channel
	MatTypeCV8SC1 = MatTypeCV8S + MatChannels1

	// MatTypeCV8SC2 is a Mat of 8-bit signed int with 2 channels
	MatTypeCV8SC2 = MatTypeCV8S + MatChannels2

	// MatTypeCV8SC3 is a Mat of 8-bit signed int with 3 channels
	MatTypeCV8SC3 = MatTypeCV8S + MatChannels3

	// MatTypeCV8SC4 is a Mat of 8-bit signed int with 4 channels
	MatTypeCV8SC4 = MatTypeCV8S + MatChannels4

	// MatTypeCV16UC1 is a Mat of 16-bit unsigned int with a single channel
	MatTypeCV16UC1 = MatTypeCV16U + MatChannels1

	// MatTypeCV16UC2 is a Mat of 16-bit unsigned int with 2 channels
	MatTypeCV16UC2 = MatTypeCV16U + MatChannels2

	// MatTypeCV16UC3 is a Mat of 16-bit unsigned int with 3 channels
	MatTypeCV16UC3 = MatTypeCV16U + MatChannels3

	// MatTypeCV16UC4 is a Mat of 16-bit unsigned int with 4 channels
	MatTypeCV16UC4 = MatTypeCV16U + MatChannels4

	// MatTypeCV16SC1 is a Mat of 16-bit signed int with a single channel
	MatTypeCV16SC1 = MatTypeCV16S + MatChannels1

	// MatTypeCV16SC3 is a Mat of 16-bit signed int with 3 channels
	MatTypeCV16SC3 = MatTypeCV16S + MatChannels3

	// MatTypeCV16SC4 is a Mat of 16-bit signed int with 4 channels
	MatTypeCV16SC4 = MatTypeCV16S + MatChannels4

	// MatTypeCV32SC1 is a Mat of 32-bit signed int with a single channel
	MatTypeCV32SC1 = MatTypeCV32S + MatChannels1

	// MatTypeCV32SC2 is a Mat of 32-bit signed int with 2 channels
	MatTypeCV32SC2 = MatTypeCV32S + MatChannels2

	// MatTypeCV32SC3 is a Mat of 32-bit signed int with 3 channels
	MatTypeCV32SC3 = MatTypeCV32S + MatChannels3

	// MatTypeCV32SC4 is a Mat of 32-bit signed int with 4 channels
	MatTypeCV32SC4 = MatTypeCV32S + MatChannels4

	// MatTypeCV32FC1 is a Mat of 32-bit float int with a single channel
	MatTypeCV32FC1 = MatTypeCV32F + MatChannels1

	// MatTypeCV32FC2 is a Mat of 32-bit float int with 2 channels
	MatTypeCV32FC2 = MatTypeCV32F + MatChannels2

	// MatTypeCV32FC3 is a Mat of 32-bit float int with 3 channels
	MatTypeCV32FC3 = MatTypeCV32F + MatChannels3

	// MatTypeCV32FC4 is a Mat of 32-bit float int with 4 channels
	MatTypeCV32FC4 = MatTypeCV32F + MatChannels4

	// MatTypeCV64FC1 is a Mat of 64-bit float int with a single channel
	MatTypeCV64FC1 = MatTypeCV64F + MatChannels1

	// MatTypeCV64FC2 is a Mat of 64-bit float int with 2 channels
	MatTypeCV64FC2 = MatTypeCV64F + MatChannels2

	// MatTypeCV64FC3 is a Mat of 64-bit float int with 3 channels
	MatTypeCV64FC3 = MatTypeCV64F + MatChannels3

	// MatTypeCV64FC4 is a Mat of 64-bit float int with 4 channels
	MatTypeCV64FC4 = MatTypeCV64F + MatChannels4
)

// CompareType is used for Compare operations to indicate which kind of
// comparison to use.
type CompareType int

const (
	// CompareEQ src1 is equal to src2.
	CompareEQ CompareType = 0

	// CompareGT src1 is greater than src2.
	CompareGT CompareType = 1

	// CompareGE src1 is greater than or equal to src2.
	CompareGE CompareType = 2

	// CompareLT src1 is less than src2.
	CompareLT CompareType = 3

	// CompareLE src1 is less than or equal to src2.
	CompareLE CompareType = 4

	// CompareNE src1 is unequal to src2.
	CompareNE CompareType = 5
)

type Point2f struct {
	X float32
	Y float32
}

var ErrEmptyByteSlice = errors.New("empty byte array")

// Mat represents an n-dimensional dense numerical single-channel
// or multi-channel array. It can be used to store real or complex-valued
// vectors and matrices, grayscale or color images, voxel volumes,
// vector fields, point clouds, tensors, and histograms.
//
// For further details, please see:
// http://docs.opencv.org/master/d3/d63/classcv_1_1Mat.html
//
type Mat struct {
	p C.Mat

	// Non-nil if Mat was created with a []byte (using NewMatFromBytes()). Nil otherwise.
	d []byte
}

// NewMat returns a new empty Mat.
func NewMat() Mat {
	return newMat(C.Mat_New())
}

// NewMatWithSize returns a new Mat with a specific size and type.
func NewMatWithSize(rows int, cols int, mt MatType) Mat {
	return newMat(C.Mat_NewWithSize(C.int(rows), C.int(cols), C.int(mt)))
}

// NewMatFromBytes returns a new Mat with a specific size and type, initialized from a []byte.
func NewMatFromBytes(rows int, cols int, mt MatType, data []byte) (Mat, error) {
	cBytes, err := toByteArray(data)
	if err != nil {
		return Mat{}, err
	}
	mat := newMat(C.Mat_NewFromBytes(C.int(rows), C.int(cols), C.int(mt), *cBytes))

	// Store a reference to the backing data slice. This is needed because we pass the backing
	// array directly to C code and without keeping a Go reference to it, it might end up
	// garbage collected which would result in crashes.
	//
	// TODO(bga): This could live in newMat() but I wanted to reduce the change surface.
	// TODO(bga): Code that needs access to the array from Go could use this directly.
	mat.d = data

	return mat, nil
}

// Ptr returns the Mat's underlying object pointer.
func (m *Mat) Ptr() C.Mat {
	return m.p
}

// Clone returns a cloned full copy of the Mat.
func (m *Mat) Clone() Mat {
	return newMat(C.Mat_Clone(m.p))
}

// CopyTo copies Mat into destination Mat.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d63/classcv_1_1Mat.html#a33fd5d125b4c302b0c9aa86980791a77
//
func (m *Mat) CopyTo(dst *Mat) {
	C.Mat_CopyTo(m.p, dst.p)
	return
}

// Empty determines if the Mat is empty or not.
func (m *Mat) Empty() bool {
	isEmpty := C.Mat_Empty(m.p)
	return isEmpty != 0
}

// Rows returns the number of rows for this Mat.
func (m *Mat) Rows() int {
	return int(C.Mat_Rows(m.p))
}

// Cols returns the number of columns for this Mat.
func (m *Mat) Cols() int {
	return int(C.Mat_Cols(m.p))
}

// Channels returns the number of channels for this Mat.
func (m *Mat) Channels() int {
	return int(C.Mat_Channels(m.p))
}

// Type returns the type for this Mat.
func (m *Mat) Type() MatType {
	return MatType(C.Mat_Type(m.p))
}

// Step returns the number of bytes each matrix row occupies.
func (m *Mat) Step() int {
	return int(C.Mat_Step(m.p))
}

// SetDoubleAt sets a value at a specific row/col
// in this Mat expecting it to be of type double aka CV_64F.
func (m *Mat) SetDoubleAt(row int, col int, val float64) {
	C.Mat_SetDouble(m.p, C.int(row), C.int(col), C.double(val))
}

// GetDoubleAt returns a value from a specific row/col
// in this Mat expecting it to be of type double aka CV_64F.
func (m *Mat) GetDoubleAt(row int, col int) float64 {
	return float64(C.Mat_GetDouble(m.p, C.int(row), C.int(col)))
}

// ToBytes copies the underlying Mat data to a byte array.
//
// For further details, please see:
// https://docs.opencv.org/3.3.1/d3/d63/classcv_1_1Mat.html#a4d33bed1c850265370d2af0ff02e1564
func (m *Mat) ToBytes() []byte {
	b := C.Mat_DataPtr(m.p)
	return toGoBytes(b)
}

// IsContinuous determines if the Mat is continuous.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d63/classcv_1_1Mat.html#aa90cea495029c7d1ee0a41361ccecdf3
//
func (m *Mat) IsContinuous() bool {
	return bool(C.Mat_IsContinuous(m.p))
}

// DataPtrUint8 returns a slice that references the OpenCV allocated data.
//
// The data is no longer valid once the Mat has been closed. Any data that
// needs to be accessed after the Mat is closed must be copied into Go memory.
func (m *Mat) DataPtrUint8() ([]uint8, error) {
	if !m.IsContinuous() {
		return nil, errors.New("DataPtrUint8 requires continuous Mat")
	}

	p := C.Mat_DataPtr(m.p)
	h := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(p.data)),
		Len:  int(p.length),
		Cap:  int(p.length),
	}
	return *(*[]uint8)(unsafe.Pointer(h)), nil
}

func toByteArray(b []byte) (*C.struct_ByteArray, error) {
	if len(b) == 0 {
		return nil, ErrEmptyByteSlice
	}
	return &C.struct_ByteArray{
		data:   (*C.char)(unsafe.Pointer(&b[0])),
		length: C.int(len(b)),
	}, nil
}

func toGoBytes(b C.struct_ByteArray) []byte {
	return C.GoBytes(unsafe.Pointer(b.data), b.length)
}

func toRect(rect C.Rect) image.Rectangle {
	return image.Rect(int(rect.x), int(rect.y), int(rect.x+rect.width), int(rect.y+rect.height))
}

// Region returns a new Mat that points to a region of this Mat. Changes made to the
// region Mat will affect the original Mat, since they are pointers to the underlying
// OpenCV Mat object.
func (m *Mat) Region(rio image.Rectangle) Mat {
	cRect := C.struct_Rect{
		x:      C.int(rio.Min.X),
		y:      C.int(rio.Min.Y),
		width:  C.int(rio.Size().X),
		height: C.int(rio.Size().Y),
	}

	return newMat(C.Mat_Region(m.p, cRect))
}

// PointVector is a wrapper around a std::vector< cv::Point >*
// This is needed anytime that you need to pass or receive a collection of points.
type PointVector struct {
	p C.PointVector
}

// NewPointVector returns a new empty PointVector.
func NewPointVector() PointVector {
	return PointVector{p: C.PointVector_New()}
}

// NewPointVectorFromPoints returns a new PointVector that has been
// initialized to a slice of image.Point.
func NewPointVectorFromPoints(pts []image.Point) PointVector {
	p := (*C.struct_Point)(C.malloc(C.size_t(C.sizeof_struct_Point * len(pts))))
	defer C.free(unsafe.Pointer(p))

	h := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(p)),
		Len:  len(pts),
		Cap:  len(pts),
	}
	pa := *(*[]C.Point)(unsafe.Pointer(h))

	for j, point := range pts {
		pa[j] = C.struct_Point{
			x: C.int(point.X),
			y: C.int(point.Y),
		}
	}

	cpoints := C.struct_Points{
		points: (*C.Point)(p),
		length: C.int(len(pts)),
	}

	return PointVector{p: C.PointVector_NewFromPoints(cpoints)}
}

// IsNil checks the CGo pointer in the PointVector.
func (pv PointVector) IsNil() bool {
	return pv.p == nil
}

// Size returns how many Point are in the PointVector.
func (pv PointVector) Size() int {
	return int(C.PointVector_Size(pv.p))
}

// At returns the image.Point
func (pv PointVector) At(idx int) image.Point {
	if idx > pv.Size() {
		return image.Point{}
	}

	cp := C.PointVector_At(pv.p, C.int(idx))
	return image.Pt(int(cp.x), int(cp.y))
}

// Append appends an image.Point at end of the PointVector.
func (pv PointVector) Append(point image.Point) {
	p := C.struct_Point{
		x: C.int(point.X),
		y: C.int(point.Y),
	}

	C.PointVector_Append(pv.p, p)

	return
}

// ToPoints returns a slice of image.Point for the data in this PointVector.
func (pv PointVector) ToPoints() []image.Point {
	points := make([]image.Point, pv.Size())

	for j := 0; j < pv.Size(); j++ {
		points[j] = pv.At(j)
	}
	return points
}

// Close closes and frees memory for this PointVector.
func (pv PointVector) Close() {
	C.PointVector_Close(pv.p)
}

// Point2fVector is a wrapper around a std::vector< cv::Point2f >*
// This is needed anytime that you need to pass or receive a collection of points.
type Point2fVector struct {
	p C.Point2fVector
}

// NewPoint2fVector returns a new empty Point2fVector.
func NewPoint2fVector() Point2fVector {
	return Point2fVector{p: C.Point2fVector_New()}
}

// NewPoint2fVectorFromPoints returns a new Point2fVector that has been
// initialized to a slice of image.Point.
func NewPoint2fVectorFromPoints(pts []Point2f) Point2fVector {
	p := (*C.struct_Point2f)(C.malloc(C.size_t(C.sizeof_struct_Point2f * len(pts))))
	defer C.free(unsafe.Pointer(p))

	h := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(p)),
		Len:  len(pts),
		Cap:  len(pts),
	}
	pa := *(*[]C.Point2f)(unsafe.Pointer(h))

	for j, point := range pts {
		pa[j] = C.struct_Point2f{
			x: C.float(point.X),
			y: C.float(point.Y),
		}
	}

	cpoints := C.struct_Points2f{
		points: (*C.Point2f)(p),
		length: C.int(len(pts)),
	}

	return Point2fVector{p: C.Point2fVector_NewFromPoints(cpoints)}
}

// IsNil checks the CGo pointer in the Point2fVector.
func (pfv Point2fVector) IsNil() bool {
	return pfv.p == nil
}

// Size returns how many Point are in the PointVector.
func (pfv Point2fVector) Size() int {
	return int(C.Point2fVector_Size(pfv.p))
}

// At returns the image.Point
func (pfv Point2fVector) At(idx int) Point2f {
	if idx > pfv.Size() {
		return Point2f{}
	}

	cp := C.Point2fVector_At(pfv.p, C.int(idx))
	return Point2f{float32(cp.x), float32(cp.y)}
}

// ToPoints returns a slice of image.Point for the data in this PointVector.
func (pfv Point2fVector) ToPoints() []Point2f {
	points := make([]Point2f, pfv.Size())

	for j := 0; j < pfv.Size(); j++ {
		points[j] = pfv.At(j)
	}
	return points
}

// Close closes and frees memory for this Point2fVector.
func (pfv Point2fVector) Close() {
	C.Point2fVector_Close(pfv.p)
}

// PointsVector is a wrapper around a std::vector< std::vector< cv::Point > >*
type PointsVector struct {
	p C.PointsVector
}

// NewPointsVector returns a new empty PointsVector.
func NewPointsVector() PointsVector {
	return PointsVector{p: C.PointsVector_New()}
}

// NewPointsVectorFromPoints returns a new PointsVector that has been
// initialized to a slice of slices of image.Point.
func NewPointsVectorFromPoints(pts [][]image.Point) PointsVector {
	points := make([]C.struct_Points, len(pts))

	for i, pt := range pts {
		p := (*C.struct_Point)(C.malloc(C.size_t(C.sizeof_struct_Point * len(pt))))
		defer C.free(unsafe.Pointer(p))

		h := &reflect.SliceHeader{
			Data: uintptr(unsafe.Pointer(p)),
			Len:  len(pt),
			Cap:  len(pt),
		}
		pa := *(*[]C.Point)(unsafe.Pointer(h))

		for j, point := range pt {
			pa[j] = C.struct_Point{
				x: C.int(point.X),
				y: C.int(point.Y),
			}
		}

		points[i] = C.struct_Points{
			points: (*C.Point)(p),
			length: C.int(len(pt)),
		}
	}

	cPoints := C.struct_Contours{
		contours: (*C.struct_Points)(&points[0]),
		length:   C.int(len(pts)),
	}

	return PointsVector{p: C.PointsVector_NewFromPoints(cPoints)}
}

// ToPoints returns a slice of slices of image.Point for the data in this PointsVector.
func (pvs PointsVector) ToPoints() [][]image.Point {
	ppoints := make([][]image.Point, pvs.Size())
	for i := 0; i < pvs.Size(); i++ {
		pts := pvs.At(i)
		points := make([]image.Point, pts.Size())

		for j := 0; j < pts.Size(); j++ {
			points[j] = pts.At(j)
		}
		ppoints[i] = points
	}

	return ppoints
}

// IsNil checks the CGo pointer in the PointsVector.
func (pvs PointsVector) IsNil() bool {
	return pvs.p == nil
}

// Size returns how many vectors of Points are in the PointsVector.
func (pvs PointsVector) Size() int {
	return int(C.PointsVector_Size(pvs.p))
}

// At returns the PointVector at that index of the PointsVector.
func (pvs PointsVector) At(idx int) PointVector {
	if idx > pvs.Size() {
		return PointVector{}
	}

	return PointVector{p: C.PointsVector_At(pvs.p, C.int(idx))}
}

// Append appends a PointVector at end of the PointsVector.
func (pvs PointsVector) Append(pv PointVector) {
	if !pv.IsNil() {
		C.PointsVector_Append(pvs.p, pv.p)
	}

	return
}

// Close closes and frees memory for this PointsVector.
func (pvs PointsVector) Close() {
	C.PointsVector_Close(pvs.p)
}

// PerspectiveTransform performs the perspective matrix transformation of vectors.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#gad327659ac03e5fd6894b90025e6900a7
//
func PerspectiveTransform(src Mat, dst *Mat, tm Mat) {
	C.Mat_PerspectiveTransform(src.p, dst.p, tm.p)
}

// ConvertTo converts Mat into destination Mat.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d63/classcv_1_1Mat.html#adf88c60c5b4980e05bb556080916978b
//
func (m *Mat) ConvertTo(dst *Mat, mt MatType) {
	C.Mat_ConvertTo(m.p, dst.p, C.int(mt))
	return
}

// Size returns an array with one element for each dimension containing the size of that dimension for the Mat.
//
// For further details, please see:
// https://docs.opencv.org/master/d3/d63/classcv_1_1Mat.html#aa4d317d43fb0cba9c2503f3c61b866c8
//
func (m *Mat) Size() (dims []int) {
	cdims := C.IntVector{}
	C.Mat_Size(m.p, &cdims)
	defer C.IntVector_Close(cdims)

	h := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cdims.val)),
		Len:  int(cdims.length),
		Cap:  int(cdims.length),
	}
	pdims := *(*[]C.int)(unsafe.Pointer(h))

	for i := 0; i < int(cdims.length); i++ {
		dims = append(dims, int(pdims[i]))
	}
	return
}

// Split creates an array of single channel images from a multi-channel image
// Created images should be closed manualy to avoid memory leaks.
//
// For further details, please see:
// https://docs.opencv.org/master/d2/de8/group__core__array.html#ga0547c7fed86152d7e9d0096029c8518a
//
func Split(src Mat) (mv []Mat) {
	cMats := C.struct_Mats{}
	C.Mat_Split(src.p, &(cMats))
	defer C.Mats_Close(cMats)
	mv = make([]Mat, cMats.length)
	for i := C.int(0); i < cMats.length; i++ {
		mv[i].p = C.Mats_get(cMats, i)
		addMatToProfile(mv[i].p)
	}
	return
}
