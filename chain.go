/*
@Time : 2022/8/2 9:01 PM
@Author : wujianping
@File : chain
@Software: GoLand
*/
package cricchainsdk

import (
	"cricchainsdk/crypto"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	"math/big"
)

type Chain struct {
	Privatekey *ecdsa.PrivateKey `json:"privatekey"`
	BaseUrl    string            `json:"base_url"`
	Abi        *ChainAbi
	Address    common.Address `json:"address"`
}

func NewChain(url string, privatekey *ecdsa.PrivateKey, address common.Address) (*Chain, error) {
	//加载abi
	abi, err := AbiNew()
	if err != nil {
		return nil, err
	}
	return &Chain{
		Privatekey: privatekey,
		BaseUrl:    url,
		Abi:        abi,
		Address:    address,
	}, nil
}

//铸造
func (c *Chain) Mint(address common.Address, tokenId string, url string) (map[string]interface{}, error) {
	//加载abi
	amount, _ := new(big.Int).SetString(tokenId, 10) //10
	mite, err := c.Abi.Mite(address, amount, url)
	if err != nil {
		return nil, err
	}
	//对abi签名
	//sign, err := crypto.Sign(mite, c.Privatekey)
	//if err != nil {
	//	return nil, err
	//}
	fmt.Println(hexutil.Encode(mite))
	if err != nil {
		return nil, err
	}
	req, err := c.Req("/chain/mint.json", "POST", map[string]string{
		"x-singer":       "0xFeBbC825f4e7195b951D0B670447968D251e2F89",
		"x-authrization": "123",
	}, map[string]interface{}{
		"txData": hexutil.Encode(mite),
	})
	if err != nil {
		return nil, err
	}
	//发送后端铸造
	return req, nil
}

////转账
//func (c *Chain) SafeTransfer(fromaddress common.Address, toaddress common.Address, tokenid string) (map[string]interface{}, error) {
//	//加载abi
//	amount, _ := new(big.Int).SetString(tokenid, 10) //10
//	transfer, err := c.Abi.SafeTransfer(fromaddress, toaddress, amount)
//	if err != nil {
//		return nil, err
//	}
//	fmt.Println(hexutil.Encode(transfer))
//	return nil, nil
//}

//销毁
func (c *Chain) Burn(tokenId string) (map[string]interface{}, error) {
	//加载abi
	amount, _ := new(big.Int).SetString(tokenId, 10) //10
	burn, err := c.Abi.Burn(amount)
	if err != nil {
		return nil, err
	}
	fmt.Println(hexutil.Encode(burn))
	return nil, nil
}

func (c *Chain) TokenURI(tokenId string) (map[string]interface{}, error) {
	amount, _ := new(big.Int).SetString(tokenId, 10) //10
	uri, err := c.Abi.TokenURI(amount)
	if err != nil {
		return nil, err
	}
	fmt.Println(hexutil.Encode(uri))

	return nil, nil
}

//获取测试币
func (c *Chain) GetChainFaucet(address common.Address) (map[string]interface{}, error) {
	fmt.Println(address.String())
	req, err := c.Req("/chain/faucet.json", "GET", nil, map[string]interface{}{
		"address": address.String(),
	})
	return req, err
}

//查询账户信息
func (c *Chain) GetAccount() (map[string]interface{}, error) {
	req, err := c.Req("/chain/account.json", "GET", nil, map[string]interface{}{
		"address": c.Address.String(),
	})
	return req, err
}

//实名认证
//realName:姓名
//idCardNo:身份证号
//address:地址
func (c *Chain) PostRealAuth(realName string, idCardNo string, address string) (map[string]interface{}, error) {
	return c.Req("/chain/realAuth.json", "POST", nil, map[string]interface{}{
		"address":  address,
		"idCardNo": idCardNo,
		"realName": realName,
	})
}

//转账cric
func (c *Chain) TransferCric(transacetioninfo *TransactionInfo) (map[string]interface{}, error) {

	marshal, err := proto.Marshal(transacetioninfo.Body)

	sign := crypto.Sign(marshal, c.Privatekey)

	signature, err := hex.DecodeString(sign)

	transacetioninfo.Signature = signature

	if err != nil {
		return nil, err
	}

	infobyte, err := proto.Marshal(transacetioninfo)

	req, err := c.Req("/chain/transferCric.json", "POST", map[string]string{}, map[string]interface{}{
		"txData": hex.EncodeToString(infobyte),
	})
	if err != nil {

		return nil, err
	}
	return req, nil

}

// 获取交易详情
func (c *Chain) TransactionInfo(hash string) (map[string]interface{}, error) {
	resp, err := c.Req("/chain/transaction.json", "GET", nil, map[string]interface{}{
		"hash": hash,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Chain) Req(url string, method string, header map[string]string, data map[string]interface{}) (map[string]interface{}, error) {
	// 默认是application/x-www-form-urlencoded
	url = fmt.Sprintf("%s%s", c.BaseUrl, url)
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req) // 用完需要释放资源
	req.Header.SetContentType("application/json")
	req.Header.SetMethod(method)
	timestamp := uuid.New().String()
	for key, v := range header {
		req.Header.Set(key, v)
	}
	if method == "GET" {
		url = fmt.Sprintf("%s?", url)
		for key, value := range data {
			url = fmt.Sprintf("%s&%s=%s", url, key, value)
		}
	} else {
		requestBody, _ := json.Marshal(data)
		req.SetBody(requestBody)
	}
	req.Header.Set("x-request-id", timestamp)

	req.SetRequestURI(url)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp) // 用完需要释放资源

	if err := fasthttp.Do(req, resp); err != nil {
		return nil, err

	}
	if resp.StatusCode() != 200 {
		return nil, errors.New("请求失败")
	}
	b := resp.Body()

	respdata := map[string]interface{}{}
	err := json.Unmarshal(b, &respdata)
	if err != nil {
		return nil, err
	}
	fmt.Println("result:\r\n", string(b))
	return respdata, nil

}
