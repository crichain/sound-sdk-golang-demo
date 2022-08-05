/*
@Time : 2022/8/2 9:39 PM
@Author : wujianping
@File : chain_test
@Software: GoLand
*/
package cricchainsdk

import (
	"cricchainsdk/crypto"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"testing"
	"time"
)

//私钥:4e4e8c93e1774c4100a2edda3a11960551ba6f083672f88fdbd9863a7f66cbc9
//公钥:2cbad470c3260b49a366d5010b6672ac21cacb91bd0def30cadb40a9c07ecd50b2353626338182a229902426e3dc947b6d57e841692e71c3be0a743f2cbe9d0f
//地址:0x06e81b2bc890f56d496e9938f1a8769518496d24
var Url = "http://172.16.60.138:3001"

func TsNewChain() *Chain {
	//加载私钥
	toECDSA, err := crypto.HexToECDSA("4e4e8c93e1774c4100a2edda3a11960551ba6f083672f88fdbd9863a7f66cbc9")
	if err != nil {
		panic("加载私钥报错")
	}
	formaddress := crypto.PubkeyToAddress(toECDSA.PublicKey)
	chain, err := NewChain(Url, toECDSA, formaddress)
	if err != nil {
		return nil
	}
	return chain
}
func TestMint(t *testing.T) {
	chain := TsNewChain()
	//marshal, err := protojson.Marshal(bbbb)
	//if err != nil {
	//	return
	//}

	//fmt.Println(marshal)
	//hexs := hexutil.Encode(b)
	//datas := map[string]interface{}{
	//	"body":      data,
	//	"signature": sign,
	//}
	//marshal, err := json.Marshal(datas)
	//if err != nil {
	//	return
	//}
	//encode := hexutil.Encode(marshal)
	//fmt.Println(encode)
	//marshal, err := json.Marshal(hexs)
	//if err != nil {
	//	return
	//}

	//sign := crypto.Sign(marshal, chain.Privatekey)

	//实名认证
	//auth, err := chain.PostRealAuth("邬建平", "230125198911254813", formaddress.String())
	//if err != nil {
	//	return
	//}
	//fmt.Println(auth)
	//获取账户信息
	account, err := chain.GetAccount()
	if err != nil {
		return
	}
	fmt.Println(account)
	//铸造
	//mint, err := chain.Mint(common.HexToAddress("0x1bF569aC852d52605850DF5f28E8a038cCcFd330"), "10", "https://json.meta-world.com/asset/blockchain/metadata/zNBuLBkPFqkbyZfXoCKZ")
	//if err != nil {
	//	return
	//}
	//fmt.Println(mint)
	//转账
	//transfer, err := chain.SafeTransfer(formaddress, common.HexToAddress("0xff01929ed3e0019b5f146606effab315a8da6ef8"), "6")
	//if err != nil {
	//	return
	//}
	//fmt.Println(transfer)

	////销毁
	//burn, err := chain.Burn("6")
	//if err != nil {
	//	return
	//}
	//fmt.Println(burn)
	//获取token
	//uri, err := chain.TokenURI("10")
	//if err != nil {
	//	return
	//}
	//fmt.Println(uri)
	//获取测试币
	//faucet, err := chain.GetChainFaucet(formaddress)
	//fmt.Println(faucet)
	//if err != nil {
	//	return
	//}

}

//转账
func TestChain_TransferCric(t *testing.T) {

	chain := TsNewChain()

	account, err := chain.GetAccount()
	if err != nil {
		return
	}
	nonce := int64(account["data"].(map[string]interface{})["nonce"].(float64))

	data := TransactionBody{}
	decode, err := hexutil.Decode("0x06e81b2bc890f56d496e9938f1a8769518496d24")
	if err != nil {
		return
	}
	toaddress, err := hexutil.Decode("0xff01929ed3e0019b5f146606effab315a8da6ef8")
	if err != nil {
		return
	}
	data.Address = decode
	data.Recipient = toaddress
	//"0.0000000000001"
	//"0.0000000000001"
	fmt.Println(hexutil.Encode(big.NewInt(10000000).Bytes()))
	amount, err := hexutil.Decode(hexutil.Encode(big.NewInt(1000000).Bytes()))
	if err != nil {
		return
	}
	if err != nil {
		return
	}
	if err != nil {
		return
	}
	data.Amount = amount
	data.Nonce = nonce + 1
	data.InnerCodetype = 0
	data.Timestamp = time.Now().Unix()
	decodeString, err := hexutil.Decode("0x01")
	if err != nil {
		return
	}
	data.Version = decodeString
	data.ChainId = 168

	info := TransactionInfo{
		Body: &data,
	}
	cric, err := chain.TransferCric(&info)
	fmt.Println(cric)
	if err != nil {
		return
	}

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
