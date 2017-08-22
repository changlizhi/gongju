package tests

import (
	"gongju"
	"testing"
	"log"
)
func TestBiaolie(t *testing.T){
	log.Println(gongju.Biaolies())
}

func TestBiaojiegou(t *testing.T){
	log.Println(gongju.Suoyoubiaojiegou())
}
func TestBiao(t *testing.T){
	log.Println(gongju.Biao("Juese"))
}
func TestLieleixing(t *testing.T){
	log.Println(gongju.Lieleixing("Id"))
}
func TestLiechangdu(t *testing.T){
	log.Println(gongju.Liechangdu("Id"))
}