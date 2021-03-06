package account

import (
	"github.com/QOSGroup/qbase/account"
	go_amino "github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto/encoding/amino"
)

var cdc = go_amino.NewCodec()

// 包初始化，注册codec
func init() {
	cryptoAmino.RegisterAmino(cdc)
	RegisterCodec(cdc)
}

// 为包内定义结构注册codec
func RegisterCodec(cdc *go_amino.Codec) {
	cdc.RegisterInterface((*account.Account)(nil), nil)
	//cdc.RegisterInterface((*btypes.Coin)(nil), nil)
	cdc.RegisterConcrete(&QOSAccount{}, "qbase/account/QOSAccount", nil)
	//cdc.RegisterConcrete(&types.QSC{}, "qbase/coin/QSC", nil)
}
