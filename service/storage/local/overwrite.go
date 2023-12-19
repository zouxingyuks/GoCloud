package local

import (
	"github.com/pkg/errors"
	"os"
	"path"
	"strings"
	"time"
)

// Check 检查文件是否存在，如果存在则根据 overwrite 参数决定是否覆盖
func (l *Store) Check(overwrite bool, filepath string) (string, error) {
	// 检查文件是否存在
	if _, err := os.Stat(filepath); err == nil {
		// 文件存在
		if !overwrite {
			// 如果不覆盖，则生成一个新的文件名
			filepath = generateNewFileName(filepath)
		}
	} else if !os.IsNotExist(err) {
		// 发生了其他错误
		return "", errors.Wrap(err, "error checking if file exists in Store")
	}
	return filepath, nil
}

// generateNewFileName 生成一个新的文件名，避免与现有文件冲突
func generateNewFileName(originalPath string) string {
	// 将文件路径分解为目录和文件名
	dir, file := path.Split(originalPath)

	// 获取文件的扩展名
	ext := path.Ext(file)

	// 从文件名中移除扩展名，得到纯文件名
	name := strings.TrimSuffix(file, ext)

	// 生成一个时间戳，格式为年月日时分秒（如 20230102150405）
	timestamp := time.Now().Format("20060102150405")

	// 将纯文件名、时间戳和扩展名拼接，形成新的文件名
	// 例如，如果原文件名为 'example.txt'，新文件名可能是 'example_20230102150405.txt'
	newName := name + "_" + timestamp + ext

	// 将目录和新文件名拼接，形成完整的新文件路径
	return path.Join(dir, newName)
}
