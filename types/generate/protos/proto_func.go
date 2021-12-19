package protos

import (
	"time"

	"sungora/lib/typ"

	"github.com/shopspring/decimal"
	"github.com/volatiletech/null/v8"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// type conversion support functions
// вспомогательные функции реализующие уникальную обработку свойств для определяемых рабочих типов

// pbFromNullBytes перевод из примитива grpc в рабочий тип
func pbFromNullBytes(b []byte) null.Bytes {
	return null.BytesFrom(b)
}

// pbFromNullJSON перевод из примитива grpc в рабочий тип
func pbFromNullJSON(b []byte) null.JSON {
	return null.JSONFrom(b)
}

// pbFromNullString перевод из примитива grpc в рабочий
func pbFromNullString(s string) null.String {
	if s == "" {
		return null.String{}
	}
	return null.StringFrom(s)
}

// pbFromNullTime перевод из примитива grpc в рабочий тип
func pbFromNullTime(d *timestamppb.Timestamp) null.Time {
	return null.TimeFrom(d.AsTime())
}

// pbToTime перевод в примитив grpc из рабочего типа
func pbToTime(d time.Time) *timestamppb.Timestamp {
	return timestamppb.New(d)
}

// pbFromTime перевод из примитива grpc в рабочий тип
func pbFromTime(d *timestamppb.Timestamp) time.Time {
	return d.AsTime()
}

// pbFromDecimal перевод из примитива grpc в рабочий тип
func pbFromDecimal(v string) decimal.Decimal {
	d, _ := decimal.NewFromString(v)
	return d
}

func pbToUUIDS(list []typ.UUID) []string {
	uu := make([]string, len(list))
	for i := range list {
		uu[i] = list[i].String()
	}
	return uu
}

func pbFromUUIDS(list []string) []typ.UUID {
	uu := make([]typ.UUID, len(list))
	for i := range list {
		uu[i] = typ.UUIDMustParse(list[i])
	}
	return uu
}
