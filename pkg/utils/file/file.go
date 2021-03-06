package fileUtils

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

func FileExist(path string) bool {
	var exist = true
	if _, err := os.Stat(path); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func MkDirIfNeeded(dir string) error {
	if !FileExist(dir) {
		err := os.MkdirAll(dir, os.ModePerm)
		return err
	}

	return nil
}

func ReadFile(filePath string) string {
	buf := ReadFileBuf(filePath)
	str := string(buf)
	str = RemoveBlankLine(str)
	return str
}

func ReadFileBuf(filePath string) []byte {
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []byte(err.Error())
	}

	return buf
}
func RemoveBlankLine(str string) string {
	myExp := regexp.MustCompile(`\n{3,}`) // 连续换行
	ret := myExp.ReplaceAllString(str, "\n\n")

	ret = strings.Trim(ret, "\n")
	ret = strings.TrimSpace(ret)

	return ret
}

func RmDir(dir string) error {
	if FileExist(dir) {
		err := os.RemoveAll(dir)
		return err
	}

	return nil
}

func GetAbsolutePath(pth string) string {
	if !IsAbsolutePath(pth) {
		pth, _ = filepath.Abs(pth)
	}

	pth = AddSepIfNeeded(pth)

	return pth
}
func IsAbsolutePath(pth string) bool {
	return path.IsAbs(pth) ||
		strings.Index(pth, ":") == 1 // windows
}

func AddSepIfNeeded(pth string) string {
	sepa := string(os.PathSeparator)

	if strings.LastIndex(pth, sepa) < len(pth)-1 {
		pth += sepa
	}
	return pth
}
