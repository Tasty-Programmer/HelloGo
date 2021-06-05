package main


import "fmt"
import "24lab.net/testlib"
func main() {

	song := testlib.GetMusic("Alicia Keys")
	println(song)

}


package testlib


var pop map[string]string

func init() {
	pop = make(map[string]string)
	pop["Adele"] = "Hello"
	pop["Alicia Keys"] = "Fallin'"
	pop["John Legend"] = "All of Me"
}

// GetMusic : Popular music by singer (Outer call able)
func GetMusic(singer string) string {
	return pop[singer]
}

func getKeys() {  // inside call able
	for _, kv := range pop {
		fmt.Println(kv)
	}
}