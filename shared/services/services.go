package services

import (
    "fmt"
    "sync"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/rocket-pool/rocketpool-go/rocketpool"
    "github.com/urfave/cli"

    "github.com/rocket-pool/smartnode/shared/services/accounts"
    "github.com/rocket-pool/smartnode/shared/services/beacon"
    "github.com/rocket-pool/smartnode/shared/services/beacon/lighthouse"
    "github.com/rocket-pool/smartnode/shared/services/beacon/prysm"
    "github.com/rocket-pool/smartnode/shared/services/config"
    "github.com/rocket-pool/smartnode/shared/services/passwords"
)


// Service instances & initializers
var (
    cfg config.RocketPoolConfig
    passwordManager *passwords.PasswordManager
    accountManager *accounts.AccountManager
    ethClient *ethclient.Client
    rocketPool *rocketpool.RocketPool
    beaconClient beacon.Client

    initCfg sync.Once
    initPasswordManager sync.Once
    initAccountManager sync.Once
    initEthClient sync.Once
    initRocketPool sync.Once
    initBeaconClient sync.Once
)


//
// Service providers
//


func GetConfig(c *cli.Context) (config.RocketPoolConfig, error) {
    return getConfig(c)
}


func GetPasswordManager(c *cli.Context) (*passwords.PasswordManager, error) {
    cfg, err := getConfig(c)
    if err != nil {
        return nil, err
    }
    return getPasswordManager(cfg), nil
}


func GetAccountManager(c *cli.Context) (*accounts.AccountManager, error) {
    cfg, err := getConfig(c)
    if err != nil {
        return nil, err
    }
    pm := getPasswordManager(cfg)
    return getAccountManager(cfg, pm), nil
}


func GetEthClient(c *cli.Context) (*ethclient.Client, error) {
    cfg, err := getConfig(c)
    if err != nil {
        return nil, err
    }
    return getEthClient(cfg)
}


func GetRocketPool(c *cli.Context) (*rocketpool.RocketPool, error) {
    cfg, err := getConfig(c)
    if err != nil {
        return nil, err
    }
    ec, err := getEthClient(cfg)
    if err != nil {
        return nil, err
    }
    return getRocketPool(cfg, ec)
}


func GetBeaconClient(c *cli.Context) (beacon.Client, error) {
    cfg, err := getConfig(c)
    if err != nil {
        return nil, err
    }
    return getBeaconClient(cfg)
}


//
// Service instance getters
//


func getConfig(c *cli.Context) (config.RocketPoolConfig, error) {
    var err error
    initCfg.Do(func() {
        _, cfg, err = config.Load(c)
    })
    return cfg, err
}


func getPasswordManager(cfg config.RocketPoolConfig) *passwords.PasswordManager {
    initPasswordManager.Do(func() {
        passwordManager = passwords.NewPasswordManager(cfg.Smartnode.PasswordPath)
    })
    return passwordManager
}


func getAccountManager(cfg config.RocketPoolConfig, pm *passwords.PasswordManager) *accounts.AccountManager {
    initAccountManager.Do(func() {
        accountManager = accounts.NewAccountManager(cfg.Smartnode.NodeKeychainPath, pm)
    })
    return accountManager
}


func getEthClient(cfg config.RocketPoolConfig) (*ethclient.Client, error) {
    var err error
    initEthClient.Do(func() {
        ethClient, err = ethclient.Dial(cfg.Chains.Eth1.Provider)
    })
    return ethClient, err
}


func getRocketPool(cfg config.RocketPoolConfig, client *ethclient.Client) (*rocketpool.RocketPool, error) {
    var err error
    initRocketPool.Do(func() {
        rocketPool, err = rocketpool.NewRocketPool(client, common.HexToAddress(cfg.Rocketpool.StorageAddress))
    })
    return rocketPool, err
}


func getBeaconClient(cfg config.RocketPoolConfig) (beacon.Client, error) {
    var err error
    initBeaconClient.Do(func() {
        switch cfg.Chains.Eth2.Client.Selected {
            case "lighthouse":
                beaconClient = lighthouse.NewClient(cfg.Chains.Eth2.Provider)
            case "prysm":
                beaconClient, err = prysm.NewClient(cfg.Chains.Eth2.Provider)
            default:
                err = fmt.Errorf("Unknown Eth 2.0 client '%s' selected", cfg.Chains.Eth2.Client.Selected)
        }
    })
    return beaconClient, err
}

