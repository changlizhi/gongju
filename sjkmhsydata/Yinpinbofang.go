package sjkmhsydata

import "changliang/zf"

type Yinpinbofang struct{}

func (ypbf *Yinpinbofang) Bofangshijian() string {
	return zf.Zfs.Bofangshijian(false)
}
func (ypbf *Yinpinbofang) Yinpinbianma() string {
	return zf.Zfs.Yinpinbianma(false)
}
func (ypbf *Yinpinbofang) Weiyibianma() string {
	return zf.Zfs.Weiyibianma(false)
}
