# 模块

## 编写模块

### 准备
    是否开启gomod
    go env GO111MODULE
    go env -w GO111MODULE=on
    

### 创建
    gitlab 创建项目
    将项目克隆到本地
    git clone git@xxxx/project.git
    cd project
    go mod init git/xxx/project
### 编写
    编写代码
    编写测试文件进行丹云测试

### 测试
    go fmt // 格式化
    go vet ./...
    go test -v .
    go run ./xx/aaa -name aaa

### 文档

### 上传
    git add .
    git commit -m "xxx"
    git tag v1.0.0
    git tag -l

### 引用


## 模块管理


## 工作区