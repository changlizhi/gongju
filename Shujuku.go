package gongju

import (
	"changliang/zf"
	"changliang/zfzhi"
)

type Jicuziduan struct{}
type Ziduanbiaoji struct{}
type Biaobiaoji struct{}

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

func (bbj *Biaobiaoji) Juesebiaoji() string {
	return zf.Zfs.Yiban(false)
}
func (bbj *Biaobiaoji) Dtziyuanbiaoji() string {
	return zf.Zfs.Yiban(false)
}
func (bbj *Biaobiaoji) Juesequanxianbiaoji() string {
	return zf.Zfs.Yiban(false)
}
func (bbj *Biaobiaoji) Jueseziyuanbiaoji() string {
	return zf.Zfs.Yiban(false)
}
func (bbj *Biaobiaoji) Quanxianbiaoji() string {
	return zf.Zfs.Yiban(false)
}
func (bbj *Biaobiaoji) Xinxibiaoji() string {
	return zf.Zfs.Yiban(false)
}
func (bbj *Biaobiaoji) Xinxijuesebiaoji() string {
	return zf.Zfs.Yiban(false)
}
func (bbj *Biaobiaoji) Yanzhengbiaoji() string {
	return zf.Zfs.Yiban(false)
}
func (bbj *Biaobiaoji) Yanzhengleixingbiaoji() string {
	return zf.Zfs.Yiban(false)
}
func (bbj *Biaobiaoji) Zhanghaojuesebiaoji() string {
	return zf.Zfs.Yiban(false)
}
func (bbj *Biaobiaoji) Ziyuanbiaoji() string {
	return zf.Zfs.Yemianlianjie(false)
}
func (bbj *Biaobiaoji) Shebeibiaoji() string {
	return zf.Zfs.Yiban(false)
}
func (bbj *Biaobiaoji) Yinpinbiaoji() string {
	return zf.Zfs.Yiban(false)
}
func (bbj *Biaobiaoji) Yinpinbofangbiaoji() string {
	return zf.Zfs.Yiban(false)
}
func (bbj *Biaobiaoji) Yinpinxiazaibiaoji() string {
	return zf.Zfs.Yiban(false)
}
func (bbj *Biaobiaoji) Xiangmubiaoji() string {
	return zf.Zfs.Yiban(false)
}
func (bbj *Biaobiaoji) Xiangmuzubiaoji() string {
	return zf.Zfs.Yiban(false)
}
func (bbj *Biaobiaoji) Xiangmuzuxiangmubiaoji() string {
	return zf.Zfs.Yiban(false)
}
func (bbj *Biaobiaoji) Shijianbiaoji() string {
	return zf.Zfs.Yiban(false)
}
func (bbj *Biaobiaoji) Zhanghaoshijianbiaoji() string {
	return zf.Zfs.Yiban(false)
}
func (bbj *Biaobiaoji) Zhanghaobiaoji() string {
	return zf.Zfs.Yiban(false)
}
func (bbj *Biaobiaoji) Zidianbiaoji() string {
	return zf.Zfs.Yiban(false)
}
func (bbj *Biaobiaoji) Yonghubiaoji() string {
	return zf.Zfs.Yiban(false)
}
func (bbj *Biaobiaoji) Yonghujibingbiaoji() string {
	return zf.Zfs.Yiban(false)
}
func (bbj *Biaobiaoji) Wenzhangbiaoji() string {
	return zf.Zfs.Yiban(false)
}
func (bbj *Biaobiaoji) Liuyanbiaoji() string {
	return zf.Zfs.Yiban(false)
}
func (bbj *Biaobiaoji) Fuwufankuibiaoji() string {
	return zf.Zfs.Yiban(false)
}
func (bbj *Biaobiaoji) Xiangmushujubiaoji() string {
	return zf.Zfs.Yiban(false)
}
func (bbj *Biaobiaoji) Xiangmuzushijianbiaoji() string {
	return zf.Zfs.Yiban(false)
}
func (bbj *Biaobiaoji) Jibingbiaoji() string {
	return zf.Zfs.Yiban(false)
}
func (bbj *Biaobiaoji) Cujinfanganbiaoji() string {
	return zf.Zfs.Yiban(false)
}

