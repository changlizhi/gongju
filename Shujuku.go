package gongju

import (
	"changliang/zf"
	"changliang/zfzhi"
)

type Jicuziduan struct{}
type Ziduanbiaoji struct{}

func (jczd *Jicuziduan) Id() string {
	return zf.Zfs.Id(false)
}
func (jczd *Jicuziduan) Bianma() string {
	return zf.Zfs.Bianma(false)
}
func (jczd *Jicuziduan) Youxiaoxing() string {
	return zf.Zfs.Youxiaoxing(false)
}
func (jczd *Jicuziduan) Caozuoren() string {
	return zf.Zfs.Caozuoren(false)
}
func (jczd *Jicuziduan) Biaoji() string {
	return zf.Zfs.Biaoji(false)
}
func (jczd *Jicuziduan) Chuangjianriqi() string {
	return zf.Zfs.Chuangjianriqi(false)
}
func (jczd *Jicuziduan) Caozuoriqi() string {
	return zf.Zfs.Caozuoriqi(false)
}

func (zdbj *Ziduanbiaoji) Idbiaoji() string {
	return zf.Zfs.Zhujian(false)
}
func (zdbj *Ziduanbiaoji) Bianmabiaoji() string {
	return zf.Zfs.Sql(false) + zf.Zfs.Bianma(false)//DT_所有的编码都以DT_开头
}
func (zdbj *Ziduanbiaoji) Youxiaoxingbiaoji() string {
	return zfzhi.Zhi.Shuzi1w()
}
func (zdbj *Ziduanbiaoji) Caozuorenbiaoji() string {
	//mhsy_clz
	ret := zf.Zfs.Mhsy(true) + zfzhi.Zhi.Xhx()
	return ret
}
func (zdbj *Ziduanbiaoji) Biaojibiaoji() string {
	//标记信息"Jiance":1,"Zhengchang":1,"Wancheng":1,
	ret := zfzhi.Zhi.Syh() + zf.Zfs.Jiance(false) + zfzhi.Zhi.Syh() +
		zfzhi.Zhi.Mh() + zfzhi.Zhi.Shuzi1w() + zfzhi.Zhi.Dou() +
		zfzhi.Zhi.Syh() + zf.Zfs.Zhengchang(false) + zfzhi.Zhi.Syh() +
		zfzhi.Zhi.Mh() + zfzhi.Zhi.Shuzi1w() + zfzhi.Zhi.Dou() +
		zfzhi.Zhi.Syh() + zf.Zfs.Wancheng(false) + zfzhi.Zhi.Syh() +
		zfzhi.Zhi.Mh() + zfzhi.Zhi.Shuzi1w() + zfzhi.Zhi.Dou()
	return ret
}
func (zdbj *Ziduanbiaoji) Chuangjianriqibiaoji() string {
	//2017-01-01 00:00:00
	ret := zfzhi.Zhi.Shuzi2w() + zfzhi.Zhi.Shuzi0w() + zfzhi.Zhi.Shuzi1w() + zfzhi.Zhi.Shuzi7w() +
		zfzhi.Zhi.Jian() + zfzhi.Zhi.Shuzi0w() + zfzhi.Zhi.Shuzi1w() +
		zfzhi.Zhi.Jian() + zfzhi.Zhi.Shuzi0w() + zfzhi.Zhi.Shuzi1w() +
		zfzhi.Zhi.Kgf() + zfzhi.Zhi.Shuzi0w() + zfzhi.Zhi.Shuzi0w() +
		zfzhi.Zhi.Mh() + zfzhi.Zhi.Shuzi0w() + zfzhi.Zhi.Shuzi0w() +
		zfzhi.Zhi.Mh() + zfzhi.Zhi.Shuzi0w() + zfzhi.Zhi.Shuzi0w()

	return ret
}
func (zdbj *Ziduanbiaoji) Caozuoriqibiaoji() string {
	return zdbj.Chuangjianriqibiaoji()
}