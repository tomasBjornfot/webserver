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
	// prepare the STL files
	deck, _ := prepare_stl("c:\\tmp\\testfile.stl")
	// calculating the cross sections
    cs_deck := deck.CalculateCrossSections(settings.YresFine, 1.0)
    
    //fmt.Println(settings)
    fmt.Println(deck.No_tri)
    fmt.Println(cs_deck.No_rows)
    fmt.Println(cs_deck.No_cols[:cs_deck.No_rows])
    j := 50
    for i := range(cs_deck.X[j][:cs_deck.No_cols[j]]) {
		fmt.Println(cs_deck.X[j][i], cs_deck.Y[j][i], cs_deck.Z[j][i])
	}
}
