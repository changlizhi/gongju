package gongju

import "time"

func Shengchengzhujian(workerId int64, datacenterId int64) int64 {
	/**
 * 开始时间截 (2015-01-01)
 */
	twepoch := 1420041600000
	/**
	 * 机器id所占的位数
	 */
	workerIdBits := 5

	/**
	 * 数据标识id所占的位数
	 */
	datacenterIdBits := 5

	/**
	 * 支持的最大机器id，结果是31 (这个移位算法可以很快的计算出几位二进制数所能表示的最大十进制数)
	 */
	maxWorkerId := -1 ^ (-1 << workerIdBits)

	/**
	 * 支持的最大数据标识id，结果是31
	 */
	maxDatacenterId := -1 ^ (-1 << datacenterIdBits)

	/**
	 * 序列在id中占的位数
	 */
	sequenceBits := 12

	/**
	 * 机器ID向左移12位
	 */
	workerIdShift := sequenceBits

	/**
	 * 数据标识id向左移17位(12+5)
	 */
	datacenterIdShift := sequenceBits + workerIdBits

	/**
	 * 时间截向左移22位(5+5+12)
	 */
	timestampLeftShift := sequenceBits + workerIdBits + datacenterIdBits

	/**
	 * 生成序列的掩码，这里为4095 (0b111111111111:=0xfff:=4095)
	 */
	sequenceMask := -1 ^ (-1 << sequenceBits)
	sequence := 0
	lastTimestamp := -1

	if workerId > maxWorkerId || workerId < 0 {
		panic(error("worker Id can't be greater than %d or less than 0"))
	}
	if datacenterId > maxDatacenterId || datacenterId < 0 {
		panic(error("datacenter Id can't be greater than %d or less than 0"))
	}

	timestamp := currentmil();

	//如果当前时间小于上一次ID生成的时间戳，说明系统时钟回退过这个时候应当抛出异常
	if timestamp < lastTimestamp {
		panic(error("Clock moved backwards.  Refusing to generate id for %d milliseconds"))
	}

	//如果是同一时间生成的，则进行毫秒内序列
	if lastTimestamp == timestamp {
		sequence = (sequence + 1) & sequenceMask
		//毫秒内序列溢出
		if sequence == 0 {
			//阻塞到下一个毫秒,获得新的时间戳
			timestamp = tilNextMillis;
		}
	} else {
		sequence = 0
	}

	//上次生成ID的时间截
	lastTimestamp = timestamp;

	//移位并通过或运算拼到一起组成64位的ID
	return ((timestamp - twepoch) << timestampLeftShift) |
		(datacenterId << datacenterIdShift) |
		(workerId << workerIdShift) | sequence

}

func tilNextMillis(lastTimestamp int64) int64 {
	timestamp := currentmil();
	for timestamp <= lastTimestamp {
		timestamp = currentmil();
	}
	return timestamp;
}

func currentmil() int64 {
	t := time.Now()
	return t.Unix()
}