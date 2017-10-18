package gongju

import (
	"changliang/zf"
	"changliang/zfzhi"
	"changliang/zw"
	"log"
	"reflect"
	"strings"
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
