/**
* @Author:wjp
* @Date:2022/8/6 18:16
 */
package cricchainsdk

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

//铸造
func TestMint(t *testing.T) {
	chain := TsNewChain()
	mint, err := chain.Mint("0xce7e273ed4081e6309664734dc7a162e2e20e6cd", "cbc657068a8e34d905cd7af6b06c9859133814a7", "https://ipfs.infura.io/ipfs/QmbApAkdkGj4jFu6Jr2thcNHraRBYJ7nEL7cvpabM7bLcK")
	if err != nil {
		return
	}
	fmt.Println(mint)

}

//转账
func TestChain_SafeTransfer(t *testing.T) {
	chain := TsNewChain()
	transfer, err := chain.SafeTransfer(common.HexToAddress("0x61d4c124df65ba081992ff2a8c77c67a8b3cb77c"), "0xce7e273ed4081e6309664734dc7a162e2e20e6cd", "cbc657068a8e34d905cd7af6b06c9859133814a7")
	if err != nil {
		return
	}
	fmt.Println(transfer)
}

//销毁
func TestChain_Burn(t *testing.T) {
	chain := TsNewChain()
	burn, err := chain.Burn("cbc657068a8e34d905cd7af6b06c9859133814a7", "0xce7e273ed4081e6309664734dc7a162e2e20e6cd")
	if err != nil {
		return
	}
	fmt.Println(burn)
}

//获取tokenid
func TestChain_TokenURI(t *testing.T) {
	chain := TsNewChain()
	uri, err := chain.TokenURI("cbc657068a8e34d905cd7af6b06c9859133814a7", "0xce7e273ed4081e6309664734dc7a162e2e20e6cd")
	if err != nil {
		return
	}
	fmt.Println(uri)
}

//添加白名单
func TestChain_AddWhiteList(t *testing.T) {
	chain := TsNewChain()
	whilelist, err := chain.AddWhiteList(common.HexToAddress("0x61d4c124df65ba081992ff2a8c77c67a8b3cb77c"), "0xce7e273ed4081e6309664734dc7a162e2e20e6cd")
	if err != nil {
		return
	}
	fmt.Println(whilelist)
}

//获取白名单
func TestChain_GetWhiteList(t *testing.T) {
	chain := TsNewChain()
	list, err := chain.GetWhiteList("1", "0xce7e273ed4081e6309664734dc7a162e2e20e6cd")
	if err != nil {
		return
	}
	fmt.Println(list)
}

func TestChain_InWhiteList(t *testing.T) {
	chain := TsNewChain()
	inwhite, err := chain.InWhiteList(common.HexToAddress("0x61d4c124df65ba081992ff2a8c77c67a8b3cb77c"), "0xce7e273ed4081e6309664734dc7a162e2e20e6cd")
	if err != nil {
		return
	}
	fmt.Println(inwhite)
}

//删除白名单
func TestChain_DelWhiteList(t *testing.T) {
	chain := TsNewChain()
	inwhite, err := chain.DelWhiteList(common.HexToAddress("0x61d4c124df65ba081992ff2a8c77c67a8b3cb77c"), "0xce7e273ed4081e6309664734dc7a162e2e20e6cd")
	if err != nil {
		return
	}
	fmt.Println(inwhite)
}

//授权
func TestChainAbi_SetApprovalForAll(t *testing.T) {
	approval, err := TsNewChain().SetApprovalForAll(common.HexToAddress("0x61d4c124df65ba081992ff2a8c77c67a8b3cb77c"), true, "0xce7e273ed4081e6309664734dc7a162e2e20e6cd")
	if err != nil {
		return
	}
	fmt.Println(approval)
}