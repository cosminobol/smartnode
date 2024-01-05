package state

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	v110rc1_rewards "github.com/rocket-pool/rocketpool-go/legacy/v1.1.0-rc1/rewards"
	"github.com/rocket-pool/rocketpool-go/rewards"
	"github.com/rocket-pool/rocketpool-go/rocketpool"
	"github.com/rocket-pool/smartnode/shared/services/config"
	cfgtypes "github.com/rocket-pool/smartnode/shared/types/config"
)

// TODO: temp until rocketpool-go supports RocketStorage contract address lookups per block
func GetClaimIntervalTime(cfg *config.RocketPoolConfig, index uint64, rp *rocketpool.RocketPool, opts *bind.CallOpts) (time.Duration, error) {
	switch cfg.Smartnode.Network.Value.(cfgtypes.Network) {
	case cfgtypes.Network_Prater:
		if index < 2 {
			contractAddress := cfg.Smartnode.GetPreviousRewardsPoolAddresses()[0]
			return v110rc1_rewards.GetClaimIntervalTime(rp, opts, &contractAddress)
		}
	}

	return rewards.GetClaimIntervalTime(rp, opts)
}

// TODO: temp until rocketpool-go supports RocketStorage contract address lookups per block
func GetNodeOperatorRewardsPercent(cfg *config.RocketPoolConfig, index uint64, rp *rocketpool.RocketPool, opts *bind.CallOpts) (*big.Int, error) {
	switch cfg.Smartnode.Network.Value.(cfgtypes.Network) {
	case cfgtypes.Network_Prater:
		if index < 2 {
			contractAddress := cfg.Smartnode.GetPreviousRewardsPoolAddresses()[0]
			return v110rc1_rewards.GetNodeOperatorRewardsPercent(rp, opts, &contractAddress)
		}
	}

	return rewards.GetNodeOperatorRewardsPercent(rp, opts)
}

// TODO: temp until rocketpool-go supports RocketStorage contract address lookups per block
func GetTrustedNodeOperatorRewardsPercent(cfg *config.RocketPoolConfig, index uint64, rp *rocketpool.RocketPool, opts *bind.CallOpts) (*big.Int, error) {
	switch cfg.Smartnode.Network.Value.(cfgtypes.Network) {
	case cfgtypes.Network_Prater:
		if index < 2 {
			contractAddress := cfg.Smartnode.GetPreviousRewardsPoolAddresses()[0]
			return v110rc1_rewards.GetTrustedNodeOperatorRewardsPercent(rp, opts, &contractAddress)
		}
	}

	return rewards.GetTrustedNodeOperatorRewardsPercent(rp, opts)
}

// TODO: temp until rocketpool-go supports RocketStorage contract address lookups per block
func GetProtocolDaoRewardsPercent(cfg *config.RocketPoolConfig, index uint64, rp *rocketpool.RocketPool, opts *bind.CallOpts) (*big.Int, error) {
	switch cfg.Smartnode.Network.Value.(cfgtypes.Network) {
	case cfgtypes.Network_Prater:
		if index < 2 {
			contractAddress := cfg.Smartnode.GetPreviousRewardsPoolAddresses()[0]
			return v110rc1_rewards.GetProtocolDaoRewardsPercent(rp, opts, &contractAddress)
		}
	}

	return rewards.GetProtocolDaoRewardsPercent(rp, opts)
}

// TODO: temp until rocketpool-go supports RocketStorage contract address lookups per block
func GetPendingRPLRewards(cfg *config.RocketPoolConfig, index uint64, rp *rocketpool.RocketPool, opts *bind.CallOpts) (*big.Int, error) {
	switch cfg.Smartnode.Network.Value.(cfgtypes.Network) {
	case cfgtypes.Network_Prater:
		if index < 2 {
			contractAddress := cfg.Smartnode.GetPreviousRewardsPoolAddresses()[0]
			return v110rc1_rewards.GetPendingRPLRewards(rp, opts, &contractAddress)
		}
	}

	return rewards.GetPendingRPLRewards(rp, opts)
}

// Returns the index of the Most Significant Bit of n, or UINT_MAX if the input is 0
// The index of the Least Significant Bit is 0.
func indexOfMSB(n *big.Int) uint {
	copyN := big.NewInt(0).Set(n)
	var out uint
	for copyN.Cmp(big.NewInt(0)) > 0 {
		copyN.Rsh(copyN, 1)
		out++
	}

	// 0-index
	return out - 1
}

func log2(x *big.Int) *big.Int {
	out := big.NewInt(0)

	// Calculate the integer part of the logarithm
	copyX := big.NewInt(0).Set(x)
	copyX.Quo(x, oneEth)
	// The input is always over 2 Eth, so we do not need to worry about
	// overflowing indexOfMSB
	n := indexOfMSB(copyX)

	// Add integer part of the logarithm
	out.Mul(oneEth, big.NewInt(int64(n)))

	// Calculate y = x * 2**-n
	y := big.NewInt(0).Rsh(big.NewInt(0).Set(x), n)

	// If y is the unit number, the fractional part is zero.
	if y.Cmp(oneEth) == 0 {
		return out
	}

	doubleUnit := big.NewInt(0).Mul(big.NewInt(2), oneEth)
	delta := big.NewInt(0).Rsh(oneEth, 1)
	for i := 0; i < 60; i++ {
		y.Mul(y, y)
		y.Quo(y, oneEth)

		if y.Cmp(doubleUnit) >= 0 {
			out.Add(out, delta)
			y.Rsh(y, 1)
		}

		delta.Rsh(delta, 1)
	}

	return out
}

func ethNaturalLog(x *big.Int) *big.Int {
	log2e := big.NewInt(1_442695040888963407)
	log2x := log2(x)

	numerator := big.NewInt(0).Mul(oneEth, log2x)
	return numerator.Quo(numerator, log2e)
}
