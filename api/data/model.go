package data

import (
	"context"
	"errors"
	"io"
	"os"
	"time"

	"sungora/lib/errs"
	"sungora/lib/minio"
	"sungora/lib/response"
	"sungora/lib/storage"
	"sungora/lib/typ"
	"sungora/services/mdsungora"

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

func (mm *Model) UploadFile(filePath, bucket string, objID typ.UUID) (*mdsungora.Minio, error) {
	return nil, nil
}

func (mm *Model) UploadRequest(rw *response.Response, bucket string, objID typ.UUID) (mdsungora.MinioSlice, error) {
	fileData, _, err := rw.UploadBuffer()
	if err != nil {
		return nil, err
	}
	us, err := rw.GetUserTest()
	if err != nil {
		return nil, err
	}

	res := make([]*mdsungora.Minio, 0, len(fileData))
	for fName, buf := range fileData {
		if buf.Len() == 0 {
			return nil, errs.NewBadRequest(errors.New("file size zero"), "ошибка получения файла")
		}
		stM := &mdsungora.Minio{}
		stM.Bucket = bucket
		stM.ObjectID = objID
		stM.Name = fName
		stM.FileType = rw.Request.Header.Get("content-type")
		stM.FileSize = buf.Len()
		stM.UserLogin = us.Login
		stM.CreatedAt = time.Now()
		err = minio.PutFile(rw.Request.Context(), stM.Bucket, stM.ObjectID.String(), buf, int64(stM.FileSize))
		if err != nil {
			return nil, err
		}

		err = stM.Insert(rw.Request.Context(), mm.st.DB(), boil.Infer())
		if err != nil {
			return nil, errs.NewBadRequest(err, "couldn't insert file info")
		}

		res = append(res, stM)
	}

	return res, nil
}

func (mm *Model) Confirm(ctx context.Context, obj *mdsungora.Minio) error {
	obj.IsConfirm = true
	if _, err := obj.Update(ctx, mm.st.DB(), boil.Whitelist(mdsungora.MinioColumns.IsConfirm)); err != nil {
		return errs.NewBadRequest(err)
	}
	return nil
}

func (mm *Model) SaveFS(ctx context.Context, obj *mdsungora.Minio) error {
	if err := os.MkdirAll(mm.dir, 0777); err != nil {
		return errs.NewBadRequest(err, "ошибка создания хранилища")
	}
	data, err := minio.GetFile(ctx, obj.Bucket, obj.ObjectID.String())
	if err != nil {
		return err
	}
	fp, err := os.OpenFile(mm.dir+"/"+obj.Name, os.O_RDWR|os.O_CREATE, 0x0755)
	if err != nil {
		return errs.NewBadRequest(err)
	}
	if _, err := io.Copy(fp, data); err != nil {
		return errs.NewBadRequest(err)
	}
	return fp.Close()
}

func (mm *Model) GetFiles(ctx context.Context, bucket string, objID typ.UUID) (mdsungora.MinioSlice, error) {
	return mdsungora.Minios(
		mdsungora.MinioWhere.Bucket.EQ(bucket),
		mdsungora.MinioWhere.ObjectID.EQ(objID),
		qm.OrderBy(mdsungora.MinioColumns.CreatedAt+" DESC"),
		qm.Offset(0), qm.Limit(100),
	).All(ctx, mm.st.DB())
}

func (mm *Model) GetFilesBucket(ctx context.Context, bucket string) (mdsungora.MinioSlice, error) {
	return mdsungora.Minios(
		mdsungora.MinioWhere.Bucket.EQ(bucket),
		qm.OrderBy(mdsungora.MinioColumns.CreatedAt+" DESC"),
		qm.Offset(0), qm.Limit(100),
	).All(ctx, mm.st.DB())
}

func (self *Model) RemoveNotConfirm(ctx context.Context) error {
	list, err := mdsungora.Minios(
		mdsungora.MinioWhere.IsConfirm.EQ(false),
	).All(ctx, self.st.DB())
	if err != nil {
		return errs.NewBadRequest(err)
	}
	for i := range list {
		if err := minio.RemoveFile(list[i].Bucket, list[i].ObjectID.String()); err != nil {
			return err
		}
		if _, err := list[i].Delete(ctx, self.st.DB()); err != nil {
			return errs.NewBadRequest(err)
		}
	}
	return nil
}
