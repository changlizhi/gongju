package sjkmh

import "changliang/zf"

type Mhsydata struct {
	Juesemh                   *Juese
	Yinpinmh                  *Yinpin
	Tizhengzhibiaomh          *Tizhengzhibiao
	Shebeimh                  *Shebei
	Ziyuanmh                  *Ziyuan
	Zhanghaomh                *Zhanghao
	Yonghumh                  *Yonghu
	Tizhengzhibiaozumh        *Tizhengzhibiaozu
	Yinpinxiazaimh            *Yinpinxiazai
	Tizhengzhibiaozushijianmh *Tizhengzhibiaozushijian
	Tizhengzhibiaoshujumh     *Tizhengzhibiaoshuju
	Fuwufankuimh              *Fuwufankui
	Zhanghaoshijianmh         *Zhanghaoshijian
	Yinpinbofangmh            *Yinpinbofang
	Zidianmh                  *Zidian
	Zhanghaojuesemh           *Zhanghaojuese
	Liuyanmh                  *Liuyan
	Jibingmh                  *Jibing
	Tizhengzhibiaozuxiangmumh *Tizhengzhibiaozuxiangmu
	Shijianmh                 *Shijian
	Wenzhangmh                *Wenzhang
	Jueseziyuanmh             *Jueseziyuan
}
type Yinpin struct{}

func (yinpin *Yinpin) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (yinpin *Yinpin) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (yinpin *Yinpin) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (yinpin *Yinpin) Id() string {
	return zf.Zfs.Id(false)
}

type Tizhengzhibiao struct{}

func (xiangmu *Tizhengzhibiao) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (xiangmu *Tizhengzhibiao) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (xiangmu *Tizhengzhibiao) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (xiangmu *Tizhengzhibiao) Id() string {
	return zf.Zfs.Id(false)
}

type Shebei struct{}

func (shebei *Shebei) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (shebei *Shebei) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (shebei *Shebei) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (shebei *Shebei) Id() string {
	return zf.Zfs.Id(false)
}

type Ziyuan struct{}

func (ziyuan *Ziyuan) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (ziyuan *Ziyuan) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (ziyuan *Ziyuan) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (ziyuan *Ziyuan) Id() string {
	return zf.Zfs.Id(false)
}

type Zhanghao struct{}

func (zhanghao *Zhanghao) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (zhanghao *Zhanghao) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (zhanghao *Zhanghao) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (zhanghao *Zhanghao) Id() string {
	return zf.Zfs.Id(false)
}

type Yonghu struct{}

func (yonghu *Yonghu) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (yonghu *Yonghu) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (yonghu *Yonghu) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (yonghu *Yonghu) Id() string {
	return zf.Zfs.Id(false)
}

type Tizhengzhibiaozu struct{}

func (xiangmuzu *Tizhengzhibiaozu) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (xiangmuzu *Tizhengzhibiaozu) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (xiangmuzu *Tizhengzhibiaozu) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (xiangmuzu *Tizhengzhibiaozu) Id() string {
	return zf.Zfs.Id(false)
}

type Yinpinxiazai struct{}

func (yinpinxiazai *Yinpinxiazai) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (yinpinxiazai *Yinpinxiazai) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (yinpinxiazai *Yinpinxiazai) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (yinpinxiazai *Yinpinxiazai) Id() string {
	return zf.Zfs.Id(false)
}

type Tizhengzhibiaozushijian struct{}

func (xiangmuzushijian *Tizhengzhibiaozushijian) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (xiangmuzushijian *Tizhengzhibiaozushijian) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (xiangmuzushijian *Tizhengzhibiaozushijian) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (xiangmuzushijian *Tizhengzhibiaozushijian) Id() string {
	return zf.Zfs.Id(false)
}

type Tizhengzhibiaoshuju struct{}

func (xiangmushuju *Tizhengzhibiaoshuju) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (xiangmushuju *Tizhengzhibiaoshuju) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (xiangmushuju *Tizhengzhibiaoshuju) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (xiangmushuju *Tizhengzhibiaoshuju) Id() string {
	return zf.Zfs.Id(false)
}

type Fuwufankui struct{}

func (fuwufankui *Fuwufankui) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (fuwufankui *Fuwufankui) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (fuwufankui *Fuwufankui) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (fuwufankui *Fuwufankui) Id() string {
	return zf.Zfs.Id(false)
}

type Zhanghaoshijian struct{}

func (zhanghaoshijian *Zhanghaoshijian) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (zhanghaoshijian *Zhanghaoshijian) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (zhanghaoshijian *Zhanghaoshijian) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (zhanghaoshijian *Zhanghaoshijian) Id() string {
	return zf.Zfs.Id(false)
}

type Yinpinbofang struct{}

func (yinpinbofang *Yinpinbofang) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (yinpinbofang *Yinpinbofang) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (yinpinbofang *Yinpinbofang) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (yinpinbofang *Yinpinbofang) Id() string {
	return zf.Zfs.Id(false)
}

type Zidian struct{}

func (zidian *Zidian) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (zidian *Zidian) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (zidian *Zidian) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (zidian *Zidian) Id() string {
	return zf.Zfs.Id(false)
}

type Zhanghaojuese struct{}

func (zhanghaojuese *Zhanghaojuese) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (zhanghaojuese *Zhanghaojuese) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (zhanghaojuese *Zhanghaojuese) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (zhanghaojuese *Zhanghaojuese) Id() string {
	return zf.Zfs.Id(false)
}

type Liuyan struct{}

func (liuyan *Liuyan) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (liuyan *Liuyan) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (liuyan *Liuyan) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (liuyan *Liuyan) Id() string {
	return zf.Zfs.Id(false)
}

type Jibing struct{}

func (jibing *Jibing) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (jibing *Jibing) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (jibing *Jibing) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (jibing *Jibing) Id() string {
	return zf.Zfs.Id(false)
}

type Tizhengzhibiaozuxiangmu struct{}

func (xiangmuzuxiangmu *Tizhengzhibiaozuxiangmu) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (xiangmuzuxiangmu *Tizhengzhibiaozuxiangmu) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (xiangmuzuxiangmu *Tizhengzhibiaozuxiangmu) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (xiangmuzuxiangmu *Tizhengzhibiaozuxiangmu) Id() string {
	return zf.Zfs.Id(false)
}

type Juese struct{}

func (juese *Juese) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (juese *Juese) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (juese *Juese) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (juese *Juese) Id() string {
	return zf.Zfs.Id(false)
}

type Shijian struct{}

func (shijian *Shijian) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (shijian *Shijian) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (shijian *Shijian) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (shijian *Shijian) Id() string {
	return zf.Zfs.Id(false)
}

type Wenzhang struct{}

func (wenzhang *Wenzhang) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (wenzhang *Wenzhang) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (wenzhang *Wenzhang) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (wenzhang *Wenzhang) Id() string {
	return zf.Zfs.Id(false)
}

type Jueseziyuan struct{}

func (jueseziyuan *Jueseziyuan) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (jueseziyuan *Jueseziyuan) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (jueseziyuan *Jueseziyuan) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (jueseziyuan *Jueseziyuan) Id() string {
	return zf.Zfs.Id(false)
}
