/**
* @Author:wjp
* @Date:2022/8/2 18:10
 */
package crypto

import (
	"fmt"
	"testing"
)

var privkey, _ = HexToECDSA("b046347a995c3131c99fc5cf9e29ee4f7721e9b5ff06397df4eab597d08a9ef1")
var senderPriKey = "b046347a995c3131c99fc5cf9e29ee4f7721e9b5ff06397df4eab597d08a9ef1"
var toAddress = "0x76891bf2c57a813afd85d391fa87910153f18857"

func TestTokenIds(t *testing.T) {

	uri := TokenURI("11231231")

	fmt.Println(uri)
	fmt.Println(Sign(uri, privkey))
}
func TestSafeMint(t *testing.T) {
	//address := PubkeyToAddress(privkey.PublicKey)
	//fmt.Println(strings.ToLower(address.String()))
	//acc, _ := new(big.Int).SetString("1000000", 10)
	//mint := SafeMint(toAddress, acc, "hello")
	//fmt.Println(Sign("123123", privkey))
}
