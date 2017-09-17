package gongju

import (
	"bytes"
	"changliang/zfzhi"
	"changliang/zf"
	"log"
)

func Shengchengmoban() {
	mkarr := Mokuaimingsarr
	mks := Mokuaimings
	for _, mkvo := range mkarr {
		mkv := mks[mkvo].Zhi
		_, biaos, _ := Biaolies(mkv)
		for bk, _ := range biaos {
			buffer := &bytes.Buffer{}
			buffer.WriteString(mkv)
			dir := Getgopath() + zfzhi.Zhi.Xx() + mkv +
				zfzhi.Zhi.Xx() + zf.Zfs.Zd(true) + zf.Zfs.Fortests(true)
			path := dir + zfzhi.Zhi.Xx() + bk + zf.Zfs.Fortests(true) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
			log.Println(buffer.String())
			log.Println("path", path)
		}

	}
}
