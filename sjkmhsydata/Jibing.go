package sjkmhsydata

import "changliang/zf"

type Jibing struct{}

func (jb *Jibing) Weiyibianma() string {
	return zf.Zfs.Weiyibianma(false)
}
func (jb *Jibing) Tangniaobing() string {
	//这个字段空指着下面四个字段是否有值
	return zf.Zfs.Tangniaobing(false)
}
func (jb *Jibing) Kongfu() string {
	return zf.Zfs.Kongfu(false)
}
func (jb *Jibing) Canhouerxiaoshi() string {
	return zf.Zfs.Canhouerxiaoshi(false)
}
func (jb *Jibing) Tanghuaxuehongdanbai() string {
	return zf.Zfs.Tanghuaxuehongdanbai(false)
}
func (jb *Jibing) Xuetangchuli() string {
	//这个字段又控制下一个字段的显示
	return zf.Zfs.Xuetangchuli(false)
}
func (jb *Jibing) Jiliang() string {
	return zf.Zfs.Jiliang(false)
}
func (jb *Jibing) Xinnaoxueguanbing() string {
	//是一个字符数组，心脑血管疾病的类型逗号分割
	return zf.Zfs.Xinnaoxueguanbing(false)
}
func (jb *Jibing) Xinnaoxueguanbingqita() string {
	//心脑血管疾病的其他类型
	return zf.Zfs.Xinnaoxueguanbing(false) + zf.Zfs.Qita(true)
}
func (jb *Jibing) Pifubing() string {
	//是一个字符数组，皮肤疾病的类型逗号分割
	return zf.Zfs.Pifubing(false)
}
func (jb *Jibing) Pifubingqita() string {
	//皮肤病的其他类型
	return zf.Zfs.Pifubing(false) + zf.Zfs.Qita(true)
}
func (jb *Jibing) Ganran() string {
	//是一个字符数组，感染病的类型逗号分割
	return zf.Zfs.Ganran(false)
}
func (jb *Jibing) Ganranqita() string {
	//感染类疾病的其他类型
	return zf.Zfs.Ganran(false) + zf.Zfs.Qita(true)
}
func (jb *Jibing) Exingzhongliu() string {
	//是一个字符数组，恶性肿瘤的类型逗号分割
	return zf.Zfs.Exingzhongliu(false)
}
func (jb *Jibing) Qita() string {
	return zf.Zfs.Qita(false)
}
func (jb *Jibing) Buliangfanying() string {
	return zf.Zfs.Buliangfanying(false)
}
func (jb *Jibing) Jiazubingshi() string { //家族病史字符串逗号分割
	return zf.Zfs.Jiazubingshi(false)
}
