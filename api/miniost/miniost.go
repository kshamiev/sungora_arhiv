package miniost

import (
	"context"
	"errors"
	"io"
	"os"
	"time"

	"sungora/lib/errs"
	"sungora/lib/response"
	"sungora/lib/storage"
	"sungora/lib/typ"
	"sungora/services/mdsample"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type MinioST struct {
	st  storage.Face
	dir string
}

func NewMinioST(st storage.Face, dir string) *MinioST {
	return &MinioST{
		st:  st,
		dir: dir,
	}
}

func (self *MinioST) UploadRequest(rw *response.Response, bucket string, objID typ.UUID) ([]*mdsample.MinioST, error) {
	fileData, _, err := rw.UploadBuffer()
	if err != nil {
		return nil, err
	}
	us, err := rw.GetUser()
	if err != nil {
		return nil, err
	}

	res := make([]*mdsample.MinioST, 0, len(fileData))
	for fName, buf := range fileData {
		if buf.Len() == 0 {
			return nil, errs.NewBadRequest(errors.New("file size zero"), "ошибка получения файла")
		}
		stM := &mdsample.MinioST{}
		stM.Bucket = bucket
		stM.ObjectID = objID
		stM.Name = fName
		stM.FileType = rw.Request.Header.Get("content-type")
		stM.FileSize = buf.Len()
		stM.UserLogin = us.Login
		stM.CreatedAt = time.Now()
		err = PutFile(rw.Request.Context(), stM.Bucket, stM.ObjectID.String(), buf, int64(stM.FileSize))
		if err != nil {
			return nil, err
		}

		err = stM.Insert(rw.Request.Context(), self.st.DB(), boil.Infer())
		if err != nil {
			return nil, errs.NewBadRequest(err, "couldn't insert file info")
		}

		res = append(res, stM)
	}

	return res, nil
}

func (self *MinioST) Confirm(ctx context.Context, obj *mdsample.MinioST) error {
	obj.IsConfirm = true
	if _, err := obj.Update(ctx, self.st.DB(), boil.Whitelist(mdsample.MinioSTColumns.IsConfirm)); err != nil {
		return errs.NewBadRequest(err)
	}
	return nil
}

func (self *MinioST) SaveFS(ctx context.Context, obj *mdsample.MinioST) error {
	if err := os.MkdirAll(self.dir, 0777); err != nil {
		return errs.NewBadRequest(err, "ошибка создания хранилища")
	}
	data, err := GetFile(ctx, obj.Bucket, obj.ObjectID.String())
	if err != nil {
		return err
	}
	fp, err := os.OpenFile(self.dir+"/"+obj.Name, os.O_RDWR|os.O_CREATE, 0x0755)
	if err != nil {
		return errs.NewBadRequest(err)
	}
	if _, err := io.Copy(fp, data); err != nil {
		return errs.NewBadRequest(err)
	}
	return fp.Close()
}

func (self *MinioST) GetFiles(ctx context.Context, bucket string, objID typ.UUID) (mdsample.MinioSTSlice, error) {
	return mdsample.MinioSTS(
		mdsample.MinioSTWhere.Bucket.EQ(bucket),
		mdsample.MinioSTWhere.ObjectID.EQ(objID),
		qm.OrderBy(mdsample.MinioSTColumns.CreatedAt+" DESC"),
		qm.Offset(0), qm.Limit(100),
	).All(ctx, self.st.DB())
}

func (self *MinioST) GetFilesBucket(ctx context.Context, bucket string) (mdsample.MinioSTSlice, error) {
	return mdsample.MinioSTS(
		mdsample.MinioSTWhere.Bucket.EQ(bucket),
		qm.OrderBy(mdsample.MinioSTColumns.CreatedAt+" DESC"),
		qm.Offset(0), qm.Limit(100),
	).All(ctx, self.st.DB())
}

func (self *MinioST) RemoveNotConfirm(ctx context.Context) error {
	list, err := mdsample.MinioSTS(
		mdsample.MinioSTWhere.IsConfirm.EQ(false),
	).All(ctx, self.st.DB())
	if err != nil {
		return errs.NewBadRequest(err)
	}
	for i := range list {
		if err := RemoveFile(list[i].Bucket, list[i].ObjectID.String()); err != nil {
			return err
		}
		if _, err := list[i].Delete(ctx, self.st.DB()); err != nil {
			return errs.NewBadRequest(err)
		}
	}
	return nil
}
