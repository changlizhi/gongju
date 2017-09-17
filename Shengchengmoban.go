package gongju

import (
	"bytes"
	"changliang/zfzhi"
	"changliang/zf"
	"log"
)

func importsmoban(mkv string, buffer *bytes.Buffer) {
	buffer.WriteString(zf.Zfs.Import(true))
	buffer.WriteString(zfzhi.Zhi.Xkhz() + zfzhi.Zhi.Hhf())

	//"xxxxxxxx"
	nstr := zfzhi.Zhi.Syh() + zfzhi.Zhi.Syh() + zfzhi.Zhi.Hhf()
	buffer.WriteString(nstr)

	buffer.WriteString(zfzhi.Zhi.Hhf() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Hhf())
}
func fangfamoban(bianma string, buffer *bytes.Buffer) {

}
func Shengchengmoban() {
	mkarr := Mokuaimingsarr
	mks := Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		_, biaos, _ := Biaolies(mkv)
		for bk, _ := range biaos {
			buffer := &bytes.Buffer{}
			buffer.WriteString(mkv)
			importsmoban(mkv, buffer)
			fangfamoban(bk, buffer)

			dir := Getgopath() + zfzhi.Zhi.Xx() + mkv +
				zfzhi.Zhi.Xx() + zf.Zfs.Zd(true) + zf.Zfs.Fortests(true)
			path := dir + zfzhi.Zhi.Xx() + bk + zf.Zfs.Fortests(true) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
			log.Println(buffer.String())
			log.Println("path----------------", path)
		}

	}
}
