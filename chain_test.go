/*
@Time : 2022/8/2 9:39 PM
@Author : wujianping
@File : chain_test
@Software: GoLand
*/
package cricchainsdk

import (
	"cricchainsdk/chainproto"
	"cricchainsdk/crypto"
	"cricchainsdk/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"testing"
)

var Url = "http://test.open-api.crichain.cn"

func TsNewChain() *Chain {

	toECDSA, err := crypto.HexToECDSA("")
	if err != nil {
		panic("加载私钥报错")
	}
	formaddress := crypto.PubkeyToAddress(toECDSA.PublicKey)
	chain, err := NewChain(Url, toECDSA, formaddress, "abi.abi")
	if err != nil {
		return nil
	}
	return chain
}

// 获取测试币
func TestChain_GetChainFaucet(t *testing.T) {
	chain := TsNewChain()
	faucet, err := chain.GetChainFaucet()
	if err != nil {
		return
	}
	fmt.Println(faucet)
}

// 转账cric
func TestChain_TransferCric(t *testing.T) {

	chain := TsNewChain()
	data := chainproto.TransactionBody{}
	decode, err := hexutil.Decode(chain.Address.String())
	if err != nil {
		return
	}
	toaddress, err := hexutil.Decode("")
	if err != nil {
		return
	}
	data.Address = decode
	data.Recipient = toaddress
	data.To = toaddress

	amount, err := hexutil.Decode(hexutil.Encode(utils.ToMoney(1000).Bytes()))
	if err != nil {
		return
	}

	data.Amount = amount

	info := chainproto.TransactionInfo{
		Body: &data,
	}
	cric, err := chain.TransferCric(&info, utils.OperateId())
	fmt.Println(cric)
	if err != nil {
		return
	}

}

// 获取用户信息
func TestChain_GetAccount(t *testing.T) {
	chain := TsNewChain()
	account, err := chain.GetAccount()
	if err != nil {
		return
	}
	fmt.Println(account)
}

// 获取交易详情
func TestChain_TransactionInfo(t *testing.T) {
	ch := TsNewChain()
	info, err := ch.TransactionInfo("")
	if err != nil {
		return
	}
	fmt.Println(info)
}

// 实名认证
func TestChain_PostRealAuth(t *testing.T) {
	chain := TsNewChain()
	auth, err := chain.PostRealAuth("", "", "")
	if err != nil {
		return
	}
	fmt.Println(auth)
}
