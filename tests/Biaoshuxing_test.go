package tests

import (
	"changliang/zf"
	"gongju"
	"log"
	"testing"
)

func TestFanshebiaolie(t *testing.T) {
	log.Println(gongju.Fanshebiaolies(zf.Zfs.Mhsydata(true)))
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
func TestJichuziduanbiaoji(t *testing.T) {
	log.Println(gongju.Ziduanbiaoji("Id"))
}
