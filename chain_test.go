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

//私钥:4e4e8c93e1774c4100a2edda3a11960551ba6f083672f88fdbd9863a7f66cbc9
//公钥:2cbad470c3260b49a366d5010b6672ac21cacb91bd0def30cadb40a9c07ecd50b2353626338182a229902426e3dc947b6d57e841692e71c3be0a743f2cbe9d0f
//地址:0x06e81b2bc890f56d496e9938f1a8769518496d24
var Url = "http://test.open-api.crichain.cn"

func TsNewChain() *Chain {
	//加载私钥b046347a995c3131c99fc5cf9e29ee4f7721e9b5ff06397df4eab597d08a9ef1
	//d5bb8ed7884cb7dc64a4170035f30d5589720cb02c027d86c7ac436d0f85b58c
	toECDSA, err := crypto.HexToECDSA("4e4e8c93e1774c4100a2edda3a11960551ba6f083672f88fdbd9863a7f66cbc9")
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

//获取测试币
func TestChain_GetChainFaucet(t *testing.T) {
	chain := TsNewChain()
	faucet, err := chain.GetChainFaucet()
	if err != nil {
		return
	}
	fmt.Println(faucet)
}

//转账cric
func TestChain_TransferCric(t *testing.T) {

	chain := TsNewChain()
	data := chainproto.TransactionBody{}
	decode, err := hexutil.Decode(chain.Address.String())
	if err != nil {
		return
	}
	toaddress, err := hexutil.Decode("0x61d4c124df65ba081992ff2a8c77c67a8b3cb77c")
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

//获取用户信息
func TestChain_GetAccount(t *testing.T) {
	chain := TsNewChain()
	account, err := chain.GetAccount()
	if err != nil {
		return
	}
	fmt.Println(account)
}

//获取交易详情
func TestChain_TransactionInfo(t *testing.T) {
	ch := TsNewChain()
	info, err := ch.TransactionInfo("0x8478b00a085bcbc3400bfc39e49bfc0f2402ca7ea8188b4ce63555901ddfe67c")
	if err != nil {
		return
	}
	fmt.Println(info)
}

//实名认证
func TestChain_PostRealAuth(t *testing.T) {
	chain := TsNewChain()
	auth, err := chain.PostRealAuth("邬建平", "230125198911254813", "0x06e81b2bc890f56d496e9938f1a8769518496d24")
	if err != nil {
		return
	}
	fmt.Println(auth)
}
