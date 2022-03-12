package protos

import (
	"log"
	"os"
)

func Generate1(dir, md, pb string) {
	_ = os.Mkdir(dir+"/"+md, 0o700)
	_ = os.Mkdir(dir+"/"+pb, 0o700)

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
