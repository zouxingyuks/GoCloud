package cache

import (
	"GoCloud/pkg/util"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func tempConnect() Driver {
	config := RedisConfig{
		Addr:     "100.76.246.116:6379",
		Password: "redis123",
		DB:       0,
		PoolSize: 10,
	}
	d, err := New(KindRedis, WithConf(config))
	if err != nil {
		panic(err)
	}
	return d
}

// hash 类的前置数据
func tempH() Driver {
	d := tempConnect()
	type data struct {
		Name string
		Age  int
	}
	data1 := data{
		Name: "test",
		Age:  18,
	}
	mapData, err := util.StructToMapF1(data1)
	if err != nil {
		panic(errors.Wrap(err, "struct to map"))
	}

	err = d.HMSet("test", mapData)

	if err != nil {
		panic(err)
	}
	return d
}
func TestRedisStore_HGet(t *testing.T) {
	type args struct {
		key   string
		field string
		want  any
	}
	d := tempH()
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				key:   "test",
				field: "Age",
				want:  "18",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := d.HGet(tt.args.key, tt.args.field)
			if (err != nil) != tt.wantErr {
				t.Errorf("HGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//使用 assert.EqualValues() 比较两个值是否相等
			if !assert.Equal(t, tt.args.want, got) {
				t.Errorf("HGet() got = %v, want %v", got, tt.args.want)
			}
		})
	}
}

func TestRedisStore_HSet(t *testing.T) {
	type args struct {
		key   string
		field string
		value interface{}
	}
	d := tempH()
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				key:   "test",
				field: "Age",
				value: 29,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := d.HSet(tt.args.key, tt.args.field, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("HSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRedisStore_HGetAll(t *testing.T) {
	type args struct {
		key   string
		want  map[string]string
		want1 any
	}
	d := tempH()
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				key:  "test",
				want: map[string]string{"Name": "test", "Age": "18"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := d.HGetAll(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("HGetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.EqualValues(t, tt.args.want, got) {
				t.Errorf("HGetAll() got = %v, want %v", got, tt.args.want)
			}
		})
	}
}

func TestRedisStore_HMSet(t *testing.T) {
	type args struct {
		key    string
		values map[string]interface{}
	}
	d := tempConnect()
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				key: "user:test:alice",
				values: map[string]interface{}{
					"Name": "Alice",
					"Age":  20,
				},
			},
		},
		{
			name: "test2",
			args: args{
				key: "user:test:bob",
				values: map[string]interface{}{
					"Name": "Bob",
					"Age":  19,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := d.HMSet(tt.args.key, tt.args.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("HMSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
