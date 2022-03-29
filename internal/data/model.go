package data

import (
	"context"
	"errors"
	"io"
	"os"
	"time"

	"sample/lib/app/response"
	"sample/lib/errs"
	"sample/lib/minio"
	"sample/lib/storage"
	"sample/services/mdsample"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Model struct {
	st  storage.Face
	dir string
}

func NewModel(st storage.Face, dir string) *Model {
	return &Model{
		st:  st,
		dir: dir,
	}
}

func (mm *Model) UploadRequest(rw *response.Response, bucket string) (mdsample.MinioSlice, error) {
	fileData, _, err := rw.UploadBuffer()
	if err != nil {
		return nil, err
	}
	us, err := rw.GetUser()
	if err != nil {
		return nil, err
	}

	res := make([]*mdsample.Minio, 0, len(fileData))
	for fName, buf := range fileData {
		if buf.Len() == 0 {
			return nil, errs.New(errors.New("file size zero"), "ошибка получения файла")
		}
		stM := &mdsample.Minio{}
		stM.Bucket = bucket
		stM.Name = fName
		stM.FileType = rw.Request.Header.Get("content-type")
		stM.FileSize = buf.Len()
		stM.UserLogin = us.Login
		stM.CreatedAt = time.Now()

		err = stM.Insert(rw.Request.Context(), mm.st.DB(), boil.Infer())
		if err != nil {
			return nil, errs.New(err, "couldn't insert file info")
		}

		err = minio.PutFile(rw.Request.Context(), stM.Bucket, stM.ID.String(), buf, int64(stM.FileSize))
		if err != nil {
			return nil, err
		}

		res = append(res, stM)
	}

	return res, nil
}

func (mm *Model) Confirm(ctx context.Context, obj *mdsample.Minio) error {
	obj.IsConfirm = true
	if _, err := obj.Update(ctx, mm.st.DB(), boil.Whitelist(
		mdsample.MinioColumns.IsConfirm,
		mdsample.MinioColumns.ObjectID,
	)); err != nil {
		return errs.New(err)
	}
	return nil
}

func (mm *Model) SaveFS(ctx context.Context, obj *mdsample.Minio) error {
	if err := os.MkdirAll(mm.dir, 0o777); err != nil {
		return errs.New(err, "ошибка создания хранилища")
	}
	data, err := minio.GetFile(ctx, obj.Bucket, obj.ID.String())
	if err != nil {
		return err
	}
	fp, err := os.OpenFile(mm.dir+"/"+obj.Name, os.O_RDWR|os.O_CREATE, 0x0755)
	if err != nil {
		return errs.New(err)
	}
	if _, err := io.Copy(fp, data); err != nil {
		return errs.New(err)
	}
	return fp.Close()
}

func (mm *Model) GetFiles(ctx context.Context, bucket string, objID int64) (mdsample.MinioSlice, error) {
	return mdsample.Minios(
		mdsample.MinioWhere.Bucket.EQ(bucket),
		mdsample.MinioWhere.ObjectID.EQ(objID),
		qm.OrderBy(mdsample.MinioColumns.CreatedAt+" DESC"),
		qm.Offset(0), qm.Limit(100),
	).All(ctx, mm.st.DB())
}

func (mm *Model) GetFilesBucket(ctx context.Context, bucket string) (mdsample.MinioSlice, error) {
	return mdsample.Minios(
		mdsample.MinioWhere.Bucket.EQ(bucket),
		qm.OrderBy(mdsample.MinioColumns.CreatedAt+" DESC"),
		qm.Offset(0), qm.Limit(100),
	).All(ctx, mm.st.DB())
}

func (self *Model) RemoveNotConfirm(ctx context.Context) error {
	list, err := mdsample.Minios(
		mdsample.MinioWhere.IsConfirm.EQ(false),
	).All(ctx, self.st.DB())
	if err != nil {
		return errs.New(err)
	}
	for i := range list {
		if err := minio.RemoveFile(list[i].Bucket, list[i].ID.String()); err != nil {
			return err
		}
		if _, err := list[i].Delete(ctx, self.st.DB()); err != nil {
			return errs.New(err)
		}
	}
	return nil
}
