package util

import (
	"encoding/json"
	"github.com/pkg/errors"
)

//本文件用于处理数据结构的转换

// StructToMapF1  将struct转换为map[string]any
// 使用json.Marshal和json.Unmarshal
// 优点：简单
// 缺点：1. 无法处理struct中的嵌套struct
//  2. 无法处理struct中的map
//  3. 无法处理struct中的slice
//  4. 无法处理struct中的interface
//  5. 无法处理struct中的指针
//  6. 无法处理struct中的私有字段
//  7. 无法处理struct中的非基本类型字段
//  8. 在处理数字时，解析的结果为float64 类型，在使用时需要注意
func StructToMapF1[T any](obj T) (map[string]any, error) {
	objMap := make(map[string]any)
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return nil, errors.Wrap(err, "struct to map error")
	}
	err = json.Unmarshal(jsonBytes, &objMap)
	return objMap, nil
}
