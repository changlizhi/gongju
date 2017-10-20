package sjkmhsydata

import "changliang/zf"

type Fuwufankui struct{}

func (fwfk *Fuwufankui) Fuwuganshou() string {
	return zf.Zfs.Fuwuganshou(false)
}
func (fwfk *Fuwufankui) Shentiganshou() string {
	return zf.Zfs.Shentiganshou(false)
}
func (fwfk *Fuwufankui) Xiaoguokeping() string {
	return zf.Zfs.Xiaoguokeping(false)
}
func (fwfk *Fuwufankui) Weiyibianma() string {
	return zf.Zfs.Weiyibianma(false)
}
