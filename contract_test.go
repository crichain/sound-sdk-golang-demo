/**
* @Author:wjp
* @Date:2022/8/6 18:16
 */
package cricchainsdk

import (
	"cricchainsdk/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

//铸造0xc85138f548911f15f833b7c20ee93222067ba8eff1f60fcaffce89fb681c2480
func TestMint(t *testing.T) {
	chain := TsNewChain()
	mint, err := chain.Mint(common.HexToAddress("0x06e81b2bc890f56d496e9938f1a8769518496d24"), "0xce7e273ed4081e6309664734dc7a162e2e20e6cd", big.NewInt(129), "https://ipfs.infura.io/ipfs/QmbApAkdkGj4jFu6Jr2thcNHraRBYJ7nEL7cvpabM7bLcK", utils.OperateId())
	if err != nil {
		return
	}
	fmt.Println(mint)

}

//转账0xc9df471a6fcb288ec3d9807d6af1b28b8caadec317973f9ed3065687e8e9ac8b
func TestChain_SafeTransfer(t *testing.T) {
	chain := TsNewChain()
	transfer, err := chain.SafeTransfer(common.HexToAddress("0x61d4c124df65ba081992ff2a8c77c67a8b3cb77c"), "0xce7e273ed4081e6309664734dc7a162e2e20e6cd", big.NewInt(128), utils.OperateId())
	if err != nil {
		return
	}
	fmt.Println(transfer)
}

//从某个地址转移到某个地址
func TestChain_SafeTransferForm(t *testing.T) {
	chain := TsNewChain()
	transfer, err := chain.SafeTransferForm(common.HexToAddress("0x61d4c124df65ba081992ff2a8c77c67a8b3cb77c"), common.HexToAddress("0x06e81b2bc890f56d496e9938f1a8769518496d24"), "0xce7e273ed4081e6309664734dc7a162e2e20e6cd", big.NewInt(127), "123123")
	if err != nil {
		return
	}
	fmt.Println(transfer)
}

//销毁
func TestChain_Burn(t *testing.T) {
	chain := TsNewChain()
	burn, err := chain.Burn(big.NewInt(125), "0xce7e273ed4081e6309664734dc7a162e2e20e6cd", utils.OperateId())
	if err != nil {
		return
	}
	fmt.Println(burn)
}

//获取tokenid
func TestChain_TokenURI(t *testing.T) {
	chain := TsNewChain()
	uri, err := chain.TokenURI(big.NewInt(127), "0xce7e273ed4081e6309664734dc7a162e2e20e6cd", utils.OperateId())
	if err != nil {
		return
	}
	fmt.Println(uri)
}

////添加白名单
//func TestChain_AddWhiteList(t *testing.T) {
//	chain := TsNewChain()
//	whilelist, err := chain.AddWhiteList(common.HexToAddress("0x06e81b2bc890f56d496e9938f1a8769518496d24"), "0xce7e273ed4081e6309664734dc7a162e2e20e6cd", "123123")
//	if err != nil {
//		return
//	}
//	fmt.Println(whilelist)
//}
//
////获取白名单
//func TestChain_GetWhiteList(t *testing.T) {
//	chain := TsNewChain()
//	list, err := chain.GetWhiteList(big.NewInt(1), "0xce7e273ed4081e6309664734dc7a162e2e20e6cd", "123123")
//	if err != nil {
//		return
//	}
//	fmt.Println(list)
//}
//
//func TestChain_InWhiteList(t *testing.T) {
//	chain := TsNewChain()
//	inwhite, err := chain.InWhiteList(common.HexToAddress("0x61d4c124df65ba081992ff2a8c77c67a8b3cb77c"), "0xce7e273ed4081e6309664734dc7a162e2e20e6cd", "123123")
//	if err != nil {
//		return
//	}
//	fmt.Println(inwhite)
//}

//删除白名单
//func TestChain_DelWhiteList(t *testing.T) {
//	chain := TsNewChain()
//	inwhite, err := chain.delWhiteList(common.HexToAddress("0x61d4c124df65ba081992ff2a8c77c67a8b3cb77c"), "0xce7e273ed4081e6309664734dc7a162e2e20e6cd", "123123")
//	if err != nil {
//		return
//	}
//	fmt.Println(inwhite)
//}
//
////授权
//func TestChainAbi_SetApprovalForAll(t *testing.T) {
//	approval, err := TsNewChain().SetApprovalForAll(common.HexToAddress("0x61d4c124df65ba081992ff2a8c77c67a8b3cb77c"), true, "0xce7e273ed4081e6309664734dc7a162e2e20e6cd", "123123")
//	if err != nil {
//		return
//	}
//	fmt.Println(approval)
//}
