package gongju

import (
	"changliang/zf"
	"changliang/zfzhi"
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
	"sort"
	"jichu/peizhi"
	"reflect"
)

var Chushihuas = make(map[string]Tongyong)
var Mokuaimings = make(map[string]Tongyong)
var Mokuaimingsarr = []string{}
var Jsonlies = make(map[string]Tongyong)
var Jsonliesarr = []string{}
var Jsonmojis = make(map[string]Tongyong)

func init() {
	chushihua_json()
}
func chushihua_json() {
	shezhi := Shezhijson()
	for _, c := range shezhi.Chushihua {
		Chushihuas[c.Bianma] = c
	}
	for _, c := range shezhi.Jsonlie {
		Jsonlies[c.Bianma] = c
		Jsonliesarr = append(Jsonliesarr, c.Bianma)
	}
	sort.Strings(Jsonliesarr)
	for _, c := range shezhi.Jsonmoji {
		Jsonmojis[c.Bianma] = c
	}
	for _, c := range shezhi.Mokuaiming {
		Mokuaimings[c.Bianma] = c
		Mokuaimingsarr = append(Mokuaimingsarr, c.Bianma)
	}
	sort.Strings(Mokuaimingsarr)
}

// 获取项目所在目录，这个方法无论在个系统都可以准确获取到项目目录
func Getpath(mokuaiming string) string {
	_, file, _, _ := runtime.Caller(zfzhi.Zhi.Shuzi1())
	fumulu := zfzhi.Zhi.Dh() + zfzhi.Zhi.Dh() + zfzhi.Zhi.Xx()
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, fumulu)))
	if mokuaiming != zf.Zfs.Gopath(true) {
		apppath = apppath + zfzhi.Zhi.Xx() + mokuaiming
	}
	return apppath
}
func Getgopath() string {
	return Getpath(zf.Zfs.Gopath(true))
}
func Getjichupath() string {
	return Getpath(zf.Zfs.Jichu(true))
}
// 获取文件目录，直接返回文件目录结构，不论文件是否存在
func Getwenjianmulu(mokuaiming string, mulu string, wenjian string, leixing string) string {
	path := Getpath(mokuaiming) + // apppath
		zfzhi.Zhi.Xx() + // /
		mulu +
		zfzhi.Zhi.Xx() + // /
		wenjian +
		zfzhi.Zhi.Dh() +
		leixing
	return path
}

func Shezhijson() *Shezhi {
	shezhi := Shezhi{}
	obj := Fanshejiexi(
		&shezhi,
		peizhi.Pz{},
	)
	return obj.(*Shezhi)
}

func Jiexi(path string, model interface{}) interface{} {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("读取json失败", err)
	}
	json.Unmarshal(bytes, &model)
	return model
}

func Yigejsonlies(jsonming string) []string {
	jl := peizhi.Jsonlie{}
	rtjl := reflect.TypeOf(jl)
	for i := zfzhi.Zhi.Shuzi0(); i < rtjl.NumField(); i++ {
		rtjlf := rtjl.Field(i)
		rtjft := rtjlf.Type
		if rtjlf.Type.Name() == jsonming {
			ret := []string{}
			for j := zfzhi.Zhi.Shuzi0(); j < rtjft.NumField(); j++ {
				rtjftf := rtjft.Field(j)
				ret = append(ret, rtjftf.Name)
			}
			return ret
		}
	}
	return []string{}
}
