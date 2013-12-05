package utils

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"os"
)

/**
 * 获取文件内容
 */
func GetFileBytes(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}

/**
 * 获取文件内容，string类型
 */
func GetFileContent(path string) (string, error) {
	file, err := GetFileBytes(path)
	if err != nil {
		return "", err
	}
	return string(file[:]), err
}

/**
 * 将文件内容直接转为Json对象
 */
func GetFileJson(path string, v interface{}) error {
	data, err := GetFileBytes(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

/**
 * 将文件内容直接转为xml对象
 */
func GetFileXml(path string, v interface{}) error {
	data, err := GetFileBytes(path)
	if err != nil {
		return nil
	}
	return xml.Unmarshal(data, v)
}
