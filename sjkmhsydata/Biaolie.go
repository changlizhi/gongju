package sjkmhsydata

import (
	"changliang/zf"
)

func (sjkmhsydata *Sjkmhsydata) JueseMiaoshu(xiaoxie bool) string {
	return zf.Zfs.Juese(xiaoxie) + zf.Zfs.Miaoshu(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JueseMingcheng(xiaoxie bool) string {
	return zf.Zfs.Juese(xiaoxie) + zf.Zfs.Mingcheng(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) JueseId(xiaoxie bool) string {
	return zf.Zfs.Juese(xiaoxie) + zf.Zfs.Id(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JueseBianma(xiaoxie bool) string {
	return zf.Zfs.Juese(xiaoxie) + zf.Zfs.Bianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JuesePaixu(xiaoxie bool) string {
	return zf.Zfs.Juese(xiaoxie) + zf.Zfs.Paixu(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JueseBiaoji(xiaoxie bool) string {
	return zf.Zfs.Juese(xiaoxie) + zf.Zfs.Biaoji(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JueseChuangjianriqi(xiaoxie bool) string {
	return zf.Zfs.Juese(xiaoxie) + zf.Zfs.Chuangjianriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JueseCaozuoriqi(xiaoxie bool) string {
	return zf.Zfs.Juese(xiaoxie) + zf.Zfs.Caozuoriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JueseYouxiaoxing(xiaoxie bool) string {
	return zf.Zfs.Juese(xiaoxie) + zf.Zfs.Youxiaoxing(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JueseCaozuoren(xiaoxie bool) string {
	return zf.Zfs.Juese(xiaoxie) + zf.Zfs.Caozuoren(xiaoxie)

}

//2017/10/11 22:49:22
func (sjkmhsydata *Sjkmhsydata) ZhanghaoWeiyibianma(xiaoxie bool) string {
	return zf.Zfs.Zhanghao(xiaoxie) + zf.Zfs.Weiyibianma(xiaoxie)
}
func (sjkmhsydata *Sjkmhsydata) ZhanghaoZhi(xiaoxie bool) string {
	return zf.Zfs.Zhanghao(xiaoxie) + zf.Zfs.Zhi(xiaoxie)
}
func (sjkmhsydata *Sjkmhsydata) ZhanghaoLeixing(xiaoxie bool) string {
	return zf.Zfs.Zhanghao(xiaoxie) + zf.Zfs.Leixing(xiaoxie)
}
func (sjkmhsydata *Sjkmhsydata) ZhanghaoId(xiaoxie bool) string {
	return zf.Zfs.Zhanghao(xiaoxie) + zf.Zfs.Id(xiaoxie)
}
func (sjkmhsydata *Sjkmhsydata) ZhanghaoBianma(xiaoxie bool) string {
	return zf.Zfs.Zhanghao(xiaoxie) + zf.Zfs.Bianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZhanghaoPaixu(xiaoxie bool) string {
	return zf.Zfs.Zhanghao(xiaoxie) + zf.Zfs.Paixu(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZhanghaoBiaoji(xiaoxie bool) string {
	return zf.Zfs.Zhanghao(xiaoxie) + zf.Zfs.Biaoji(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZhanghaoChuangjianriqi(xiaoxie bool) string {
	return zf.Zfs.Zhanghao(xiaoxie) + zf.Zfs.Chuangjianriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZhanghaoCaozuoriqi(xiaoxie bool) string {
	return zf.Zfs.Zhanghao(xiaoxie) + zf.Zfs.Caozuoriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZhanghaoYouxiaoxing(xiaoxie bool) string {
	return zf.Zfs.Zhanghao(xiaoxie) + zf.Zfs.Youxiaoxing(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZhanghaoCaozuoren(xiaoxie bool) string {
	return zf.Zfs.Zhanghao(xiaoxie) + zf.Zfs.Caozuoren(xiaoxie)

}

//2017/10/11 22:49:22
func (sjkmhsydata *Sjkmhsydata) ZiyuanLeixing(xiaoxie bool) string {
	return zf.Zfs.Ziyuan(xiaoxie) + zf.Zfs.Leixing(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZiyuanLujing(xiaoxie bool) string {
	return zf.Zfs.Ziyuan(xiaoxie) + zf.Zfs.Lujing(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZiyuanFubianma(xiaoxie bool) string {
	return zf.Zfs.Ziyuan(xiaoxie) + zf.Zfs.Fubianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZiyuanMingcheng(xiaoxie bool) string {
	return zf.Zfs.Ziyuan(xiaoxie) + zf.Zfs.Mingcheng(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZiyuanId(xiaoxie bool) string {
	return zf.Zfs.Ziyuan(xiaoxie) + zf.Zfs.Id(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZiyuanBianma(xiaoxie bool) string {
	return zf.Zfs.Ziyuan(xiaoxie) + zf.Zfs.Bianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZiyuanPaixu(xiaoxie bool) string {
	return zf.Zfs.Ziyuan(xiaoxie) + zf.Zfs.Paixu(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZiyuanBiaoji(xiaoxie bool) string {
	return zf.Zfs.Ziyuan(xiaoxie) + zf.Zfs.Biaoji(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZiyuanChuangjianriqi(xiaoxie bool) string {
	return zf.Zfs.Ziyuan(xiaoxie) + zf.Zfs.Chuangjianriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZiyuanCaozuoriqi(xiaoxie bool) string {
	return zf.Zfs.Ziyuan(xiaoxie) + zf.Zfs.Caozuoriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZiyuanYouxiaoxing(xiaoxie bool) string {
	return zf.Zfs.Ziyuan(xiaoxie) + zf.Zfs.Youxiaoxing(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZiyuanCaozuoren(xiaoxie bool) string {
	return zf.Zfs.Ziyuan(xiaoxie) + zf.Zfs.Caozuoren(xiaoxie)

}

//2017/10/11 22:49:22
func (sjkmhsydata *Sjkmhsydata) ZhanghaojueseJuesebianma(xiaoxie bool) string {
	return zf.Zfs.Zhanghaojuese(xiaoxie) + zf.Zfs.Juesebianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZhanghaojueseWeiyibianma(xiaoxie bool) string {
	return zf.Zfs.Zhanghaojuese(xiaoxie) + zf.Zfs.Weiyibianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZhanghaojueseId(xiaoxie bool) string {
	return zf.Zfs.Zhanghaojuese(xiaoxie) + zf.Zfs.Id(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZhanghaojueseBianma(xiaoxie bool) string {
	return zf.Zfs.Zhanghaojuese(xiaoxie) + zf.Zfs.Bianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZhanghaojuesePaixu(xiaoxie bool) string {
	return zf.Zfs.Zhanghaojuese(xiaoxie) + zf.Zfs.Paixu(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZhanghaojueseBiaoji(xiaoxie bool) string {
	return zf.Zfs.Zhanghaojuese(xiaoxie) + zf.Zfs.Biaoji(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZhanghaojueseChuangjianriqi(xiaoxie bool) string {
	return zf.Zfs.Zhanghaojuese(xiaoxie) + zf.Zfs.Chuangjianriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZhanghaojueseCaozuoriqi(xiaoxie bool) string {
	return zf.Zfs.Zhanghaojuese(xiaoxie) + zf.Zfs.Caozuoriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZhanghaojueseYouxiaoxing(xiaoxie bool) string {
	return zf.Zfs.Zhanghaojuese(xiaoxie) + zf.Zfs.Youxiaoxing(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZhanghaojueseCaozuoren(xiaoxie bool) string {
	return zf.Zfs.Zhanghaojuese(xiaoxie) + zf.Zfs.Caozuoren(xiaoxie)

}

//2017/10/11 22:49:22
func (sjkmhsydata *Sjkmhsydata) JueseziyuanZiyuanbianma(xiaoxie bool) string {
	return zf.Zfs.Jueseziyuan(xiaoxie) + zf.Zfs.Ziyuanbianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JueseziyuanJuesebianma(xiaoxie bool) string {
	return zf.Zfs.Jueseziyuan(xiaoxie) + zf.Zfs.Juesebianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JueseziyuanId(xiaoxie bool) string {
	return zf.Zfs.Jueseziyuan(xiaoxie) + zf.Zfs.Id(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JueseziyuanBianma(xiaoxie bool) string {
	return zf.Zfs.Jueseziyuan(xiaoxie) + zf.Zfs.Bianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JueseziyuanPaixu(xiaoxie bool) string {
	return zf.Zfs.Jueseziyuan(xiaoxie) + zf.Zfs.Paixu(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JueseziyuanBiaoji(xiaoxie bool) string {
	return zf.Zfs.Jueseziyuan(xiaoxie) + zf.Zfs.Biaoji(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JueseziyuanChuangjianriqi(xiaoxie bool) string {
	return zf.Zfs.Jueseziyuan(xiaoxie) + zf.Zfs.Chuangjianriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JueseziyuanCaozuoriqi(xiaoxie bool) string {
	return zf.Zfs.Jueseziyuan(xiaoxie) + zf.Zfs.Caozuoriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JueseziyuanYouxiaoxing(xiaoxie bool) string {
	return zf.Zfs.Jueseziyuan(xiaoxie) + zf.Zfs.Youxiaoxing(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JueseziyuanCaozuoren(xiaoxie bool) string {
	return zf.Zfs.Jueseziyuan(xiaoxie) + zf.Zfs.Caozuoren(xiaoxie)

}

//2017/10/11 22:49:22
func (sjkmhsydata *Sjkmhsydata) XiangmuYixuebianma(xiaoxie bool) string {
	return zf.Zfs.Xiangmu(xiaoxie) + zf.Zfs.Yixuebianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuMiaoshu(xiaoxie bool) string {
	return zf.Zfs.Xiangmu(xiaoxie) + zf.Zfs.Miaoshu(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuLeixing(xiaoxie bool) string {
	return zf.Zfs.Xiangmu(xiaoxie) + zf.Zfs.Leixing(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuDanwei(xiaoxie bool) string {
	return zf.Zfs.Xiangmu(xiaoxie) + zf.Zfs.Danwei(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuMingcheng(xiaoxie bool) string {
	return zf.Zfs.Xiangmu(xiaoxie) + zf.Zfs.Mingcheng(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuId(xiaoxie bool) string {
	return zf.Zfs.Xiangmu(xiaoxie) + zf.Zfs.Id(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuBianma(xiaoxie bool) string {
	return zf.Zfs.Xiangmu(xiaoxie) + zf.Zfs.Bianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuPaixu(xiaoxie bool) string {
	return zf.Zfs.Xiangmu(xiaoxie) + zf.Zfs.Paixu(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuBiaoji(xiaoxie bool) string {
	return zf.Zfs.Xiangmu(xiaoxie) + zf.Zfs.Biaoji(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuChuangjianriqi(xiaoxie bool) string {
	return zf.Zfs.Xiangmu(xiaoxie) + zf.Zfs.Chuangjianriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuCaozuoriqi(xiaoxie bool) string {
	return zf.Zfs.Xiangmu(xiaoxie) + zf.Zfs.Caozuoriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuYouxiaoxing(xiaoxie bool) string {
	return zf.Zfs.Xiangmu(xiaoxie) + zf.Zfs.Youxiaoxing(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuCaozuoren(xiaoxie bool) string {
	return zf.Zfs.Xiangmu(xiaoxie) + zf.Zfs.Caozuoren(xiaoxie)

}

//2017/10/11 22:49:22
func (sjkmhsydata *Sjkmhsydata) XiangmuzuMiaoshu(xiaoxie bool) string {
	return zf.Zfs.Xiangmuzu(xiaoxie) + zf.Zfs.Miaoshu(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuzuMingcheng(xiaoxie bool) string {
	return zf.Zfs.Xiangmuzu(xiaoxie) + zf.Zfs.Mingcheng(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuzuId(xiaoxie bool) string {
	return zf.Zfs.Xiangmuzu(xiaoxie) + zf.Zfs.Id(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuzuBianma(xiaoxie bool) string {
	return zf.Zfs.Xiangmuzu(xiaoxie) + zf.Zfs.Bianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuzuPaixu(xiaoxie bool) string {
	return zf.Zfs.Xiangmuzu(xiaoxie) + zf.Zfs.Paixu(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuzuBiaoji(xiaoxie bool) string {
	return zf.Zfs.Xiangmuzu(xiaoxie) + zf.Zfs.Biaoji(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuzuChuangjianriqi(xiaoxie bool) string {
	return zf.Zfs.Xiangmuzu(xiaoxie) + zf.Zfs.Chuangjianriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuzuCaozuoriqi(xiaoxie bool) string {
	return zf.Zfs.Xiangmuzu(xiaoxie) + zf.Zfs.Caozuoriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuzuYouxiaoxing(xiaoxie bool) string {
	return zf.Zfs.Xiangmuzu(xiaoxie) + zf.Zfs.Youxiaoxing(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuzuCaozuoren(xiaoxie bool) string {
	return zf.Zfs.Xiangmuzu(xiaoxie) + zf.Zfs.Caozuoren(xiaoxie)

}

//2017/10/11 22:49:22
func (sjkmhsydata *Sjkmhsydata) XiangmuzuxiangmuXiangmubianma(xiaoxie bool) string {
	return zf.Zfs.Xiangmuzuxiangmu(xiaoxie) + zf.Zfs.Xiangmubianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuzuxiangmuXiangmuzubianma(xiaoxie bool) string {
	return zf.Zfs.Xiangmuzuxiangmu(xiaoxie) + zf.Zfs.Xiangmuzubianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuzuxiangmuId(xiaoxie bool) string {
	return zf.Zfs.Xiangmuzuxiangmu(xiaoxie) + zf.Zfs.Id(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuzuxiangmuBianma(xiaoxie bool) string {
	return zf.Zfs.Xiangmuzuxiangmu(xiaoxie) + zf.Zfs.Bianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuzuxiangmuPaixu(xiaoxie bool) string {
	return zf.Zfs.Xiangmuzuxiangmu(xiaoxie) + zf.Zfs.Paixu(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuzuxiangmuBiaoji(xiaoxie bool) string {
	return zf.Zfs.Xiangmuzuxiangmu(xiaoxie) + zf.Zfs.Biaoji(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuzuxiangmuChuangjianriqi(xiaoxie bool) string {
	return zf.Zfs.Xiangmuzuxiangmu(xiaoxie) + zf.Zfs.Chuangjianriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuzuxiangmuCaozuoriqi(xiaoxie bool) string {
	return zf.Zfs.Xiangmuzuxiangmu(xiaoxie) + zf.Zfs.Caozuoriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuzuxiangmuYouxiaoxing(xiaoxie bool) string {
	return zf.Zfs.Xiangmuzuxiangmu(xiaoxie) + zf.Zfs.Youxiaoxing(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuzuxiangmuCaozuoren(xiaoxie bool) string {
	return zf.Zfs.Xiangmuzuxiangmu(xiaoxie) + zf.Zfs.Caozuoren(xiaoxie)

}

//2017/10/11 22:49:22
func (sjkmhsydata *Sjkmhsydata) ShijianXiangmuzubianma(xiaoxie bool) string {
	return zf.Zfs.Shijian(xiaoxie) + zf.Zfs.Xiangmuzubianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ShijianQiyong(xiaoxie bool) string {
	return zf.Zfs.Shijian(xiaoxie) + zf.Zfs.Qiyong(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ShijianMiaoshu(xiaoxie bool) string {
	return zf.Zfs.Shijian(xiaoxie) + zf.Zfs.Miaoshu(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ShijianMingcheng(xiaoxie bool) string {
	return zf.Zfs.Shijian(xiaoxie) + zf.Zfs.Mingcheng(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ShijianId(xiaoxie bool) string {
	return zf.Zfs.Shijian(xiaoxie) + zf.Zfs.Id(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ShijianBianma(xiaoxie bool) string {
	return zf.Zfs.Shijian(xiaoxie) + zf.Zfs.Bianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ShijianPaixu(xiaoxie bool) string {
	return zf.Zfs.Shijian(xiaoxie) + zf.Zfs.Paixu(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ShijianBiaoji(xiaoxie bool) string {
	return zf.Zfs.Shijian(xiaoxie) + zf.Zfs.Biaoji(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ShijianChuangjianriqi(xiaoxie bool) string {
	return zf.Zfs.Shijian(xiaoxie) + zf.Zfs.Chuangjianriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ShijianCaozuoriqi(xiaoxie bool) string {
	return zf.Zfs.Shijian(xiaoxie) + zf.Zfs.Caozuoriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ShijianYouxiaoxing(xiaoxie bool) string {
	return zf.Zfs.Shijian(xiaoxie) + zf.Zfs.Youxiaoxing(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ShijianCaozuoren(xiaoxie bool) string {
	return zf.Zfs.Shijian(xiaoxie) + zf.Zfs.Caozuoren(xiaoxie)

}

//2017/10/11 22:49:22
func (sjkmhsydata *Sjkmhsydata) XiangmuzushijianShijianbianma(xiaoxie bool) string {
	return zf.Zfs.Xiangmuzushijian(xiaoxie) + zf.Zfs.Shijianbianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuzushijianXiangmuzubianma(xiaoxie bool) string {
	return zf.Zfs.Xiangmuzushijian(xiaoxie) + zf.Zfs.Xiangmuzubianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuzushijianId(xiaoxie bool) string {
	return zf.Zfs.Xiangmuzushijian(xiaoxie) + zf.Zfs.Id(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuzushijianBianma(xiaoxie bool) string {
	return zf.Zfs.Xiangmuzushijian(xiaoxie) + zf.Zfs.Bianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuzushijianPaixu(xiaoxie bool) string {
	return zf.Zfs.Xiangmuzushijian(xiaoxie) + zf.Zfs.Paixu(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuzushijianBiaoji(xiaoxie bool) string {
	return zf.Zfs.Xiangmuzushijian(xiaoxie) + zf.Zfs.Biaoji(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuzushijianChuangjianriqi(xiaoxie bool) string {
	return zf.Zfs.Xiangmuzushijian(xiaoxie) + zf.Zfs.Chuangjianriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuzushijianCaozuoriqi(xiaoxie bool) string {
	return zf.Zfs.Xiangmuzushijian(xiaoxie) + zf.Zfs.Caozuoriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuzushijianYouxiaoxing(xiaoxie bool) string {
	return zf.Zfs.Xiangmuzushijian(xiaoxie) + zf.Zfs.Youxiaoxing(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmuzushijianCaozuoren(xiaoxie bool) string {
	return zf.Zfs.Xiangmuzushijian(xiaoxie) + zf.Zfs.Caozuoren(xiaoxie)

}

//2017/10/11 22:49:22
func (sjkmhsydata *Sjkmhsydata) ZhanghaoshijianShijianbianma(xiaoxie bool) string {
	return zf.Zfs.Zhanghaoshijian(xiaoxie) + zf.Zfs.Shijianbianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZhanghaoshijianWeiyibianma(xiaoxie bool) string {
	return zf.Zfs.Zhanghaoshijian(xiaoxie) + zf.Zfs.Weiyibianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZhanghaoshijianId(xiaoxie bool) string {
	return zf.Zfs.Zhanghaoshijian(xiaoxie) + zf.Zfs.Id(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZhanghaoshijianBianma(xiaoxie bool) string {
	return zf.Zfs.Zhanghaoshijian(xiaoxie) + zf.Zfs.Bianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZhanghaoshijianPaixu(xiaoxie bool) string {
	return zf.Zfs.Zhanghaoshijian(xiaoxie) + zf.Zfs.Paixu(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZhanghaoshijianBiaoji(xiaoxie bool) string {
	return zf.Zfs.Zhanghaoshijian(xiaoxie) + zf.Zfs.Biaoji(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZhanghaoshijianChuangjianriqi(xiaoxie bool) string {
	return zf.Zfs.Zhanghaoshijian(xiaoxie) + zf.Zfs.Chuangjianriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZhanghaoshijianCaozuoriqi(xiaoxie bool) string {
	return zf.Zfs.Zhanghaoshijian(xiaoxie) + zf.Zfs.Caozuoriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZhanghaoshijianYouxiaoxing(xiaoxie bool) string {
	return zf.Zfs.Zhanghaoshijian(xiaoxie) + zf.Zfs.Youxiaoxing(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZhanghaoshijianCaozuoren(xiaoxie bool) string {
	return zf.Zfs.Zhanghaoshijian(xiaoxie) + zf.Zfs.Caozuoren(xiaoxie)

}

//2017/10/11 22:49:22
func (sjkmhsydata *Sjkmhsydata) ZidianFubianma(xiaoxie bool) string {
	return zf.Zfs.Zidian(xiaoxie) + zf.Zfs.Fubianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZidianMiaoshu(xiaoxie bool) string {
	return zf.Zfs.Zidian(xiaoxie) + zf.Zfs.Miaoshu(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZidianMingcheng(xiaoxie bool) string {
	return zf.Zfs.Zidian(xiaoxie) + zf.Zfs.Mingcheng(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZidianId(xiaoxie bool) string {
	return zf.Zfs.Zidian(xiaoxie) + zf.Zfs.Id(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZidianBianma(xiaoxie bool) string {
	return zf.Zfs.Zidian(xiaoxie) + zf.Zfs.Bianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZidianPaixu(xiaoxie bool) string {
	return zf.Zfs.Zidian(xiaoxie) + zf.Zfs.Paixu(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZidianBiaoji(xiaoxie bool) string {
	return zf.Zfs.Zidian(xiaoxie) + zf.Zfs.Biaoji(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZidianChuangjianriqi(xiaoxie bool) string {
	return zf.Zfs.Zidian(xiaoxie) + zf.Zfs.Chuangjianriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZidianCaozuoriqi(xiaoxie bool) string {
	return zf.Zfs.Zidian(xiaoxie) + zf.Zfs.Caozuoriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZidianYouxiaoxing(xiaoxie bool) string {
	return zf.Zfs.Zidian(xiaoxie) + zf.Zfs.Youxiaoxing(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ZidianCaozuoren(xiaoxie bool) string {
	return zf.Zfs.Zidian(xiaoxie) + zf.Zfs.Caozuoren(xiaoxie)

}

//2017/10/11 22:49:22
func (sjkmhsydata *Sjkmhsydata) YonghuLinchuangzhenduan(xiaoxie bool) string {
	return zf.Zfs.Yonghu(xiaoxie) + zf.Zfs.Linchuangzhenduan(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YonghuJiankangzhuangkuang(xiaoxie bool) string {
	return zf.Zfs.Yonghu(xiaoxie) + zf.Zfs.Jiankangzhuangkuang(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YonghuJiankangtouzi(xiaoxie bool) string {
	return zf.Zfs.Yonghu(xiaoxie) + zf.Zfs.Jiankangtouzi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YonghuYundongshijian(xiaoxie bool) string {
	return zf.Zfs.Yonghu(xiaoxie) + zf.Zfs.Yundongshijian(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YonghuYundongcishu(xiaoxie bool) string {
	return zf.Zfs.Yonghu(xiaoxie) + zf.Zfs.Yundongcishu(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YonghuYundongleixing(xiaoxie bool) string {
	return zf.Zfs.Yonghu(xiaoxie) + zf.Zfs.Yundongleixing(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YonghuShifouhejiu(xiaoxie bool) string {
	return zf.Zfs.Yonghu(xiaoxie) + zf.Zfs.Shifouhejiu(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YonghuShifouxiyan(xiaoxie bool) string {
	return zf.Zfs.Yonghu(xiaoxie) + zf.Zfs.Shifouxiyan(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YonghuShuimianzhuangtai(xiaoxie bool) string {
	return zf.Zfs.Yonghu(xiaoxie) + zf.Zfs.Shuimianzhuangtai(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YonghuYinshixiguan(xiaoxie bool) string {
	return zf.Zfs.Yonghu(xiaoxie) + zf.Zfs.Yinshixiguan(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YonghuHunfou(xiaoxie bool) string {
	return zf.Zfs.Yonghu(xiaoxie) + zf.Zfs.Hunfou(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YonghuWenhuachengdu(xiaoxie bool) string {
	return zf.Zfs.Yonghu(xiaoxie) + zf.Zfs.Wenhuachengdu(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YonghuChushengriqi(xiaoxie bool) string {
	return zf.Zfs.Yonghu(xiaoxie) + zf.Zfs.Chushengriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YonghuMinzu(xiaoxie bool) string {
	return zf.Zfs.Yonghu(xiaoxie) + zf.Zfs.Minzu(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YonghuTizhong(xiaoxie bool) string {
	return zf.Zfs.Yonghu(xiaoxie) + zf.Zfs.Tizhong(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YonghuShengao(xiaoxie bool) string {
	return zf.Zfs.Yonghu(xiaoxie) + zf.Zfs.Shengao(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YonghuZhiye(xiaoxie bool) string {
	return zf.Zfs.Yonghu(xiaoxie) + zf.Zfs.Zhiye(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YonghuNianling(xiaoxie bool) string {
	return zf.Zfs.Yonghu(xiaoxie) + zf.Zfs.Nianling(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YonghuXingbie(xiaoxie bool) string {
	return zf.Zfs.Yonghu(xiaoxie) + zf.Zfs.Xingbie(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YonghuWeiyibianma(xiaoxie bool) string {
	return zf.Zfs.Yonghu(xiaoxie) + zf.Zfs.Weiyibianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YonghuQidongriqi(xiaoxie bool) string {
	return zf.Zfs.Yonghu(xiaoxie) + zf.Zfs.Qidongriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YonghuTouxiang(xiaoxie bool) string {
	return zf.Zfs.Yonghu(xiaoxie) + zf.Zfs.Touxiang(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YonghuMingcheng(xiaoxie bool) string {
	return zf.Zfs.Yonghu(xiaoxie) + zf.Zfs.Mingcheng(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YonghuId(xiaoxie bool) string {
	return zf.Zfs.Yonghu(xiaoxie) + zf.Zfs.Id(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YonghuBianma(xiaoxie bool) string {
	return zf.Zfs.Yonghu(xiaoxie) + zf.Zfs.Bianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YonghuPaixu(xiaoxie bool) string {
	return zf.Zfs.Yonghu(xiaoxie) + zf.Zfs.Paixu(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YonghuBiaoji(xiaoxie bool) string {
	return zf.Zfs.Yonghu(xiaoxie) + zf.Zfs.Biaoji(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YonghuChuangjianriqi(xiaoxie bool) string {
	return zf.Zfs.Yonghu(xiaoxie) + zf.Zfs.Chuangjianriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YonghuCaozuoriqi(xiaoxie bool) string {
	return zf.Zfs.Yonghu(xiaoxie) + zf.Zfs.Caozuoriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YonghuYouxiaoxing(xiaoxie bool) string {
	return zf.Zfs.Yonghu(xiaoxie) + zf.Zfs.Youxiaoxing(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YonghuCaozuoren(xiaoxie bool) string {
	return zf.Zfs.Yonghu(xiaoxie) + zf.Zfs.Caozuoren(xiaoxie)

}

//2017/10/11 22:49:22
func (sjkmhsydata *Sjkmhsydata) JibingJiazubingshi(xiaoxie bool) string {
	return zf.Zfs.Jibing(xiaoxie) + zf.Zfs.Jiazubingshi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JibingExingzhongliu(xiaoxie bool) string {
	return zf.Zfs.Jibing(xiaoxie) + zf.Zfs.Exingzhongliu(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JibingGanran(xiaoxie bool) string {
	return zf.Zfs.Jibing(xiaoxie) + zf.Zfs.Ganran(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JibingPifubing(xiaoxie bool) string {
	return zf.Zfs.Jibing(xiaoxie) + zf.Zfs.Pifubing(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JibingXinnaoxueguanbing(xiaoxie bool) string {
	return zf.Zfs.Jibing(xiaoxie) + zf.Zfs.Xinnaoxueguanbing(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JibingJiliang(xiaoxie bool) string {
	return zf.Zfs.Jibing(xiaoxie) + zf.Zfs.Jiliang(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JibingXuetangchuli(xiaoxie bool) string {
	return zf.Zfs.Jibing(xiaoxie) + zf.Zfs.Xuetangchuli(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JibingXuetang(xiaoxie bool) string {
	return zf.Zfs.Jibing(xiaoxie) + zf.Zfs.Xuetang(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JibingTangniaobing(xiaoxie bool) string {
	return zf.Zfs.Jibing(xiaoxie) + zf.Zfs.Tangniaobing(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JibingWeiyibianma(xiaoxie bool) string {
	return zf.Zfs.Jibing(xiaoxie) + zf.Zfs.Weiyibianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JibingId(xiaoxie bool) string {
	return zf.Zfs.Jibing(xiaoxie) + zf.Zfs.Id(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JibingBianma(xiaoxie bool) string {
	return zf.Zfs.Jibing(xiaoxie) + zf.Zfs.Bianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JibingPaixu(xiaoxie bool) string {
	return zf.Zfs.Jibing(xiaoxie) + zf.Zfs.Paixu(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JibingBiaoji(xiaoxie bool) string {
	return zf.Zfs.Jibing(xiaoxie) + zf.Zfs.Biaoji(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JibingChuangjianriqi(xiaoxie bool) string {
	return zf.Zfs.Jibing(xiaoxie) + zf.Zfs.Chuangjianriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JibingCaozuoriqi(xiaoxie bool) string {
	return zf.Zfs.Jibing(xiaoxie) + zf.Zfs.Caozuoriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JibingYouxiaoxing(xiaoxie bool) string {
	return zf.Zfs.Jibing(xiaoxie) + zf.Zfs.Youxiaoxing(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) JibingCaozuoren(xiaoxie bool) string {
	return zf.Zfs.Jibing(xiaoxie) + zf.Zfs.Caozuoren(xiaoxie)

}

//2017/10/11 22:49:22
func (sjkmhsydata *Sjkmhsydata) WenzhangShifoufabu(xiaoxie bool) string {
	return zf.Zfs.Wenzhang(xiaoxie) + zf.Zfs.Shifoufabu(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) WenzhangNeirong(xiaoxie bool) string {
	return zf.Zfs.Wenzhang(xiaoxie) + zf.Zfs.Neirong(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) WenzhangMingcheng(xiaoxie bool) string {
	return zf.Zfs.Wenzhang(xiaoxie) + zf.Zfs.Mingcheng(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) WenzhangId(xiaoxie bool) string {
	return zf.Zfs.Wenzhang(xiaoxie) + zf.Zfs.Id(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) WenzhangBianma(xiaoxie bool) string {
	return zf.Zfs.Wenzhang(xiaoxie) + zf.Zfs.Bianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) WenzhangPaixu(xiaoxie bool) string {
	return zf.Zfs.Wenzhang(xiaoxie) + zf.Zfs.Paixu(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) WenzhangBiaoji(xiaoxie bool) string {
	return zf.Zfs.Wenzhang(xiaoxie) + zf.Zfs.Biaoji(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) WenzhangChuangjianriqi(xiaoxie bool) string {
	return zf.Zfs.Wenzhang(xiaoxie) + zf.Zfs.Chuangjianriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) WenzhangCaozuoriqi(xiaoxie bool) string {
	return zf.Zfs.Wenzhang(xiaoxie) + zf.Zfs.Caozuoriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) WenzhangYouxiaoxing(xiaoxie bool) string {
	return zf.Zfs.Wenzhang(xiaoxie) + zf.Zfs.Youxiaoxing(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) WenzhangCaozuoren(xiaoxie bool) string {
	return zf.Zfs.Wenzhang(xiaoxie) + zf.Zfs.Caozuoren(xiaoxie)

}

//2017/10/11 22:49:22
func (sjkmhsydata *Sjkmhsydata) LiuyanShijian(xiaoxie bool) string {
	return zf.Zfs.Liuyan(xiaoxie) + zf.Zfs.Shijian(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) LiuyanNeirong(xiaoxie bool) string {
	return zf.Zfs.Liuyan(xiaoxie) + zf.Zfs.Neirong(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) LiuyanBiaoti(xiaoxie bool) string {
	return zf.Zfs.Liuyan(xiaoxie) + zf.Zfs.Biaoti(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) LiuyanHuifuyonghu(xiaoxie bool) string {
	return zf.Zfs.Liuyan(xiaoxie) + zf.Zfs.Huifuyonghu(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) LiuyanLiuyanyonghu(xiaoxie bool) string {
	return zf.Zfs.Liuyan(xiaoxie) + zf.Zfs.Liuyanyonghu(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) LiuyanId(xiaoxie bool) string {
	return zf.Zfs.Liuyan(xiaoxie) + zf.Zfs.Id(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) LiuyanBianma(xiaoxie bool) string {
	return zf.Zfs.Liuyan(xiaoxie) + zf.Zfs.Bianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) LiuyanPaixu(xiaoxie bool) string {
	return zf.Zfs.Liuyan(xiaoxie) + zf.Zfs.Paixu(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) LiuyanBiaoji(xiaoxie bool) string {
	return zf.Zfs.Liuyan(xiaoxie) + zf.Zfs.Biaoji(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) LiuyanChuangjianriqi(xiaoxie bool) string {
	return zf.Zfs.Liuyan(xiaoxie) + zf.Zfs.Chuangjianriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) LiuyanCaozuoriqi(xiaoxie bool) string {
	return zf.Zfs.Liuyan(xiaoxie) + zf.Zfs.Caozuoriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) LiuyanYouxiaoxing(xiaoxie bool) string {
	return zf.Zfs.Liuyan(xiaoxie) + zf.Zfs.Youxiaoxing(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) LiuyanCaozuoren(xiaoxie bool) string {
	return zf.Zfs.Liuyan(xiaoxie) + zf.Zfs.Caozuoren(xiaoxie)

}

//2017/10/11 22:49:22
func (sjkmhsydata *Sjkmhsydata) FuwufankuiWeiyibianma(xiaoxie bool) string {
	return zf.Zfs.Fuwufankui(xiaoxie) + zf.Zfs.Weiyibianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) FuwufankuiXiaoguokeping(xiaoxie bool) string {
	return zf.Zfs.Fuwufankui(xiaoxie) + zf.Zfs.Xiaoguokeping(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) FuwufankuiShentiganshou(xiaoxie bool) string {
	return zf.Zfs.Fuwufankui(xiaoxie) + zf.Zfs.Shentiganshou(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) FuwufankuiFuwuganshou(xiaoxie bool) string {
	return zf.Zfs.Fuwufankui(xiaoxie) + zf.Zfs.Fuwuganshou(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) FuwufankuiId(xiaoxie bool) string {
	return zf.Zfs.Fuwufankui(xiaoxie) + zf.Zfs.Id(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) FuwufankuiBianma(xiaoxie bool) string {
	return zf.Zfs.Fuwufankui(xiaoxie) + zf.Zfs.Bianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) FuwufankuiPaixu(xiaoxie bool) string {
	return zf.Zfs.Fuwufankui(xiaoxie) + zf.Zfs.Paixu(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) FuwufankuiBiaoji(xiaoxie bool) string {
	return zf.Zfs.Fuwufankui(xiaoxie) + zf.Zfs.Biaoji(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) FuwufankuiChuangjianriqi(xiaoxie bool) string {
	return zf.Zfs.Fuwufankui(xiaoxie) + zf.Zfs.Chuangjianriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) FuwufankuiCaozuoriqi(xiaoxie bool) string {
	return zf.Zfs.Fuwufankui(xiaoxie) + zf.Zfs.Caozuoriqi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) FuwufankuiYouxiaoxing(xiaoxie bool) string {
	return zf.Zfs.Fuwufankui(xiaoxie) + zf.Zfs.Youxiaoxing(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) FuwufankuiCaozuoren(xiaoxie bool) string {
	return zf.Zfs.Fuwufankui(xiaoxie) + zf.Zfs.Caozuoren(xiaoxie)

}

//2017/10/11 22:49:22
func (sjkmhsydata *Sjkmhsydata) XiangmushujuWeiyibianma(xiaoxie bool) string {
	return zf.Zfs.Xiangmushuju(xiaoxie) + zf.Zfs.Weiyibianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmushujuZhidingren(xiaoxie bool) string {
	return zf.Zfs.Xiangmushuju(xiaoxie) + zf.Zfs.Zhidingren(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmushujuCankao(xiaoxie bool) string {
	return zf.Zfs.Xiangmushuju(xiaoxie) + zf.Zfs.Cankao(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmushujuNeirong(xiaoxie bool) string {
	return zf.Zfs.Xiangmushuju(xiaoxie) + zf.Zfs.Neirong(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmushujuMingcheng(xiaoxie bool) string {
	return zf.Zfs.Xiangmushuju(xiaoxie) + zf.Zfs.Mingcheng(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmushujuId(xiaoxie bool) string {
	return zf.Zfs.Xiangmushuju(xiaoxie) + zf.Zfs.Id(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmushujuBianma(xiaoxie bool) string {
	return zf.Zfs.Xiangmushuju(xiaoxie) + zf.Zfs.Bianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) XiangmushujuPaixu(xiaoxie bool) string {

	return zf.Zfs.Xiangmushuju(xiaoxie) + zf.Zfs.Paixu(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) XiangmushujuBiaoji(xiaoxie bool) string {

	return zf.Zfs.Xiangmushuju(xiaoxie) + zf.Zfs.Biaoji(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) XiangmushujuChuangjianriqi(xiaoxie bool) string {

	return zf.Zfs.Xiangmushuju(xiaoxie) + zf.Zfs.Chuangjianriqi(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) XiangmushujuCaozuoriqi(xiaoxie bool) string {

	return zf.Zfs.Xiangmushuju(xiaoxie) + zf.Zfs.Caozuoriqi(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) XiangmushujuYouxiaoxing(xiaoxie bool) string {

	return zf.Zfs.Xiangmushuju(xiaoxie) + zf.Zfs.Youxiaoxing(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) XiangmushujuCaozuoren(xiaoxie bool) string {

	return zf.Zfs.Xiangmushuju(xiaoxie) + zf.Zfs.Caozuoren(xiaoxie)

}

//2017/10/11 22:49:22
func (sjkmhsydata *Sjkmhsydata) CujinfanganZhi(xiaoxie bool) string {

	return zf.Zfs.Cujinfangan(xiaoxie) + zf.Zfs.Zhi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) CujinfanganShijianbianma(xiaoxie bool) string {

	return zf.Zfs.Cujinfangan(xiaoxie) + zf.Zfs.Shijianbianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) CujinfanganXiangmubianma(xiaoxie bool) string {

	return zf.Zfs.Cujinfangan(xiaoxie) + zf.Zfs.Xiangmubianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) CujinfanganWeiyibianma(xiaoxie bool) string {

	return zf.Zfs.Cujinfangan(xiaoxie) + zf.Zfs.Weiyibianma(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) CujinfanganId(xiaoxie bool) string {

	return zf.Zfs.Cujinfangan(xiaoxie) + zf.Zfs.Id(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) CujinfanganBianma(xiaoxie bool) string {

	return zf.Zfs.Cujinfangan(xiaoxie) + zf.Zfs.Bianma(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) CujinfanganPaixu(xiaoxie bool) string {

	return zf.Zfs.Cujinfangan(xiaoxie) + zf.Zfs.Paixu(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) CujinfanganBiaoji(xiaoxie bool) string {

	return zf.Zfs.Cujinfangan(xiaoxie) + zf.Zfs.Biaoji(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) CujinfanganChuangjianriqi(xiaoxie bool) string {

	return zf.Zfs.Cujinfangan(xiaoxie) + zf.Zfs.Chuangjianriqi(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) CujinfanganCaozuoriqi(xiaoxie bool) string {

	return zf.Zfs.Cujinfangan(xiaoxie) + zf.Zfs.Caozuoriqi(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) CujinfanganYouxiaoxing(xiaoxie bool) string {

	return zf.Zfs.Cujinfangan(xiaoxie) + zf.Zfs.Youxiaoxing(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) CujinfanganCaozuoren(xiaoxie bool) string {

	return zf.Zfs.Cujinfangan(xiaoxie) + zf.Zfs.Caozuoren(xiaoxie)

}

//2017/10/11 22:49:22
func (sjkmhsydata *Sjkmhsydata) ShebeiMiaoshu(xiaoxie bool) string {

	return zf.Zfs.Shebei(xiaoxie) + zf.Zfs.Miaoshu(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ShebeiShebeihao(xiaoxie bool) string {

	return zf.Zfs.Shebei(xiaoxie) + zf.Zfs.Shebeihao(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ShebeiMingcheng(xiaoxie bool) string {

	return zf.Zfs.Shebei(xiaoxie) + zf.Zfs.Mingcheng(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) ShebeiId(xiaoxie bool) string {

	return zf.Zfs.Shebei(xiaoxie) + zf.Zfs.Id(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) ShebeiBianma(xiaoxie bool) string {

	return zf.Zfs.Shebei(xiaoxie) + zf.Zfs.Bianma(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) ShebeiPaixu(xiaoxie bool) string {

	return zf.Zfs.Shebei(xiaoxie) + zf.Zfs.Paixu(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) ShebeiBiaoji(xiaoxie bool) string {

	return zf.Zfs.Shebei(xiaoxie) + zf.Zfs.Biaoji(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) ShebeiChuangjianriqi(xiaoxie bool) string {

	return zf.Zfs.Shebei(xiaoxie) + zf.Zfs.Chuangjianriqi(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) ShebeiCaozuoriqi(xiaoxie bool) string {

	return zf.Zfs.Shebei(xiaoxie) + zf.Zfs.Caozuoriqi(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) ShebeiYouxiaoxing(xiaoxie bool) string {

	return zf.Zfs.Shebei(xiaoxie) + zf.Zfs.Youxiaoxing(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) ShebeiCaozuoren(xiaoxie bool) string {

	return zf.Zfs.Shebei(xiaoxie) + zf.Zfs.Caozuoren(xiaoxie)

}

//2017/10/11 22:49:22
func (sjkmhsydata *Sjkmhsydata) YinpinMiaoshu(xiaoxie bool) string {

	return zf.Zfs.Yinpin(xiaoxie) + zf.Zfs.Miaoshu(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YinpinMingcheng(xiaoxie bool) string {

	return zf.Zfs.Yinpin(xiaoxie) + zf.Zfs.Mingcheng(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YinpinJiamifangshi(xiaoxie bool) string {

	return zf.Zfs.Yinpin(xiaoxie) + zf.Zfs.Jiamifangshi(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YinpinLujing(xiaoxie bool) string {

	return zf.Zfs.Yinpin(xiaoxie) + zf.Zfs.Id(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YinpinId(xiaoxie bool) string {

	return zf.Zfs.Yinpin(xiaoxie) + zf.Zfs.Id(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) YinpinBianma(xiaoxie bool) string {

	return zf.Zfs.Yinpin(xiaoxie) + zf.Zfs.Bianma(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) YinpinPaixu(xiaoxie bool) string {

	return zf.Zfs.Yinpin(xiaoxie) + zf.Zfs.Paixu(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) YinpinBiaoji(xiaoxie bool) string {

	return zf.Zfs.Yinpin(xiaoxie) + zf.Zfs.Biaoji(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) YinpinChuangjianriqi(xiaoxie bool) string {

	return zf.Zfs.Yinpin(xiaoxie) + zf.Zfs.Chuangjianriqi(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) YinpinCaozuoriqi(xiaoxie bool) string {

	return zf.Zfs.Yinpin(xiaoxie) + zf.Zfs.Caozuoriqi(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) YinpinYouxiaoxing(xiaoxie bool) string {

	return zf.Zfs.Yinpin(xiaoxie) + zf.Zfs.Youxiaoxing(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) YinpinCaozuoren(xiaoxie bool) string {

	return zf.Zfs.Yinpin(xiaoxie) + zf.Zfs.Caozuoren(xiaoxie)

}

//2017/10/12 11:55:03
func (sjkmhsydata *Sjkmhsydata) YinpinbofangYinpinbianma(xiaoxie bool) string {

	return zf.Zfs.Yinpinbofang(xiaoxie) + zf.Zfs.Bofangshijian(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YinpinbofangWeiyibianma(xiaoxie bool) string {

	return zf.Zfs.Yinpinbofang(xiaoxie) + zf.Zfs.Bofangshijian(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YinpinbofangBofangshijian(xiaoxie bool) string {

	return zf.Zfs.Yinpinbofang(xiaoxie) + zf.Zfs.Bofangshijian(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YinpinbofangId(xiaoxie bool) string {

	return zf.Zfs.Yinpinbofang(xiaoxie) + zf.Zfs.Id(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) YinpinbofangBianma(xiaoxie bool) string {

	return zf.Zfs.Yinpinbofang(xiaoxie) + zf.Zfs.Bianma(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) YinpinbofangPaixu(xiaoxie bool) string {

	return zf.Zfs.Yinpinbofang(xiaoxie) + zf.Zfs.Paixu(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) YinpinbofangBiaoji(xiaoxie bool) string {

	return zf.Zfs.Yinpinbofang(xiaoxie) + zf.Zfs.Biaoji(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) YinpinbofangChuangjianriqi(xiaoxie bool) string {

	return zf.Zfs.Yinpinbofang(xiaoxie) + zf.Zfs.Chuangjianriqi(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) YinpinbofangCaozuoriqi(xiaoxie bool) string {

	return zf.Zfs.Yinpinbofang(xiaoxie) + zf.Zfs.Caozuoriqi(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) YinpinbofangYouxiaoxing(xiaoxie bool) string {

	return zf.Zfs.Yinpinbofang(xiaoxie) + zf.Zfs.Youxiaoxing(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) YinpinbofangCaozuoren(xiaoxie bool) string {

	return zf.Zfs.Yinpinbofang(xiaoxie) + zf.Zfs.Caozuoren(xiaoxie)

}

//2017/10/12 11:55:03
func (sjkmhsydata *Sjkmhsydata) YinpinxiazaiJinzhixiazai(xiaoxie bool) string {

	return zf.Zfs.Yinpinxiazai(xiaoxie) + zf.Zfs.Id(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YinpinxiazaiYinpinbianma(xiaoxie bool) string {

	return zf.Zfs.Yinpinxiazai(xiaoxie) + zf.Zfs.Id(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YinpinxiazaiWeiyibianma(xiaoxie bool) string {

	return zf.Zfs.Yinpinxiazai(xiaoxie) + zf.Zfs.Id(xiaoxie)

}
func (sjkmhsydata *Sjkmhsydata) YinpinxiazaiId(xiaoxie bool) string {

	return zf.Zfs.Yinpinxiazai(xiaoxie) + zf.Zfs.Id(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) YinpinxiazaiBianma(xiaoxie bool) string {

	return zf.Zfs.Yinpinxiazai(xiaoxie) + zf.Zfs.Bianma(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) YinpinxiazaiPaixu(xiaoxie bool) string {

	return zf.Zfs.Yinpinxiazai(xiaoxie) + zf.Zfs.Paixu(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) YinpinxiazaiBiaoji(xiaoxie bool) string {

	return zf.Zfs.Yinpinxiazai(xiaoxie) + zf.Zfs.Biaoji(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) YinpinxiazaiChuangjianriqi(xiaoxie bool) string {

	return zf.Zfs.Yinpinxiazai(xiaoxie) + zf.Zfs.Chuangjianriqi(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) YinpinxiazaiCaozuoriqi(xiaoxie bool) string {

	return zf.Zfs.Yinpinxiazai(xiaoxie) + zf.Zfs.Caozuoriqi(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) YinpinxiazaiYouxiaoxing(xiaoxie bool) string {

	return zf.Zfs.Yinpinxiazai(xiaoxie) + zf.Zfs.Youxiaoxing(xiaoxie)

}

func (sjkmhsydata *Sjkmhsydata) YinpinxiazaiCaozuoren(xiaoxie bool) string {

	return zf.Zfs.Yinpinxiazai(xiaoxie) + zf.Zfs.Caozuoren(xiaoxie)

}
