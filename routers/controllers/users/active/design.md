# 基于 Token 的用户激活业务设计

## 操作流程

1. 用户调用激活接口。
2. 系统生成 token 并发送给用户的邮箱。
3. 用户点击邮箱中的链接，激活账号。
4. 系统验证 token 的有效性。
5. 如果 token 有效，系统激活用户账号。
6. 如果 token 无效，系统提示用户重新激活。
7. 用户成功激活账号。

## 功能要求

### Token 生成与验证：

交由下层服务实现。

### 日志记录：

#### 激活链接生成日志

1. 记录激活链接的生成时间。

#### 激活链接使用日志

1. 记录 token 的使用时间。
2. 记录 token 的使用者。
3. 记录 token 的使用情况。
4. 记录 token 的过期时间。
5. 记录 token 的有效性。

### 业务实现

- [ ] [检测用户是否需要激活](./status/design.md)
- [ ] 发送 token
    - [x] [构造激活链接](./generateURL.go)
        - [x] 生成 token
            - [x] [token 生成](./tokenGenerate.go)
            - [x] [token 解析](./tokenCheck.go)
            - [ ] token 配置化
        - [x] 配置化
    - [ ] [发送邮件](./sendEmail.go)
        - [x] [邮件模板](emailTemplate.go)
            - [ ] 配置化
        - [ ] 邮件内容
- [x] 激活用户
    - [ ] 验证 token
- [x] 记录日志
    - [x] [生成日志](generateURL.go)
    - [x] [使用日志](active.go)
