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
func (abi *ChainAbi) SafeMint(address common.Address, tokenId *big.Int, url string) ([]byte, error) {

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

//授权
// operator:授权人
// approved:是否授权
func (abi *ChainAbi) SetApprovalForAll(address common.Address, approved bool) ([]byte, error) {
	return abi.Abi.Pack("setApprovalForAll", address, approved)
}

//添加白名单
func (abi *ChainAbi) AddWhiteList(address common.Address) ([]byte, error) {
	return abi.Abi.Pack("addWhiteList", address)
}

//删除白名单
func (abi *ChainAbi) DelWhiteList(address common.Address) ([]byte, error) {
	return abi.Abi.Pack("delWhiteList", address)
}

//获取白名单
func (abi *ChainAbi) GetWhiteList(index *big.Int) ([]byte, error) {
	return abi.Abi.Pack("getWhiteList", index)
}

//
func (abi *ChainAbi) InWhiteList(address common.Address) ([]byte, error) {
	return abi.Abi.Pack("inWhiteList", address)
}
