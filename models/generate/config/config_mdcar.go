package config

import (
	"sungora/models/generate/protos"
	"sungora/models/mdcar"
)

func init() {
	protos.GenerateConfig["mdcar"] = []interface{}{
		&mdcar.Area{},
		&mdcar.CarrierMoneyRatio{},
		&mdcar.CarrierMoving{},
		&mdcar.CarrierPacking{},
		&mdcar.CarrierTransport{},
		&mdcar.Carrier{},
		&mdcar.CarriersAPI{},
		&mdcar.CarriersFD{},
		&mdcar.CarriersUpload{},
		&mdcar.City{},
		&mdcar.CombinedMoving{},
		&mdcar.CombinedPacking{},
		&mdcar.CombinedPrice{},
		&mdcar.CombinedRate{},
		&mdcar.DedicatedDistance{},
		&mdcar.DedicatedMGH{},
		&mdcar.DedicatedMoving{},
		&mdcar.DedicatedPassMoscow{},
		&mdcar.DedicatedTransport{},
		&mdcar.EgisCache{},
		&mdcar.FederalDistrict{},
		&mdcar.File{},
		&mdcar.GooseDBVersion{},
		&mdcar.MinioST{},
		&mdcar.Packing{},
		&mdcar.ShippedStatus{},
		&mdcar.TariffDelivery{},
		&mdcar.TariffMoving{},
		&mdcar.TariffPacking{},
		&mdcar.TariffPrice{},
	}
}
