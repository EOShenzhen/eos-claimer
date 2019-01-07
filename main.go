package main

import (
	"fmt"
	"log"
	"time"
	eos "github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/system"
	"encoding/hex"
)

var bp_name string = "eosstorebest"
var claim_key string = "your bp claimer private key"
var claim_key_permission_name string = "claimer"

var end_points = []string{
	"http://peer1.eoshuobipool.com:8181",
	"http://api-mainnet.starteos.io",
	"http://publicapi-mainnet.eosauthority.com",
	"https://nodes.eos42.io",
	"https://mainnet.eoscanada.com",
	"http://api.eos.store",
	"https://api.eosnewyork.io",
	"http://bp.cryptolions.io",
	"http://127.0.0.1:8888",
}

var chain_id string = "aca376f206b8fc25a6ed44dbdc66547c36c6c33e3a119ffbeaef943642f0e906"


func NewClaimRewards(owner eos.AccountName) *eos.Action {
	a := &eos.Action{
		Account: system.AN("eosio"),
		Name:    system.ActN("claimrewards"),
		Authorization: []eos.PermissionLevel{
			{Actor: owner, Permission: eos.PermissionName(claim_key_permission_name)},
		},
		ActionData: eos.NewActionData(system.ClaimRewards{
			Owner: owner,
		}),
	}
	return a
}

func main() {
	fmt.Println("current time :"  + time.Now().Format("2006-01-02:15-04-05"))

	// add private key
	keyBag := eos.NewKeyBag()
	if err := keyBag.Add(claim_key); err != nil {
		log.Fatalln("Couldn't load private key!", err)
	}

	// Find best api
	var api *eos.API
	for _, end_point := range end_points{
		api = eos.New(end_point)
		r, err := api.GetInfo()
		if err != nil {
			fmt.Println( "get info from" + end_point + "failed", err)
		} else if hex.EncodeToString(r.ChainID) == chain_id {
			break
		}
	}
	fmt.Println(api.BaseURL)

	// SetSigner
	api.SetSigner(keyBag)

	// Claim rewards
	act := NewClaimRewards( eos.AccountName( bp_name ) )
	sleep := 5 * time.Second
	for i := 0; i < 12; i++ {
		_, err := api.SignPushActions(act)
		if err != nil {
			fmt.Println("获取收益失败:", err)
		} else {
			fmt.Println("获取收益成功!")
		}
		time.Sleep(sleep)
	}

	// Get balance of eosstorebest and calculate claim amount
	out, err := api.GetCurrencyBalance( eos.AccountName( bp_name ) , "EOS", "eosio.token")
	if err != nil {
		fmt.Println("failed to get balance", err)
	} else {
		fmt.Println( bp_name + " balance :", out)
	}
}
