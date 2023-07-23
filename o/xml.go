package o

import (
	`encoding/xml`
	`os`
	
	`github.com/gookit/goutil/fsutil`
)

// ReadXML 读取XML文件
func ReadXML(file string, data interface{}) (err error) {
	var content []byte
	content, err = os.ReadFile(os.ExpandEnv(file))
	if err != nil {
		return err
	}
	err = xml.Unmarshal(content, &data)
	if err != nil {
		return err
	}
	return nil
}

// SaveXML 存储XML文件
func SaveXML(data interface{}, file string) error {
	xmlByte, err := xml.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	headerBytes := []byte(xml.Header)
	xmlData := append(headerBytes, xmlByte...)
	_, err = fsutil.PutContents(os.ExpandEnv(file), string(xmlData))
	if err != nil {
		return err
	}
	return nil
}
