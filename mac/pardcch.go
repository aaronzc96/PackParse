package mac

import (
	"../byteop"
	"fmt"
)

func ParDcch(pack []byte) {
	if len(pack) < 2 {
		fmt.Println("长度错误！")
	} else {
		ParHeader(pack[0], pack[1])
		fmt.Printf("主地址：%x%x\n", pack[2], pack[3])
		for i := 4; i < len(pack); i++ {
			msgType, devnum := byteop.GetBits(pack[i], 7, 5), byteop.GetBits(pack[i], 4, 0)
			switch msgType {
			case 0:
				fmt.Println("USCH调度：")
				for j := 0; j < devnum; j++ {
					fmt.Printf("从地址：%x%x 起始时隙：%d 结束时隙：%d\n", pack[i], pack[i+1], pack[i+2], pack[i+3])
					i += 4
				}
			case 1:
				fmt.Println("DRX调度：")
				for j := 0; j < devnum; j++ {
					fmt.Printf("从地址：%x%x DRX帧数：%d\n", pack[i], pack[i+1], int(pack[i+2])*256+int(pack[i+3]))
					i += 4
				}
			case 2:
				fmt.Println("注册成功应答：")
				for j := 0; j < devnum; j++ {
					fmt.Printf("设备标识：%x%x%x%x%x%x 通信地址：%x%x\n", pack[i], pack[i+1], pack[i+2], pack[i+3],
						pack[i+4], pack[i+5], pack[i+6], pack[i+7])
					i += 8
				}
			case 3:
				fmt.Println("上行接收应答：")

			}
		}
	}
}
