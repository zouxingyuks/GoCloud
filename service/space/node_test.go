package space

import (
	"fmt"
	"testing"
)

func TestNewNode(t *testing.T) {
	tests := []struct {
		name string
		data *MetaData
	}{
		{
			name: "1. 正常创建文件",
			data: NewMetaData("test", "test", false),
		},
		{
			name: "2. 正常创建文件夹",
			data: NewMetaData("test", "test", true),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("%+v", tt.data)
			err := newNode(*tt.data)
			if err != nil {
				t.Errorf("NewNode() error = %v", err)
			}
		})
	}
}
func TestDeleteNode(t *testing.T) {
	//todo
}
func TestCheckNodeExit(t *testing.T) {
	//1.在根目录下存在同名文件
	//2. 在子目录下存在同名文件
	//3. 在根目录下存在同名文件夹
	//4. 在子目录下存在同名文件夹
	//5. 在根目录下存在同名的文件与文件夹
	//6. 在子目录下存在同名的文件与文件夹
	//8. 不存在同名文件
	//7. 不存在同名文件夹
	//8.不存在同名文件与文件夹
	type args struct {
		prepareData []*MetaData
		testData    []*MetaData
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1. 在根目录下存在同名文件",
			args: args{
				testData: []*MetaData{
					NewMetaData("test", "testFile", false),
					NewMetaData("test", "testFile", false),
				},
			},
			want: true,
		},
		{
			name: "2. 在子目录下存在同名文件",
			args: args{
				prepareData: []*MetaData{
					NewMetaData("test", "testDir", true),
				},
				testData: []*MetaData{
					NewMetaData("test", "testFile", false, WithParent("/testDir")),
					NewMetaData("test", "testFile", false, WithParent("/testDir")),
				},
			},
			want: true,
		},
		{
			name: "3. 在根目录下存在同名文件夹",
			args: args{
				testData: []*MetaData{
					NewMetaData("test", "testDir", true),
					NewMetaData("test", "testDir", true),
				},
			},
			want: true,
		},
		{
			name: "4. 在子目录下存在同名文件夹",
			args: args{
				prepareData: []*MetaData{
					NewMetaData("test", "testDir", true),
				},
				testData: []*MetaData{
					NewMetaData("test", "testDir", true, WithParent("/testDir")),
					NewMetaData("test", "testDir", true, WithParent("/testDir")),
				},
			}, want: true,
		},
		{
			name: "5. 在根目录下存在同名的文件与文件夹",
			args: args{
				testData: []*MetaData{
					NewMetaData("test", "testDir", true),
					NewMetaData("test", "testDir", false),
				},
			},
			want: false,
		},
		{
			name: "6. 在子目录下存在同名的文件与文件夹",
			args: args{
				prepareData: []*MetaData{
					NewMetaData("test", "testDir", true),
				},
				testData: []*MetaData{
					NewMetaData("test", "testDir", true, WithParent("/testDir")),
					NewMetaData("test", "testDir", false, WithParent("/testDir")),
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, data := range tt.args.prepareData {
				err := newNode(*data)
				if err != nil {
					t.Errorf("newNode() error = %v", err)
				}
			}
			err := newNode(*tt.args.testData[0])
			if err != nil {
				t.Errorf("newNode() error = %v", err)
			}
			result, err := CheckNodeExit(*tt.args.testData[1])
			if result != tt.want || err != nil {

				t.Errorf("CheckNodeExit() = %v, want %v,err %v", result, tt.want, err)
			}
		})
	}
}
