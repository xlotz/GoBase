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
    git push --tags

### 引用
    go get xxx/xxx/xxx@latest // 引用
    go install xxx/xxx/xxx@latest //安装
    go run -mod=mod xxx/xxx/hello -name jack // 直接运行


### 设置代理
    go env -w GOPROXY=https://goproxy.cn,direct

### 下载依赖
    go git github.com/gin-gonic/gin
    cat go.mod
    go git github.com/gin-gonic/gin@latest //升级依赖
    go get github.com/gin-gonic/gin@none // 删除依赖

    go install github.com/go-delve/delve/cmd/dlv@latest //安装命令行
    
    

## 模块管理
    常用命令

    go mod download //下载当前项目的依赖包
    go mod edit	    //编辑go.mod文件
    go mod graph    //输出模块依赖图
    go mod init	    //在当前目录初始化go mod
    go mod tidy	    //清理项目模块
    go mod verify	//验证项目的依赖合法性
    go mod why	    //解释项目哪些地方用到了依赖
    go clean -modcache	//用于删除项目模块依赖缓存
    go list -m	    //列出模块

## 工作区