package gongju

import (
	"bytes"
	"changliang/zf"
	"changliang/zfzhi"
	"changliang/zh"
)

func Pinjieziduan(mokuai string, bianma string, buffer *bytes.Buffer) {

	for _, lk := range Biao(mokuai, bianma) {
		buffer.WriteString(lk + zfzhi.Zhi.Mh())
		zhi := shengchengzhi(lk, Lieleixing(lk))
		buffer.WriteString(zhi + zfzhi.Zhi.Dou() + zfzhi.Zhi.Hhf())
	}

}
func shengchengzhi(lieming string, leixing string) string {
	if leixing == zf.Zfs.String(true) {
		// leixing + zf.Zfs.Xxx(false) + fangfa + istring,
		ret := zf.Zfs.Leixing(true) + zfzhi.Zhi.Jia() +
			zh.Zhs.Zfszhfalse(lieming) + zfzhi.Zhi.Jia() +
			zf.Zfs.Fangfa(true) + zfzhi.Zhi.Jia() + zf.Zfs.I(true) + zf.Zfs.String(true)
		return ret
	}
	// i
	if leixing == zf.Zfs.Int(true) {
		ret := zf.Zfs.I(true) // i
		return ret
	}
	// time.Now()
	if leixing == zf.Zfs.Time(true) {
		// time.Now()
		ret := zf.Zfs.Time(true) + zfzhi.Zhi.Dh() + zf.Zfs.Now(false) + zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Xkhy()
		return ret
	}
	// 1.0
	if leixing == zf.Zfs.Float32(true) {
		ret := zf.Zfs.Float32(true) + zfzhi.Zhi.Xkhz() + zf.Zfs.I(true) + zfzhi.Zhi.Xkhy()
		return ret
	}
	if leixing == zf.Zfs.Float64(true) {
		// 1.0
		ret := zf.Zfs.Float64(true) + zfzhi.Zhi.Xkhz() + zf.Zfs.I(true) + zfzhi.Zhi.Xkhy()
		return ret
	}
	// Null
	return zf.Zfs.Null(false)
}
