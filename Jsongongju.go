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
)

var Chushihuas = make(map[string]Tongyong)
var Mokuaimings = make(map[string]Tongyong)
var Mokuaimingsarr = []string{}
var Mulus = make(map[string]Tongyong)
var Jsonlies0 = make(map[string]Tongyong)
var Jsonlies1 = make(map[string]Tongyong)
var Jsonmojis = make(map[string]Tongyong)

func init() {
	chushihua_json()
}
func chushihua_json() {
	shezhi := Shezhijson()
	for _, c := range shezhi.Chushihua {
		Chushihuas[c.Bianma] = c
	}
	for _, c := range shezhi.Mulu {
		Mulus[c.Bianma] = c
	}
	for _, c := range shezhi.Jsonlie {
		if Jsonliejibie(c.Bianma) == zfzhi.Zhi.Shuzi0() {
			Jsonlies0[c.Bianma] = c
		}
		if Jsonliejibie(c.Bianma) == zfzhi.Zhi.Shuzi1() {
			Jsonlies1[c.Bianma] = c
		}
	}
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
func Getchangliangpath() string {
	return Getpath(zf.Zfs.Changliang(true))
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

// 获取jichu/peizhi.json的目录
func Getjichujsonpath(bianma string) string {
	path := Getwenjianmulu(zf.Zfs.Jichu(true), zf.Zfs.Peizhi(true), bianma, zf.Zfs.Json(true))
	return path
}

func Shezhipath() string {
	return Getjichujsonpath(zf.Zfs.Shezhi(false)) //yingyong
}

func Shezhijson() *Shezhi {
	shezhi := Shezhi{}
	obj := Jiexi(Shezhipath(), &shezhi)
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
