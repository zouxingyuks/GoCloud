package space

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

// NewSpace 创建新的集合
// collection 集合名称，集合名称是唯一的。
func NewSpace(Name string) error {

	// 1.1 检查集合名称是否为空
	if Name == "" {
		return errors.New("元数据名称不能为空")
	}
	// 2.2 检查集合长度是否超过 255
	if len(Name) > 255 {
		return errors.New("元数据名称长度不能超过 255 个字符")
	}
	// 2.3 检查元数据名称是否合法
	if !CheckName(Name) {
		return errors.New("元数据名称不合法")
	}
	// 2.4 检查元数据是否重复
	result, err := FilesDB().ListCollectionNames(context.Background(), bson.M{
		"name": Name,
	})
	if err != nil {
		return errors.Wrap(err, "判断集合是否存在失败")
	}
	if len(result) != 0 {
		return errors.New("集合已存在")
	}
	// 2.检查元数据名称长度是否合法
	// 写入数据
	err = FilesDB().CreateCollection(context.Background(), Name)
	if err != nil {
		return errors.Wrap(err, "创建集合失败")
	}
	return nil
}

// DeleteSpace 删除指定的集合
func DeleteSpace(Name string) error {
	err := FilesDB().Collection(Name).Drop(context.Background())
	if err != nil {
		return errors.Wrap(err, "删除集合失败")
	}
	return nil
}

// ExistSpace 检查集合是否存在
func ExistSpace(Name string) (bool, error) {
	result, err := FilesDB().ListCollectionNames(context.Background(), bson.M{
		"name": Name,
	})
	if err != nil {
		return false, errors.Wrap(err, "判断集合是否存在失败")
	}
	if len(result) != 0 {
		return true, nil
	}
	return false, nil
}
