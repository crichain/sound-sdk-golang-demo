/**
* @Author:wjp
* @Date:2022/8/6 13:16
 */
package cricchainsdk

import (
	"cricchainsdk/chainproto"
	"cricchainsdk/chainresp"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

//铸造
func (c *Chain) Mint(contract, tokenid string, uri string) (resp *chainresp.TxData, err error) {
	nonce, err := c.GetNonce()
	if err != nil {
		return
	}
	setString, ok := new(big.Int).SetString(tokenid, 16)
	if !ok {
		fmt.Println("不ok")
		return nil, errors.New("tokenid转换失败")
	}
	mint, err := c.Abi.SafeMint(c.Address, setString, uri)
	if err != nil {
		return
	}
	data := chainproto.TransactionBody{}
	data.CodeData = mint
	data.Address = c.Address.Bytes()
	data.Recipient, _ = hexutil.Decode(contract)
	data.Nonce = nonce
	info := chainproto.TransactionInfo{
		Body: &data,
	}
	mintbytes, err := c.ChainCall("/chain/mint.json", "POST", &info)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(mintbytes, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

//转账
func (c *Chain) SafeTransfer(toaddress common.Address, contract string, tokenId string) (tx *chainresp.TxData, err error) {
	setString, ok := new(big.Int).SetString(tokenId, 16)
	if !ok {
		fmt.Println("不ok")
		return nil, errors.New("tokenid转换失败")
	}
	hexcode, err := c.Abi.SafeTransfer(c.Address, toaddress, setString)
	if err != nil {
		return nil, err
	}
	data := chainproto.TransactionBody{}
	data.CodeData = hexcode

	data.To = toaddress.Bytes()
	data.Address = c.Address.Bytes()
	data.Recipient, _ = hexutil.Decode(contract)
	data.Nonce, _ = c.GetNonce()
	info := chainproto.TransactionInfo{
		Body: &data,
	}
	resp, err := c.ChainCall("", "POST", &info)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &tx)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

//销毁
func (c *Chain) Burn(tokenId string, contract string) (tx *chainresp.TxData, err error) {
	setString, ok := new(big.Int).SetString(tokenId, 16)
	if !ok {
		fmt.Println("不ok")
		return nil, errors.New("tokenid转换失败")
	}
	hexcode, err := c.Abi.Burn(setString)
	if err != nil {
		return nil, err
	}
	data := chainproto.TransactionBody{}
	data.CodeData = hexcode
	data.Address = c.Address.Bytes()
	data.Recipient, _ = hexutil.Decode(contract)
	data.Nonce, _ = c.GetNonce()
	info := chainproto.TransactionInfo{
		Body: &data,
	}
	resp, err := c.ChainCall("", "POST", &info)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &tx)
	if err != nil {
		return nil, err
	}
	return tx, nil

}

//获取tokenuri
func (c *Chain) TokenURI(tokenId string, contract string) (tx *chainresp.TxData, err error) {
	setString, ok := new(big.Int).SetString(tokenId, 16)
	if !ok {
		fmt.Println("tokenid转换失败")
		return nil, errors.New("tokenid转换失败")
	}
	hexcode, err := c.Abi.TokenURI(setString)
	if err != nil {
		return nil, err
	}

	data := chainproto.TransactionBody{}
	data.CodeData = hexcode
	data.Address = c.Address.Bytes()
	data.Recipient, _ = hexutil.Decode(contract)
	data.Nonce, _ = c.GetNonce()
	info := chainproto.TransactionInfo{
		Body: &data,
	}
	resp, err := c.ChainCall("", "POST", &info)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &tx)
	if err != nil {
		return nil, err
	}
	return tx, nil

}

//合约授权
//operator: 操作人员
//approved:是否核准
func (c *Chain) SetApprovalForAll(operator common.Address, approved bool, contract string) (tx *chainresp.TxData, err error) {

	hexcode, err := c.Abi.SetApprovalForAll(operator, approved)
	if err != nil {
		return nil, err
	}
	data := chainproto.TransactionBody{}
	data.CodeData = hexcode
	data.Address = c.Address.Bytes()
	data.Recipient, _ = hexutil.Decode(contract)
	data.Nonce, _ = c.GetNonce()
	info := chainproto.TransactionInfo{
		Body: &data,
	}
	resp, err := c.ChainCall("", "POST", &info)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &tx)
	if err != nil {
		return nil, err
	}
	return tx, nil

}

//添加白名单
//address 白名单地址
func (c *Chain) AddWhiteList(address common.Address, contract string) (tx *chainresp.TxData, err error) {

	hexcode, err := c.Abi.AddWhiteList(address)
	if err != nil {
		return nil, err
	}
	data := chainproto.TransactionBody{}
	data.CodeData = hexcode
	data.Address = c.Address.Bytes()
	data.Recipient, _ = hexutil.Decode(contract)
	data.Nonce, _ = c.GetNonce()
	info := chainproto.TransactionInfo{
		Body: &data,
	}
	resp, err := c.ChainCall("", "POST", &info)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &tx)
	if err != nil {
		return nil, err
	}
	return tx, nil

}

//删除白名单
func (c *Chain) DelWhiteList(address common.Address, contract string) (tx *chainresp.TxData, err error) {
	hexcode, err := c.Abi.DelWhiteList(address)
	if err != nil {
		return nil, err
	}
	data := chainproto.TransactionBody{}
	data.CodeData = hexcode
	data.Address = c.Address.Bytes()
	data.Recipient, _ = hexutil.Decode(contract)
	data.Nonce, _ = c.GetNonce()
	info := chainproto.TransactionInfo{
		Body: &data,
	}
	resp, err := c.ChainCall("", "POST", &info)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &tx)
	if err != nil {
		return nil, err
	}
	return tx, nil

}

//获取白名单
//index 位置
//contract 合约地址
func (c *Chain) GetWhiteList(index string, contract string) (tx *chainresp.TxData, err error) {
	indexint, ok := new(big.Int).SetString(index, 16)
	if !ok {
		fmt.Println("tokenid转换失败")
		return nil, errors.New("tokenid转换失败")
	}
	hexcode, err := c.Abi.GetWhiteList(indexint)
	if err != nil {
		return nil, err
	}
	data := chainproto.TransactionBody{}
	data.CodeData = hexcode
	data.Address = c.Address.Bytes()
	data.Recipient, _ = hexutil.Decode(contract)
	data.Nonce, _ = c.GetNonce()
	info := chainproto.TransactionInfo{
		Body: &data,
	}
	resp, err := c.ChainCall("", "POST", &info)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &tx)
	if err != nil {
		return nil, err
	}
	return tx, nil

}

//inWhiteList
func (c *Chain) InWhiteList(address common.Address, contract string) (tx *chainresp.TxData, err error) {
	hexcode, err := c.Abi.InWhiteList(address)
	if err != nil {
		return nil, err
	}

	data := chainproto.TransactionBody{}
	data.CodeData = hexcode
	data.Address = c.Address.Bytes()
	data.Recipient, _ = hexutil.Decode(contract)
	data.Nonce, _ = c.GetNonce()
	info := chainproto.TransactionInfo{
		Body: &data,
	}
	resp, err := c.ChainCall("", "POST", &info)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &tx)
	if err != nil {
		return nil, err
	}
	return tx, nil

}
