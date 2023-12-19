package upload

import (
	"GoCloud/pkg/log"
	"GoCloud/pkg/serializer"
)

//	auditLogging 文件上传接口的日志审计
/*
需要记录下列信息
1. 系统响应结果： HTTP状态码。
2. 用户身份信息： 用户ID或用户名，以识别执行上传操作的用户。 用户的IP地址和设备信息（如果可用），用于审计和安全分析。
3. 文件信息： 上传的文件名和类型。 文件大小。 文件的哈希值或摘要，用于验证完整性。
4. 操作详情： 时间戳：记录每个操作发生的确切时间，包括日期和时间。 上传成功或失败的状态。 如果失败，记录失败的原因（例如，文件太大、格式不支持等）。
*/
func auditLogging(Code int, Msg string, userInfo any, fileInfo any, action string) serializer.Response {
	entry := log.NewEntry("controller.files")
	return serializer.NewResponse(entry, Code, serializer.WithMsg(Msg), serializer.WithField(log.Field{
		Key:   "userInfo",
		Value: userInfo,
	}), serializer.WithField(log.Field{
		Key:   "fileInfo",
		Value: fileInfo,
	}), serializer.WithField(log.Field{
		Key:   "action",
		Value: action,
	}))
}
