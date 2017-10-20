package sjkmhsydata

import "changliang/zf"

type Cujinfangan struct{}

func (cjfa *Cujinfangan) Mingcheng() string {
	return zf.Zfs.Mingcheng(false)
}
func (cjfa *Cujinfangan) Cankao() string {
	return zf.Zfs.Cankao(false)
}
func (cjfa *Cujinfangan) Neirong() string {
	return zf.Zfs.Neirong(false)
}
func (cjfa *Cujinfangan) Zhidingren() string {
	return zf.Zfs.Zhidingren(false)
}
func (cjfa *Cujinfangan) Weiyibianma() string {
	return zf.Zfs.Weiyibianma(false)
}
