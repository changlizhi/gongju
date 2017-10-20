package sjkmhsydata

import "changliang/zf"

type Zhanghaoshijian struct{}

func (zhsj *Zhanghaoshijian) Weiyibianma() string {
	return zf.Zfs.Weiyibianma(false)
}
func (zhsj *Zhanghaoshijian) Shijianbianma() string {
	return zf.Zfs.Shijianbianma(false)
}
