package attributes

import (
	"log"
	"os"
)

// GetAttribute get attribute of file
func GetAttribute(f os.File, attribute string) string {
	attribs, err := f.Stat()
	checkError(err)

}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Error %v", err)
	}
}
