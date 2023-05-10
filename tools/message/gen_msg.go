package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	PATH               = "../../message"
	PROTO [2][1]string = [2][1]string{
		{
			"protoc.exe --plugin=protoc-gen-go=protoc-gen-go.exe  --go_out=%s  --proto_path=%s	",
		}, //win
		{
			"protoc --go_out=%s  --proto_path=%s	",
		}, //linux
	}
)

// args[1] : proto file path
func main() {
	args := os.Args
	if len(args) >= 2 {
		PATH = args[1]
	}
	files, err := filepath.Glob(PATH + "/*.proto")
	str := ""
	if err == nil {
		files1 := []string{}
		for _, v := range files {
			v = strings.Replace(v, "\\", "/", -1)
			if strings.LastIndex(v, "message.proto") != -1 {
				str += v + "	"
				continue
			}
			files1 = append(files1, v)
		}

		for _, v := range files1 {
			str += v + "	"
		}
	}

	index := strings.LastIndex(str, "	")
	if index != -1 {
		str = str[:index]
	}

	//生成bat文件
	{
		stream := bytes.NewBuffer([]byte{})
		file, err := os.Create("gen_msg.bat")
		if err == nil {
			for _, v := range PROTO[0] {
				stream.WriteString(fmt.Sprintf(v, PATH, PATH))
				stream.WriteString(str)
				stream.WriteString("\n")
			}
			file.Write(stream.Bytes())
			file.Close()
		}
	}
	{
		stream := bytes.NewBuffer([]byte{})
		file, err := os.Create("gen_msg.sh")
		if err == nil {
			for _, v := range PROTO[1] {
				stream.WriteString(fmt.Sprintf(v, PATH, PATH))
				stream.WriteString(str)
				stream.WriteString("\n")
			}
			file.Write(stream.Bytes())
			file.Close()
		}
	}
}
