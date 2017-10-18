package tests

import (
	"changliang/zf"
	"gongju"
	"log"
	"testing"
)

func TestBiaolie(t *testing.T) {
	log.Println(gongju.Biaolies(zf.Zfs.Hfxyonghu(true)))
	log.Println(gongju.Biaolies(zf.Zfs.Mhsyyonghu(true)))
}

func TestBiaojiegou(t *testing.T) {
	log.Println(gongju.Fanshesuoyoubiaojiegou(zf.Zfs.Mhsydata(true)))
}
func TestBiao(t *testing.T) {
	log.Println(gongju.Fanshebiao(zf.Zfs.Mhsydata(true), "Cujinfangan"))
}
func TestLieleixing(t *testing.T) {
	log.Println(gongju.Lieleixing("Id"))
}
func TestLiechangdu(t *testing.T) {
	log.Println(gongju.Liechangdu("Id"))
}
func TestZhongwen(t *testing.T) {
	log.Println(gongju.Zhongwen("Id"))
}
