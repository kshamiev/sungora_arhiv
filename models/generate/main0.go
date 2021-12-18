package main

import (
	"log"
	"os"

	"sungora/models/generate/protos"
)

func main() {
	dir, md, pb := protos.Init()

	_ = os.Mkdir(dir+"/"+md, 0700)
	_ = os.Mkdir(dir+"/"+pb, 0700)

	for i := range protos.FileListCustom {
		_ = os.Remove(dir + "/" + protos.FileListCustom[i])
		fi, err := os.Stat(dir + "/" + md + "/" + protos.FileListCustom[i])
		if err == nil && fi.Mode().IsRegular() {
			if err := os.Rename(
				dir+"/"+md+"/"+protos.FileListCustom[i],
				dir+"/"+protos.FileListCustom[i],
			); err != nil {
				log.Fatal(err)
			}
		}
	}
}
