package main

import (
	"fmt"
	"github.com/miguelmota/go-ethereum-hdwallet"
	"log"
)

func main() {
	// 實作 BIP-0039 產生一組隨機的助記詞(BIP-0039 mnemonic code)
	mnemonic, _ := hdwallet.NewMnemonic(256)
	fmt.Println("mnemonic code: " + mnemonic)
	// 利用 mnemonic code 產生一組錢包物件
	// 從原始碼可以得知 NewFromMnemonic 其實實作了
	// 1. 驗證 mnemonic code
	// 2. 用 mnemonic code 產生 seed
	// 3. 用 seed 產出 masterKey，並生成空錢包物件，根據 BIP-0039 我們只需要記得 mnemonic code 就可以了
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}
	// 產生第一組錢包
	// 解析 BIP-0044 格式，轉不出來會噴錯, 最後一碼為 address index
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	// 在 BIP-0044 的路徑下產生一個新帳戶，如果 pin 設置為 true，則將該帳戶添加到跟蹤帳戶列表中 (wallet)
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}
	// 這個是這個第一組錢包的私鑰，但其實不用記，只是印出來說明一下
	// 因為從 mnemonic code 產生的 seed 然後再產生的 masterKey 和 account 的 BIP-0044 資訊就可以推導出來
	// 所以要好好記得且不能遺失你的 mnemonic code
	pkh, _ := wallet.PrivateKeyHex(account)
	fmt.Println("First Wallet Private Key: " + pkh)
	fmt.Println("First Wallet: " + account.Address.Hex()) // 0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947
	// 產生第二組錢包
	// 解析 BIP-0044 格式，轉不出來會噴錯, 最後一碼改成 1
	path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1")
	// 在 BIP-0044 的路徑下產生一個新帳戶，如果 pin 設置為 true，則將該帳戶添加到跟蹤帳戶列表中 (wallet)
	account, err = wallet.Derive(path, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Second Wallet: " + account.Address.Hex()) // 0x8230645aC28A4EdD1b0B53E7Cd8019744E9dD559
	for key, item := range wallet.Accounts() {
		// 印出有 pin 的地址
		fmt.Println("Pin Address " + fmt.Sprintf("%v: %s", key+1, item.Address))
	}
}
