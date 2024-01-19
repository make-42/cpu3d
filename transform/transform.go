package transform

import (
	"math"

	"github.com/make-42/cpu3d/utils"
)

/* Edges */
func EdgesRotateWorldCoords(edges *[]utils.SpaceEdge, rotation *utils.RotationCoords) []utils.SpaceEdge {
	newCoords := make([]utils.SpaceEdge, len(*edges))
	rotationMatrix := GenRotationMatrix(rotation)
	for i, coords := range *edges {
		newCoords[i].A = Rotate(&coords.A, &rotationMatrix)
		newCoords[i].B = Rotate(&coords.B, &rotationMatrix)
	}
	return newCoords
}

func EdgesTranslateWorldCoords(edges *[]utils.SpaceEdge, offset *utils.SpaceCoords) []utils.SpaceEdge {
	newCoords := make([]utils.SpaceEdge, len(*edges))
	for i, coords := range *edges {
		newCoords[i].A = Translate(&coords.A, offset)
		newCoords[i].B = Translate(&coords.B, offset)
	}
	return newCoords
}

func EdgesScaleWorldCoords(edges *[]utils.SpaceEdge, scale float64) []utils.SpaceEdge {
	newCoords := make([]utils.SpaceEdge, len(*edges))
	for i, coords := range *edges {
		newCoords[i].A = Scale(&coords.A, scale)
		newCoords[i].B = Scale(&coords.B, scale)
	}
	return newCoords
}

/* Point clouds */
func PointCloudRotateWorldCoords(pointCloud *[]utils.SpaceCoords, rotation *utils.RotationCoords) []utils.SpaceCoords {
	newCoords := make([]utils.SpaceCoords, len(*pointCloud))
	rotationMatrix := GenRotationMatrix(rotation)
	for i, coords := range *pointCloud {
		newCoords[i] = Rotate(&coords, &rotationMatrix)
	}
	return newCoords
}

func PointCloudTranslateWorldCoords(pointCloud *[]utils.SpaceCoords, offset *utils.SpaceCoords) []utils.SpaceCoords {
	newCoords := make([]utils.SpaceCoords, len(*pointCloud))
	for i, coords := range *pointCloud {
		newCoords[i] = Translate(&coords, offset)
	}
	return newCoords
}

func PointCloudScaleWorldCoords(pointCloud *[]utils.SpaceCoords, scale float64) []utils.SpaceCoords {
	newCoords := make([]utils.SpaceCoords, len(*pointCloud))
	for i, coords := range *pointCloud {
		newCoords[i] = Scale(&coords, scale)
	}
	return newCoords
}

// Generic

func Translate(coords *utils.SpaceCoords, offset *utils.SpaceCoords) utils.SpaceCoords {
	return utils.SpaceCoords{
		X: coords.X + offset.X,
		Y: coords.Y + offset.Y,
		Z: coords.Z + offset.Z,
	}
}

func GenRotationMatrix(rotation *utils.RotationCoords) utils.RotationMatrix {
	return utils.RotationMatrix{
		A: math.Cos(rotation.X) * math.Cos(rotation.Y), B: math.Cos(rotation.X)*math.Sin(rotation.Y)*math.Sin(rotation.Z) - math.Sin(rotation.X)*math.Cos(rotation.Z), C: math.Cos(rotation.X)*math.Sin(rotation.Y)*math.Cos(rotation.Z) + math.Sin(rotation.X)*math.Sin(rotation.Z),
		D: math.Sin(rotation.X) * math.Cos(rotation.Y), E: math.Sin(rotation.X)*math.Sin(rotation.Y)*math.Sin(rotation.Z) + math.Cos(rotation.X)*math.Cos(rotation.Z), F: math.Sin(rotation.X)*math.Sin(rotation.Y)*math.Cos(rotation.Z) - math.Cos(rotation.X)*math.Sin(rotation.Z),
		G: -math.Sin(rotation.Y), H: math.Cos(rotation.Y) * math.Sin(rotation.Z), I: math.Cos(rotation.Y) * math.Cos(rotation.Z),
	}
}

func Rotate(coords *utils.SpaceCoords, rotationMatrix *utils.RotationMatrix) utils.SpaceCoords {
	return utils.SpaceCoords{
		X: rotationMatrix.A*coords.X + rotationMatrix.B*coords.Y + rotationMatrix.C*coords.Z,

		Y: rotationMatrix.D*coords.X + rotationMatrix.E*coords.Y + rotationMatrix.F*coords.Z,

		Z: rotationMatrix.G*coords.X + rotationMatrix.H*coords.Y + rotationMatrix.I*coords.Z,
	}
}

func Scale(coords *utils.SpaceCoords, scale float64) utils.SpaceCoords {
	return utils.SpaceCoords{
		X: coords.X * scale,
		Y: coords.Y * scale,
		Z: coords.Z * scale,
	}
}
