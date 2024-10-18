# 文件

## 介绍
    os库 负责OS文件系统的具体实现
    io库 读写IO的抽象层
    fs库 文件系统的抽象层

## 打开
    os包 Open 返回值一个文件指针和一个错误 
    例：
    file, err := os.Open("README.txt")

    os包 OpenFile可以修改文件描述符和文件权限
    例：
    file, err := os.OpenFile("README.txt", os.O_RDWR|os.O_CREATE, 0666)
    
    注意：
        打开文件后永远要记得关闭。
        defer file.Close()

## 读取
    Read
        file.Read 读取文件到字节切片
        func (f *File) Read(b []byte) (n int, err error) 
    ReadAt
        file.ReadAt 从指定偏移量读取文件到字节切片
        func (f *File) ReadAt(b []byte, off int64) (n int, err error) 
    
    os.ReadFile
        func ReadFile(name string) ([]byte, error)
    io.ReadAll
        func ReadAll(r Reader) ([]byte, error)

## 写入
    Write 写入字节切片
        func (f *File) Write(b []byte) (n int, err error)
    WriteString 写入字符串
        func (f *File) WriteString(s string) (n int, err error)
    WriteAt 从指定位置写入，当以只读模式打开时报错
        func (f *File) WriteAt(b []byte, off int64) (n int, err error)
    
    os.WriteFile 写入文件
        func WriteFile(name string, data []byte, perm FileMode) error
    io.WriteString
        func WriteString(w Writer, s string) (n int, err error) 

    os.Create 用于创建文件，对OpenFile的封装，当父目录不存在时，创建失败。

## 复制
    第一种：打开2个文件，读出写入
        os.ReadFile
        os.WriteFile
    第二种： os.File.ReadFrom
        func (f *File) ReadFrom(r io.Reader) (n int64, err error)
    第三种： io.Copy
        将文件写入缓存区，在写到目标文件，缓冲区默认32KB
        func Copy(dst Writer, src Reader) (written int64, err error)

## 重命名
    os.Rename
        func Rename(oldpath, newpath string) error

## 删除
    os.Remove
        // 删除单个文件或者空目录，当目录不为空时会返回错误
        func Remove(name string) error
    os.RemoveAll
        // 删除指定目录的所有文件和目录包括子目录与子文件
        func RemoveAll(path string) error

## 刷新
    os.Sync 封装底层系统调用的Fsync,将缓存的IO落盘
    

## 文件夹
    和文件类似

### 读取
    os.ReadDir
        func ReadDir(name string) ([]DirEntry, error)
    os.File.ReadDir
        func (f *File) ReadDir(n int) ([]DirEntry, error)

### 创建
    os.Mkdir
        // 用指定的权限创建指定名称的目录
        func Mkdir(name string, perm FileMode) error

    os.MkdirAll
        // 相较于前者该函数会创建一切必要的父目录
        func MkdirAll(path string, perm FileMode) error

### 复制
    第一种：写函数变量目录
    第二种： filepath