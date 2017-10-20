package sjkmhsydata

import "changliang/zf"

type Shijian struct{}

func (sj *Shijian) Mingcheng() string {
	return zf.Zfs.Mingcheng(false)
}
func (sj *Shijian) Miaoshu() string {
	return zf.Zfs.Miaoshu(false)
}
func (sj *Shijian) Qiyong() string {
	return zf.Zfs.Qiyong(false)
}
func (sj *Shijian) Xiangmuzubianma() string {
	return zf.Zfs.Xiangmuzubianma(false)
}
