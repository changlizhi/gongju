package gongju

import (
	"reflect"
	"changliang/zf"
	"changliang/zfzhi"
	"changliang/jsonlie"
)

func Jsonliejibie(lieming string) int {
	jl := jsonlie.Jl{}
	ffm := lieming + zf.Zfs.Jibie(true)
	v := reflect.ValueOf(&jl)
	ret := v.MethodByName(ffm).Call(nil)[zfzhi.Zhi.Shuzi0()].Interface().(int)
	return ret
}
func Jsonliezhiding(lieming string) string {
	jl := jsonlie.Jl{}
	ffm := lieming + zf.Zfs.Zhiding(true)
	v := reflect.ValueOf(&jl)
	ret := v.MethodByName(ffm).Call(nil)[zfzhi.Zhi.Shuzi0()].Interface().(string)
	return ret
}
func Yigejsonlies(jsonming string) []string {
	ret := Pipeifangfa(zf.Zfs.Jl(true), zf.Zfs.Jl(false), zf.Zfs.Jsonlie(true), jsonming)
	return ret
}