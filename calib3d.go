package gocv

/*
#include <stdlib.h>
#include "calib3d.h"
*/
import "C"
import "image"

// EstimateNewCameraMatrixForUndistortRectify estimates new camera matrix for undistortion or rectification.
//
// For further details, please see:
// https://docs.opencv.org/master/db/d58/group__calib3d__fisheye.html#ga384940fdf04c03e362e94b6eb9b673c9
func EstimateNewCameraMatrixForUndistortRectify(k, d Mat, imgSize image.Point, r Mat, p *Mat, balance float64, newSize image.Point, fovScale float64) {
	imgSz := C.struct_Size{
		width:  C.int(imgSize.X),
		height: C.int(imgSize.Y),
	}
	newSz := C.struct_Size{
		width:  C.int(newSize.X),
		height: C.int(newSize.Y),
	}
	C.Fisheye_EstimateNewCameraMatrixForUndistortRectify(k.Ptr(), d.Ptr(), imgSz, r.Ptr(), p.Ptr(), C.double(balance), newSz, C.double(fovScale))
}

// GetOptimalNewCameraMatrixWithParams computes and returns the optimal new camera matrix based on the free scaling parameter.
//
// For further details, please see:
// https://docs.opencv.org/master/d9/d0c/group__calib3d.html#ga7a6c4e032c97f03ba747966e6ad862b1
//
func GetOptimalNewCameraMatrixWithParams(cameraMatrix Mat, distCoeffs Mat, imageSize image.Point, alpha float64, newImgSize image.Point, centerPrincipalPoint bool) (Mat, image.Rectangle) {
	sz := C.struct_Size{
		width:  C.int(imageSize.X),
		height: C.int(imageSize.Y),
	}
	newSize := C.struct_Size{
		width:  C.int(newImgSize.X),
		height: C.int(newImgSize.Y),
	}
	rt := C.struct_Rect{}
	return newMat(C.GetOptimalNewCameraMatrixWithParams(cameraMatrix.Ptr(), distCoeffs.Ptr(), sz, C.double(alpha), newSize, &rt, C.bool(centerPrincipalPoint))), toRect(rt)
}

// FisheyeUndistortPoints transforms points to compensate for fisheye lens distortion
//
// For further details, please see:
// https://docs.opencv.org/master/db/d58/group__calib3d__fisheye.html#gab738cdf90ceee97b2b52b0d0e7511541
func FisheyeUndistortPoints(distorted Mat, undistorted *Mat, k, d, r, p Mat) {
	C.Fisheye_UndistortPoints(distorted.Ptr(), undistorted.Ptr(), k.Ptr(), d.Ptr(), r.Ptr(), p.Ptr())
}

func Undistort(src Mat, dst *Mat, cameraMatrix Mat, distCoeffs Mat, newCameraMatrix Mat) {
	C.Undistort(src.Ptr(), dst.Ptr(), cameraMatrix.Ptr(), distCoeffs.Ptr(), newCameraMatrix.Ptr())
}

// UndistortPoints transforms points to compensate for lens distortion
//
// For further details, please see:
// https://docs.opencv.org/master/d9/d0c/group__calib3d.html#ga55c716492470bfe86b0ee9bf3a1f0f7e
func UndistortPoints(src Mat, dst *Mat, cameraMatrix, distCoeffs, rectificationTransform, newCameraMatrix Mat) {
	C.UndistortPoints(src.Ptr(), dst.Ptr(), cameraMatrix.Ptr(), distCoeffs.Ptr(), rectificationTransform.Ptr(), newCameraMatrix.Ptr())
}

// FisheyeUndistortImageWithParams transforms an image to compensate for fisheye lens distortion with Knew matrix
func FisheyeUndistortImageWithParams(distorted Mat, undistorted *Mat, k, d, knew Mat, size image.Point) {
	sz := C.struct_Size{
		width:  C.int(size.X),
		height: C.int(size.Y),
	}
	C.Fisheye_UndistortImageWithParams(distorted.Ptr(), undistorted.Ptr(), k.Ptr(), d.Ptr(), knew.Ptr(), sz)
}
