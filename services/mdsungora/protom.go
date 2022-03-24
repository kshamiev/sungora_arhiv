package mdsungora

import (
	"sungora/services/pbsungora"
	"time"

	"github.com/google/uuid"
)

func NewGooseDBVersionToProto(tt *GooseDBVersion) *pbsungora.GooseDBVersion {
	if tt == nil {
		return nil
	}
	return &pbsungora.GooseDBVersion{
		Id:        int64(tt.ID),
		VersionId: tt.VersionID,
		IsApplied: tt.IsApplied,
		Tstamp:    pbToTime(tt.Tstamp.Time),
	}
}

func NewGooseDBVersionSliceToProto(tt []*GooseDBVersion) []*pbsungora.GooseDBVersion {
	res := make([]*pbsungora.GooseDBVersion, len(tt))
	for i := range tt {
		res[i] = NewGooseDBVersionToProto(tt[i])
	}
	return res
}

func NewGooseDBVersionFromProto(proto *pbsungora.GooseDBVersion) *GooseDBVersion {
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

func NewGooseDBVersionSliceFromProto(list []*pbsungora.GooseDBVersion) []*GooseDBVersion {
	res := make([]*GooseDBVersion, len(list))
	for i := range list {
		res[i] = NewGooseDBVersionFromProto(list[i])
	}
	return res
}

func NewMinioToProto(tt *Minio) *pbsungora.Minio {
	if tt == nil {
		return nil
	}
	return &pbsungora.Minio{
		Id:        tt.ID.String(),
		Bucket:    tt.Bucket,
		ObjectId:  tt.ObjectID,
		Name:      tt.Name,
		FileType:  tt.FileType,
		FileSize:  int64(tt.FileSize),
		Label:     tt.Label.JSON,
		UserLogin: tt.UserLogin,
		CreatedAt: pbToTime(tt.CreatedAt),
		IsConfirm: tt.IsConfirm,
	}
}

func NewMinioSliceToProto(tt []*Minio) []*pbsungora.Minio {
	res := make([]*pbsungora.Minio, len(tt))
	for i := range tt {
		res[i] = NewMinioToProto(tt[i])
	}
	return res
}

func NewMinioFromProto(proto *pbsungora.Minio) *Minio {
	if proto == nil {
		return nil
	}
	return &Minio{
		ID:        uuid.MustParse(proto.Id),
		Bucket:    proto.Bucket,
		ObjectID:  proto.ObjectId,
		Name:      proto.Name,
		FileType:  proto.FileType,
		FileSize:  int(proto.FileSize),
		Label:     pbFromNullJSON(proto.Label),
		UserLogin: proto.UserLogin,
		CreatedAt: pbFromTime(proto.CreatedAt),
		IsConfirm: proto.IsConfirm,
	}
}

func NewMinioSliceFromProto(list []*pbsungora.Minio) []*Minio {
	res := make([]*Minio, len(list))
	for i := range list {
		res[i] = NewMinioFromProto(list[i])
	}
	return res
}

func NewOrderToProto(tt *Order) *pbsungora.Order {
	if tt == nil {
		return nil
	}
	return &pbsungora.Order{
		Id:        tt.ID,
		UserId:    tt.UserID.Int64,
		Number:    int64(tt.Number),
		Status:    tt.Status,
		CreatedAt: pbToTime(tt.CreatedAt),
		UpdatedAt: pbToTime(tt.UpdatedAt),
		DeletedAt: pbToTime(tt.DeletedAt.Time),
	}
}

func NewOrderSliceToProto(tt []*Order) []*pbsungora.Order {
	res := make([]*pbsungora.Order, len(tt))
	for i := range tt {
		res[i] = NewOrderToProto(tt[i])
	}
	return res
}

func NewOrderFromProto(proto *pbsungora.Order) *Order {
	if proto == nil {
		return nil
	}
	return &Order{
		ID:        proto.Id,
		UserID:    pbFromNullInt64(proto.UserId),
		Number:    int(proto.Number),
		Status:    proto.Status,
		CreatedAt: pbFromTime(proto.CreatedAt),
		UpdatedAt: pbFromTime(proto.UpdatedAt),
		DeletedAt: pbFromNullTime(proto.DeletedAt),
	}
}

func NewOrderSliceFromProto(list []*pbsungora.Order) []*Order {
	res := make([]*Order, len(list))
	for i := range list {
		res[i] = NewOrderFromProto(list[i])
	}
	return res
}

func NewRoleToProto(tt *Role) *pbsungora.Role {
	if tt == nil {
		return nil
	}
	return &pbsungora.Role{
		Id:          tt.ID,
		Code:        tt.Code,
		Description: tt.Description,
	}
}

func NewRoleSliceToProto(tt []*Role) []*pbsungora.Role {
	res := make([]*pbsungora.Role, len(tt))
	for i := range tt {
		res[i] = NewRoleToProto(tt[i])
	}
	return res
}

func NewRoleFromProto(proto *pbsungora.Role) *Role {
	if proto == nil {
		return nil
	}
	return &Role{
		ID:          proto.Id,
		Code:        proto.Code,
		Description: proto.Description,
	}
}

func NewRoleSliceFromProto(list []*pbsungora.Role) []*Role {
	res := make([]*Role, len(list))
	for i := range list {
		res[i] = NewRoleFromProto(list[i])
	}
	return res
}

func NewUserToProto(tt *User) *pbsungora.User {
	if tt == nil {
		return nil
	}
	return &pbsungora.User{
		Id:          tt.ID,
		Login:       tt.Login,
		Description: tt.Description.String,
		Price:       tt.Price.String(),
		SummaOne:    tt.SummaOne,
		SummaTwo:    tt.SummaTwo,
		Cnt:         int64(tt.CNT),
		Cnt2:        int32(tt.CNT2),
		Cnt4:        int64(tt.CNT4),
		Cnt8:        tt.CNT8,
		ShardingId:  tt.ShardingID.String(),
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

func NewUserSliceToProto(tt []*User) []*pbsungora.User {
	res := make([]*pbsungora.User, len(tt))
	for i := range tt {
		res[i] = NewUserToProto(tt[i])
	}
	return res
}

func NewUserFromProto(proto *pbsungora.User) *User {
	if proto == nil {
		return nil
	}
	return &User{
		ID:          proto.Id,
		Login:       proto.Login,
		Description: pbFromNullString(proto.Description),
		Price:       pbFromDecimal(proto.Price),
		SummaOne:    proto.SummaOne,
		SummaTwo:    proto.SummaTwo,
		CNT:         int(proto.Cnt),
		CNT2:        int16(proto.Cnt2),
		CNT4:        int(proto.Cnt4),
		CNT8:        proto.Cnt8,
		ShardingID:  uuid.MustParse(proto.ShardingId),
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

func NewUserSliceFromProto(list []*pbsungora.User) []*User {
	res := make([]*User, len(list))
	for i := range list {
		res[i] = NewUserFromProto(list[i])
	}
	return res
}
