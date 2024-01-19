package model

import (
	"io"
	"os"

	"github.com/make-42/cpu3d/utils"

	"github.com/hschendel/stl"
	"golang.org/x/exp/slices"
)

func ReadSTLFileFromPathToEdges(STLFilePath string) []utils.SpaceEdge {
	STLFile, err := os.Open("assets/gif/komi-san-48.gif")
	utils.CheckError(err)
	defer STLFile.Close()
	return ReadSTLFileToEdges(STLFile)
}

func ReadSTLFileToEdges(STLFile io.ReadSeeker) []utils.SpaceEdge {
	solid, err := stl.ReadAll(STLFile)
	utils.CheckError(err)
	tris := solid.Triangles
	edges := []utils.SpaceEdge{}
	for _, tri := range tris {
		edgesToAdd := [3]utils.SpaceEdge{{
			A: utils.SpaceCoords{X: float64(tri.Vertices[0][0]), Y: float64(tri.Vertices[0][1]), Z: float64(tri.Vertices[0][2])},
			B: utils.SpaceCoords{X: float64(tri.Vertices[1][0]), Y: float64(tri.Vertices[1][1]), Z: float64(tri.Vertices[1][2])},
		}, {
			A: utils.SpaceCoords{X: float64(tri.Vertices[0][0]), Y: float64(tri.Vertices[0][1]), Z: float64(tri.Vertices[0][2])},
			B: utils.SpaceCoords{X: float64(tri.Vertices[2][0]), Y: float64(tri.Vertices[2][1]), Z: float64(tri.Vertices[2][2])},
		}, {
			A: utils.SpaceCoords{X: float64(tri.Vertices[1][0]), Y: float64(tri.Vertices[1][1]), Z: float64(tri.Vertices[1][2])},
			B: utils.SpaceCoords{X: float64(tri.Vertices[2][0]), Y: float64(tri.Vertices[2][1]), Z: float64(tri.Vertices[2][2])},
		}}
		for _, edge := range edgesToAdd {
			if !slices.Contains(edges, edge) {
				edges = append(edges, edge)
			}
		}
	}
	return edges
}
