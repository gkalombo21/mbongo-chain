package bank

import (
	"fmt"
	"github.com/gkalombo21/mbongo-chain/x/bank/keeper"
	"github.com/gkalombo21/mbongo-chain/x/bank/types"
)

// HandleMsgSend executes the Send message.
func HandleMsgSend(k *keeper.Keeper, msg types.MsgSend) error {
	if err := msg.ValidateBasic(); err != nil {
		return err
	}
	if err := k.Send(msg.From, msg.To, msg.Amount); err != nil {
		return err
	}
	fmt.Printf("Transferred %d tokens from %s to %s\n", msg.Amount, msg.From, msg.To)
	return nil
}
