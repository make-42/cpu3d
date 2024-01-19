package main

import (
	"fmt"
	"math"

	"github.com/make-42/cpu3d/camera"
	"github.com/make-42/cpu3d/model"
	"github.com/make-42/cpu3d/render"
	"github.com/make-42/cpu3d/transform"
	"github.com/make-42/cpu3d/utils"
)

func main() {
	displayResolution := utils.IntPair{X: 2156 * 2, Y: 1440 * 2}
	worldCamera := camera.Camera{
		FocalLength: 0.06, //m
		SensorSize: utils.PlaneCoords{
			X: 22.3e-3, //m
			Y: 14.9e-3, //m
		},
	}
	frameCount := 60

	/*
		pointCloudOrigin := []utils.SpaceCoords{
			{X: 0, Y: 0, Z: 0},
			{X: 0, Y: 0, Z: 1},
			{X: 0, Y: 1, Z: 0},
			{X: 0, Y: 1, Z: 1},
			{X: 1, Y: 0, Z: 0},
			{X: 1, Y: 0, Z: 1},
			{X: 1, Y: 1, Z: 0},
			{X: 1, Y: 1, Z: 1}}

		pointCloudCentered := transform.PointCloudTranslateWorldCoords(&pointCloudOrigin, &utils.SpaceCoords{
			X: -0.5,
			Y: -0.5,
			Z: -0.5,
		})
			for i := 0; i < frameCount; i++ {
				pointCloudRotated := transform.PointCloudRotateWorldCoords(&pointCloudCentered, &utils.RotationCoords{
					X: float64(i) / 8,
					Y: float64(i) / 16,
					Z: float64(i) / 32,
				})

				pointCloud := transform.PointCloudTranslateWorldCoords(&pointCloudRotated, &utils.SpaceCoords{
					X: 0,
					Y: 0,
					Z: 2,
				})
				//screenCoords := camera.PointCloudWorldCoordsToSensorCoords(&pointCloud, &worldCamera)
				//render.RenderScreenPoints(4, displayResolution, &screenCoords, "output-pointcloud.png")

				edges := []utils.SpaceEdge{
					{A: pointCloud[0], B: pointCloud[1]},
					{A: pointCloud[0], B: pointCloud[2]},
					{A: pointCloud[0], B: pointCloud[4]},
					{A: pointCloud[7], B: pointCloud[3]},
					{A: pointCloud[7], B: pointCloud[5]},
					{A: pointCloud[7], B: pointCloud[6]},

					{A: pointCloud[1], B: pointCloud[5]},
					{A: pointCloud[1], B: pointCloud[3]},
					{A: pointCloud[2], B: pointCloud[6]},
					{A: pointCloud[2], B: pointCloud[3]},
					{A: pointCloud[4], B: pointCloud[6]},
					{A: pointCloud[4], B: pointCloud[5]},
				}
				lines := camera.EdgesWorldCoordsToSensorCoords(&edges, &worldCamera)
				render.RenderScreenLines(1, displayResolution, &lines, fmt.Sprintf("output/cube/output-lines-%04d.png", i))
			}
	*/
	/*
		edgesTeapot := model.ReadSTLFileToEdges("input/teapot.stl")
		edgesTeapotCentered := transform.EdgesTranslateWorldCoords(&edgesTeapot, &utils.SpaceCoords{X: 0, Y: 0, Z: -3})
		for i := 0; i < frameCount; i++ {
			progress := ease.InOutQuart(float64(i) / float64(frameCount-1))
			edgesTeapotRotatedFirst := transform.EdgesRotateWorldCoords(&edgesTeapotCentered, &utils.RotationCoords{
				X: 2 * math.Pi * progress,
				Y: 0,
				Z: 0,
			})
			edgesTeapotRotated := transform.EdgesRotateWorldCoords(&edgesTeapotRotatedFirst, &utils.RotationCoords{X: 0, Y: 0, Z: -math.Pi / 4})

			edgesTeapotScaled := transform.EdgesScaleWorldCoords(&edgesTeapotRotated, 0.2)
			edges := transform.EdgesTranslateWorldCoords(&edgesTeapotScaled, &utils.SpaceCoords{
				X: 0,
				Y: 0,
				Z: 10,
			})
			lines := camera.EdgesWorldCoordsToSensorCoords(&edges, &worldCamera)
			render.RenderScreenLines(1, displayResolution, &lines, fmt.Sprintf("output/teacup/output-lines-%04d.png", i))
		}
	*/
	edgesWM := model.ReadSTLFileToEdges("input/venus - simplified med.stl")
	edgesWMCentered := transform.EdgesTranslateWorldCoords(&edgesWM, &utils.SpaceCoords{X: 0, Y: 0, Z: 0})
	for i := 0; i < frameCount; i++ {
		progress := float64(i) / float64(frameCount-1)
		edgesWMRotatedFirst := transform.EdgesRotateWorldCoords(&edgesWMCentered, &utils.RotationCoords{
			X: 2 * math.Pi * progress,
			Y: 0,
			Z: 0,
		})
		edgesWMRotated := transform.EdgesRotateWorldCoords(&edgesWMRotatedFirst, &utils.RotationCoords{X: 0, Y: 0, Z: -math.Pi / 4})

		edgesWMScaled := transform.EdgesScaleWorldCoords(&edgesWMRotated, 0.01)
		edges := transform.EdgesTranslateWorldCoords(&edgesWMScaled, &utils.SpaceCoords{
			X: 0,
			Y: 0,
			Z: 5,
		})
		lines := camera.EdgesWorldCoordsToSensorCoords(&edges, &worldCamera)
		render.RenderScreenLines(displayResolution, &lines, fmt.Sprintf("output/venus/output-lines-%04d.png", i))
	}
}
