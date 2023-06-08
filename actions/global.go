package actions

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Type kind of dispatch action

type Type string

// Area means select-boxes .Multi-boxes can be specified
// if Areas is empty ,all ofthe select-boxes will be deleted
// The first area.

type Areas struct {

	// BrushType Optional : 'polygon' 'rect' 'lineX' ' lineY'
	BrushType string `json:"brushType,omitempty"`

	//CoordRange Only for "coordinate system area" define the area with the coordinates
	CoordRange []string `json:"corrdRange ,omitempty"`

	// XAxisIndex Assigns which of the xAxisIndex can sue Area selecting
	XAxisIndex interface{} `json:"xAxisIndex,omitempty"`
}
