package bridge

import (
	"bytes"
	"encoding/binary"
	"math/big"
	"strconv"
	"strings"

	"github.com/ChainSafe/chainbridge-core/chains/evm/calls"
	"github.com/ChainSafe/chainbridge-core/chains/evm/calls/consts"
	"github.com/ChainSafe/chainbridge-core/chains/evm/calls/contracts"
	"github.com/ChainSafe/chainbridge-core/chains/evm/calls/contracts/deposit"
	"github.com/ChainSafe/chainbridge-core/chains/evm/calls/transactor"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ChainSafe/chainbridge-core/chains/evm/executor/proposal"
	"github.com/ChainSafe/chainbridge-core/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
)

type BridgeContract struct {
	contracts.Contract
}

func NewBridgeContract(
	client calls.ContractCallerDispatcher,
	bridgeContractAddress common.Address,
	transactor transactor.Transactor,
) *BridgeContract {
	a, _ := abi.JSON(strings.NewReader(consts.BridgeABI))
	b := common.FromHex(consts.BridgeBin)
	return &BridgeContract{contracts.NewContract(bridgeContractAddress, a, b, client, transactor)}
}

func (c *BridgeContract) AdminSetGenericResource(
	handler common.Address,
	rID types.ResourceID,
	addr common.Address,
	depositFunctionSig [4]byte,
	depositerOffset *big.Int,
	executeFunctionSig [4]byte,
	opts transactor.TransactOptions,
) (*common.Hash, error) {
	log.Debug().Msgf("Setting generic resource %s", hexutil.Encode(rID[:]))
	return c.ExecuteTransaction(
		"adminSetGenericResource",
		opts,
		handler, rID, addr, depositFunctionSig, depositerOffset, executeFunctionSig,
	)
}

func (c *BridgeContract) AdminSetResource(
	handlerAddr common.Address,
	rID types.ResourceID,
	targetContractAddr common.Address,
	opts transactor.TransactOptions,
) (*common.Hash, error) {
	log.Debug().Msgf("Setting resource %s", hexutil.Encode(rID[:]))
	return c.ExecuteTransaction(
		"adminSetResource",
		opts,
		handlerAddr, rID, targetContractAddr,
	)
}

func (c *BridgeContract) SetDepositNonce(
	domainId uint8,
	depositNonce uint64,
	opts transactor.TransactOptions,
) (*common.Hash, error) {
	log.Debug().Msgf("Setting deposit nonce %d for %d", depositNonce, domainId)
	return c.ExecuteTransaction(
		"adminSetDepositNonce",
		opts,
		domainId, depositNonce,
	)
}

func (c *BridgeContract) SetBurnableInput(
	handlerAddr common.Address,
	tokenContractAddr common.Address,
	opts transactor.TransactOptions,
) (*common.Hash, error) {
	log.Debug().Msgf("Setting burnable input for %s", tokenContractAddr.String())
	return c.ExecuteTransaction(
		"adminSetBurnable",
		opts,
		handlerAddr, tokenContractAddr,
	)
}

func (c *BridgeContract) deposit(
	resourceID types.ResourceID,
	destDomainID uint8,
	data []byte,
	feeData []byte,
	opts transactor.TransactOptions,
) (*common.Hash, error) {
	return c.ExecuteTransaction(
		"deposit",
		opts,
		destDomainID, resourceID, data, feeData,
	)
}

func (c *BridgeContract) Erc20Deposit(
	recipient common.Address,
	amount *big.Int,
	resourceID types.ResourceID,
	destDomainID uint8,
	feeData []byte,
	opts transactor.TransactOptions,
) (*common.Hash, error) {
	log.Debug().
		Str("recipient", recipient.String()).
		Str("resourceID", hexutil.Encode(resourceID[:])).
		Str("amount", amount.String()).
		Uint8("destDomainID", destDomainID).
		Hex("feeData", feeData).
		Msgf("ERC20 deposit")
	var data []byte
	if opts.Priority == 0 {
		data = deposit.ConstructErc20DepositData(recipient.Bytes(), amount)
	} else {
		data = deposit.ConstructErc20DepositDataWithPriority(recipient.Bytes(), amount, opts.Priority)
	}

	txHash, err := c.deposit(resourceID, destDomainID, data, feeData, opts)
	if err != nil {
		log.Error().Err(err)
		return nil, err
	}
	return txHash, err
}

func (c *BridgeContract) Erc721Deposit(
	tokenId *big.Int,
	metadata string,
	recipient common.Address,
	resourceID types.ResourceID,
	destDomainID uint8,
	feeData []byte,
	opts transactor.TransactOptions,
) (*common.Hash, error) {
	log.Debug().
		Str("recipient", recipient.String()).
		Str("resourceID", hexutil.Encode(resourceID[:])).
		Str("tokenID", tokenId.String()).
		Uint8("destDomainID", destDomainID).
		Hex("feeData", feeData).
		Msgf("ERC721 deposit")
	var data []byte
	if opts.Priority == 0 {
		data = deposit.ConstructErc721DepositData(recipient.Bytes(), tokenId, []byte(metadata))
	} else {
		data = deposit.ConstructErc721DepositDataWithPriority(recipient.Bytes(), tokenId, []byte(metadata), opts.Priority)
	}

	txHash, err := c.deposit(resourceID, destDomainID, data, feeData, opts)
	if err != nil {
		log.Error().Err(err)
		return nil, err
	}
	return txHash, err
}

func (c *BridgeContract) GenericDeposit(
	metadata []byte,
	resourceID types.ResourceID,
	destDomainID uint8,
	feeData []byte,
	opts transactor.TransactOptions,
) (*common.Hash, error) {
	log.Debug().
		Str("resourceID", hexutil.Encode(resourceID[:])).
		Uint8("destDomainID", destDomainID).
		Hex("feeData", feeData).
		Msgf("Generic deposit")
	data := deposit.ConstructGenericDepositData(metadata)

	txHash, err := c.deposit(resourceID, destDomainID, data, feeData, opts)
	if err != nil {
		log.Error().Err(err)
		return nil, err
	}
	return txHash, err
}

func (c *BridgeContract) ExecuteProposal(
	proposal *proposal.Proposal,
	signature []byte,
	revertOnFail bool,
	opts transactor.TransactOptions,
) (*common.Hash, error) {
	log.Debug().
		Str("depositNonce", strconv.FormatUint(proposal.DepositNonce, 10)).
		Str("resourceID", hexutil.Encode(proposal.ResourceId[:])).
		Str("handler", proposal.HandlerAddress.String()).
		Msgf("Execute proposal")
	return c.ExecuteTransaction(
		"executeProposal",
		opts,
		proposal.Source, proposal.DepositNonce, proposal.Data, proposal.ResourceId, signature,
	)
}

func (c *BridgeContract) ProposalHash(proposal *proposal.Proposal) ([]byte, error) {
	nonceBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(nonceBytes, proposal.DepositNonce)
	proposalBytes := bytes.Join(
		[][]byte{
			{proposal.Source},
			{proposal.Destination},
			nonceBytes,
			proposal.Data,
			proposal.ResourceId[:],
		},
		nil,
	)
	hash := crypto.Keccak256Hash(proposalBytes)
	return hash.Bytes(), nil
}

func (c *BridgeContract) Pause(opts transactor.TransactOptions) (*common.Hash, error) {
	log.Debug().Msg("Pause transfers")
	return c.ExecuteTransaction(
		"adminPauseTransfers",
		opts,
	)
}

func (c *BridgeContract) Unpause(opts transactor.TransactOptions) (*common.Hash, error) {
	log.Debug().Msg("Unpause transfers")
	return c.ExecuteTransaction(
		"adminUnpauseTransfers",
		opts,
	)
}

func (c *BridgeContract) EndKeygen(mpcAddress common.Address, opts transactor.TransactOptions) (*common.Hash, error) {
	log.Debug().Msg("Ending keygen process")
	return c.ExecuteTransaction(
		"endKeygen",
		opts,
		mpcAddress,
	)
}

func (c *BridgeContract) Withdraw(
	handlerAddress,
	tokenAddress,
	recipientAddress common.Address,
	amountOrTokenId *big.Int,
	opts transactor.TransactOptions,
) (*common.Hash, error) {
	// @dev withdrawal data should include:
	// tokenAddress
	// recipientAddress
	// realAmount
	data := bytes.Buffer{}
	data.Write(common.LeftPadBytes(tokenAddress.Bytes(), 32))
	data.Write(common.LeftPadBytes(recipientAddress.Bytes(), 32))
	data.Write(common.LeftPadBytes(amountOrTokenId.Bytes(), 32))

	return c.ExecuteTransaction("adminWithdraw", opts, handlerAddress, data.Bytes())
}

func (c *BridgeContract) IsProposalExecuted(p *proposal.Proposal) (bool, error) {
	log.Debug().
		Str("depositNonce", strconv.FormatUint(p.DepositNonce, 10)).
		Str("resourceID", hexutil.Encode(p.ResourceId[:])).
		Str("handler", p.HandlerAddress.String()).
		Msg("Getting is proposal exectued")
	res, err := c.CallContract("isProposalExecuted", p.Source, big.NewInt(int64(p.DepositNonce)))
	if err != nil {
		return false, err
	}
	out := *abi.ConvertType(res[0], new(bool)).(*bool)
	return out, nil
}

func (c *BridgeContract) GetHandlerAddressForResourceID(
	resourceID types.ResourceID,
) (common.Address, error) {
	log.Debug().Msgf("Getting handler address for resource %s", hexutil.Encode(resourceID[:]))
	res, err := c.CallContract("_resourceIDToHandlerAddress", resourceID)
	if err != nil {
		return common.Address{}, err
	}
	out := *abi.ConvertType(res[0], new(common.Address)).(*common.Address)
	return out, nil
}

func (c *BridgeContract) AdminChangeFeeHandler(
	feeHandlerAddr common.Address,
	opts transactor.TransactOptions,
) (*common.Hash, error) {
	log.Debug().Msgf("Setting fee handler %s", feeHandlerAddr.String())
	return c.ExecuteTransaction(
		"adminChangeFeeHandler",
		opts,
		feeHandlerAddr,
	)
}