/**
* @Author:wjp
* @Date:2022/8/2 18:04
 */
package crypto

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"io/ioutil"
	"math/big"
)

func TokenURI(tokenId string) []byte {
	abiData, err := ioutil.ReadFile("/Users/wujianping/go/src/DemoTest/abi/abi.abi")
	if err != nil {
		fmt.Println("read file: ", err)
		return nil
	}
	contractABI, err := abi.JSON(bytes.NewReader(abiData))
	if err != nil {
		fmt.Println("abi json: ", err)
		return nil
	}
	acc, _ := new(big.Int).SetString(tokenId, 10)
	callData, err := contractABI.Pack("tokenURI", acc)

	return callData
}
func SafeMint(to string, int2 *big.Int, uri string) string {
	abiData, err := ioutil.ReadFile("/Users/wujianping/go/src/DemoTest/abi/abi.abi")
	if err != nil {
		fmt.Println("read file: ", err)
		return ""
	}
	contractABI, err := abi.JSON(bytes.NewReader(abiData))
	if err != nil {
		fmt.Println("abi json: ", err)
		return ""
	}

	callData, err := contractABI.Pack("safeMint", to, int2, uri)
	fmt.Println(hexutil.Encode(callData))
	return hexutil.Encode(callData)
}
