/*
@Time : 2022/8/2 9:01 PM
@Author : wujianping
@File : chain
@Software: GoLand
*/
package cricchainsdk

import (
	"cricchainsdk/chainproto"
	"cricchainsdk/chainresp"
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
	"log"
	"regexp"
	"time"
)

type Chain struct {
	Privatekey *ecdsa.PrivateKey `json:"privatekey"`
	BaseUrl    string            `json:"base_url"`
	Abi        *ChainAbi
	Address    common.Address `json:"address"`
}

func NewChain(url string, privatekey *ecdsa.PrivateKey, address common.Address, path string) (*Chain, error) {
	//加载abi
	abi, err := AbiNew(path)
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

//获取测试币
func (c *Chain) GetChainFaucet() (data *chainresp.TxData, err error) {

	resp, err := c.Req("/chain/faucet.json", "GET", nil, map[string]interface{}{
		"address": c.Address.String(),
	})
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
	//err = json.Unmarshal(resp, &account)
	//if err != nil {
	//	return nil, err
	//}
	//return account, err
}

//查询账户信息
func (c *Chain) GetAccount() (account *chainresp.Account, err error) {
	resp, err := c.Req("/chain/account.json", "GET", nil, map[string]interface{}{
		"address": c.Address.String(),
	})
	err = json.Unmarshal(resp, &account)
	if err != nil {
		return nil, err
	}
	return account, nil
}

//实名认证
//realName:姓名
//idCardNo:身份证号
//address:地址
func (c *Chain) PostRealAuth(realName string, idCardNo string, address string) (map[string]interface{}, error) {
	resp, err := c.Req("/chain/realAuth.json", "POST", nil, map[string]interface{}{
		"address":  address,
		"idCardNo": idCardNo,
		"realName": realName,
	})
	if err != nil {
		return nil, err
	}
	data := map[string]interface{}{}
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//转账cric
func (c *Chain) TransferCric(transacetioninfo *chainproto.TransactionInfo, operateId string) (txdata *chainresp.TxData, err error) {
	//时间戳
	transacetioninfo.Body.Timestamp = time.Now().UnixMilli()
	//版本转[]byte
	decodeString, err := hexutil.Decode("0x562e32303232")
	//版本
	transacetioninfo.Body.Version = decodeString
	//节点
	transacetioninfo.Body.ChainId = 168

	transacetioninfo.Body.Nonce, _ = c.GetNonce()

	marshal, err := proto.Marshal(transacetioninfo.Body)

	sign := crypto.Sign(marshal, c.Privatekey)

	signature, _ := hex.DecodeString(sign)

	transacetioninfo.Signature = signature

	infobyte, err := proto.Marshal(transacetioninfo)

	resp, err := c.Req("/chain/transferCric.json", "POST", map[string]string{}, map[string]interface{}{
		"txData": hex.EncodeToString(infobyte),
	})
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &txdata)
	if err != nil {
		return nil, err
	}
	return txdata, nil
}

// 获取交易详情
func (c *Chain) TransactionInfo(hash string) (info *chainresp.TransactionInfo, err error) {
	resp, err := c.Req("/chain/transaction.json", "GET", nil, map[string]interface{}{
		"hash": hash,
	})
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (c *Chain) ChainCall(method string, info *chainproto.TransactionInfo, handler map[string]string, data map[string]interface{}) ([]byte, error) {
	//时间戳
	info.Body.Timestamp = time.Now().UnixMilli()
	//版本转[]byte
	decodeString, err := hexutil.Decode("0x562e32303232")
	//版本
	info.Body.Version = decodeString
	//节点
	info.Body.ChainId = 168
	info.Body.Nonce, _ = c.GetNonce()
	marshal, err := proto.Marshal(info.Body)

	sign := crypto.Sign(marshal, c.Privatekey)

	signature, _ := hex.DecodeString(sign)

	info.Signature = signature

	infobyte, err := proto.Marshal(info)
	data["data"] = hex.EncodeToString(infobyte)
	req, err := c.Req("/chain/callcontract.json", method, handler, data)
	if err != nil {

		return nil, err
	}
	return req, nil
}

//获取noner
func (c *Chain) GetNonce() (int64, error) {
	account, err := c.GetAccount()
	if err != nil {
		return 0, err
	}
	return account.Data.Nonce, nil
}

//请求类
func (c *Chain) Req(url string, method string, header map[string]string, data map[string]interface{}) ([]byte, error) {
	// 默认是application/x-www-form-urlencoded
	url = fmt.Sprintf("%s%s", c.BaseUrl, url)
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req) // 用完需要释放资源
	req.Header.SetContentType("application/json")
	req.Header.SetMethod(method)

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
	request_id := uuid.New().String()
	fmt.Println(request_id)
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	request_id = reg.ReplaceAllString(request_id, "")
	fmt.Println("x-request-id:", request_id)
	req.Header.Set("x-request-id", request_id)

	req.SetRequestURI(url)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp) // 用完需要释放资源

	if err := fasthttp.Do(req, resp); err != nil {
		return nil, err

	}
	if resp.StatusCode() != 200 {
		return nil, errors.New("请求失败")
	}
	fmt.Println("请求参数为:", data)
	fmt.Println("返回参数为:", string(resp.Body()))
	return resp.Body(), nil

}
