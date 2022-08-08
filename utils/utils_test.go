/**
* @Author:wjp
* @Date:2022/8/7 18:23
 */
package utils

import (
	"fmt"
	"testing"
)

func TestOperateId(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(OperateId())
	}

}
