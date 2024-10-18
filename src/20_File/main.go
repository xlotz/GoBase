package main

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

//
//func checkFile(file, err string) (string,int) {
//	var msg string
//	var code int = 200
//	if os.IsNotExist(err) {
//		msg = "文件不存在"
//		code = 100
//		fmt.Println("")
//	} else if err != nil {
//		msg = "访问异常"
//		code = 101
//		fmt.Println("")
//	} else {
//		msg = "读取成功"
//		fmt.Println("", file)
//	}
//	return msg, code
//}

func main() {
	fmt.Println("使用os.Open打开文件...")
	file, err := os.Open("README.md")
	defer file.Close()
	if os.IsNotExist(err) {
		fmt.Println("文件不存在")
	} else if err != nil {
		fmt.Println("访问异常")
	} else {
		fmt.Println("读取成功", file)
	}
	fmt.Println("使用os.OpenFile打开文件...")
	file2, err := os.OpenFile("README.txt", os.O_RDWR|os.O_CREATE, 0666)
	defer file2.Close()
	if os.IsNotExist(err) {
		fmt.Println("文件不存在")
	} else if err != nil {
		fmt.Println("访问异常")
	} else {
		fmt.Println("读取成功", file2.Name())

		// 读取
		//fmt.Println("使用ReadFile读取文件")
		//bytes, err := ReadFile(file2)
		//if err != nil {
		//	fmt.Println("读取异常", err)
		//} else {
		//	fmt.Println(string(bytes))
		//}

		fmt.Println("使用io.ReadAll读取文件...")
		bytes3, err := io.ReadAll(file2)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(bytes3))
		}

	}

	fmt.Println("使用os.ReadFile读取文件...")
	bytes2, err := os.ReadFile("README.md")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(bytes2))
	}

	// 写入
	// WriteString
	wfile, err := os.OpenFile("TestW.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("文件打开异常, ", err)
	} else {
		fmt.Println("文件打开成功，", wfile.Name())
		for i := 0; i < 5; i++ {
			// WriteString
			//offset, err := wfile.WriteString("hello world!\n")

			// io.WriteString
			offset, err := io.WriteString(wfile, "io.WriteString!\n")
			if err != nil {
				fmt.Println(offset, err)
			}
		}
		fmt.Println(wfile.Close())

		// 复制
		fmt.Println("文件复制测试")
		// 以只读的方式打开原文件
		origin, err := os.OpenFile("README.txt", os.O_RDONLY, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer origin.Close()
		// 以只写的方式打开副本文件
		target, err := os.OpenFile("DREADME.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer target.Close()

		// 复制1，从源文件读取文件，写入目标文件
		//offset, err := target.ReadFrom(origin)
		//if err != nil {
		//	fmt.Println(err)
		//	return
		//}
		//fmt.Println("复制成功", offset)

		// 复制2
		written, err := io.Copy(target, origin)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(written)
		}
	}

	// os.WriteFile
	//err3 := os.WriteFile("TestW.txt", []byte("os.WriteFile!\n"), 0666)
	//if err != nil {
	//	fmt.Println(err3)
	//}

	// 重命名
	fmt.Println("重命名")
	err4 := os.Rename("DREADME.txt", "R_DREADME.txt")
	if err4 != nil {
		fmt.Println(err4)
	} else {
		fmt.Println("重命名成功")
	}

	// 删除
	fmt.Println("删除")
	err5 := os.Remove("R_DREADME.txt")
	if err5 != nil {
		fmt.Println(err5)
	} else {
		fmt.Println("删除成功")
	}

	// 刷新
	fmt.Println("os.Sync")
	create, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer create.Close()
	_, err = create.Write([]byte("Hello"))
	if err != nil {
		panic(err)
	}
	if err := create.Sync(); err != nil {
		return
	}

	// 读取文件夹
	fmt.Println("读取文件夹, os.ReadDir")
	// 当前文件夹
	dir, err := os.ReadDir(".")
	if err != nil {
		fmt.Println(err)
	} else {
		for _, entry := range dir {
			fmt.Println(entry.Name())
		}
	}

	fmt.Println("读取文件夹, os.Open")
	dir2, err6 := os.Open(".")
	if err6 != nil {
		fmt.Println(dir2)
	}
	defer dir2.Close()
	dirs, err := dir2.ReadDir(-1)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, entry := range dirs {
			fmt.Println(entry.Name())
		}
	}

	// 创建文件夹
	err7 := os.Mkdir("Test", 0666)
	if err7 != nil {
		fmt.Println(err7)
	} else {
		fmt.Println("目录创建成功")
	}

}

func ReadFile(file *os.File) ([]byte, error) {
	buffer := make([]byte, 0, 512)
	for {
		// 当容量不足时
		if len(buffer) == cap(buffer) {
			// 扩容
			buffer = append(buffer, 0)[:len(buffer)]
		}
		// 继续读取文件
		offset, err := file.Read(buffer[len(buffer):cap(buffer)])
		// 将已写入的数据归入切片
		buffer = buffer[:len(buffer)+offset]
		// 发生错误时
		if err != nil {
			if errors.Is(err, io.EOF) {
				err = nil
			}
			return buffer, err
		}
	}
}

// 复制
func CopyDir(src, dst string) error {
	// 检测源文件夹的状态
	_, err := os.Stat(src)
	if err != nil {
		return err
	}

	return filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 计算相对路径
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		// 拼接目标路径
		destpath := filepath.Join(dst, rel)

		// 创建文件夹
		var dirpath string
		var mode os.FileMode = 0755
		if info.IsDir() {
			dirpath = destpath
			mode = info.Mode()
		} else if info.Mode().IsRegular() {
			dirpath = filepath.Dir(destpath)
		}

		if err := os.MkdirAll(dirpath, mode); err != nil {
			return err
		}
		// 创建文件
		if info.Mode().IsRegular() {
			srcfile, err := os.Open(path)
			if err != nil {
				return err
			}

			// 关闭文件
			defer srcfile.Close()
			destfile, err := os.OpenFile(destpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, info.Mode())
			if err != nil {
				return err
			}
			defer destfile.Close()
			//复制文件内容

			if _, err := io.Copy(destfile, srcfile); err != nil {
				return err
			}
			return nil
		}
		return nil
	})
}
