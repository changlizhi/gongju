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
	Jianceshujumh             *Jianceshuju
	Fuwufankuimh              *Fuwufankui
	Zhanghaoshijianmh         *Zhanghaoshijian
	Yinpinbofangmh            *Yinpinbofang
	Zidianmh                  *Zidian
	Zhanghaojuesemh           *Zhanghaojuese
	Liuyanmh                  *Liuyan
	Jibingmh                  *Jibing
	Tizhengzhibiaozuzhibiaomh *Tizhengzhibiaozuzhibiao
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

func (tizhengzhibiao *Tizhengzhibiao) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (tizhengzhibiao *Tizhengzhibiao) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (tizhengzhibiao *Tizhengzhibiao) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (tizhengzhibiao *Tizhengzhibiao) Id() string {
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

func (tizhengzhibiaozu *Tizhengzhibiaozu) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (tizhengzhibiaozu *Tizhengzhibiaozu) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (tizhengzhibiaozu *Tizhengzhibiaozu) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (tizhengzhibiaozu *Tizhengzhibiaozu) Id() string {
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

func (tizhengzhibiaozushijian *Tizhengzhibiaozushijian) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (tizhengzhibiaozushijian *Tizhengzhibiaozushijian) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (tizhengzhibiaozushijian *Tizhengzhibiaozushijian) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (tizhengzhibiaozushijian *Tizhengzhibiaozushijian) Id() string {
	return zf.Zfs.Id(false)
}

type Jianceshuju struct{}

func (jianceshuju *Jianceshuju) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (jianceshuju *Jianceshuju) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (jianceshuju *Jianceshuju) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (jianceshuju *Jianceshuju) Id() string {
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

type Tizhengzhibiaozuzhibiao struct{}

func (tizhengzhibiaozutizhengzhibiao *Tizhengzhibiaozuzhibiao) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (tizhengzhibiaozutizhengzhibiao *Tizhengzhibiaozuzhibiao) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (tizhengzhibiaozutizhengzhibiao *Tizhengzhibiaozuzhibiao) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (tizhengzhibiaozutizhengzhibiao *Tizhengzhibiaozuzhibiao) Id() string {
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
