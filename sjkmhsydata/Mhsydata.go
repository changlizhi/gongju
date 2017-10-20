package sjkmhsydata

import "changliang/zf"

type Sjkmhsydata struct{}

var Sjkmhsydatas = &Sjkmhsydata{}

func (mhsydata *Mhsydata) Mingcheng(xiaoxie bool) string {
	return zf.Zfs.Mhsydata(xiaoxie)
}

var Mhsydatas = &Mhsydata{}

type Mhsydata struct {
	Dtzidian           *Zidian
	Dtyinpinxiazai     *Yinpinxiazai
	Dtyonghu           *Yonghu
	Dtxiangmuzuxiangmu *Xiangmuzuxiangmu
	Dtzhanghaojuese    *Zhanghaojuese
	Dtjibing           *Jibing
	Dtzhanghaoshijian  *Zhanghaoshijian
	Dtfuwufankui       *Fuwufankui
	Dtchujinfangan     *Cujinfangan
	Dtyinpinbofang     *Yinpinbofang
	Dtshebei           *Shebei
	Dtzhanghao         *Zhanghao
	Dtshijian          *Shijian
	Dtxiangmu          *Xiangmu
	Dtjueseziyuan      *Jueseziyuan
	Dtjuese            *Juese
	Dtyinpin           *Yinpin
	Dtziyuan           *Ziyuan
	Dtliuyan           *Liuyan
	Dtwenzhang         *Wenzhang
	Dtxiangmuzu        *Xiangmuzu
	Dtxiangmuzushijian *Xiangmuzushijian
	Dtxiangmushuju     *Xiangmushuju
}
