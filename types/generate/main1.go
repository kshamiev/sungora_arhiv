package main

import (
	"log"
	"os"

	"sungora/types/generate/protos"
)

func main() {
	dir, md, _ := protos.Init()

	for i := range protos.FileListCustom {
		fi, err := os.Stat(dir + "/" + protos.FileListCustom[i])
		if err == nil && fi.Mode().IsRegular() {
			if err := os.Rename(
				dir+"/"+protos.FileListCustom[i],
				dir+"/"+md+"/"+protos.FileListCustom[i],
			); err != nil {
				log.Fatal(err)
			}
		}
	}
}
