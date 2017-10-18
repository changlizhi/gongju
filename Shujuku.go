package gongju

import "changliang/zf"

type Jicuziduan struct{}

func (jczd *Jicuziduan) Bianma() string {
	return zf.Zfs.Bianma(false)
}
func (jczd *Jicuziduan) Paixu() string {
	return zf.Zfs.Paixu(false)
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
func (jczd *Jicuziduan) Youxiaoxing() string {
	return zf.Zfs.Youxiaoxing(false)
}
func (jczd *Jicuziduan) Id() string {
	return zf.Zfs.Id(false)
}
func (jczd *Jicuziduan) Caozuoren() string {
	return zf.Zfs.Caozuoren(false)
}
