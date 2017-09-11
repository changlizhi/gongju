package tests

import (
	"gongju"
	"testing"
	"log"
	"changliang/zf"
)
func TestBiaolie(t *testing.T){
	log.Println(gongju.Biaolies(zf.Zfs.Hfxyonghu(true)))
	log.Println(gongju.Biaolies(zf.Zfs.Mhsyyonghu(true)))
}

func TestBiaojiegou(t *testing.T){
	log.Println(gongju.Suoyoubiaojiegou(zf.Zfs.Hfxyonghu(true)))
}
func TestBiao(t *testing.T){
	log.Println(gongju.Biao(zf.Zfs.Hfxyonghu(true),"Juese"))
}
func TestLieleixing(t *testing.T){
	log.Println(gongju.Lieleixing("Id"))
}
func TestLiechangdu(t *testing.T){
	log.Println(gongju.Liechangdu("Id"))
}