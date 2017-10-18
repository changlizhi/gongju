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
	bufferzhong := &bytes.Buffer{}
	for i := zfzhi.Zhi.Shuzi0(); i < mn; i++ {
		rff := rt.Field(i)
		pz := rff.Type
		rvf := rv.FieldByName(rff.Name)

		pztype := strings.Split(pz.String(), zfzhi.Zhi.Dh())[zfzhi.Zhi.Shuzi1()]

		//"xxx":
		pzstr := zfzhi.Zhi.Syh() + pztype + zfzhi.Zhi.Syh() + zfzhi.Zhi.Mh()
		bufferzhong.WriteString(pzstr)
		bufferzhong.WriteString(zfzhi.Zhi.Zkhz())
		pzmn := pz.NumMethod()
		bufferin := &bytes.Buffer{}
		for j := zfzhi.Zhi.Shuzi0(); j < pzmn; j++ {
			bufferin.WriteString(zfzhi.Zhi.Dkhz())
			//"Guilei": "Shujuku",
			guilei := zfzhi.Zhi.Syh() + zf.Zfs.Guilei(false) + zfzhi.Zhi.Syh() +
				zfzhi.Zhi.Mh() + zfzhi.Zhi.Syh() + pztype + zfzhi.Zhi.Syh() +
				zfzhi.Zhi.Dou()
			bufferin.WriteString(guilei)

			//"Bianma": "Mingcheng",
			bianma := zfzhi.Zhi.Syh() + zf.Zfs.Bianma(false) + zfzhi.Zhi.Syh() +
				zfzhi.Zhi.Mh() + zfzhi.Zhi.Syh() + pz.Method(j).Name + zfzhi.Zhi.Syh() +
				zfzhi.Zhi.Dou()
			bufferin.WriteString(bianma)

			//"Mingcheng": "数据库名称",
			mingcheng := zfzhi.Zhi.Syh() + zf.Zfs.Mingcheng(false) + zfzhi.Zhi.Syh() +
				zfzhi.Zhi.Mh() + zfzhi.Zhi.Syh() + Zhongwen(pz.Method(j).Name) + zfzhi.Zhi.Syh() +
				zfzhi.Zhi.Dou()
			bufferin.WriteString(mingcheng)

			rvfm := rvf.MethodByName(pz.Method(j).Name).Call(nil)[zfzhi.Zhi.Shuzi0()].String()
			//"Zhi": "mhsy_data"
			zhi := zfzhi.Zhi.Syh() + zf.Zfs.Zhi(false) + zfzhi.Zhi.Syh() +
				zfzhi.Zhi.Mh() + zfzhi.Zhi.Syh() + rvfm + zfzhi.Zhi.Syh()
			bufferin.WriteString(zhi)
			bufferin.WriteString(zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Dou())
		}
		instr := bufferin.String()
		instrnew := instr[zfzhi.Zhi.Shuzi0():(len(instr) - zfzhi.Zhi.Shuzi1())]
		bufferzhong.WriteString(instrnew)
		bufferzhong.WriteString(zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Dou())
	}
	zhongstr := bufferzhong.String()
	zhongstrnew := zhongstr[zfzhi.Zhi.Shuzi0():(len(zhongstr) - zfzhi.Zhi.Shuzi1())]
	buffer.WriteString(zhongstrnew)
	buffer.WriteString(zfzhi.Zhi.Dkhy())
	json.Unmarshal(buffer.Bytes(), &modelret)
	return modelret
}
