package sjkmhsydata

import "changliang/zf"

type Xiangmuzushijian struct{}

func (smzsj *Xiangmuzushijian) Xiangmuzubianma() string {
	return zf.Zfs.Xiangmuzubianma(false)
}
func (smzsj *Xiangmuzushijian) Shijianbianma() string {
	return zf.Zfs.Shijianbianma(false)
}
