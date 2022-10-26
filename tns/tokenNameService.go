package tns

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	lru "github.com/hashicorp/golang-lru"
)

var (
	cc       *lru.ARCCache
	upstream = "https://rpc.ankr.com/eth"
)

func init() {
	var err error
	cc, err = lru.NewARC(30000)
	if err != nil {
		panic(err)
	}
}

func GetName(addr common.Address) (string, error) {
	if v, ok := cc.Get(addr); ok {
		return v.(string), nil
	}
	cl, err := ethclient.Dial(upstream)
	if err != nil {
		return "", err
	}
	erc20, err := NewERC20(addr, cl)
	if err != nil {
		return "", err
	}
	name, err := erc20.Name(&bind.CallOpts{})
	if err != nil {
		return "", err
	}
	cc.Add(addr, name)
	return name, nil
}
