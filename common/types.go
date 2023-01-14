package common

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type BlocknativePayload struct {
	From                          string
	To                            string
	Nonce                         int
	Gas                           int
	GasPrice                      string
	GasPriceGwei                  int
	BasUsed                       int
	BaseFeePerGas                 string
	BaseFeePerGasGwei             int
	MaxPriorityFeePerGas          string
	MaxPriorityFeePerGasGwei      int
	MaxFeePerGas                  string
	MaxFeePerGasGwei              int
	TxType                        int
	Value                         string
	Hash                          string `json:"hash" binding:"required"`
	Input                         string
	V                             string
	R                             string
	S                             string
	BlockHash                     string
	Blockint                      string
	EstimatedBlocksUntilConfirmed int
}

type TransactionLogs struct {
	Event            *abi.Event
	IndexedParams    []common.Address
	NonIndexedParams []byte
}

type USDCQueuedEvent struct {
	Depositor         common.Address
	Amount            *big.Int
	DepositorsBalance *big.Int
	ReceiptIndex      *big.Int
}

type USDCDeQueuedEvent struct {
	Depositor         common.Address
	Amount            *big.Int
	DepositorsBalance *big.Int
}

type CrabQueuedEvent struct {
	Withdrawer         common.Address
	Amount             *big.Int
	WithdrawersBalance *big.Int
	ReceiptIndex       *big.Int
}

type CrabDeQueuedEvent struct {
	Withdrawer         common.Address
	Amount             *big.Int
	WithdrawersBalance *big.Int
}

type USDCDepositedEvent struct {
	Depositor    common.Address
	UsdcAmount   *big.Int
	CrabAmount   *big.Int
	ReceiptIndex *big.Int
	refundedETH  *big.Int
}

type CrabWithdrawnEvent struct {
	Withdrawer   common.Address
	CrabAmount   *big.Int
	UsdcAmount   *big.Int
	ReceiptIndex *big.Int
}

// TODO: add struct to those events
// event WithdrawRejected(address indexed withdrawer, uint256 crabAmount, uint256 index);

// event BidTraded(uint256 indexed bidId, address indexed trader, uint256 quantity, uint256 price, bool isBuying);

// event SetAuctionTwapPeriod(uint32 previousTwap, uint32 newTwap);
// event SetOTCPriceTolerance(uint256 previousTolerance, uint256 newOtcPriceTolerance);
// event SetMinCrab(uint256 amount);
// event SetMinUSDC(uint256 amount);
// event SetDepositsIndex(uint256 newDepositsIndex);
// event SetWithdrawsIndex(uint256 newWithdrawsIndex);
// event NonceTrue(address sender, uint256 nonce);
// event ToggledAuctionLive(bool isAuctionLive);
