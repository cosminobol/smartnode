package api

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/rocket-pool/rocketpool-go/core"
	"github.com/rocket-pool/rocketpool-go/types"
)

// Encrypted validator keystore following the EIP-2335 standard
// (https://eips.ethereum.org/EIPS/eip-2335)
type ValidatorKeystore struct {
	Crypto  map[string]interface{} `json:"crypto"`
	Version uint                   `json:"version"`
	UUID    uuid.UUID              `json:"uuid"`
	Path    string                 `json:"path"`
	Pubkey  types.ValidatorPubkey  `json:"pubkey"`
}

type WalletStatusData struct {
	PasswordSet       bool           `json:"passwordSet"`
	WalletInitialized bool           `json:"walletInitialized"`
	AccountAddress    common.Address `json:"accountAddress"`
}

type WalletInitData struct {
	Mnemonic       string         `json:"mnemonic"`
	AccountAddress common.Address `json:"accountAddress"`
}

type WalletRecoverData struct {
	AccountAddress common.Address          `json:"accountAddress"`
	ValidatorKeys  []types.ValidatorPubkey `json:"validatorKeys"`
}

type WalletSearchAndRecoverData struct {
	FoundWallet    bool                    `json:"foundWallet"`
	AccountAddress common.Address          `json:"accountAddress"`
	DerivationPath string                  `json:"derivationPath"`
	Index          uint                    `json:"index"`
	ValidatorKeys  []types.ValidatorPubkey `json:"validatorKeys"`
}

type WalletRebuildData struct {
	ValidatorKeys []types.ValidatorPubkey `json:"validatorKeys"`
}

type WalletExportData struct {
	Password          string `json:"password"`
	Wallet            string `json:"wallet"`
	AccountPrivateKey string `json:"accountPrivateKey"`
}

type WalletSetEnsNameData struct {
	Address common.Address        `json:"address"`
	EnsName string                `json:"ensName"`
	TxInfo  *core.TransactionInfo `json:"txInfo"`
}

type WalletTestMnemonicData struct {
	CurrentAddress   common.Address `json:"currentAddress"`
	RecoveredAddress common.Address `json:"recoveredAddress"`
}
