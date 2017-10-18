package gongju

import (
	"changliang/zf"
	"changliang/zfzhi"
	"changliang/zw"
	"log"
	"reflect"
	"strings"
	"gongju/peizhi"
)

func Fanshebiaolies(mokuaiming string) (biaolie map[string]string, biao map[string]string, lie map[string]string) {
	biaolie = make(map[string]string)
	biao = make(map[string]string)
	lie = make(map[string]string)
	bjg := Fanshebiaojiegou(mokuaiming)
	for k, vs := range bjg {
		for _, v := range vs {
			biaolie[k+v] = strings.ToLower(k + v)
			biao[k] = strings.ToLower(k)
			lie[v] = strings.ToLower(v)
		}
	}
	return
}

func Fanshebiao(mokuaiming string, mingcheng string) []string {
	biaojiegou := Fanshebiaojiegou(mokuaiming)
	return biaojiegou[mingcheng]
}

func Liechangdu(lieming string) int {
	a := zfzhi.Zfzhi{}
	ffm := lieming + zf.Zfs.Changdu(true)
	v := reflect.ValueOf(&a)
	ret := v.MethodByName(ffm).Call(nil)[zfzhi.Zhi.Shuzi0()].Interface().(int)
	return ret
}

func Lieleixing(lieming string) string {
	a := zfzhi.Zfzhi{}
	ffm := lieming + zf.Zfs.Leixingzhi(true)
	v := reflect.ValueOf(&a)
	ret := v.MethodByName(ffm).Call(nil)[zfzhi.Zhi.Shuzi0()].Interface().(string)
	if ret == zfzhi.Zhi.Kzf() {
		log.Println("类型错误-----空")
		return zfzhi.Zhi.Kzf()
	}
	return ret
}

func Zhongwen(lieming string) string {
	a := zw.Zw{}
	v := reflect.ValueOf(&a)
	ret := v.MethodByName(lieming).Call(nil)[zfzhi.Zhi.Shuzi0()].Interface().(string)
	if ret == zfzhi.Zhi.Kzf() {
		log.Println("类型错误-----空")
		return zfzhi.Zhi.Kzf()
	}
	return ret
}

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
