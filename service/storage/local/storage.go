package local

import (
	"github.com/pkg/errors"
	"io"
	"os"
	"path"
)

type Storage interface {
	GetRoot() string
	Save(vPath string, reader io.Reader, overwrite bool) (string, error)
	Del(vPath string) error
	Check(overwrite bool, filepath string) (string, error)
}
type Store struct {
	RootPath string
}

func (l *Store) GetRoot() string {
	return l.RootPath
}

func (l *Store) Save(vPath string, reader io.Reader, overwrite bool) (string, error) {

	// 指定文件要保存的路径
	filepath := path.Join(l.GetRoot(), vPath)
	// 检查文件是否存在
	filepath, err := l.Check(overwrite, filepath)
	if err != nil {
		return "", err
	}
	// 创建文件
	out, err := os.Create(filepath)
	_, err = io.Copy(out, reader)
	if err != nil {
		return "", errors.Wrap(err, "Unable to save the file To Store")
	}
	err = out.Close()
	return filepath, err
}
func (l *Store) Del(vPath string) error {
	// 指定文件要保存的路径
	filepath := path.Join(l.GetRoot(), vPath)
	err := os.Remove(filepath)
	return err
}
