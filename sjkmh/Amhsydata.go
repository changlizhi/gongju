package sjkmh

import "changliang/zf"

type Mhsydata struct {
	Jueseshujumh                    *Jueseshuju
	Juesejiegoumh                   *Juesejiegou
	Yinpinshujumh                   *Yinpinshuju
	Yinpinjiegoumh                  *Yinpinjiegou
	Tizhengzhibiaoshujumh           *Tizhengzhibiaoshuju
	Tizhengzhibiaojiegoumh          *Tizhengzhibiaojiegou
	Shebeishujumh                   *Shebeishuju
	Shebeijiegoumh                  *Shebeijiegou
	Ziyuanshujumh                   *Ziyuanshuju
	Ziyuanjiegoumh                  *Ziyuanjiegou
	Zhanghaoshujumh                 *Zhanghaoshuju
	Zhanghaojiegoumh                *Zhanghaojiegou
	Yonghushujumh                   *Yonghushuju
	Yonghujiegoumh                  *Yonghujiegou
	Tizhengzhibiaozushujumh         *Tizhengzhibiaozushuju
	Tizhengzhibiaozujiegoumh        *Tizhengzhibiaozujiegou
	Yinpinxiazaishujumh             *Yinpinxiazaishuju
	Yinpinxiazaijiegoumh            *Yinpinxiazaijiegou
	Tizhengzhibiaozushijianshujumh  *Tizhengzhibiaozushijianshuju
	Tizhengzhibiaozushijianjiegoumh *Tizhengzhibiaozushijianjiegou
	Jiancefenxishujumh              *Jiancefenxishuju
	Jiancefenxijiegoumh             *Jiancefenxijiegou
	Fuwufankuishujumh               *Fuwufankuishuju
	Fuwufankuijiegoumh              *Fuwufankuijiegou
	Zhanghaoshijianshujumh          *Zhanghaoshijianshuju
	Zhanghaoshijianjiegoumh         *Zhanghaoshijianjiegou
	Yinpinbofangshujumh             *Yinpinbofangshuju
	Yinpinbofangjiegoumh            *Yinpinbofangjiegou
	Zidianshujumh                   *Zidianshuju
	Zidianjiegoumh                  *Zidianjiegou
	Zhanghaojueseshujumh            *Zhanghaojueseshuju
	Zhanghaojuesejiegoumh           *Zhanghaojuesejiegou
	Liuyanshujumh                   *Liuyanshuju
	Liuyanjiegoumh                  *Liuyanjiegou
	Jibingshujumh                   *Jibingshuju
	Jibingjiegoumh                  *Jibingjiegou
	Tizhengzhibiaozuzhibiaoshujumh  *Tizhengzhibiaozuzhibiaoshuju
	Tizhengzhibiaozuzhibiaojiegoumh *Tizhengzhibiaozuzhibiaojiegou
	Shijianshujumh                  *Shijianshuju
	Shijianjiegoumh                 *Shijianjiegou
	Wenzhangshujumh                 *Wenzhangshuju
	Wenzhangjiegoumh                *Wenzhangjiegou
	Jueseziyuanshujumh              *Jueseziyuanshuju
	Jueseziyuanjiegoumh             *Jueseziyuanjiegou
}
type Jueseshuju struct{}
type Yinpinshuju struct{}
type Tizhengzhibiaoshuju struct{}
type Shebeishuju struct{}
type Ziyuanshuju struct{}
type Zhanghaoshuju struct{}
type Yonghushuju struct{}
type Tizhengzhibiaozushuju struct{}
type Yinpinxiazaishuju struct{}
type Tizhengzhibiaozushijianshuju struct{}
type Jiancefenxishuju struct{}
type Fuwufankuishuju struct{}
type Zhanghaoshijianshuju struct{}
type Yinpinbofangshuju struct{}
type Zidianshuju struct{}
type Zhanghaojueseshuju struct{}
type Liuyanshuju struct{}
type Jibingshuju struct{}
type Tizhengzhibiaozuzhibiaoshuju struct{}
type Shijianshuju struct{}
type Wenzhangshuju struct{}
type Jueseziyuanshuju struct{}

type Juesejiegou struct{}
type Yinpinjiegou struct{}
type Tizhengzhibiaojiegou struct{}
type Shebeijiegou struct{}
type Ziyuanjiegou struct{}

type Zhanghaojiegou struct{}

type Yonghujiegou struct{}

type Tizhengzhibiaozujiegou struct{}

type Yinpinxiazaijiegou struct{}

type Tizhengzhibiaozushijianjiegou struct{}

type Jiancefenxijiegou struct{}

type Fuwufankuijiegou struct{}

type Zhanghaoshijianjiegou struct{}

type Yinpinbofangjiegou struct{}

type Zidianjiegou struct{}

type Zhanghaojuesejiegou struct{}

type Liuyanjiegou struct{}

type Jibingjiegou struct{}

type Tizhengzhibiaozuzhibiaojiegou struct{}

type Shijianjiegou struct{}

type Wenzhangjiegou struct{}

type Jueseziyuanjiegou struct{}

func (jueseshuju *Jueseshuju) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (jueseshuju *Jueseshuju) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (jueseshuju *Jueseshuju) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (jueseshuju *Jueseshuju) Id() string {
	return zf.Zfs.Id(false)
}
func (yinpinshuju *Yinpinshuju) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (yinpinshuju *Yinpinshuju) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (yinpinshuju *Yinpinshuju) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (yinpinshuju *Yinpinshuju) Id() string {
	return zf.Zfs.Id(false)
}
func (tizhengzhibiaoshuju *Tizhengzhibiaoshuju) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (tizhengzhibiaoshuju *Tizhengzhibiaoshuju) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (tizhengzhibiaoshuju *Tizhengzhibiaoshuju) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (tizhengzhibiaoshuju *Tizhengzhibiaoshuju) Id() string {
	return zf.Zfs.Id(false)
}
func (shebeishuju *Shebeishuju) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (shebeishuju *Shebeishuju) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (shebeishuju *Shebeishuju) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (shebeishuju *Shebeishuju) Id() string {
	return zf.Zfs.Id(false)
}
func (ziyuanshuju *Ziyuanshuju) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (ziyuanshuju *Ziyuanshuju) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (ziyuanshuju *Ziyuanshuju) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (ziyuanshuju *Ziyuanshuju) Id() string {
	return zf.Zfs.Id(false)
}
func (zhanghaoshuju *Zhanghaoshuju) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (zhanghaoshuju *Zhanghaoshuju) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (zhanghaoshuju *Zhanghaoshuju) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (zhanghaoshuju *Zhanghaoshuju) Id() string {
	return zf.Zfs.Id(false)
}
func (yonghushuju *Yonghushuju) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (yonghushuju *Yonghushuju) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (yonghushuju *Yonghushuju) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (yonghushuju *Yonghushuju) Id() string {
	return zf.Zfs.Id(false)
}
func (tizhengzhibiaozushuju *Tizhengzhibiaozushuju) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (tizhengzhibiaozushuju *Tizhengzhibiaozushuju) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (tizhengzhibiaozushuju *Tizhengzhibiaozushuju) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (tizhengzhibiaozushuju *Tizhengzhibiaozushuju) Id() string {
	return zf.Zfs.Id(false)
}
func (yinpinxiazaishuju *Yinpinxiazaishuju) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (yinpinxiazaishuju *Yinpinxiazaishuju) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (yinpinxiazaishuju *Yinpinxiazaishuju) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (yinpinxiazaishuju *Yinpinxiazaishuju) Id() string {
	return zf.Zfs.Id(false)
}
func (tizhengzhibiaozushijianshuju *Tizhengzhibiaozushijianshuju) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (tizhengzhibiaozushijianshuju *Tizhengzhibiaozushijianshuju) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (tizhengzhibiaozushijianshuju *Tizhengzhibiaozushijianshuju) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (tizhengzhibiaozushijianshuju *Tizhengzhibiaozushijianshuju) Id() string {
	return zf.Zfs.Id(false)
}
func (jiancefenxishuju *Jiancefenxishuju) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (jiancefenxishuju *Jiancefenxishuju) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (jiancefenxishuju *Jiancefenxishuju) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (jiancefenxishuju *Jiancefenxishuju) Id() string {
	return zf.Zfs.Id(false)
}
func (fuwufankuishuju *Fuwufankuishuju) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (fuwufankuishuju *Fuwufankuishuju) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (fuwufankuishuju *Fuwufankuishuju) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (fuwufankuishuju *Fuwufankuishuju) Id() string {
	return zf.Zfs.Id(false)
}
func (zhanghaoshijianshuju *Zhanghaoshijianshuju) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (zhanghaoshijianshuju *Zhanghaoshijianshuju) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (zhanghaoshijianshuju *Zhanghaoshijianshuju) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (zhanghaoshijianshuju *Zhanghaoshijianshuju) Id() string {
	return zf.Zfs.Id(false)
}
func (yinpinbofangshuju *Yinpinbofangshuju) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (yinpinbofangshuju *Yinpinbofangshuju) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (yinpinbofangshuju *Yinpinbofangshuju) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (yinpinbofangshuju *Yinpinbofangshuju) Id() string {
	return zf.Zfs.Id(false)
}
func (zidianshuju *Zidianshuju) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (zidianshuju *Zidianshuju) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (zidianshuju *Zidianshuju) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (zidianshuju *Zidianshuju) Id() string {
	return zf.Zfs.Id(false)
}
func (zhanghaojueseshuju *Zhanghaojueseshuju) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (zhanghaojueseshuju *Zhanghaojueseshuju) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (zhanghaojueseshuju *Zhanghaojueseshuju) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (zhanghaojueseshuju *Zhanghaojueseshuju) Id() string {
	return zf.Zfs.Id(false)
}
func (liuyanshuju *Liuyanshuju) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (liuyanshuju *Liuyanshuju) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (liuyanshuju *Liuyanshuju) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (liuyanshuju *Liuyanshuju) Id() string {
	return zf.Zfs.Id(false)
}
func (jibingshuju *Jibingshuju) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (jibingshuju *Jibingshuju) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (jibingshuju *Jibingshuju) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (jibingshuju *Jibingshuju) Id() string {
	return zf.Zfs.Id(false)
}
func (tizhengzhibiaozuzhibiaoshuju *Tizhengzhibiaozuzhibiaoshuju) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (tizhengzhibiaozuzhibiaoshuju *Tizhengzhibiaozuzhibiaoshuju) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (tizhengzhibiaozuzhibiaoshuju *Tizhengzhibiaozuzhibiaoshuju) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (tizhengzhibiaozuzhibiaoshuju *Tizhengzhibiaozuzhibiaoshuju) Id() string {
	return zf.Zfs.Id(false)
}
func (shijianshuju *Shijianshuju) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (shijianshuju *Shijianshuju) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (shijianshuju *Shijianshuju) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (shijianshuju *Shijianshuju) Id() string {
	return zf.Zfs.Id(false)
}
func (wenzhangshuju *Wenzhangshuju) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (wenzhangshuju *Wenzhangshuju) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (wenzhangshuju *Wenzhangshuju) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (wenzhangshuju *Wenzhangshuju) Id() string {
	return zf.Zfs.Id(false)
}
func (jueseziyuanshuju *Jueseziyuanshuju) Zhutibianma() string {
	return zf.Zfs.Zhutibianma(false)
}
func (jueseziyuanshuju *Jueseziyuanshuju) Zhi() string {
	return zf.Zfs.Zhi(false)
}
func (jueseziyuanshuju *Jueseziyuanshuju) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (jueseziyuanshuju *Jueseziyuanshuju) Id() string {
	return zf.Zfs.Id(false)
}
func (juesejiegou *Juesejiegou) Id() string {
	return zf.Zfs.Id(false)
}
func (juesejiegou *Juesejiegou) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (juesejiegou *Juesejiegou) Ziduanleixing() string {
	return zf.Zfs.Ziduanleixing(false)
}
func (juesejiegou *Juesejiegou) Ziduanmingcheng() string {
	return zf.Zfs.Ziduanmingcheng(false)
}
func (juesejiegou *Juesejiegou) Ziduanchangdu() string {
	return zf.Zfs.Ziduanchangdu(false)
}
func (juesejiegou *Juesejiegou) Ziduanyanzheng() string {
	return zf.Zfs.Ziduanyanzheng(false)
}
func (yinpinjiegou *Yinpinjiegou) Id() string {
	return zf.Zfs.Id(false)
}
func (yinpinjiegou *Yinpinjiegou) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (yinpinjiegou *Yinpinjiegou) Ziduanleixing() string {
	return zf.Zfs.Ziduanleixing(false)
}
func (yinpinjiegou *Yinpinjiegou) Ziduanmingcheng() string {
	return zf.Zfs.Ziduanmingcheng(false)
}
func (yinpinjiegou *Yinpinjiegou) Ziduanchangdu() string {
	return zf.Zfs.Ziduanchangdu(false)
}
func (yinpinjiegou *Yinpinjiegou) Ziduanyanzheng() string {
	return zf.Zfs.Ziduanyanzheng(false)
}
func (tizhengzhibiaojiegou *Tizhengzhibiaojiegou) Id() string {
	return zf.Zfs.Id(false)
}
func (tizhengzhibiaojiegou *Tizhengzhibiaojiegou) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (tizhengzhibiaojiegou *Tizhengzhibiaojiegou) Ziduanleixing() string {
	return zf.Zfs.Ziduanleixing(false)
}
func (tizhengzhibiaojiegou *Tizhengzhibiaojiegou) Ziduanmingcheng() string {
	return zf.Zfs.Ziduanmingcheng(false)
}
func (tizhengzhibiaojiegou *Tizhengzhibiaojiegou) Ziduanchangdu() string {
	return zf.Zfs.Ziduanchangdu(false)
}
func (tizhengzhibiaojiegou *Tizhengzhibiaojiegou) Ziduanyanzheng() string {
	return zf.Zfs.Ziduanyanzheng(false)
}
func (shebeijiegou *Shebeijiegou) Id() string {
	return zf.Zfs.Id(false)
}
func (shebeijiegou *Shebeijiegou) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (shebeijiegou *Shebeijiegou) Ziduanleixing() string {
	return zf.Zfs.Ziduanleixing(false)
}
func (shebeijiegou *Shebeijiegou) Ziduanmingcheng() string {
	return zf.Zfs.Ziduanmingcheng(false)
}
func (shebeijiegou *Shebeijiegou) Ziduanchangdu() string {
	return zf.Zfs.Ziduanchangdu(false)
}
func (shebeijiegou *Shebeijiegou) Ziduanyanzheng() string {
	return zf.Zfs.Ziduanyanzheng(false)
}
func (ziyuanjiegou *Ziyuanjiegou) Id() string {
	return zf.Zfs.Id(false)
}
func (ziyuanjiegou *Ziyuanjiegou) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (ziyuanjiegou *Ziyuanjiegou) Ziduanleixing() string {
	return zf.Zfs.Ziduanleixing(false)
}
func (ziyuanjiegou *Ziyuanjiegou) Ziduanmingcheng() string {
	return zf.Zfs.Ziduanmingcheng(false)
}
func (ziyuanjiegou *Ziyuanjiegou) Ziduanchangdu() string {
	return zf.Zfs.Ziduanchangdu(false)
}
func (ziyuanjiegou *Ziyuanjiegou) Ziduanyanzheng() string {
	return zf.Zfs.Ziduanyanzheng(false)
}
func (zhanghaojiegou *Zhanghaojiegou) Id() string {
	return zf.Zfs.Id(false)
}
func (zhanghaojiegou *Zhanghaojiegou) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (zhanghaojiegou *Zhanghaojiegou) Ziduanleixing() string {
	return zf.Zfs.Ziduanleixing(false)
}
func (zhanghaojiegou *Zhanghaojiegou) Ziduanmingcheng() string {
	return zf.Zfs.Ziduanmingcheng(false)
}
func (zhanghaojiegou *Zhanghaojiegou) Ziduanchangdu() string {
	return zf.Zfs.Ziduanchangdu(false)
}
func (zhanghaojiegou *Zhanghaojiegou) Ziduanyanzheng() string {
	return zf.Zfs.Ziduanyanzheng(false)
}
func (yonghujiegou *Yonghujiegou) Id() string {
	return zf.Zfs.Id(false)
}
func (yonghujiegou *Yonghujiegou) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (yonghujiegou *Yonghujiegou) Ziduanleixing() string {
	return zf.Zfs.Ziduanleixing(false)
}
func (yonghujiegou *Yonghujiegou) Ziduanmingcheng() string {
	return zf.Zfs.Ziduanmingcheng(false)
}
func (yonghujiegou *Yonghujiegou) Ziduanchangdu() string {
	return zf.Zfs.Ziduanchangdu(false)
}
func (yonghujiegou *Yonghujiegou) Ziduanyanzheng() string {
	return zf.Zfs.Ziduanyanzheng(false)
}
func (tizhengzhibiaozujiegou *Tizhengzhibiaozujiegou) Id() string {
	return zf.Zfs.Id(false)
}
func (tizhengzhibiaozujiegou *Tizhengzhibiaozujiegou) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (tizhengzhibiaozujiegou *Tizhengzhibiaozujiegou) Ziduanleixing() string {
	return zf.Zfs.Ziduanleixing(false)
}
func (tizhengzhibiaozujiegou *Tizhengzhibiaozujiegou) Ziduanmingcheng() string {
	return zf.Zfs.Ziduanmingcheng(false)
}
func (tizhengzhibiaozujiegou *Tizhengzhibiaozujiegou) Ziduanchangdu() string {
	return zf.Zfs.Ziduanchangdu(false)
}
func (tizhengzhibiaozujiegou *Tizhengzhibiaozujiegou) Ziduanyanzheng() string {
	return zf.Zfs.Ziduanyanzheng(false)
}
func (yinpinxiazaijiegou *Yinpinxiazaijiegou) Id() string {
	return zf.Zfs.Id(false)
}
func (yinpinxiazaijiegou *Yinpinxiazaijiegou) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (yinpinxiazaijiegou *Yinpinxiazaijiegou) Ziduanleixing() string {
	return zf.Zfs.Ziduanleixing(false)
}
func (yinpinxiazaijiegou *Yinpinxiazaijiegou) Ziduanmingcheng() string {
	return zf.Zfs.Ziduanmingcheng(false)
}
func (yinpinxiazaijiegou *Yinpinxiazaijiegou) Ziduanchangdu() string {
	return zf.Zfs.Ziduanchangdu(false)
}
func (yinpinxiazaijiegou *Yinpinxiazaijiegou) Ziduanyanzheng() string {
	return zf.Zfs.Ziduanyanzheng(false)
}
func (tizhengzhibiaozushijianjiegou *Tizhengzhibiaozushijianjiegou) Id() string {
	return zf.Zfs.Id(false)
}
func (tizhengzhibiaozushijianjiegou *Tizhengzhibiaozushijianjiegou) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (tizhengzhibiaozushijianjiegou *Tizhengzhibiaozushijianjiegou) Ziduanleixing() string {
	return zf.Zfs.Ziduanleixing(false)
}
func (tizhengzhibiaozushijianjiegou *Tizhengzhibiaozushijianjiegou) Ziduanmingcheng() string {
	return zf.Zfs.Ziduanmingcheng(false)
}
func (tizhengzhibiaozushijianjiegou *Tizhengzhibiaozushijianjiegou) Ziduanchangdu() string {
	return zf.Zfs.Ziduanchangdu(false)
}
func (tizhengzhibiaozushijianjiegou *Tizhengzhibiaozushijianjiegou) Ziduanyanzheng() string {
	return zf.Zfs.Ziduanyanzheng(false)
}
func (jiancefenxijiegou *Jiancefenxijiegou) Id() string {
	return zf.Zfs.Id(false)
}
func (jiancefenxijiegou *Jiancefenxijiegou) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (jiancefenxijiegou *Jiancefenxijiegou) Ziduanleixing() string {
	return zf.Zfs.Ziduanleixing(false)
}
func (jiancefenxijiegou *Jiancefenxijiegou) Ziduanmingcheng() string {
	return zf.Zfs.Ziduanmingcheng(false)
}
func (jiancefenxijiegou *Jiancefenxijiegou) Ziduanchangdu() string {
	return zf.Zfs.Ziduanchangdu(false)
}
func (jiancefenxijiegou *Jiancefenxijiegou) Ziduanyanzheng() string {
	return zf.Zfs.Ziduanyanzheng(false)
}
func (fuwufankuijiegou *Fuwufankuijiegou) Id() string {
	return zf.Zfs.Id(false)
}
func (fuwufankuijiegou *Fuwufankuijiegou) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (fuwufankuijiegou *Fuwufankuijiegou) Ziduanleixing() string {
	return zf.Zfs.Ziduanleixing(false)
}
func (fuwufankuijiegou *Fuwufankuijiegou) Ziduanmingcheng() string {
	return zf.Zfs.Ziduanmingcheng(false)
}
func (fuwufankuijiegou *Fuwufankuijiegou) Ziduanchangdu() string {
	return zf.Zfs.Ziduanchangdu(false)
}
func (fuwufankuijiegou *Fuwufankuijiegou) Ziduanyanzheng() string {
	return zf.Zfs.Ziduanyanzheng(false)
}
func (zhanghaoshijianjiegou *Zhanghaoshijianjiegou) Id() string {
	return zf.Zfs.Id(false)
}
func (zhanghaoshijianjiegou *Zhanghaoshijianjiegou) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (zhanghaoshijianjiegou *Zhanghaoshijianjiegou) Ziduanleixing() string {
	return zf.Zfs.Ziduanleixing(false)
}
func (zhanghaoshijianjiegou *Zhanghaoshijianjiegou) Ziduanmingcheng() string {
	return zf.Zfs.Ziduanmingcheng(false)
}
func (zhanghaoshijianjiegou *Zhanghaoshijianjiegou) Ziduanchangdu() string {
	return zf.Zfs.Ziduanchangdu(false)
}
func (zhanghaoshijianjiegou *Zhanghaoshijianjiegou) Ziduanyanzheng() string {
	return zf.Zfs.Ziduanyanzheng(false)
}
func (yinpinbofangjiegou *Yinpinbofangjiegou) Id() string {
	return zf.Zfs.Id(false)
}
func (yinpinbofangjiegou *Yinpinbofangjiegou) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (yinpinbofangjiegou *Yinpinbofangjiegou) Ziduanleixing() string {
	return zf.Zfs.Ziduanleixing(false)
}
func (yinpinbofangjiegou *Yinpinbofangjiegou) Ziduanmingcheng() string {
	return zf.Zfs.Ziduanmingcheng(false)
}
func (yinpinbofangjiegou *Yinpinbofangjiegou) Ziduanchangdu() string {
	return zf.Zfs.Ziduanchangdu(false)
}
func (yinpinbofangjiegou *Yinpinbofangjiegou) Ziduanyanzheng() string {
	return zf.Zfs.Ziduanyanzheng(false)
}
func (zidianjiegou *Zidianjiegou) Id() string {
	return zf.Zfs.Id(false)
}
func (zidianjiegou *Zidianjiegou) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (zidianjiegou *Zidianjiegou) Ziduanleixing() string {
	return zf.Zfs.Ziduanleixing(false)
}
func (zidianjiegou *Zidianjiegou) Ziduanmingcheng() string {
	return zf.Zfs.Ziduanmingcheng(false)
}
func (zidianjiegou *Zidianjiegou) Ziduanchangdu() string {
	return zf.Zfs.Ziduanchangdu(false)
}
func (zidianjiegou *Zidianjiegou) Ziduanyanzheng() string {
	return zf.Zfs.Ziduanyanzheng(false)
}
func (zhanghaojuesejiegou *Zhanghaojuesejiegou) Id() string {
	return zf.Zfs.Id(false)
}
func (zhanghaojuesejiegou *Zhanghaojuesejiegou) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (zhanghaojuesejiegou *Zhanghaojuesejiegou) Ziduanleixing() string {
	return zf.Zfs.Ziduanleixing(false)
}
func (zhanghaojuesejiegou *Zhanghaojuesejiegou) Ziduanmingcheng() string {
	return zf.Zfs.Ziduanmingcheng(false)
}
func (zhanghaojuesejiegou *Zhanghaojuesejiegou) Ziduanchangdu() string {
	return zf.Zfs.Ziduanchangdu(false)
}
func (zhanghaojuesejiegou *Zhanghaojuesejiegou) Ziduanyanzheng() string {
	return zf.Zfs.Ziduanyanzheng(false)
}
func (liuyanjiegou *Liuyanjiegou) Id() string {
	return zf.Zfs.Id(false)
}
func (liuyanjiegou *Liuyanjiegou) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (liuyanjiegou *Liuyanjiegou) Ziduanleixing() string {
	return zf.Zfs.Ziduanleixing(false)
}
func (liuyanjiegou *Liuyanjiegou) Ziduanmingcheng() string {
	return zf.Zfs.Ziduanmingcheng(false)
}
func (liuyanjiegou *Liuyanjiegou) Ziduanchangdu() string {
	return zf.Zfs.Ziduanchangdu(false)
}
func (liuyanjiegou *Liuyanjiegou) Ziduanyanzheng() string {
	return zf.Zfs.Ziduanyanzheng(false)
}
func (jibingjiegou *Jibingjiegou) Id() string {
	return zf.Zfs.Id(false)
}
func (jibingjiegou *Jibingjiegou) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (jibingjiegou *Jibingjiegou) Ziduanleixing() string {
	return zf.Zfs.Ziduanleixing(false)
}
func (jibingjiegou *Jibingjiegou) Ziduanmingcheng() string {
	return zf.Zfs.Ziduanmingcheng(false)
}
func (jibingjiegou *Jibingjiegou) Ziduanchangdu() string {
	return zf.Zfs.Ziduanchangdu(false)
}
func (jibingjiegou *Jibingjiegou) Ziduanyanzheng() string {
	return zf.Zfs.Ziduanyanzheng(false)
}
func (tizhengzhibiaozuzhibiaojiegou *Tizhengzhibiaozuzhibiaojiegou) Id() string {
	return zf.Zfs.Id(false)
}
func (tizhengzhibiaozuzhibiaojiegou *Tizhengzhibiaozuzhibiaojiegou) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (tizhengzhibiaozuzhibiaojiegou *Tizhengzhibiaozuzhibiaojiegou) Ziduanleixing() string {
	return zf.Zfs.Ziduanleixing(false)
}
func (tizhengzhibiaozuzhibiaojiegou *Tizhengzhibiaozuzhibiaojiegou) Ziduanmingcheng() string {
	return zf.Zfs.Ziduanmingcheng(false)
}
func (tizhengzhibiaozuzhibiaojiegou *Tizhengzhibiaozuzhibiaojiegou) Ziduanchangdu() string {
	return zf.Zfs.Ziduanchangdu(false)
}
func (tizhengzhibiaozuzhibiaojiegou *Tizhengzhibiaozuzhibiaojiegou) Ziduanyanzheng() string {
	return zf.Zfs.Ziduanyanzheng(false)
}
func (shijianjiegou *Shijianjiegou) Id() string {
	return zf.Zfs.Id(false)
}
func (shijianjiegou *Shijianjiegou) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (shijianjiegou *Shijianjiegou) Ziduanleixing() string {
	return zf.Zfs.Ziduanleixing(false)
}
func (shijianjiegou *Shijianjiegou) Ziduanmingcheng() string {
	return zf.Zfs.Ziduanmingcheng(false)
}
func (shijianjiegou *Shijianjiegou) Ziduanchangdu() string {
	return zf.Zfs.Ziduanchangdu(false)
}
func (shijianjiegou *Shijianjiegou) Ziduanyanzheng() string {
	return zf.Zfs.Ziduanyanzheng(false)
}
func (wenzhangjiegou *Wenzhangjiegou) Id() string {
	return zf.Zfs.Id(false)
}
func (wenzhangjiegou *Wenzhangjiegou) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (wenzhangjiegou *Wenzhangjiegou) Ziduanleixing() string {
	return zf.Zfs.Ziduanleixing(false)
}
func (wenzhangjiegou *Wenzhangjiegou) Ziduanmingcheng() string {
	return zf.Zfs.Ziduanmingcheng(false)
}
func (wenzhangjiegou *Wenzhangjiegou) Ziduanchangdu() string {
	return zf.Zfs.Ziduanchangdu(false)
}
func (wenzhangjiegou *Wenzhangjiegou) Ziduanyanzheng() string {
	return zf.Zfs.Ziduanyanzheng(false)
}
func (jueseziyuanjiegou *Jueseziyuanjiegou) Id() string {
	return zf.Zfs.Id(false)
}
func (jueseziyuanjiegou *Jueseziyuanjiegou) Ziduanbianma() string {
	return zf.Zfs.Ziduanbianma(false)
}
func (jueseziyuanjiegou *Jueseziyuanjiegou) Ziduanleixing() string {
	return zf.Zfs.Ziduanleixing(false)
}
func (jueseziyuanjiegou *Jueseziyuanjiegou) Ziduanmingcheng() string {
	return zf.Zfs.Ziduanmingcheng(false)
}
func (jueseziyuanjiegou *Jueseziyuanjiegou) Ziduanchangdu() string {
	return zf.Zfs.Ziduanchangdu(false)
}
func (jueseziyuanjiegou *Jueseziyuanjiegou) Ziduanyanzheng() string {
	return zf.Zfs.Ziduanyanzheng(false)
}
