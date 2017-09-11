package tests

import (
	"gongju"
	"testing"
	"log"
	"changliang/zf"
)
func TestBiaolie(t *testing.T){
	log.Println(gongju.Biaolies(zf.Zfs.Hfxyonghu(false)))
}

func TestBiaojiegou(t *testing.T){
	log.Println(gongju.Suoyoubiaojiegou(zf.Zfs.Hfxyonghu(false)))
}
func TestBiao(t *testing.T){
	log.Println(gongju.Biao(zf.Zfs.Hfxyonghu(false),"Juese"))
}
func TestLieleixing(t *testing.T){
	log.Println(gongju.Lieleixing("Id"))
}
func TestLiechangdu(t *testing.T){
	log.Println(gongju.Liechangdu("Id"))
}