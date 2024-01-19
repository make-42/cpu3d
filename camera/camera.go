package camera

import "github.com/make-42/cpu3d/utils"

type Camera struct {
	FocalLength float64 // m

	SensorSize utils.PlaneCoords
}

func EdgesWorldCoordsToSensorCoords(edges *[]utils.SpaceEdge, camera *Camera) []utils.CameraLine {
	cameraLines := make([]utils.CameraLine, len(*edges))
	for i, edge := range *edges {
		cameraLines[i].A = WorldCoordsToSensorCoords(&edge.A, camera)
		cameraLines[i].B = WorldCoordsToSensorCoords(&edge.B, camera)
	}
	return cameraLines
}

func PointCloudWorldCoordsToSensorCoords(pcCoords *[]utils.SpaceCoords, camera *Camera) []utils.CameraCoords {
	cameraCoords := make([]utils.CameraCoords, len(*pcCoords))
	for i, coords := range *pcCoords {
		cameraCoords[i] = WorldCoordsToSensorCoords(&coords, camera)
	}
	return cameraCoords
}

func WorldCoordsToSensorCoords(coords *utils.SpaceCoords, camera *Camera) utils.CameraCoords {
	return utils.CameraCoords{
		X: (coords.X*camera.FocalLength/coords.Z)/camera.SensorSize.X + 0.5,
		Y: 0.5 - (coords.Y*camera.FocalLength/coords.Z)/camera.SensorSize.Y,
	}
}
