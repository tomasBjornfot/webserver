package main

import (
	"github.com/tomasBjornfot/stl"
	"os"
	"log"
	"encoding/json"
	"fmt"
)
/*
 * STRUCTS 
 */
// Strukt för hantering av settings från settings.json filen
type Settings struct {
	MachineRotCenter    [2]float64 `json:"MachineRotCenter"`
	MachineLength       float64    `json:"MachineLength"`
	MachineHolderDepth  float64    `json:"MachineHolderDepth"`
	HomingOffset        [3]float64 `json:"HomingOffset"`
	BlockThickness      float64    `json:"BlockThickness"`
	BlockSize           [3]float64 `json:"BlockSize"`
	ToolRadius          float64    `json:"ToolRadius"`
	ToolCuttingLength   float64    `json:"ToolCuttingLength"`
	ToolShaftLength     float64    `json:"ToolShaftLength"`
	XresRough           float64    `json:"XresRough"`
	XresFine            float64    `json:"XresFine"`
	YresRough           float64    `json:"YresRough"`
	YresFine            float64    `json:"YresFine"`
	FeedrateStringer    float64    `json:"FeedrateStringer"`
	FeedrateMax         float64    `json:"FeedrateMax"`
	FeedrateMin         float64    `json:"FeedrateMin"`
	FeedrateChangeLimit float64    `json:"FeedrateChangeLimit"`
	HandlePos           float64    `json:"HandlePos"`
	HandleWidth         int        `json:"HandleWidth"`
	HandleHeightOffset  float64    `json:"HandleHeightOffset"`
	InFolder            string     `json:"InFolder"`
	CamFolder           string     `json:"CamFolder"`
	OutFolder           string     `json:"OutFolder"`
}
type CrossSection struct {
	X       [][1000]float64
	Y       [][1000]float64
	Z       [][1000]float64
	No_rows int
	No_cols [1000]int
}
/*
 * JSON FUNCTIONS
 */
func read_settings(dir string) *Settings {
	// läser in settingsfilen och skriver på Settings strukten
	file, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	finfo, _ := file.Stat()

	bytes := make([]byte, finfo.Size())
	file.Read(bytes)

	s := new(Settings)
	json.Unmarshal(bytes, &s)
	return s
}
/*
 * FUNCTIONS AND METHODS
 */
 func calc_cs(mesh *stl.Mesh, max_distance float64) {
	y_cs := []float64
	index := 0
	md2 := max_distance*max_distance
	p0 := mesh.Profile[0]
	for (p0[1] < mesh.Profile[len(mesh.Profile)-1][1])){
		// set next point as line edge
		p1 := mesh.Profile[index + 1] 
		// calc if line edge is too far away
		dx := sp[0]-np[0]
		dy := sp[1]-np[1]
		if (dx*dx+dy*dy < md2) {
			// ok, use it
			p0 = p1
			y_cs = append(y_cs, p0[1])
			index ++
			continue
		}
		// nok, take the half distance to edge of line
		p0[1] = p0[1] + (p1[1] - p0[1])/2.0
		k := dy/dx
		m := sp[1] - k*sp[0]
		p0[0] = (y - k)/m
	}
}

/*
 * COMPOSITE FUNCTIONS
 */
func prepare_stl(path string) (*stl.Mesh, *stl.Mesh) {
	board := new(stl.Mesh)
	board.Read(path, 1)
	board.AlignMesh("boardcad")
	board.AlignMeshX()
	board.MoveToCenter2()
	deck, bottom := board.Split()
	bottom.Rotate("y", 180.0)
	return deck, bottom
}
/*
 * MAIN FUNCTION
 */
func main() {
	// reads the settings from JSON file
	settings := read_settings("settings.json")
	// prepare the STL file
	deck, bottom := prepare_stl("c:\\tmp\\test.stl")
	// calculates the deck
	calc_cs(deck, 10.0)
	// 
	fmt.Println(settings.ToolRadius)
	fmt.Println(deck.No_tri)
	fmt.Println(bottom.No_tri)
}
