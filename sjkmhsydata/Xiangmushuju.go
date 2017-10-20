package sjkmhsydata

import "changliang/zf"

type Xiangmushuju struct{}

func (xmsj *Xiangmushuju) Weiyibianma() string {
	return zf.Zfs.Weiyibianma(false)
}
func (xmsj *Xiangmushuju) Xiangmubianma() string {
	return zf.Zfs.Xiangmubianma(false)
}
func (xmsj *Xiangmushuju) Shijianbianma() string {
	return zf.Zfs.Shijianbianma(false)
}
func (xmsj *Xiangmushuju) Zhi() string {
	return zf.Zfs.Zhi(false)
}
