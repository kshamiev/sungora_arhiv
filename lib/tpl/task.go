package tpl

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"sungora/lib/errs"
)

const TaskTemplateParseName = "TaskTemplateParse"

// TaskTemplateParse Поддержание html шаблонов в актуальном состоянии.
type TaskTemplateParse struct {
	dir string
}

func NewTaskTemplateParse(dir string) *TaskTemplateParse {
	return &TaskTemplateParse{dir: dir}
}

func (task *TaskTemplateParse) Action(ctx context.Context) error {
	dir := task.dir
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Name() == "temp" {
			return filepath.SkipDir
		}
		if info.IsDir() || filepath.Ext(path) != ".html" {
			return nil
		}
		if _, ok := tplStoreInfo[path]; !ok || info.ModTime().Unix() != tplStoreInfo[path].Unix() {
			if err := task.parseFiles(dir, path); err != nil {
				return err
			}
			tplStoreInfo[path] = info.ModTime()
		}
		return nil
	})
	return err
}

// parseFiles компиляция html шаблонов
func (task *TaskTemplateParse) parseFiles(dir, path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return errs.NewBadRequest(err)
	}

	index := strings.ReplaceAll(path, dir+"/", "")

	tpl, err := template.New(index).Funcs(functions).Parse(string(data))
	if err != nil {
		return errs.NewBadRequest(err)
	}

	tplStore[index] = tpl
	return nil
}

func (task *TaskTemplateParse) WaitFor() time.Duration {
	return time.Minute
}

func (task *TaskTemplateParse) Name() string {
	return TaskTemplateParseName
}
