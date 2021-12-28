package mdsun

import (
	"sungora/lib/typ"
	"sungora/services/pbsun"
	"time"
)

func NewGooseDBVersionToProto(tt *GooseDBVersion) *pbsun.GooseDBVersion {
	if tt == nil {
		return nil
	}
	return &pbsun.GooseDBVersion{
		Id:        int64(tt.ID),
		VersionId: tt.VersionID,
		IsApplied: tt.IsApplied,
		Tstamp:    pbToTime(tt.Tstamp.Time),
	}
}

func NewGooseDBVersionSliceToProto(tt []*GooseDBVersion) []*pbsun.GooseDBVersion {
	res := make([]*pbsun.GooseDBVersion, len(tt))
	for i := range tt {
		res[i] = NewGooseDBVersionToProto(tt[i])
	}
	return res
}

func NewGooseDBVersionFromProto(proto *pbsun.GooseDBVersion) *GooseDBVersion {
	if proto == nil {
		return nil
	}
	return &GooseDBVersion{
		ID:        int(proto.Id),
		VersionID: proto.VersionId,
		IsApplied: proto.IsApplied,
		Tstamp:    pbFromNullTime(proto.Tstamp),
	}
}

func NewGooseDBVersionSliceFromProto(list []*pbsun.GooseDBVersion) []*GooseDBVersion {
	res := make([]*GooseDBVersion, len(list))
	for i := range list {
		res[i] = NewGooseDBVersionFromProto(list[i])
	}
	return res
}

func NewOrderToProto(tt *Order) *pbsun.Order {
	if tt == nil {
		return nil
	}
	return &pbsun.Order{
		Id:        tt.ID.String(),
		UserId:    tt.UserID.String(),
		Number:    int64(tt.Number),
		Status:    tt.Status,
		CreatedAt: pbToTime(tt.CreatedAt),
		UpdatedAt: pbToTime(tt.UpdatedAt),
		DeletedAt: pbToTime(tt.DeletedAt.Time),
	}
}

func NewOrderSliceToProto(tt []*Order) []*pbsun.Order {
	res := make([]*pbsun.Order, len(tt))
	for i := range tt {
		res[i] = NewOrderToProto(tt[i])
	}
	return res
}

func NewOrderFromProto(proto *pbsun.Order) *Order {
	if proto == nil {
		return nil
	}
	return &Order{
		ID:        typ.UUIDMustParse(proto.Id),
		UserID:    typ.UUIDMustParse(proto.UserId),
		Number:    int(proto.Number),
		Status:    proto.Status,
		CreatedAt: pbFromTime(proto.CreatedAt),
		UpdatedAt: pbFromTime(proto.UpdatedAt),
		DeletedAt: pbFromNullTime(proto.DeletedAt),
	}
}

func NewOrderSliceFromProto(list []*pbsun.Order) []*Order {
	res := make([]*Order, len(list))
	for i := range list {
		res[i] = NewOrderFromProto(list[i])
	}
	return res
}

func NewRoleToProto(tt *Role) *pbsun.Role {
	if tt == nil {
		return nil
	}
	return &pbsun.Role{
		Id:          tt.ID.String(),
		Code:        tt.Code,
		Description: tt.Description,
	}
}

func NewRoleSliceToProto(tt []*Role) []*pbsun.Role {
	res := make([]*pbsun.Role, len(tt))
	for i := range tt {
		res[i] = NewRoleToProto(tt[i])
	}
	return res
}

func NewRoleFromProto(proto *pbsun.Role) *Role {
	if proto == nil {
		return nil
	}
	return &Role{
		ID:          typ.UUIDMustParse(proto.Id),
		Code:        proto.Code,
		Description: proto.Description,
	}
}

func NewRoleSliceFromProto(list []*pbsun.Role) []*Role {
	res := make([]*Role, len(list))
	for i := range list {
		res[i] = NewRoleFromProto(list[i])
	}
	return res
}

func NewUserToProto(tt *User) *pbsun.User {
	if tt == nil {
		return nil
	}
	return &pbsun.User{
		Id:          tt.ID.String(),
		Login:       tt.Login,
		Description: tt.Description.String,
		Price:       tt.Price.String(),
		SummaOne:    tt.SummaOne,
		SummaTwo:    tt.SummaTwo,
		Cnt:         int64(tt.CNT),
		Cnt2:        int32(tt.CNT2),
		Cnt4:        int64(tt.CNT4),
		Cnt8:        tt.CNT8,
		IsOnline:    tt.IsOnline,
		Metrika:     tt.Metrika.JSON,
		Duration:    tt.Duration.Nanoseconds(),
		Data:        tt.Data.Bytes,
		Alias:       tt.Alias,
		CreatedAt:   pbToTime(tt.CreatedAt),
		UpdatedAt:   pbToTime(tt.UpdatedAt),
		DeletedAt:   pbToTime(tt.DeletedAt.Time),
	}
}

func NewUserSliceToProto(tt []*User) []*pbsun.User {
	res := make([]*pbsun.User, len(tt))
	for i := range tt {
		res[i] = NewUserToProto(tt[i])
	}
	return res
}

func NewUserFromProto(proto *pbsun.User) *User {
	if proto == nil {
		return nil
	}
	return &User{
		ID:          typ.UUIDMustParse(proto.Id),
		Login:       proto.Login,
		Description: pbFromNullString(proto.Description),
		Price:       pbFromDecimal(proto.Price),
		SummaOne:    proto.SummaOne,
		SummaTwo:    proto.SummaTwo,
		CNT:         int(proto.Cnt),
		CNT2:        int16(proto.Cnt2),
		CNT4:        int(proto.Cnt4),
		CNT8:        proto.Cnt8,
		IsOnline:    proto.IsOnline,
		Metrika:     pbFromNullJSON(proto.Metrika),
		Duration:    time.Duration(proto.Duration),
		Data:        pbFromNullBytes(proto.Data),
		Alias:       proto.Alias,
		CreatedAt:   pbFromTime(proto.CreatedAt),
		UpdatedAt:   pbFromTime(proto.UpdatedAt),
		DeletedAt:   pbFromNullTime(proto.DeletedAt),
	}
}

func NewUserSliceFromProto(list []*pbsun.User) []*User {
	res := make([]*User, len(list))
	for i := range list {
		res[i] = NewUserFromProto(list[i])
	}
	return res
}
