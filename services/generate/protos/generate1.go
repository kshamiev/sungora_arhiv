package protos

import (
	"log"
	"os"
)

func Generate1(dir string, md string, pb string) {
	_ = os.Mkdir(dir+"/"+md, 0700)
	_ = os.Mkdir(dir+"/"+pb, 0700)

	for i := range FileListCustom {
		_ = os.Remove(dir + "/" + FileListCustom[i])
		fi, err := os.Stat(dir + "/" + md + "/" + FileListCustom[i])
		if err == nil && fi.Mode().IsRegular() {
			if err := os.Rename(
				dir+"/"+md+"/"+FileListCustom[i],
				dir+"/"+FileListCustom[i],
			); err != nil {
				log.Fatal(err)
			}
		}
	}
}
