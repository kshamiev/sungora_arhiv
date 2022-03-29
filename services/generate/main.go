package main

import (
	"sample/services/generate/protos"
)

func main() {
	step, dir, md, pb := protos.Init()
	switch step {
	case 1:
		protos.Generate1(dir, md, pb)
	case 2:
		protos.Generate2(dir, md)
	case 3:
		protos.Generate3(dir, md)
	case 4:
		protos.Generate4(dir, md, pb)
	}
}
