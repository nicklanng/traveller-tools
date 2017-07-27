package main

import (
	"github.com/nicklanng/traveller-tools/pkg/app/views"
)

func main() {
	views.InitializeUI()

	// file, e := ioutil.ReadFile("data/sectors/spinward_marches.json")
	// if e != nil {
	// 	fmt.Printf("File error: %v\n", e)
	// 	os.Exit(1)
	// }
	//
	// var s worlds.Sector
	// err := json.Unmarshal(file, &s)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// fmt.Println(s.String())
}
