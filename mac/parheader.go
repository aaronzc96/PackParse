package mac

import (
	"../byteop"
	"fmt"
)

func bitOut(b int, s0 string, s1 string) {
	if b == 0 {
		fmt.Println(s0)
	} else {
		fmt.Println(s1)
	}
}

func ParHeader(kind byte, len byte) {
	fmt.Print("网络层指示：")
	bitOut(byteop.GetBit(kind, 3), "MAC层数据", "网络层数据")
	fmt.Print("应答指示：")
	bitOut(byteop.GetBit(kind, 2), "无需应答", "需应答")
	fmt.Print("MIC指示：")
	bitOut(byteop.GetBit(kind, 1), "无MIC", "有MIC")
	fmt.Print("加密指示：")
	bitOut(byteop.GetBit(kind, 0), "未加密", "已加密")
	fmt.Printf("负载长度：%d\n", len)
}