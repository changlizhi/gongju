package gongju

import (
	"bytes"
	"changliang/zf"
	"changliang/zfzhi"
	"encoding/json"
	"reflect"
	"strings"
)

func Fanshejiexi(modelret interface{}, model interface{}) interface{} {
	rt := reflect.TypeOf(model)
	rv := reflect.ValueOf(model)

	mn := rt.NumField()
	buffer := &bytes.Buffer{}
	buffer.WriteString(zfzhi.Zhi.Dkhz())
	iss := zfzhi.Zhi.Syh() + zfzhi.Zhi.Syh() + zfzhi.Zhi.Mh() + zfzhi.Zhi.Zkhz() + zfzhi.Zhi.Zkhy()
	buffer.WriteString(iss)

	for i := zfzhi.Zhi.Shuzi0(); i < mn; i++ {
		buffer.WriteString(zfzhi.Zhi.Dou())
		rff := rt.Field(i)
		pz := rff.Type
		rvf := rv.FieldByName(rff.Name)

		pztype := strings.Split(pz.String(), zfzhi.Zhi.Dh())[zfzhi.Zhi.Shuzi1()]

		//"xxx":
		pzstr := zfzhi.Zhi.Syh() + pztype + zfzhi.Zhi.Syh() + zfzhi.Zhi.Mh()
		buffer.WriteString(pzstr)
		buffer.WriteString(zfzhi.Zhi.Zkhz())
		pzmn := pz.NumMethod()
		buffer.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Dkhy())
		for j := zfzhi.Zhi.Shuzi0(); j < pzmn; j++ {

			buffer.WriteString(zfzhi.Zhi.Dou() + zfzhi.Zhi.Dkhz())
			//"Guilei": "Shujuku",
			guilei := zfzhi.Zhi.Syh() + zf.Zfs.Guilei(false) + zfzhi.Zhi.Syh() +
				zfzhi.Zhi.Mh() + zfzhi.Zhi.Syh() + pztype + zfzhi.Zhi.Syh() +
				zfzhi.Zhi.Dou()
			buffer.WriteString(guilei)

			//"Bianma": "Mingcheng",
			bianma := zfzhi.Zhi.Syh() + zf.Zfs.Bianma(false) + zfzhi.Zhi.Syh() +
				zfzhi.Zhi.Mh() + zfzhi.Zhi.Syh() + pz.Method(j).Name + zfzhi.Zhi.Syh() +
				zfzhi.Zhi.Dou()
			buffer.WriteString(bianma)

			//"Mingcheng": "数据库名称",
			mingcheng := zfzhi.Zhi.Syh() + zf.Zfs.Mingcheng(false) + zfzhi.Zhi.Syh() +
				zfzhi.Zhi.Mh() + zfzhi.Zhi.Syh() + Zhongwen(pz.Method(j).Name) + zfzhi.Zhi.Syh() +
				zfzhi.Zhi.Dou()
			buffer.WriteString(mingcheng)

			rvfm := rvf.MethodByName(pz.Method(j).Name).Call(nil)[zfzhi.Zhi.Shuzi0()].String()
			//"Zhi": "mhsy_data"
			zhi := zfzhi.Zhi.Syh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Syh() +
				zfzhi.Zhi.Mh() + zfzhi.Zhi.Syh() + rvfm + zfzhi.Zhi.Syh()
			buffer.WriteString(zhi)
			buffer.WriteString(zfzhi.Zhi.Dkhy())
		}
		buffer.WriteString(zfzhi.Zhi.Zkhy())
	}
	buffer.WriteString(zfzhi.Zhi.Dkhy())
	json.Unmarshal(buffer.Bytes(), &modelret)
	return modelret
}
