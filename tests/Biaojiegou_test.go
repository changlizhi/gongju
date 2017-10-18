package tests

import (
	"changliang/zf"
	"gongju"
	"log"
	"testing"
)

func TestFanshebiaojiegou(t *testing.T) {
	gongju.Fanshebiaojiegou(zf.Zfs.Mhsydata(true))
}
func TestBiaoliesfanshe(t *testing.T) {
	bls, bs, ls := gongju.Fanshebiaolies(zf.Zfs.Mhsydata(true))
	log.Println(bls)
	log.Println(bs)
	log.Println(ls)
}
func TestFanshejichulie(t *testing.T) {
	log.Println(gongju.Fanshejichulie())
}
