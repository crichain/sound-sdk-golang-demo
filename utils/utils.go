/**
* @Author:wjp
* @Date:2022/8/6 16:49
 */
package utils

import (
	"fmt"
	"math/big"
	"strconv"
)

//转换金额
//s: 金额 18位
//hex 进制 默认16进制
func ToBigInt(s string, hex int) (*big.Int, bool) {
	if s == "" {
		return nil, false
	}
	if hex == 0 {
		hex = 16
	}
	return new(big.Int).SetString(s, hex)
}

//转换金额
//
func ToMoney(money string) *big.Int {
	//i := int64(32)
	//x := 0xde0b6b3a7640000

	float, err := strconv.ParseFloat(money, 16)
	if err != nil {
		return nil
	}
	fmt.Println(float)
	//fmt.Println(x*float)

	//x := big.NewFloat(money)
	//
	//i2 := x.Mul(x, big.NewFloat(1000000000000000000))
	//i3, _ := i2.Int64()
	//s := strconv.FormatInt(i3, 16)
	//fmt.Println(s)
	return new(big.Int).SetInt64(1)
}
