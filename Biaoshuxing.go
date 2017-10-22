package gongju

import (
	"changliang/zf"
	"changliang/zfzhi"
	"changliang/zw"
	"gongju/peizhi"
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

func Biaojiziduan(ziduan string) string {
	sjk := Ziduanbiaoji{}
	ffm := ziduan + zf.Zfs.Biaoji(true)
	v := reflect.ValueOf(&sjk)
	ret := v.MethodByName(ffm).Call(nil)[zfzhi.Zhi.Shuzi0()].Interface().(string)

	return ret
}
func Biaojibiao(biao string) string {
	sjk := Biaobiaoji{}
	ffm := biao + zf.Zfs.Biaoji(true)
	v := reflect.ValueOf(&sjk)
	ret := v.MethodByName(ffm).Call(nil)[zfzhi.Zhi.Shuzi0()].Interface().(string)

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
	ret := []string{}
	//m := &Jichuziduan{}
	//rvm := reflect.ValueOf(m)
	//for i := zfzhi.Zhi.Shuzi0(); i < rvm.NumMethod(); i++ {
	//	rvmv := rvm.Method(i).Call(nil)[zfzhi.Zhi.Shuzi0()].String()
	//	ret = append(ret, rvmv)
	//}
	return ret
}

func Fanshebiaojiegou(mokuaiming string) map[string][]string {
	//循环所有Mokuaiming里的字段类型如果匹配了参数则成功
	m := peizhi.Mokuaiming{}
	rvm := reflect.ValueOf(m)
	//rtm := reflect.TypeOf(m)
	ret := make(map[string][]string)

	for i := zfzhi.Zhi.Shuzi0(); i < rvm.NumField(); i++ {
		rvmf := rvm.Field(i)
		//rtmfn, _ := rtm.FieldByName(rvmf.Type().Name() + zf.Zfs.Mk(true))
		rvmftn := strings.ToLower(rvmf.Type().Name())
		if rvmftn == mokuaiming {
			for j := zfzhi.Zhi.Shuzi0(); j < rvmf.NumField(); j++ {
				rvmff := rvmf.Field(j)
				rvmffn := strings.Split(rvmff.Type().String(), zfzhi.Zhi.Dh())[zfzhi.Zhi.Shuzi1()]
				//rtmff, _ := rtmfn.Type.FieldByName(zf.Zfs.Dt(false) + strings.ToLower(rvmffn))
				vs := Fanshejichulie()
				for k := zfzhi.Zhi.Shuzi0(); k < rvmff.NumMethod(); k++ {
					rvmffm := rvmff.Method(k)
					//log.Print("rtmff.Type.Method(k).Name------", rtmff.Type.Method(k).Name)
					val := rvmffm.Call(nil)[zfzhi.Zhi.Shuzi0()].String()
					vs = append(vs, val)
				}
				ret[rvmffn] = vs
			}
		}
	}
	return ret
}

func Fanshebiaos(mokuai string) []string {
	ret := []string{}
	m := peizhi.Mokuaiming{}
	rtm := reflect.TypeOf(m)
	for i := zfzhi.Zhi.Shuzi0(); i < rtm.NumField(); i++ {
		rtmf := rtm.Field(i)
		rtmfn := strings.ToLower(rtmf.Type.Name())
		if rtmfn == mokuai {
			rtmff := rtmf.Type
			for j := zfzhi.Zhi.Shuzi0(); j < rtmff.NumField(); j++ {
				rtmfff := rtmff.Field(j).Type
				rtmfffn := strings.Split(rtmfff.String(), zfzhi.Zhi.Dh())[zfzhi.Zhi.Shuzi1()]
				ret = append(ret, rtmfffn)
			}
		}
	}
	return ret
}

func Fanshejichubiao() []string {
	ret := []string{}
	jc := Jichu{}
	rtjc := reflect.TypeOf(jc)
	for i := zfzhi.Zhi.Shuzi0(); i < rtjc.NumField(); i++ {
		rtjcf := rtjc.Field(i).Name
		ret = append(ret, rtjcf)
	}
	return ret
}
func Fanshejichushuju() []string {
	ret := []string{}
	jc := Jichushuju{}
	rtjc := reflect.TypeOf(jc)
	for i := zfzhi.Zhi.Shuzi0(); i < rtjc.NumField(); i++ {
		rtjcf := rtjc.Field(i).Name
		ret = append(ret, rtjcf)
	}
	return ret
}
