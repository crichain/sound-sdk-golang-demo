/*
@Time : 2022/8/2 9:05 PM
@Author : wujianping
@File : abi
@Software: GoLand
*/
package cricchainsdk

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"io/ioutil"
	"math/big"
)

type ChainAbi struct {
	Abi abi.ABI
}

func AbiNew() (*ChainAbi, error) {
	abiData, err := ioutil.ReadFile("abi.abi")
	if err != nil {
		fmt.Println("read file: ", err)
		return nil, err
	}
	contractABI, err := abi.JSON(bytes.NewReader(abiData))
	if err != nil {
		fmt.Println("abi json: ", err)
		return nil, err
	}
	return &ChainAbi{
		Abi: contractABI,
	}, nil
}

//铸造
func (abi *ChainAbi) Mite(address common.Address, tokenId *big.Int, url string) ([]byte, error) {

	return abi.Abi.Pack("safeMint", address, tokenId, url)
}

//转账
func (abi *ChainAbi) SafeTransfer(fromaddress common.Address, toaddress common.Address, tokenId *big.Int) ([]byte, error) {

	return abi.Abi.Pack("safeTransfer", fromaddress, toaddress, tokenId)
}

//销毁
func (abi *ChainAbi) Burn(tokenId *big.Int) ([]byte, error) {
	return abi.Abi.Pack("burn", tokenId)
}

//获取tokenuri
func (abi *ChainAbi) TokenURI(tokenId *big.Int) ([]byte, error) {
	return abi.Abi.Pack("tokenURI", tokenId)
}
