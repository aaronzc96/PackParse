package mac

import (
	"fmt"
)

func ParBch(pack []byte) {
	if len(pack) != 55 {
		fmt.Println("错误的包长度！BCH长为55B")
	} else {
		ParHeader(pack[0], pack[1])
		fmt.Printf("主地址：%x%x\n", pack[2], pack[3])
		fmt.Printf("网络ID：%d\n版本号：%d\n", pack[4], pack[5])
		fmt.Printf("跳数：%d\n时隙长度：%d\n", pack[6], pack[7])
		fmt.Printf("超帧配置：%d\n帧序号：%d\n", int(pack[8])*256+int(pack[9]), int(pack[10])*256+int(pack[11]))
		fmt.Printf("广播周期：%d\n", int(pack[12])*256+int(pack[13]))
		fmt.Printf("下行时隙数：%d\n上行时隙数：%d\n", pack[14], pack[15])
		fmt.Println("时间保护域（单位：100us，默认取10）：")
		fmt.Printf("下行相邻帧保护域：%d\n", pack[16])
		fmt.Printf("上行相邻帧保护域：%d\n", pack[17])
		fmt.Printf("下行切上行保护域：%d\n", pack[18])
		fmt.Printf("上行切下行保护域：%d\n", pack[19])
		fmt.Printf("BCH长度：%d\n频点号：%d\n", pack[20], pack[21])
		fmt.Printf("MIC校验：%x%x\n", pack[24], pack[25])
	}
}