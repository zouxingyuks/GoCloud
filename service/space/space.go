package space

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

import "github.com/pkg/errors"

func NewSpace(Collection string) error {
	// 1. 生成元数据
	data := NewMetaData(Collection, Collection, true)
	// 2. 检查元数据合法性

	// 2.1 检查元数据名称是否为空
	if data.Name == "" {
		return errors.New("元数据名称不能为空")
	}
	// 2.2 检查元数据名称长度是否合法
	if len(data.Name) > 255 {
		return errors.New("元数据名称长度不能超过 255 个字符")
	}
	// 2.3 检查元数据名称是否合法
	if !CheckName(data.Name) {
		return errors.New("元数据名称不合法")
	}
	// 2.4 检查元数据是否重复
	result, err := FilesDB().ListCollectionNames(context.Background(), bson.M{
		"name": Collection,
	})
	if err != nil {
		return errors.Wrap(err, "判断集合是否存在失败")
	}
	if len(result) != 0 {
		return errors.New("集合已存在")
	}
	// 2.检查元数据名称长度是否合法
	// 写入数据
	err = FilesDB().CreateCollection(context.Background(), data.Root)
	if err != nil {
		return errors.Wrap(err, "创建集合失败")
	}
	collection := FilesDB().Collection(data.Root)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, data)
	if err != nil {
		return errors.Wrap(err, "NewSpace 写入数据失败")
	}
	return nil
}
