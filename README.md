# chainsdk
# 生成秘钥对及地址
## 注:最好使用eth工具包,都是16进制的
## 一,生成秘钥对
* 1,使用标准secp256r1算法生成秘钥对
## 二,生成私钥字符串
* 1,把私钥转换字节数组
- 2,翻转字节数组<br>
* 例子:反转前:   67f30b565660cd169986b7c5fe71ed000523ce64d9dcff9e64162eecfa61e9ec<br>
  转换过的私钥:    ece961faec2e16649effdcd964ce230500ed71fec5b7869916cd6056560bf367<br>
## 三,生成公钥字符串:
* 1,拿到公钥的 x,y 把x,y 都转成字节数组,然后翻转<br>
  例:原始x:    d08b7cd8b2aaa97d09b15bd52005b59617941b52630a6e1577b688ae348acd47<br>
  转换后的x:   47cd8a34ae88b677156e0a63521b941796b50520d55bb1097da9aab2d87c8bd0<br>
  原始y:      24f671a785d11211ce7506dddc5d778556eb1e9979f023d09e6f9ed6babdd860<br>
  转换后的y:   60d8bdbad69e6f9ed023f079991eeb5685775ddcdd0675ce1112d185a771f624<br>
* 2, x,y拼接成一个公钥
* 例:转换后的公钥为:47cd8a34ae88b677156e0a63521b941796b50520d55bb1097da9aab2d87c8bd060d8bdbad69e6f9ed023f079991eeb5685775ddcdd0675ce1112d185a771f624
## 四,生成地址:
- 1,把公钥字符串转换成16进制[]byte
- 2,把16进制byte用sha256.sum256 方法转换
- 3,在把转换完的16进制byte 取前20个
- 4,把处理完的16进制byte转换成字符串 在加一个0x前缀就是地址了
- 生成的地址为:ec3d00c4c2b6f65548bc235ba4ebe0195022a7ea


# 调用示例
### 创建实例
```
  ftoECDSA, err := crypto.HexToECDSA("4e4e8c93e1774c4100a2edda3a11960551ba6f083672f88fdbd9863a7f66cbc9")
	if err != nil {
		panic("加载私钥报错")
	}
	formaddress := crypto.PubkeyToAddress(toECDSA.PublicKey)
	chain, err := NewChain(Url, toECDSA, formaddress, "abi.abi")
	if err != nil {
		return nil
	}
	return chain
```

### 转账cric
```
   chain := TsNewChain()
	data := chainproto.TransactionBody{}
	decode, err := hexutil.Decode(chain.Address.String())
	if err != nil {
		return
	}
	toaddress, err := hexutil.Decode("0x61d4c124df65ba081992ff2a8c77c67a8b3cb77c")
	if err != nil {
		return
	}
	data.Address = decode
	data.Recipient = toaddress
	data.To = toaddress

	amount, err := hexutil.Decode(hexutil.Encode(utils.ToMoney(1000).Bytes()))
	if err != nil {
		return
	}

	data.Amount = amount


	info := chainproto.TransactionInfo{
		Body: &data,
	}
	cric, err := chain.TransferCric(&info, utils.OperateId())
	fmt.Println(cric)
	if err != nil {
		return
	}
```
### 获取交易详情
```
info, err := chain.TransactionInfo("0x8478b00a085bcbc3400bfc39e49bfc0f2402ca7ea8188b4ce63555901ddfe67c")
	if err != nil {
		return
	}
	fmt.Println(info)
```
### 获取测试币
```
    chain := TsNewChain()
	faucet, err := chain.GetChainFaucet()
	if err != nil {
		return
	}
	fmt.Println(faucet)
```
### 获取用户信息
```go
    chain := TsNewChain()
	account, err := chain.GetAccount()
	if err != nil {
		return
	}
	fmt.Println(account)
```
### 铸造
```go
	chain := TsNewChain()
    mint, err := chain.Mint(common.HexToAddress("0x06e81b2bc890f56d496e9938f1a8769518496d24"), "0xce7e273ed4081e6309664734dc7a162e2e20e6cd", big.NewInt(129), "https://ipfs.infura.io/ipfs/QmbApAkdkGj4jFu6Jr2thcNHraRBYJ7nEL7cvpabM7bLcK", utils.OperateId())
    if err != nil {
    return
    }
    fmt.Println(mint)
```
### 转账
```go
	chain := TsNewChain()
    transfer, err := chain.SafeTransfer(common.HexToAddress("0x61d4c124df65ba081992ff2a8c77c67a8b3cb77c"), "0xce7e273ed4081e6309664734dc7a162e2e20e6cd", big.NewInt(128), utils.OperateId())
    if err != nil {
    return
    }
    fmt.Println(transfer)
```
### 销毁
```go
    chain := TsNewChain()
    burn, err := chain.Burn(big.NewInt(125), "0xce7e273ed4081e6309664734dc7a162e2e20e6cd", utils.OperateId())
    if err != nil {
    return
    }
    fmt.Println(burn)
```

### 获取token
```go
   chain := TsNewChain() 
   uri, err := chain.TokenURI(big.NewInt(127), "0xce7e273ed4081e6309664734dc7a162e2e20e6cd", utils.OperateId())
   if err != nil {
        return
   }
   fmt.Println(uri)
```

[//]: # (### 添加白名单)

[//]: # (```go)

[//]: # (    chain := TsNewChain&#40;&#41;)

[//]: # (	whilelist, err := chain.AddWhiteList&#40;common.HexToAddress&#40;"0x61d4c124df65ba081992ff2a8c77c67a8b3cb77c"&#41;, "0xce7e273ed4081e6309664734dc7a162e2e20e6cd"&#41;)

[//]: # (	if err != nil {)

[//]: # (		return)

[//]: # (	})

[//]: # (	fmt.Println&#40;whilelist&#41;)

[//]: # (```)

[//]: # (### 获取白名单)

[//]: # (```go)

[//]: # (    chain := TsNewChain&#40;&#41;)

[//]: # (	list, err := chain.GetWhiteList&#40;"1", "0xce7e273ed4081e6309664734dc7a162e2e20e6cd"&#41;)

[//]: # (	if err != nil {)

[//]: # (		return)

[//]: # (	})

[//]: # (	fmt.Println&#40;list&#41;)

[//]: # (```)

[//]: # (### inwhite)

[//]: # (```go)

[//]: # (    chain := TsNewChain&#40;&#41;)

[//]: # (	inwhite, err := chain.InWhiteList&#40;common.HexToAddress&#40;"0x61d4c124df65ba081992ff2a8c77c67a8b3cb77c"&#41;, "0xce7e273ed4081e6309664734dc7a162e2e20e6cd"&#41;)

[//]: # (	if err != nil {)

[//]: # (		return)

[//]: # (	})

[//]: # (	fmt.Println&#40;inwhite&#41;)

[//]: # (```)

[//]: # (### 删除白名单)

[//]: # (```go)

[//]: # (    chain := TsNewChain&#40;&#41;)

[//]: # (	inwhite, err := chain.DelWhiteList&#40;common.HexToAddress&#40;"0x61d4c124df65ba081992ff2a8c77c67a8b3cb77c"&#41;, "0xce7e273ed4081e6309664734dc7a162e2e20e6cd"&#41;)

[//]: # (	if err != nil {)

[//]: # (		return)

[//]: # (	})

[//]: # (	fmt.Println&#40;inwhite&#41;)

[//]: # (```)

[//]: # (### 授权)

[//]: # (```go)

[//]: # (    approval, err := TsNewChain&#40;&#41;.SetApprovalForAll&#40;common.HexToAddress&#40;"0x61d4c124df65ba081992ff2a8c77c67a8b3cb77c"&#41;, true, "0xce7e273ed4081e6309664734dc7a162e2e20e6cd"&#41;)

[//]: # (	if err != nil {)

[//]: # (		return)

[//]: # (	})

[//]: # (	fmt.Println&#40;approval&#41;)

[//]: # (```)