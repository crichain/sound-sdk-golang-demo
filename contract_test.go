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

// 铸造
func TestMint(t *testing.T) {
	chain := TsNewChain()
	// i := int64(8888)

	mint, err := chain.Mint(common.HexToAddress(""), "", big.NewInt(129), "", utils.OperateId(), nil)
	if err != nil {
		return
	}
	fmt.Println(mint)

}

// 转账
func TestChain_SafeTransfer(t *testing.T) {
	chain := TsNewChain()
	transfer, err := chain.SafeTransfer(common.HexToAddress(""), "", big.NewInt(128), utils.OperateId(), nil)
	if err != nil {
		return
	}
	fmt.Println(transfer)
}

// 从某个地址转移到某个地址
func TestChain_SafeTransferForm(t *testing.T) {
	chain := TsNewChain()
	transfer, err := chain.SafeTransferForm(common.HexToAddress(""), common.HexToAddress(""), "", big.NewInt(127), utils.OperateId())
	if err != nil {
		return
	}
	fmt.Println(transfer)
}

// 销毁
func TestChain_Burn(t *testing.T) {
	chain := TsNewChain()
	burn, err := chain.Burn(big.NewInt(125), "", utils.OperateId())
	if err != nil {
		return
	}
	fmt.Println(burn)
}

// 获取tokenid
func TestChain_TokenURI(t *testing.T) {
	chain := TsNewChain()
	uri, err := chain.TokenURI(big.NewInt(127), "", utils.OperateId())
	if err != nil {
		return
	}
	fmt.Println(uri)
}
