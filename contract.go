/**
* @Author:wjp
* @Date:2022/8/6 13:16
 */
package cricchainsdk

import (
	"cricchainsdk/chainproto"
	"cricchainsdk/chainresp"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

const functionType = "tx"

//  Mint
//  @Description: 铸造
//  @receiver c
//  @param contract 合约地址
//  @param tokenid token
//  @param uri 铸造链接
//  @return resp 返回参数
//  @return err 返回error
//
func (c *Chain) Mint(address common.Address, contract string, tokenid *big.Int, uri string, operateId string) (resp *chainresp.TxData, err error) {
	nonce, err := c.GetNonce()
	if err != nil {
		return
	}

	mint, err := c.Abi.SafeMint(address, tokenid, uri)
	if err != nil {
		return
	}
	data := chainproto.TransactionBody{}
	data.CodeData = mint
	data.To = address.Bytes()
	data.Address = c.Address.Bytes()
	data.Recipient, _ = hexutil.Decode(contract)
	data.Nonce = nonce
	info := chainproto.TransactionInfo{
		Body: &data,
	}
	mintbytes, err := c.ChainCall("POST", &info, map[string]string{}, map[string]interface{}{
		"method":       "safeMint",
		"functionType": functionType,
		"contractCode": "NFT_A",
		"operateId":    operateId,
	})
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(mintbytes, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

//
//  SafeTransfer
//  @Description: 从当前调用者地址转账到指定地址
//  @receiver c
//  @param toaddress 转账地址
//  @param contract 合约地址
//  @param tokenId tokenid
//  @return tx 返回hash
//  @return err 错误信息
//
func (c *Chain) SafeTransfer(toaddress common.Address, contract string, tokenid *big.Int, operateId string) (tx *chainresp.TxData, err error) {

	hexcode, err := c.Abi.SafeTransfer(c.Address, toaddress, tokenid)
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
	resp, err := c.ChainCall("POST", &info, map[string]string{}, map[string]interface{}{
		"method":       "safeTransfer",
		"functionType": functionType,
		"contractCode": "NFT_A",
		"operateId":    operateId,
	})
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &tx)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

//
//  SafeTransferForm
//  @Description:从某个地址转移token到指定地址
//  @receiver c
//  @param formaddress 转移地址
//  @param toaddress 转移到的地址
//  @param contract 合约地址
//  @param tokenid tokenid
//  @param operateId
//  @return tx
//  @return err
//
func (c *Chain) SafeTransferForm(formaddress, toaddress common.Address, contract string, tokenid *big.Int, operateId string) (tx *chainresp.TxData, err error) {

	hexcode, err := c.Abi.SafeTransferForm(formaddress, toaddress, tokenid)
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
	resp, err := c.ChainCall("POST", &info, map[string]string{}, map[string]interface{}{
		"method":       "safeTransferFrom",
		"functionType": functionType,
		"contractCode": "NFT_A",
		"operateId":    operateId,
	})
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &tx)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

//  Burn
//  @Description: 销毁
//  @receiver c
//  @param tokenId:tokenid
//  @param contract:合约地址
//  @return tx:返回的hash
//  @return err : 返回的error
//
func (c *Chain) Burn(tokenId *big.Int, contract string, operateId string) (tx *chainresp.TxData, err error) {

	hexcode, err := c.Abi.Burn(tokenId)
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
	resp, err := c.ChainCall("POST", &info, map[string]string{}, map[string]interface{}{
		"method":       "burn",
		"functionType": functionType,
		"contractCode": "NFT_A",
		"operateId":    operateId,
	})
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
func (c *Chain) TokenURI(tokenId *big.Int, contract string, operateId string) (tx *chainresp.TxData, err error) {

	hexcode, err := c.Abi.TokenURI(tokenId)
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
	resp, err := c.ChainCall("POST", &info, map[string]string{}, map[string]interface{}{
		"method":          "tokenURI",
		"operateId":       operateId,
		"functionType":    "view",
		"contractCode":    "NFT_A",
		"params":          []*big.Int{tokenId},
		"contractAddress": contract,
	})
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
func (c *Chain) SetApprovalForAll(operator common.Address, approved bool, contract string, operateId string) (tx *chainresp.TxData, err error) {

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
	resp, err := c.ChainCall("POST", &info, map[string]string{}, map[string]interface{}{
		"method":       "burn",
		"functionType": functionType,
		"contractCode": "NFT_A",
		"operateId":    operateId,
	})
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
func (c *Chain) addWhiteList(address common.Address, contract string, operateId string) (tx *chainresp.TxData, err error) {

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
	resp, err := c.ChainCall("POST", &info, map[string]string{}, map[string]interface{}{
		"method":       "addWhiteList",
		"operateId":    operateId,
		"functionType": functionType,
		"contractCode": "NFT_A",
	})
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
func (c *Chain) delWhiteList(address common.Address, contract string, operateId string) (tx *chainresp.TxData, err error) {
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
	resp, err := c.ChainCall("POST", &info, map[string]string{}, map[string]interface{}{
		"method":       "delWhiteList",
		"operateId":    operateId,
		"functionType": functionType,
		"contractCode": "NFT_A",
	})
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
func (c *Chain) getWhiteList(index *big.Int, contract string, operateId string) (tx *chainresp.TxData, err error) {

	hexcode, err := c.Abi.GetWhiteList(index)
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
	resp, err := c.ChainCall("POST", &info, map[string]string{}, map[string]interface{}{
		"method":          "inWhiteList",
		"functionType":    "view",
		"contractCode":    "NFT_A",
		"contractAddress": contract,
		"params":          []*big.Int{index},
		"operateId":       operateId,
	})
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
//是否在白名单中
func (c *Chain) inWhiteList(address common.Address, contract string, operateId string) (tx *chainresp.TxData, err error) {
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
	resp, err := c.ChainCall("POST", &info, map[string]string{}, map[string]interface{}{
		"method":          "getWhiteList",
		"functionType":    "view",
		"contractCode":    "NFT_A",
		"contractAddress": contract,
		"params":          []string{address.String()},
		"operateId":       operateId,
	})
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &tx)
	if err != nil {
		return nil, err
	}
	return tx, nil

}
