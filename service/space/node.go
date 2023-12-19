package space

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// 使用 updateOne 的缺点是，如果没有找到对应的文档，就会插入一条新的文档，这样会导致数据的不一致性
// 但是此问题可以从设计上解决，比如在插入文档之前，先查询是否存在，如果存在就更新，不存在就插入

//
//// NewSpace 创建新的集合
//// collection 集合名称，集合名称是唯一的。
//func NewSpace(Collection string) error {
//	// 1. 生成元数据
//	data := NewMetaData(Collection, Collection, true)
//	// 2. 检查元数据合法性
//
//	// 2.1 检查元数据名称是否为空
//	if data.Name == "" {
//		return errors.New("元数据名称不能为空")
//	}
//	// 2.2 检查元数据名称长度是否合法
//	if len(data.Name) > 255 {
//		return errors.New("元数据名称长度不能超过 255 个字符")
//	}
//	// 2.3 检查元数据名称是否合法
//	if !CheckName(data.Name) {
//		return errors.New("元数据名称不合法")
//	}
//	// 2.4 检查元数据是否重复
//	result, err := FilesDB().ListCollectionNames(context.Background(), bson.M{
//		"name": Collection,
//	})
//	if err != nil {
//		return errors.Wrap(err, "判断集合是否存在失败")
//	}
//	if len(result) != 0 {
//		return errors.New("集合已存在")
//	}
//	// 2.检查元数据名称长度是否合法
//	// 写入数据
//	err = FilesDB().CreateCollection(context.Background(), data.Space)
//	if err != nil {
//		return errors.Wrap(err, "创建集合失败")
//	}
//	collection := FilesDB().Collection(data.Space)
//
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	_, err = collection.InsertOne(ctx, data)
//	if err != nil {
//		return errors.Wrap(err, "NewSpace 写入数据失败")
//	}
//	return nil
//}

// newNode 在指定的位置中插入新的结点
// 此处不考虑结点是否存在的问题,或是结点合法性的问题，此问题交由上层业务解决
func newNode(data MetaData) error {
	//解析 node 对应的集合
	collection := FilesDB().Collection(data.Space)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 更新操作
	if data.Parent == "/" {
		// 如果是根目录，直接插入
		_, err := collection.InsertOne(ctx, data)
		if err != nil {
			return errors.Wrap(err, "newNode 写入数据失败")
		}
		return nil
	}
	//定位到对应的文件夹结点，更新其子结点
	filter := bson.M{
		"path":   data.Parent,
		"is_dir": true,
	}
	update := bson.M{"$push": bson.M{"children": data}}

	_, err := collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return errors.Wrap(err, "更新结点失败")
	}
	return nil
}

// CreateDir 在指定的集合中创建文件夹
func CreateDir(data MetaData) error {
	//解析 node 对应的集合
	collection := FilesDB().Collection(data.Space)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 更新操作
	//定位到对应的文件夹结点，更新其子结点
	filter := bson.M{
		"path":   data.Parent,
		"is_dir": true,
	}
	update := bson.M{"$push": bson.M{"children": data}}

	_, _ = collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))

	return nil
}
func CheckName(name string) bool {
	//1 长度是否合法
	if len(name) > 255 {
		return false
	}
	//2 是否包含非法字符

	//3 是否包含敏感词
	//4 是否重复
	return true
}

// CheckNodeExit 检查结点是否存在
// todo 感觉会有bug
func CheckNodeExit(data MetaData) (bool, error) {
	//解析 node 对应的集合
	collection := FilesDB().Collection(data.Space)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if data.Parent == "/" {
		// 如果是根目录，直接查询
		filter := bson.M{
			"path":   data.Path,
			"is_dir": data.IsDir,
		}
		result := collection.FindOne(ctx, filter)
		if result.Err() != nil {
			// 一般是没有找到对应的文档
			return false, nil
		}
		return true, nil
	}
	//定位到对应的文件夹结点，查询其子结点

	filter := bson.M{
		"path":   data.Parent,
		"is_dir": true,
	}
	result := collection.FindOne(ctx, filter)
	elem := new(MetaData)

	result.Decode(elem)
	if result.Err() != nil {
		// 一般是没有找到对应的文档
		return false, nil
	}
	for _, v := range elem.ChildRen {
		if v.Name == data.Name && v.IsDir == data.IsDir {
			return true, nil
		}
	}
	return false, nil
}
