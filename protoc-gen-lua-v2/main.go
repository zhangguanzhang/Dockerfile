package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoprint"
)

const (
	splitSub = " = "
)

func main() {

	for _, lua_pb_file := range os.Args[1:] {
		file, err := os.Open(lua_pb_file)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		var (
			DescriptorMap = make(map[string]*descriptorpb.DescriptorProto, 0)
			FiledMap      = make(map[string]*descriptorpb.FieldDescriptorProto, 0)
			fullNameMap   = make(map[string]string)
			// MessageNameList = make([]string, 0)
			pb = &descriptorpb.FileDescriptorProto{
				Syntax: proto.String("proto2"),
				Name:   proto.String(strings.TrimSuffix(lua_pb_file, "_pb.lua") + ".proto"),
			}
		)

		for scanner.Scan() {
			line := scanner.Text()
			linePart := strings.SplitN(line, splitSub, 2)
			if len(linePart) != 2 {
				continue
			}
			switch {
			case linePart[1] == "protobuf.Descriptor();":
				namePart := strings.Split(linePart[0], " ")
				DescriptorMap[namePart[len(namePart)-1]] = &descriptorpb.DescriptorProto{}
			case linePart[1] == "protobuf.FieldDescriptor();":
				namePart := strings.Split(linePart[0], " ")
				FiledMap[namePart[len(namePart)-1]] = &descriptorpb.FieldDescriptorProto{}
			default:
				// 最少一个点
				if !strings.Contains(linePart[0], ".") {
					continue
				}
				// 避免存在 table.xxx.name，所以取最后一个点切割
				lastDotIndex := strings.LastIndex(linePart[0], ".")
				filedKey := linePart[0][:lastDotIndex]
				filedProperty := linePart[0][lastDotIndex+1:]

				if val, ok := FiledMap[filedKey]; ok {
					switch {
					case filedProperty == "name":
						val.Name = proto.String(linePart[1][1 : len(linePart[1])-1])
					case filedProperty == "number":
						num, _ := strconv.Atoi(linePart[1])
						val.Number = proto.Int32(int32(num))
					case filedProperty == "type":
						num, _ := strconv.Atoi(linePart[1])
						// 12 是 bytes 但是 golang 无法序列化，使用字符串类型
						if os.Getenv("GO") != "" && num == 12 {
							num = 9
						}
						val.Type = descriptorpb.FieldDescriptorProto_Type(int32(num)).Enum()
					case filedProperty == "label":
						num, _ := strconv.Atoi(linePart[1])
						val.Label = descriptorpb.FieldDescriptorProto_Label(int32(num)).Enum()
					case filedProperty == "message_type":
						// message_type时候，右侧无双引号，直接使用 linePart[1]
						val.TypeName = proto.String(fullNameMap[linePart[1]])
					}
				}
				if val, ok := DescriptorMap[filedKey]; ok {
					switch {
					case filedProperty == "name":
						val.Name = proto.String(linePart[1][1 : len(linePart[1])-1])
					case filedProperty == "full_name":
						fullNameMap[filedKey] = linePart[1][1 : len(linePart[1])-1]

						full_name := linePart[1][2 : len(linePart[1])-1]
						full_namePart := strings.Split(full_name, ".")
						if len(full_namePart) >= 2 {
							pb.Package = proto.String(full_namePart[0])
						}

					case filedProperty == "fields":
						filedSlice := strings.Split(linePart[1][1:len(linePart[1])-1], ", ")
						for _, fildName := range filedSlice {
							if FiledVal, ok := FiledMap[fildName]; ok {
								val.Field = append(val.Field, FiledVal)
							}
						}
					}
				}
			}
		}

		// fmt.Printf("%v \n", DescriptorMap)

		for _, v := range DescriptorMap {
			pb.MessageType = append(pb.MessageType, v)
		}

		// for _, m := range pb.MessageType {
		// 	for j, _ := range m.Field {
		// 		if m.Field[j].Type == descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum() {
		// 			key := m.Field[j].TypeName
		// 			m.Field[j].TypeName = proto.String(fullNameMap[*key])
		// 		}
		// 	}
		// }

		// fmt.Printf("%v \n", pb)
		fd, err := desc.CreateFileDescriptor(pb)
		if err != nil {
			fmt.Println(err)
			return
		}

		printer := protoprint.Printer{}
		protoContent, err := printer.PrintProtoToString(fd)
		if err != nil {
			panic(err)
		}

		// 打印生成的 .proto 文件内容
		fmt.Println(protoContent)
	}
}
