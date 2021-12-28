package protos

import (
	"log"
	"os"
)

func Generate2(dir string, md string) {
	for i := range FileListCustom {
		fi, err := os.Stat(dir + "/" + FileListCustom[i])
		if err == nil && fi.Mode().IsRegular() {
			if err := os.Rename(
				dir+"/"+FileListCustom[i],
				dir+"/"+md+"/"+FileListCustom[i],
			); err != nil {
				log.Fatal(err)
			}
		}
	}
}
