/**
* @Author:wjp
* @Date:2022/8/6 16:49
 */
package utils

import (
	"fmt"
	"math/big"
	"time"
)

var Number = 0

// ToMoney 转换金额
//把10进制金额转换为系统金额
//如传0.01 会被转换为10000000000000000
func ToMoney(money float64) *big.Int {

	bigval := new(big.Float)
	bigval.SetFloat64(money)
	coin := new(big.Float)
	coin.SetInt(big.NewInt(1000000000000000000))

	bigval.Mul(bigval, coin)

	result := new(big.Int)
	bigval.Int(result)

	return result
}

func OperateId() int64 {
	if Number == 9999 {
		Number = 0
	} else {
		Number = Number + 1
	}
	fmt.Println(fmt.Sprintf("%d0%04d", time.Now().UnixMicro(), Number))
	return 0
}
