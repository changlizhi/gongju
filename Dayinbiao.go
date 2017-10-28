package gongju

import (
	"bytes"
	"changliang/zf"
	"changliang/zfzhi"
	"changliang/zh"
	"gongju/sjkmh"
	"log"
	"strings"
)

func dayincl(b string, leixing string,
	bufferzf *bytes.Buffer,
	bufferlx *bytes.Buffer,
	buffercd *bytes.Buffer,
	bufferzw *bytes.Buffer) {
	mingcheng := b + strings.ToLower(leixing)
	//	func (zf *Zf) Xxxshuju(xiaoxie bool) string {
	funzf := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xkhz() + zf.Zfs.Zf(true) +
		zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + zf.Zfs.Zf(false) + zfzhi.Zhi.Xkhy() +
		mingcheng + zfzhi.Zhi.Xkhz() + zf.Zfs.Xiaoxie(true) + zfzhi.Zhi.Kgf() +
		zf.Zfs.Bool(true) + zfzhi.Zhi.Xkhy() + zf.Zfs.String(true)
	bufferzf.WriteString(funzf)
	bufferzf.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//	return fanshe.Fangfaming(xiaoxie)
	zfretstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Fanshe(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.Fangfaming(false) + zfzhi.Zhi.Xkhz() + zf.Zfs.Xiaoxie(true) + zfzhi.Zhi.Xkhy()
	bufferzf.WriteString(zfretstr)
	bufferzf.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

	//	func (zw *Zw) Xxxshuju() string
	funzw := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xkhz() + zf.Zfs.Zw(true) +
		zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + zf.Zfs.Zw(false) + zfzhi.Zhi.Xkhy() +
		mingcheng + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zf.Zfs.String(true)
	bufferzw.WriteString(funzw)
	bufferzw.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//	return "Xxx数据/结构"
	zwretstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Syh() + Zhongwen(b) + Zhongwen(leixing) + zfzhi.Zhi.Syh()
	bufferzw.WriteString(zwretstr)
	bufferzw.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

	//	func (zfzhi *Zfzhi) Xxxshujuleixingzhi() string
	funlx := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xkhz() + zf.Zfs.Zfzhi(true) +
		zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + zf.Zfs.Zfzhi(false) + zfzhi.Zhi.Xkhy() +
		mingcheng + zf.Zfs.Leixingzhi(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zf.Zfs.String(true)
	bufferlx.WriteString(funlx)
	bufferlx.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	//if z := zf.Zfs.Xxxshuju(false); z != zhi.Kzf()
	lxif := zf.Zfs.If(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Z(true) + zfzhi.Zhi.Kgf() +
		zfzhi.Zhi.Mh() + zfzhi.Zhi.Dyh() + zh.Zhs.Zfszhfalse(mingcheng) + zfzhi.Zhi.Fh() +
		zf.Zfs.Z(true) + zfzhi.Zhi.Gth() + zfzhi.Zhi.Dyh() + zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.Kzf(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()
	bufferlx.WriteString(lxif)
	bufferlx.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())

	//return zf.Zfs.String(true)
	lxretstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zh.Zhs.Zfszhtrue(zf.Zfs.String(false))
	bufferlx.WriteString(lxretstr)
	bufferlx.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
	//return zfzhi.Kzf()
	retlxkzf := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.Kzf(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()
	bufferlx.WriteString(retlxkzf)
	bufferlx.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())

	//	func (zhi *Zfzhi) Xxxshujuchangdu() int
	funcd := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xkhz() + zf.Zfs.Zfzhi(true) +
		zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + zf.Zfs.Zfzhi(false) + zfzhi.Zhi.Xkhy() +
		mingcheng + zf.Zfs.Changdu(true) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zf.Zfs.Int(true)
	buffercd.WriteString(funcd)
	buffercd.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	//	return zfzhi.Shuzi5() * zfzhi.Shuzi10()
	cdretstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Dh() +
		zf.Zfs.Shuzi5(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Xh() +
		zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Dh() + zf.Zfs.Shuzi10(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()
	buffercd.WriteString(cdretstr)
	buffercd.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
}

func Dayinbiao() {
	buffershuju := &bytes.Buffer{}
	buffershujum := &bytes.Buffer{}
	bufferjiegou := &bytes.Buffer{}
	bufferjiegoum := &bytes.Buffer{}
	buffermh := &bytes.Buffer{}
	bufferzf := &bytes.Buffer{}
	bufferlx := &bytes.Buffer{}
	buffercd := &bytes.Buffer{}
	bufferzw := &bytes.Buffer{}
	//type Mhsydata struct
	tms := zf.Zfs.Type(true) + zfzhi.Zhi.Kgf() + zf.Zfs.Mhsydata(false) + zfzhi.Zhi.Kgf() + zf.Zfs.Struct(true)
	buffermh.WriteString(tms)
	buffermh.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
	for _, b := range sjkmh.Biaos {
		bnsj := b + zf.Zfs.Shuju(true)
		bnjg := b + zf.Zfs.Jiegou(true)
		bnx := strings.ToLower(b)
		bnxsj := bnx + zf.Zfs.Shuju(true)
		bnxjg := bnx + zf.Zfs.Jiegou(true)
		//Xxxshujumh *Xxxshuju
		byy := bnsj + zf.Zfs.Mh(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + bnsj + zfzhi.Zhi.Hhf()
		buffermh.WriteString(byy)
		//Xxxshujujiegoumh *Xxxshujujiegou
		bjg := bnjg + zf.Zfs.Mh(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Xh() + bnjg + zfzhi.Zhi.Hhf()
		buffermh.WriteString(bjg)

		//type Xxxshuju struct{}
		bsj := zf.Zfs.Type(true) + zfzhi.Zhi.Kgf() + bnsj + zfzhi.Zhi.Kgf() +
			zf.Zfs.Struct(true) + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf()
		buffershuju.WriteString(bsj)

		//type Xxxshujujiegou struct{}
		bsjjg := zf.Zfs.Type(true) + zfzhi.Zhi.Kgf() + bnjg + zfzhi.Zhi.Kgf() +
			zf.Zfs.Struct(true) + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf()
		bufferjiegou.WriteString(bsjjg)

		for _, sjzd := range Fanshejichushuju() {
			//func (xxxshuju *Xxxshuju) sjzd() string
			funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() +
				zfzhi.Zhi.Xkhz() + bnxsj + zfzhi.Zhi.Kgf() +
				zfzhi.Zhi.Xh() + bnsj + zfzhi.Zhi.Xkhy() + sjzd +
				zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zf.Zfs.String(true)
			buffershujum.WriteString(funstr)
			buffershujum.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
			//return zf.Zfs.Sjzd(false)
			retstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zh.Zhs.Zfszhfalse(sjzd)
			buffershujum.WriteString(retstr)
			buffershujum.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
		}
		for _, sjzd := range Fanshejichubiao() {
			//func (xxxshujujiegou *Xxxshujujiegou) sjzd() string
			funstr := zf.Zfs.Func(true) + zfzhi.Zhi.Kgf() +
				zfzhi.Zhi.Xkhz() + bnxjg + zfzhi.Zhi.Kgf() +
				zfzhi.Zhi.Xh() + bnjg + zfzhi.Zhi.Xkhy() + sjzd +
				zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy() + zf.Zfs.String(true)
			bufferjiegoum.WriteString(funstr)
			bufferjiegoum.WriteString(zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf())
			//return zf.Zfs.Sjzd(false)
			retstr := zf.Zfs.Return(true) + zfzhi.Zhi.Kgf() + zh.Zhs.Zfszhfalse(sjzd)
			bufferjiegoum.WriteString(retstr)
			bufferjiegoum.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
		}
		dayincl(b, zf.Zfs.Shuju(false), bufferzf, bufferlx, buffercd, bufferzw)
		dayincl(b, zf.Zfs.Jiegou(false), bufferzf, bufferlx, buffercd, bufferzw)
	}
	buffermh.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Dkhy() + zfzhi.Zhi.Hhf())
	//log.Println(buffershuju)
	//log.Println(bufferjiegou)
	//log.Println(buffershujum)
	log.Println(bufferjiegoum)
	//log.Println(buffermh)
	//log.Println(bufferzf)
	//log.Println(bufferlx)
	//log.Println(bufferzw)
	//log.Println(buffercd)
}
