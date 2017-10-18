package gongju

import (
	"changliang/zfzhi"
	"gongju/peizhi"
	"reflect"
	"strings"
)

func Fanshejichulie() []string {
	m := &Jicuziduan{}
	ret := []string{}
	rvm := reflect.ValueOf(m)
	for i := zfzhi.Zhi.Shuzi0(); i < rvm.NumMethod(); i++ {
		rvmv := rvm.Method(i).Call(nil)[zfzhi.Zhi.Shuzi0()].String()
		ret = append(ret, rvmv)
	}
	return ret
}

func Fanshebiaojiegou(mokuaiming string) map[string][]string {
	//循环所有Mokuaiming里的字段类型如果匹配了参数则成功
	m := peizhi.Mokuaiming{}
	rvm := reflect.ValueOf(m)
	ret := make(map[string][]string)

	for i := zfzhi.Zhi.Shuzi0(); i < rvm.NumField(); i++ {
		rvmf := rvm.Field(i)
		rvmftn := strings.ToLower(rvmf.Type().Name())
		if rvmftn == mokuaiming {
			for j := zfzhi.Zhi.Shuzi0(); j < rvmf.NumField(); j++ {
				rvmff := rvmf.Field(j)
				rvmffn := strings.Split(rvmff.Type().String(), zfzhi.Zhi.Dh())[zfzhi.Zhi.Shuzi1()]
				vs := Fanshejichulie()
				for k := zfzhi.Zhi.Shuzi0(); k < rvmff.NumMethod(); k++ {
					rvmffm := rvmff.Method(k)
					val := rvmffm.Call(nil)[zfzhi.Zhi.Shuzi0()].String()
					vs = append(vs, val)
				}
				ret[rvmffn] = vs
			}
		}
	}
	return ret
}
