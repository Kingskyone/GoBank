package files

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

func main6() {
	ReadFilesLineBufio("./go.mod")
}

// 获取文件夹内的文件
func getFiles(dir string) []string {
	f, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalln(err)
	}
	list := make([]string, 0)
	for _, f1 := range f {
		if f1.IsDir() {
			continue
		}
		list = append(list, f1.Name())
	}
	return list
}

// 复制目录
func CopyDir(sourseDir, destDir string) {
	list := getFiles(sourseDir)
	for _, f := range list {
		_, name := path.Split(f)
		sourseFilename := sourseDir + "/" + name
		destFilename := destDir + "/" + name
		file, err := copyFile(sourseFilename, destFilename)
		if err != nil {
			return
		}
		fmt.Println(sourseFilename, file)
	}
}

// 复制文件
func copyFile(sourseName, destName string) (int64, error) {
	src, err := os.Open(sourseName)
	if err != nil {
		log.Fatalln(err)
	}
	defer src.Close()
	dst, err := os.OpenFile(destName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer dst.Close()
	// 复制内容
	return io.Copy(dst, src)
}

// 边读边写文件
func copyFileOneSide(sourseName, destName string) {
	src, err := os.Open(sourseName)
	if err != nil {
		log.Fatalln(err)
	}
	defer src.Close()
	// os.O_APPEND 连续写入
	dst, err := os.OpenFile(destName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer dst.Close()
	// 复制内容
	buff := make([]byte, 1024)
	for {
		// 通过切片长度限制
		n, err := src.Read(buff)
		if err != nil && err != io.EOF {
			log.Fatalln(err)
		}
		// 读完了
		if n == 0 {
			break
		}
		//有可能填不满
		dst.Write(buff[:n])
	}

}

// 一次性读取文件内容
func ReadFiles(fileName string) {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	// 文件位置， 文件名
	a, b := path.Split(fileName)
	fmt.Println(a, b)
	fmt.Println(bytes)
}

// 通过bufio按行读取文件内容
// 缓冲器大小默认4k
func ReadFilesLineBufio(fileName string) {
	fileHander, err := os.OpenFile(fileName, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatalln(err)
	}
	defer fileHander.Close()

	reader := bufio.NewReader(fileHander)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(line)
	}
}

// 通过bufio按行读取文件内容
// 单行大小默认64k
func ReadFilesLineScanner(fileName string) {
	fileHander, err := os.OpenFile(fileName, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatalln(err)
	}
	defer fileHander.Close()

	scanner := bufio.NewScanner(fileHander)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
