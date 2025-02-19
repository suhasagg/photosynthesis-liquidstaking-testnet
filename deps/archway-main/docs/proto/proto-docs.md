<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [archway/photosynthesis/photosynthesis.proto](#archway/photosynthesis/photosynthesis.proto)
    - [Coin](#archway.photosynthesis.v1.Coin)
    - [Contract](#archway.photosynthesis.v1.Contract)
    - [DepositRecord](#archway.photosynthesis.v1.DepositRecord)
    - [DepositRecords](#archway.photosynthesis.v1.DepositRecords)
    - [Params](#archway.photosynthesis.v1.Params)
  
    - [Task](#archway.photosynthesis.v1.Task)
  
- [archway/photosynthesis/genesis.proto](#archway/photosynthesis/genesis.proto)
    - [GenesisState](#archway.photosynthesis.v1.GenesisState)
  
- [archway/photosynthesis/query.proto](#archway/photosynthesis/query.proto)
    - [QueryAllRedeemTokenRatiosRequest](#archway.photosynthesis.v1.QueryAllRedeemTokenRatiosRequest)
    - [QueryAllRedeemTokenRatiosResponse](#archway.photosynthesis.v1.QueryAllRedeemTokenRatiosResponse)
    - [QueryAllStakeRatiosRequest](#archway.photosynthesis.v1.QueryAllStakeRatiosRequest)
    - [QueryAllStakeRatiosResponse](#archway.photosynthesis.v1.QueryAllStakeRatiosResponse)
    - [QueryDepositRecordsRequest](#archway.photosynthesis.v1.QueryDepositRecordsRequest)
    - [QueryDepositRecordsResponse](#archway.photosynthesis.v1.QueryDepositRecordsResponse)
    - [QueryParamsRequest](#archway.photosynthesis.v1.QueryParamsRequest)
    - [QueryParamsResponse](#archway.photosynthesis.v1.QueryParamsResponse)
    - [QueryRedeemTokenRatiosRequest](#archway.photosynthesis.v1.QueryRedeemTokenRatiosRequest)
    - [QueryRedeemTokenRatiosResponse](#archway.photosynthesis.v1.QueryRedeemTokenRatiosResponse)
    - [QueryRedemptionIntervalThresholdRequest](#archway.photosynthesis.v1.QueryRedemptionIntervalThresholdRequest)
    - [QueryRedemptionIntervalThresholdResponse](#archway.photosynthesis.v1.QueryRedemptionIntervalThresholdResponse)
    - [QueryRedemptionRateThresholdRequest](#archway.photosynthesis.v1.QueryRedemptionRateThresholdRequest)
    - [QueryRedemptionRateThresholdResponse](#archway.photosynthesis.v1.QueryRedemptionRateThresholdResponse)
    - [QueryStakeRatioRequest](#archway.photosynthesis.v1.QueryStakeRatioRequest)
    - [QueryStakeRatioResponse](#archway.photosynthesis.v1.QueryStakeRatioResponse)
    - [QueryTotalLiquidStakeRequest](#archway.photosynthesis.v1.QueryTotalLiquidStakeRequest)
    - [QueryTotalLiquidStakeResponse](#archway.photosynthesis.v1.QueryTotalLiquidStakeResponse)
    - [QueryTotalRedemptionAmountRequest](#archway.photosynthesis.v1.QueryTotalRedemptionAmountRequest)
    - [QueryTotalRedemptionAmountResponse](#archway.photosynthesis.v1.QueryTotalRedemptionAmountResponse)
    - [RedeemTokenRatio](#archway.photosynthesis.v1.RedeemTokenRatio)
    - [StakeRatio](#archway.photosynthesis.v1.StakeRatio)
  
    - [Query](#archway.photosynthesis.v1.Query)
  
- [archway/photosynthesis/tx.proto](#archway/photosynthesis/tx.proto)
    - [MsgClearRedeemTokenData](#archway.photosynthesis.v1.MsgClearRedeemTokenData)
    - [MsgClearRedeemTokenDataResponse](#archway.photosynthesis.v1.MsgClearRedeemTokenDataResponse)
    - [MsgExecuteCronTasks](#archway.photosynthesis.v1.MsgExecuteCronTasks)
    - [MsgExecuteCronTasksResponse](#archway.photosynthesis.v1.MsgExecuteCronTasksResponse)
    - [MsgResetAllCompletedDepositRecords](#archway.photosynthesis.v1.MsgResetAllCompletedDepositRecords)
    - [MsgResetAllCompletedDepositRecordsResponse](#archway.photosynthesis.v1.MsgResetAllCompletedDepositRecordsResponse)
    - [MsgResetTotalLiquidStake](#archway.photosynthesis.v1.MsgResetTotalLiquidStake)
    - [MsgResetTotalLiquidStakeResponse](#archway.photosynthesis.v1.MsgResetTotalLiquidStakeResponse)
    - [MsgSetRedeemTokenData](#archway.photosynthesis.v1.MsgSetRedeemTokenData)
    - [MsgSetRedeemTokenDataResponse](#archway.photosynthesis.v1.MsgSetRedeemTokenDataResponse)
  
    - [Msg](#archway.photosynthesis.v1.Msg)
  
- [archway/rewards/v1beta1/rewards.proto](#archway/rewards/v1beta1/rewards.proto)
    - [BlockRewards](#archway.rewards.v1beta1.BlockRewards)
    - [ContractMetadata](#archway.rewards.v1beta1.ContractMetadata)
    - [FlatFee](#archway.rewards.v1beta1.FlatFee)
    - [Params](#archway.rewards.v1beta1.Params)
    - [RewardsRecord](#archway.rewards.v1beta1.RewardsRecord)
    - [TxRewards](#archway.rewards.v1beta1.TxRewards)
  
- [archway/rewards/v1beta1/events.proto](#archway/rewards/v1beta1/events.proto)
    - [ContractFlatFeeSetEvent](#archway.rewards.v1beta1.ContractFlatFeeSetEvent)
    - [ContractMetadataSetEvent](#archway.rewards.v1beta1.ContractMetadataSetEvent)
    - [ContractRewardCalculationEvent](#archway.rewards.v1beta1.ContractRewardCalculationEvent)
    - [MinConsensusFeeSetEvent](#archway.rewards.v1beta1.MinConsensusFeeSetEvent)
    - [RewardsWithdrawEvent](#archway.rewards.v1beta1.RewardsWithdrawEvent)
  
- [archway/rewards/v1beta1/genesis.proto](#archway/rewards/v1beta1/genesis.proto)
    - [GenesisState](#archway.rewards.v1beta1.GenesisState)
  
- [archway/rewards/v1beta1/query.proto](#archway/rewards/v1beta1/query.proto)
    - [BlockTracking](#archway.rewards.v1beta1.BlockTracking)
    - [QueryBlockRewardsTrackingRequest](#archway.rewards.v1beta1.QueryBlockRewardsTrackingRequest)
    - [QueryBlockRewardsTrackingResponse](#archway.rewards.v1beta1.QueryBlockRewardsTrackingResponse)
    - [QueryContractMetadataRequest](#archway.rewards.v1beta1.QueryContractMetadataRequest)
    - [QueryContractMetadataResponse](#archway.rewards.v1beta1.QueryContractMetadataResponse)
    - [QueryEstimateTxFeesRequest](#archway.rewards.v1beta1.QueryEstimateTxFeesRequest)
    - [QueryEstimateTxFeesResponse](#archway.rewards.v1beta1.QueryEstimateTxFeesResponse)
    - [QueryFlatFeeRequest](#archway.rewards.v1beta1.QueryFlatFeeRequest)
    - [QueryFlatFeeResponse](#archway.rewards.v1beta1.QueryFlatFeeResponse)
    - [QueryOutstandingRewardsRequest](#archway.rewards.v1beta1.QueryOutstandingRewardsRequest)
    - [QueryOutstandingRewardsResponse](#archway.rewards.v1beta1.QueryOutstandingRewardsResponse)
    - [QueryParamsRequest](#archway.rewards.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#archway.rewards.v1beta1.QueryParamsResponse)
    - [QueryRewardsPoolRequest](#archway.rewards.v1beta1.QueryRewardsPoolRequest)
    - [QueryRewardsPoolResponse](#archway.rewards.v1beta1.QueryRewardsPoolResponse)
    - [QueryRewardsRecordsRequest](#archway.rewards.v1beta1.QueryRewardsRecordsRequest)
    - [QueryRewardsRecordsResponse](#archway.rewards.v1beta1.QueryRewardsRecordsResponse)
  
    - [Query](#archway.rewards.v1beta1.Query)
  
- [archway/rewards/v1beta1/tx.proto](#archway/rewards/v1beta1/tx.proto)
    - [MsgSetContractMetadata](#archway.rewards.v1beta1.MsgSetContractMetadata)
    - [MsgSetContractMetadataResponse](#archway.rewards.v1beta1.MsgSetContractMetadataResponse)
    - [MsgSetFlatFee](#archway.rewards.v1beta1.MsgSetFlatFee)
    - [MsgSetFlatFeeResponse](#archway.rewards.v1beta1.MsgSetFlatFeeResponse)
    - [MsgWithdrawRewards](#archway.rewards.v1beta1.MsgWithdrawRewards)
    - [MsgWithdrawRewards.RecordIDs](#archway.rewards.v1beta1.MsgWithdrawRewards.RecordIDs)
    - [MsgWithdrawRewards.RecordsLimit](#archway.rewards.v1beta1.MsgWithdrawRewards.RecordsLimit)
    - [MsgWithdrawRewardsResponse](#archway.rewards.v1beta1.MsgWithdrawRewardsResponse)
  
    - [Msg](#archway.rewards.v1beta1.Msg)
  
- [archway/tracking/v1beta1/tracking.proto](#archway/tracking/v1beta1/tracking.proto)
    - [BlockTracking](#archway.tracking.v1beta1.BlockTracking)
    - [ContractOperationInfo](#archway.tracking.v1beta1.ContractOperationInfo)
    - [TxInfo](#archway.tracking.v1beta1.TxInfo)
    - [TxTracking](#archway.tracking.v1beta1.TxTracking)
  
    - [ContractOperation](#archway.tracking.v1beta1.ContractOperation)
  
- [archway/tracking/v1beta1/genesis.proto](#archway/tracking/v1beta1/genesis.proto)
    - [GenesisState](#archway.tracking.v1beta1.GenesisState)
  
- [archway/tracking/v1beta1/query.proto](#archway/tracking/v1beta1/query.proto)
    - [QueryBlockGasTrackingRequest](#archway.tracking.v1beta1.QueryBlockGasTrackingRequest)
    - [QueryBlockGasTrackingResponse](#archway.tracking.v1beta1.QueryBlockGasTrackingResponse)
  
    - [Query](#archway.tracking.v1beta1.Query)
  
- [Scalar Value Types](#scalar-value-types)



<a name="archway/photosynthesis/photosynthesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## archway/photosynthesis/photosynthesis.proto



<a name="archway.photosynthesis.v1.Coin"></a>

### Coin
Coin defines a token with a denomination and an amount.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  | Denomination of the token (e.g., "uarch", "uatom"). |
| `amount` | [int64](#int64) |  | Amount of the token. |






<a name="archway.photosynthesis.v1.Contract"></a>

### Contract
Contract represents a registered contract in the photosynthesis module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | Address of the contract. |
| `liquidity_provider_address` | [string](#string) |  | Address of the liquidity provider for the contract. |
| `rewards_address` | [string](#string) |  | Address to which rewards will be sent. |
| `minimum_reward_amount` | [uint64](#uint64) |  | Minimum reward amount required for distribution. |
| `liquid_stake_interval` | [uint64](#uint64) |  | Interval for liquid staking specific to this contract. |
| `redemption_rate_threshold` | [int64](#int64) |  | Redemption rate threshold specific to this contract. |
| `redemption_interval_threshold` | [uint64](#uint64) |  | Redemption interval threshold for this contract. |
| `rewards_withdrawal_interval` | [uint64](#uint64) |  | Interval for rewards withdrawal for this contract. |






<a name="archway.photosynthesis.v1.DepositRecord"></a>

### DepositRecord
DepositRecord represents a record of a deposit for staking or liquid staking.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  | Address of the contract associated with the deposit. |
| `amount` | [int64](#int64) |  | Amount deposited for staking. |
| `status` | [string](#string) |  | Status of the deposit (e.g., "pending", "completed"). |






<a name="archway.photosynthesis.v1.DepositRecords"></a>

### DepositRecords
DepositRecords is a collection of deposit records for a contract.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `records` | [DepositRecord](#archway.photosynthesis.v1.DepositRecord) | repeated | A list of deposit records. |






<a name="archway.photosynthesis.v1.Params"></a>

### Params
Params defines the parameters for the photosynthesis module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `liquid_staking_interval` | [int64](#int64) |  | Interval for liquid staking (in minutes). |
| `arch_liquid_stake_interval` | [int64](#int64) |  | Interval for Archway liquid staking (in minutes). |
| `redemption_rate_query_interval` | [int64](#int64) |  | Interval for querying the redemption rate (in minutes). |
| `redemption_rate_threshold` | [int64](#int64) |  | Redemption rate threshold for triggering redemption. |
| `redemption_interval_threshold` | [int64](#int64) |  | Time threshold for triggering redemption (in minutes). |
| `rewards_withdrawal_interval` | [int64](#int64) |  | Interval for withdrawing rewards (in minutes). |





 <!-- end messages -->


<a name="archway.photosynthesis.v1.Task"></a>

### Task
Task is an enumeration for identifying the different types of tasks in the
module.

| Name | Number | Description |
| ---- | ------ | ----------- |
| TASK_UNSPECIFIED | 0 | Unspecified task. |
| TASK_LIQUID_STAKING | 1 | Task related to liquid staking. |
| TASK_ARCH_LIQUID_STAKE | 2 | Task related to Archway liquid staking. |
| TASK_REDEMPTION_RATE_QUERY | 3 | Task related to querying the redemption rate. |
| TASK_REWARDS_WITHDRAWAL | 4 | Task related to rewards withdrawal. |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="archway/photosynthesis/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## archway/photosynthesis/genesis.proto



<a name="archway.photosynthesis.v1.GenesisState"></a>

### GenesisState
GenesisState defines the initial state of the photosynthesis module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#archway.photosynthesis.v1.Params) |  | Params for the photosynthesis module |
| `deposit_records` | [DepositRecord](#archway.photosynthesis.v1.DepositRecord) | repeated | deposit_records defines a list of all deposit records. |
| `contracts` | [Contract](#archway.photosynthesis.v1.Contract) | repeated | contracts defines a list of all registered contracts. |
| `cumulative_liquidity_amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | cumulative_liquidity_amount defines the cumulative liquidity amount for all contracts. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="archway/photosynthesis/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## archway/photosynthesis/query.proto



<a name="archway.photosynthesis.v1.QueryAllRedeemTokenRatiosRequest"></a>

### QueryAllRedeemTokenRatiosRequest
Query request for all redeem token ratios.






<a name="archway.photosynthesis.v1.QueryAllRedeemTokenRatiosResponse"></a>

### QueryAllRedeemTokenRatiosResponse
Query response containing all redeem token ratios.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `redeem_token_ratios` | [RedeemTokenRatio](#archway.photosynthesis.v1.RedeemTokenRatio) | repeated |  |






<a name="archway.photosynthesis.v1.QueryAllStakeRatiosRequest"></a>

### QueryAllStakeRatiosRequest







<a name="archway.photosynthesis.v1.QueryAllStakeRatiosResponse"></a>

### QueryAllStakeRatiosResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `stake_ratios` | [StakeRatio](#archway.photosynthesis.v1.StakeRatio) | repeated |  |






<a name="archway.photosynthesis.v1.QueryDepositRecordsRequest"></a>

### QueryDepositRecordsRequest
QueryDepositRecordsRequest is the request type for the Query/DepositRecords
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="archway.photosynthesis.v1.QueryDepositRecordsResponse"></a>

### QueryDepositRecordsResponse
QueryDepositRecordsResponse is the response type for the Query/DepositRecords
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deposit_records` | [DepositRecord](#archway.photosynthesis.v1.DepositRecord) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="archway.photosynthesis.v1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="archway.photosynthesis.v1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#archway.photosynthesis.v1.Params) |  |  |






<a name="archway.photosynthesis.v1.QueryRedeemTokenRatiosRequest"></a>

### QueryRedeemTokenRatiosRequest







<a name="archway.photosynthesis.v1.QueryRedeemTokenRatiosResponse"></a>

### QueryRedeemTokenRatiosResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `redeem_ratios` | [RedeemTokenRatio](#archway.photosynthesis.v1.RedeemTokenRatio) | repeated |  |






<a name="archway.photosynthesis.v1.QueryRedemptionIntervalThresholdRequest"></a>

### QueryRedemptionIntervalThresholdRequest
QueryRedemptionIntervalThresholdRequest is the request type for the
Query/RedemptionIntervalThreshold RPC method.






<a name="archway.photosynthesis.v1.QueryRedemptionIntervalThresholdResponse"></a>

### QueryRedemptionIntervalThresholdResponse
QueryRedemptionIntervalThresholdResponse is the response type for the
Query/RedemptionIntervalThreshold RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `redemption_interval_threshold` | [int64](#int64) |  |  |






<a name="archway.photosynthesis.v1.QueryRedemptionRateThresholdRequest"></a>

### QueryRedemptionRateThresholdRequest
QueryRedemptionRateThresholdRequest is the request type for the
Query/RedemptionRateThreshold RPC method.






<a name="archway.photosynthesis.v1.QueryRedemptionRateThresholdResponse"></a>

### QueryRedemptionRateThresholdResponse
QueryRedemptionRateThresholdResponse is the response type for the
Query/ RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `redemption_rate_threshold` | [int64](#int64) |  |  |






<a name="archway.photosynthesis.v1.QueryStakeRatioRequest"></a>

### QueryStakeRatioRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  |  |






<a name="archway.photosynthesis.v1.QueryStakeRatioResponse"></a>

### QueryStakeRatioResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  |  |
| `stake_ratio` | [string](#string) |  |  |






<a name="archway.photosynthesis.v1.QueryTotalLiquidStakeRequest"></a>

### QueryTotalLiquidStakeRequest
QueryTotalLiquidStakeRequest is the request type for the
Query/TotalLiquidStake RPC method.






<a name="archway.photosynthesis.v1.QueryTotalLiquidStakeResponse"></a>

### QueryTotalLiquidStakeResponse
QueryTotalLiquidStakeResponse is the response type for the
Query/TotalLiquidStake RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `total_liquid_stake` | [string](#string) |  |  |






<a name="archway.photosynthesis.v1.QueryTotalRedemptionAmountRequest"></a>

### QueryTotalRedemptionAmountRequest
Query request for total redemption amount.






<a name="archway.photosynthesis.v1.QueryTotalRedemptionAmountResponse"></a>

### QueryTotalRedemptionAmountResponse
Query response for total redemption amount.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `total_amount` | [string](#string) |  |  |






<a name="archway.photosynthesis.v1.RedeemTokenRatio"></a>

### RedeemTokenRatio



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  |  |
| `redeem_ratio` | [string](#string) |  |  |






<a name="archway.photosynthesis.v1.StakeRatio"></a>

### StakeRatio



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  |  |
| `stake_ratio` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="archway.photosynthesis.v1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#archway.photosynthesis.v1.QueryParamsRequest) | [QueryParamsResponse](#archway.photosynthesis.v1.QueryParamsResponse) |  | GET|/archway/photosynthesis/v1/params|
| `TotalLiquidStake` | [QueryTotalLiquidStakeRequest](#archway.photosynthesis.v1.QueryTotalLiquidStakeRequest) | [QueryTotalLiquidStakeResponse](#archway.photosynthesis.v1.QueryTotalLiquidStakeResponse) |  | GET|/archway/photosynthesis/v1/total_liquid_stake|
| `DepositRecords` | [QueryDepositRecordsRequest](#archway.photosynthesis.v1.QueryDepositRecordsRequest) | [QueryDepositRecordsResponse](#archway.photosynthesis.v1.QueryDepositRecordsResponse) |  | GET|/archway/photosynthesis/v1/deposit_records/{contract_address}|
| `RedemptionRateThreshold` | [QueryRedemptionRateThresholdRequest](#archway.photosynthesis.v1.QueryRedemptionRateThresholdRequest) | [QueryRedemptionRateThresholdResponse](#archway.photosynthesis.v1.QueryRedemptionRateThresholdResponse) |  | GET|/archway/photosynthesis/v1/redemption_rate_threshold|
| `RedemptionIntervalThreshold` | [QueryRedemptionIntervalThresholdRequest](#archway.photosynthesis.v1.QueryRedemptionIntervalThresholdRequest) | [QueryRedemptionIntervalThresholdResponse](#archway.photosynthesis.v1.QueryRedemptionIntervalThresholdResponse) |  | GET|/archway/photosynthesis/v1/redemption_interval_threshold|
| `StakeRatio` | [QueryStakeRatioRequest](#archway.photosynthesis.v1.QueryStakeRatioRequest) | [QueryStakeRatioResponse](#archway.photosynthesis.v1.QueryStakeRatioResponse) |  | |
| `AllStakeRatios` | [QueryAllStakeRatiosRequest](#archway.photosynthesis.v1.QueryAllStakeRatiosRequest) | [QueryAllStakeRatiosResponse](#archway.photosynthesis.v1.QueryAllStakeRatiosResponse) |  | |
| `RedeemTokenRatios` | [QueryRedeemTokenRatiosRequest](#archway.photosynthesis.v1.QueryRedeemTokenRatiosRequest) | [QueryRedeemTokenRatiosResponse](#archway.photosynthesis.v1.QueryRedeemTokenRatiosResponse) |  | |
| `TotalRedemptionAmount` | [QueryTotalRedemptionAmountRequest](#archway.photosynthesis.v1.QueryTotalRedemptionAmountRequest) | [QueryTotalRedemptionAmountResponse](#archway.photosynthesis.v1.QueryTotalRedemptionAmountResponse) | Query the total redemption amount. | |

 <!-- end services -->



<a name="archway/photosynthesis/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## archway/photosynthesis/tx.proto



<a name="archway.photosynthesis.v1.MsgClearRedeemTokenData"></a>

### MsgClearRedeemTokenData
MsgClearRedeemTokenData defines a message for clearing all redeem token data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  | Address of the creator/initiator of the transaction |






<a name="archway.photosynthesis.v1.MsgClearRedeemTokenDataResponse"></a>

### MsgClearRedeemTokenDataResponse
MsgClearRedeemTokenDataResponse defines the response after clearing redeem
token data.






<a name="archway.photosynthesis.v1.MsgExecuteCronTasks"></a>

### MsgExecuteCronTasks
Response type for MsgExecuteCronTasks.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  | Address of the sender. |






<a name="archway.photosynthesis.v1.MsgExecuteCronTasksResponse"></a>

### MsgExecuteCronTasksResponse
Response type for MsgExecuteCronTasks.






<a name="archway.photosynthesis.v1.MsgResetAllCompletedDepositRecords"></a>

### MsgResetAllCompletedDepositRecords
MsgResetAllCompletedDepositRecords defines the message for resetting all
completed deposit records.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |






<a name="archway.photosynthesis.v1.MsgResetAllCompletedDepositRecordsResponse"></a>

### MsgResetAllCompletedDepositRecordsResponse
Response type for MsgResetAllCompletedDepositRecords.






<a name="archway.photosynthesis.v1.MsgResetTotalLiquidStake"></a>

### MsgResetTotalLiquidStake
MsgResetTotalLiquidStake defines the message for resetting total liquid
stake.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  |  |






<a name="archway.photosynthesis.v1.MsgResetTotalLiquidStakeResponse"></a>

### MsgResetTotalLiquidStakeResponse
Response type for MsgResetTotalLiquidStake.






<a name="archway.photosynthesis.v1.MsgSetRedeemTokenData"></a>

### MsgSetRedeemTokenData
MsgSetRedeemTokenData defines a message for setting redeem token data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creator` | [string](#string) |  | Address of the creator/initiator of the transaction |
| `contract_address` | [string](#string) |  | Contract address for which the redeem token data is set |
| `amount` | [string](#string) |  | Amount of redeem tokens (stored as string to handle large numbers) |






<a name="archway.photosynthesis.v1.MsgSetRedeemTokenDataResponse"></a>

### MsgSetRedeemTokenDataResponse
MsgSetRedeemTokenDataResponse defines the response after setting redeem token
data.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="archway.photosynthesis.v1.Msg"></a>

### Msg
Service definition for Msg.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `ExecuteCronTasks` | [MsgExecuteCronTasks](#archway.photosynthesis.v1.MsgExecuteCronTasks) | [MsgExecuteCronTasksResponse](#archway.photosynthesis.v1.MsgExecuteCronTasksResponse) |  | |
| `ResetTotalLiquidStake` | [MsgResetTotalLiquidStake](#archway.photosynthesis.v1.MsgResetTotalLiquidStake) | [MsgResetTotalLiquidStakeResponse](#archway.photosynthesis.v1.MsgResetTotalLiquidStakeResponse) |  | |
| `ResetAllCompletedDepositRecords` | [MsgResetAllCompletedDepositRecords](#archway.photosynthesis.v1.MsgResetAllCompletedDepositRecords) | [MsgResetAllCompletedDepositRecordsResponse](#archway.photosynthesis.v1.MsgResetAllCompletedDepositRecordsResponse) |  | |
| `SetRedeemTokenData` | [MsgSetRedeemTokenData](#archway.photosynthesis.v1.MsgSetRedeemTokenData) | [MsgSetRedeemTokenDataResponse](#archway.photosynthesis.v1.MsgSetRedeemTokenDataResponse) | Set redeem token data for a contract. | |
| `ClearRedeemTokenData` | [MsgClearRedeemTokenData](#archway.photosynthesis.v1.MsgClearRedeemTokenData) | [MsgClearRedeemTokenDataResponse](#archway.photosynthesis.v1.MsgClearRedeemTokenDataResponse) | Clear all redeem token data. | |

 <!-- end services -->



<a name="archway/rewards/v1beta1/rewards.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## archway/rewards/v1beta1/rewards.proto



<a name="archway.rewards.v1beta1.BlockRewards"></a>

### BlockRewards
BlockRewards defines block related rewards distribution data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `height` | [int64](#int64) |  | height defines the block height. |
| `inflation_rewards` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | inflation_rewards is the rewards to be distributed. |
| `max_gas` | [uint64](#uint64) |  | max_gas defines the maximum gas for the block that is used to distribute inflation rewards (consensus parameter). |






<a name="archway.rewards.v1beta1.ContractMetadata"></a>

### ContractMetadata
ContractMetadata defines the contract rewards distribution options for a
particular contract.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  | contract_address defines the contract address (bech32 encoded). |
| `owner_address` | [string](#string) |  | owner_address is the contract owner address that can modify contract reward options (bech32 encoded). That could be the contract admin or the contract itself. If owner_address is set to contract address, contract can modify the metadata on its own using WASM bindings. |
| `rewards_address` | [string](#string) |  | rewards_address is an address to distribute rewards to (bech32 encoded). If not set (empty), rewards are not distributed for this contract. |
| `minimum_reward_amount` | [uint64](#uint64) |  |  |
| `liquidity_token_address` | [string](#string) |  |  |
| `liquid_stake_interval` | [uint64](#uint64) |  |  |
| `redemption_interval` | [uint64](#uint64) |  |  |
| `rewards_withdrawal_interval` | [uint64](#uint64) |  |  |
| `redemption_address` | [string](#string) |  |  |
| `redemption_rate_threshold` | [uint64](#uint64) |  |  |
| `redemption_interval_threshold` | [uint64](#uint64) |  |  |
| `maximum_threshold` | [uint64](#uint64) |  |  |
| `archway_reward_funds_transfer_address` | [string](#string) |  |  |
| `liquidity_provider_address` | [string](#string) |  |  |
| `liquidity_provider_commission` | [uint64](#uint64) |  |  |
| `airdrop_duration` | [uint64](#uint64) |  |  |
| `airdrop_recipient_address` | [string](#string) |  |  |
| `airdrop_vesting_period` | [uint64](#uint64) |  |  |






<a name="archway.rewards.v1beta1.FlatFee"></a>

### FlatFee
FlatFee defines the flat fee for a particular contract.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  | contract_address defines the contract address (bech32 encoded). |
| `flat_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | flat_fee defines the minimum flat fee set by the contract_owner |






<a name="archway.rewards.v1beta1.Params"></a>

### Params
Params defines the module parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `inflation_rewards_ratio` | [string](#string) |  | inflation_rewards_ratio defines the percentage of minted inflation tokens that are used for dApp rewards [0.0, 1.0]. If set to 0.0, no inflation rewards are distributed. |
| `tx_fee_rebate_ratio` | [string](#string) |  | tx_fee_rebate_ratio defines the percentage of tx fees that are used for dApp rewards [0.0, 1.0]. If set to 0.0, no fee rewards are distributed. |
| `max_withdraw_records` | [uint64](#uint64) |  | max_withdraw_records defines the maximum number of RewardsRecord objects used for the withdrawal operation. |






<a name="archway.rewards.v1beta1.RewardsRecord"></a>

### RewardsRecord
RewardsRecord defines a record that is used to distribute rewards later (lazy
distribution). This record is being created by the x/rewards EndBlocker and
pruned after the rewards are distributed. An actual rewards x/bank transfer
might be triggered by a Tx (via CLI for example) or by a contract via WASM
bindings. For a contract to trigger rewards transfer, contract address must
be set as the rewards_address in a corresponding ContractMetadata.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  | id is the unique ID of the record. |
| `rewards_address` | [string](#string) |  | rewards_address is the address to distribute rewards to (bech32 encoded). |
| `rewards` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | rewards are the rewards to be transferred later. |
| `calculated_height` | [int64](#int64) |  | calculated_height defines the block height of rewards calculation event. |
| `calculated_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | calculated_time defines the block time of rewards calculation event. |






<a name="archway.rewards.v1beta1.TxRewards"></a>

### TxRewards
TxRewards defines transaction related rewards distribution data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_id` | [uint64](#uint64) |  | tx_id is the tracking transaction ID (x/tracking is the data source for this value). |
| `height` | [int64](#int64) |  | height defines the block height. |
| `fee_rewards` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | fee_rewards is the rewards to be distributed. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="archway/rewards/v1beta1/events.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## archway/rewards/v1beta1/events.proto



<a name="archway.rewards.v1beta1.ContractFlatFeeSetEvent"></a>

### ContractFlatFeeSetEvent
ContractFlatFeeSetEvent is emitted when the contract flat fee is updated


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  | contract_address defines the bech32 address of the contract for which the flat fee is set |
| `flat_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | flat_fee defines the amount that has been set as the minimum fee for the contract |






<a name="archway.rewards.v1beta1.ContractMetadataSetEvent"></a>

### ContractMetadataSetEvent
ContractMetadataSetEvent is emitted when the contract metadata is created or
updated.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  | contract_address defines the contract address. |
| `metadata` | [ContractMetadata](#archway.rewards.v1beta1.ContractMetadata) |  | metadata defines the new contract metadata state. |






<a name="archway.rewards.v1beta1.ContractRewardCalculationEvent"></a>

### ContractRewardCalculationEvent
ContractRewardCalculationEvent is emitted when the contract reward is
calculated.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  | contract_address defines the contract address. |
| `gas_consumed` | [uint64](#uint64) |  | gas_consumed defines the total gas consumption by all WASM operations within one transaction. |
| `inflation_rewards` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | inflation_rewards defines the inflation rewards portions of the rewards. |
| `fee_rebate_rewards` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | fee_rebate_rewards defines the fee rebate rewards portions of the rewards. |
| `metadata` | [ContractMetadata](#archway.rewards.v1beta1.ContractMetadata) |  | metadata defines the contract metadata (if set). |






<a name="archway.rewards.v1beta1.MinConsensusFeeSetEvent"></a>

### MinConsensusFeeSetEvent
MinConsensusFeeSetEvent is emitted when the minimum consensus fee is updated.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `fee` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) |  | fee defines the updated minimum gas unit price. |






<a name="archway.rewards.v1beta1.RewardsWithdrawEvent"></a>

### RewardsWithdrawEvent
RewardsWithdrawEvent is emitted when credited rewards for a specific
rewards_address are distributed. Event could be triggered by a transaction
(via CLI for example) or by a contract via WASM bindings.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `reward_address` | [string](#string) |  | rewards_address defines the rewards address rewards are distributed to. |
| `rewards` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | rewards defines the total rewards being distributed. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="archway/rewards/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## archway/rewards/v1beta1/genesis.proto



<a name="archway.rewards.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the initial state of the tracking module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#archway.rewards.v1beta1.Params) |  | params defines all the module parameters. |
| `contracts_metadata` | [ContractMetadata](#archway.rewards.v1beta1.ContractMetadata) | repeated | contracts_metadata defines a list of all contracts metadata. |
| `block_rewards` | [BlockRewards](#archway.rewards.v1beta1.BlockRewards) | repeated | block_rewards defines a list of all block rewards objects. |
| `tx_rewards` | [TxRewards](#archway.rewards.v1beta1.TxRewards) | repeated | tx_rewards defines a list of all tx rewards objects. |
| `min_consensus_fee` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) |  | min_consensus_fee defines the minimum gas unit price. |
| `rewards_record_last_id` | [uint64](#uint64) |  | rewards_record_last_id defines the last unique ID for a RewardsRecord objs. |
| `rewards_records` | [RewardsRecord](#archway.rewards.v1beta1.RewardsRecord) | repeated | rewards_records defines a list of all active (undistributed) rewards records. |
| `flat_fees` | [FlatFee](#archway.rewards.v1beta1.FlatFee) | repeated | flat_fees defines a list of contract flat fee. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="archway/rewards/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## archway/rewards/v1beta1/query.proto



<a name="archway.rewards.v1beta1.BlockTracking"></a>

### BlockTracking
BlockTracking is the tracking information for a block.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `inflation_rewards` | [BlockRewards](#archway.rewards.v1beta1.BlockRewards) |  | inflation_rewards defines the inflation rewards for the block. |
| `tx_rewards` | [TxRewards](#archway.rewards.v1beta1.TxRewards) | repeated | tx_rewards defines the transaction rewards for the block. |






<a name="archway.rewards.v1beta1.QueryBlockRewardsTrackingRequest"></a>

### QueryBlockRewardsTrackingRequest
QueryBlockRewardsTrackingRequest is the request for
Query.BlockRewardsTracking.






<a name="archway.rewards.v1beta1.QueryBlockRewardsTrackingResponse"></a>

### QueryBlockRewardsTrackingResponse
QueryBlockRewardsTrackingResponse is the response for
Query.BlockRewardsTracking.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `block` | [BlockTracking](#archway.rewards.v1beta1.BlockTracking) |  |  |






<a name="archway.rewards.v1beta1.QueryContractMetadataRequest"></a>

### QueryContractMetadataRequest
QueryContractMetadataRequest is the request for Query.ContractMetadata.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  | contract_address is the contract address (bech32 encoded). |






<a name="archway.rewards.v1beta1.QueryContractMetadataResponse"></a>

### QueryContractMetadataResponse
QueryContractMetadataResponse is the response for Query.ContractMetadata.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `metadata` | [ContractMetadata](#archway.rewards.v1beta1.ContractMetadata) |  |  |






<a name="archway.rewards.v1beta1.QueryEstimateTxFeesRequest"></a>

### QueryEstimateTxFeesRequest
QueryEstimateTxFeesRequest is the request for Query.EstimateTxFees.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `gas_limit` | [uint64](#uint64) |  | gas_limit is the transaction gas limit. |
| `contract_address` | [string](#string) |  | contract_address whose flat fee is considered when estimating tx fees. |






<a name="archway.rewards.v1beta1.QueryEstimateTxFeesResponse"></a>

### QueryEstimateTxFeesResponse
QueryEstimateTxFeesResponse is the response for Query.EstimateTxFees.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `gas_unit_price` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) |  | gas_unit_price defines the minimum transaction fee per gas unit. |
| `estimated_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | estimated_fee is the estimated transaction fee for a given gas limit. |






<a name="archway.rewards.v1beta1.QueryFlatFeeRequest"></a>

### QueryFlatFeeRequest
QueryFlatFeeRequest is the request for Query.FlatFeet


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  | contract_address is the contract address (bech32 encoded). |






<a name="archway.rewards.v1beta1.QueryFlatFeeResponse"></a>

### QueryFlatFeeResponse
QueryFlatFeeResponse is the response for Query.FlatFee


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `flat_fee_amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | flat_fee_amount defines the minimum flat fee set by the contract_owner per contract execution. |






<a name="archway.rewards.v1beta1.QueryOutstandingRewardsRequest"></a>

### QueryOutstandingRewardsRequest
QueryOutstandingRewardsRequest is the request for Query.OutstandingRewards.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `rewards_address` | [string](#string) |  | rewards_address is the target address to query calculated rewards for (bech32 encoded). |






<a name="archway.rewards.v1beta1.QueryOutstandingRewardsResponse"></a>

### QueryOutstandingRewardsResponse
QueryOutstandingRewardsResponse is the response for Query.OutstandingRewards.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `total_rewards` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | total_rewards is the total rewards credited to the rewards_address. |
| `records_num` | [uint64](#uint64) |  | records_num is the total number of RewardsRecord objects stored for the rewards_address. |






<a name="archway.rewards.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request for Query.Params.






<a name="archway.rewards.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response for Query.Params.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#archway.rewards.v1beta1.Params) |  |  |






<a name="archway.rewards.v1beta1.QueryRewardsPoolRequest"></a>

### QueryRewardsPoolRequest
QueryRewardsPoolRequest is the request for Query.RewardsPool.






<a name="archway.rewards.v1beta1.QueryRewardsPoolResponse"></a>

### QueryRewardsPoolResponse
QueryRewardsPoolResponse is the response for Query.RewardsPool.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `undistributed_funds` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | undistributed_funds are undistributed yet tokens (ready for withdrawal). |
| `treasury_funds` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | treasury_funds are treasury tokens available (no mechanism is available to withdraw ATM). Treasury tokens are collected on a block basis. Those tokens are unused block rewards. |






<a name="archway.rewards.v1beta1.QueryRewardsRecordsRequest"></a>

### QueryRewardsRecordsRequest
QueryRewardsRecordsRequest is the request for Query.RewardsRecords.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `rewards_address` | [string](#string) |  | rewards_address is the target address to query records for (bech32 encoded). |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination is an optional pagination options for the request. |






<a name="archway.rewards.v1beta1.QueryRewardsRecordsResponse"></a>

### QueryRewardsRecordsResponse
QueryRewardsRecordsResponse is the response for Query.RewardsRecords.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `records` | [RewardsRecord](#archway.rewards.v1beta1.RewardsRecord) | repeated | records is the list of rewards records. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination is the pagination details in the response. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="archway.rewards.v1beta1.Query"></a>

### Query
Query service for the tracking module.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#archway.rewards.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#archway.rewards.v1beta1.QueryParamsResponse) | Params returns module parameters. | GET|/archway/rewards/v1/params|
| `ContractMetadata` | [QueryContractMetadataRequest](#archway.rewards.v1beta1.QueryContractMetadataRequest) | [QueryContractMetadataResponse](#archway.rewards.v1beta1.QueryContractMetadataResponse) | ContractMetadata returns the contract rewards parameters (metadata). | GET|/archway/rewards/v1/contract_metadata|
| `BlockRewardsTracking` | [QueryBlockRewardsTrackingRequest](#archway.rewards.v1beta1.QueryBlockRewardsTrackingRequest) | [QueryBlockRewardsTrackingResponse](#archway.rewards.v1beta1.QueryBlockRewardsTrackingResponse) | BlockRewardsTracking returns block rewards tracking for the current block. | GET|/archway/rewards/v1/block_rewards_tracking|
| `RewardsPool` | [QueryRewardsPoolRequest](#archway.rewards.v1beta1.QueryRewardsPoolRequest) | [QueryRewardsPoolResponse](#archway.rewards.v1beta1.QueryRewardsPoolResponse) | RewardsPool returns the current undistributed rewards pool funds. | GET|/archway/rewards/v1/rewards_pool|
| `EstimateTxFees` | [QueryEstimateTxFeesRequest](#archway.rewards.v1beta1.QueryEstimateTxFeesRequest) | [QueryEstimateTxFeesResponse](#archway.rewards.v1beta1.QueryEstimateTxFeesResponse) | EstimateTxFees returns the estimated transaction fees for the given transaction gas limit using the minimum consensus fee value for the current block. | GET|/archway/rewards/v1/estimate_tx_fees|
| `RewardsRecords` | [QueryRewardsRecordsRequest](#archway.rewards.v1beta1.QueryRewardsRecordsRequest) | [QueryRewardsRecordsResponse](#archway.rewards.v1beta1.QueryRewardsRecordsResponse) | RewardsRecords returns the paginated list of RewardsRecord objects stored for the provided rewards_address. | GET|/archway/rewards/v1/rewards_records|
| `OutstandingRewards` | [QueryOutstandingRewardsRequest](#archway.rewards.v1beta1.QueryOutstandingRewardsRequest) | [QueryOutstandingRewardsResponse](#archway.rewards.v1beta1.QueryOutstandingRewardsResponse) | OutstandingRewards returns total rewards credited from different contracts for the provided rewards_address. | GET|/archway/rewards/v1/outstanding_rewards|
| `FlatFee` | [QueryFlatFeeRequest](#archway.rewards.v1beta1.QueryFlatFeeRequest) | [QueryFlatFeeResponse](#archway.rewards.v1beta1.QueryFlatFeeResponse) | FlatFee returns the flat fee set by the contract owner for the provided contract_address | GET|/archway/rewards/v1/flat_fee|

 <!-- end services -->



<a name="archway/rewards/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## archway/rewards/v1beta1/tx.proto



<a name="archway.rewards.v1beta1.MsgSetContractMetadata"></a>

### MsgSetContractMetadata
MsgSetContractMetadata is the request for Msg.SetContractMetadata.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender_address` | [string](#string) |  | sender_address is the msg sender address (bech32 encoded). |
| `metadata` | [ContractMetadata](#archway.rewards.v1beta1.ContractMetadata) |  | metadata is the contract metadata to set / update. If metadata exists, non-empty fields will be updated. |






<a name="archway.rewards.v1beta1.MsgSetContractMetadataResponse"></a>

### MsgSetContractMetadataResponse
MsgSetContractMetadataResponse is the response for Msg.SetContractMetadata.






<a name="archway.rewards.v1beta1.MsgSetFlatFee"></a>

### MsgSetFlatFee
MsgSetFlatFee is the request for Msg.SetFlatFee.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender_address` | [string](#string) |  | sender_address is the msg sender address (bech32 encoded). |
| `contract_address` | [string](#string) |  | contract_address is the contract address (bech32 encoded). |
| `flat_fee_amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | flat_fee_amount defines the minimum flat fee set by the contract_owner |






<a name="archway.rewards.v1beta1.MsgSetFlatFeeResponse"></a>

### MsgSetFlatFeeResponse
MsgSetFlatFeeResponse is the response for Msg.SetFlatFee.






<a name="archway.rewards.v1beta1.MsgWithdrawRewards"></a>

### MsgWithdrawRewards
MsgWithdrawRewards is the request for Msg.WithdrawRewards.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `rewards_address` | [string](#string) |  | rewards_address is the address to distribute rewards to (bech32 encoded). |
| `records_limit` | [MsgWithdrawRewards.RecordsLimit](#archway.rewards.v1beta1.MsgWithdrawRewards.RecordsLimit) |  | records_limit defines the maximum number of RewardsRecord objects to process. If provided limit is 0, the default limit is used. |
| `record_ids` | [MsgWithdrawRewards.RecordIDs](#archway.rewards.v1beta1.MsgWithdrawRewards.RecordIDs) |  | record_ids defines specific RewardsRecord object IDs to process. |






<a name="archway.rewards.v1beta1.MsgWithdrawRewards.RecordIDs"></a>

### MsgWithdrawRewards.RecordIDs



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `ids` | [uint64](#uint64) | repeated |  |






<a name="archway.rewards.v1beta1.MsgWithdrawRewards.RecordsLimit"></a>

### MsgWithdrawRewards.RecordsLimit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `limit` | [uint64](#uint64) |  |  |






<a name="archway.rewards.v1beta1.MsgWithdrawRewardsResponse"></a>

### MsgWithdrawRewardsResponse
MsgWithdrawRewardsResponse is the response for Msg.WithdrawRewards.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `records_num` | [uint64](#uint64) |  | records_num is the number of RewardsRecord objects processed. |
| `total_rewards` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | rewards are the total rewards transferred. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="archway.rewards.v1beta1.Msg"></a>

### Msg
Msg defines the module messaging service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `SetContractMetadata` | [MsgSetContractMetadata](#archway.rewards.v1beta1.MsgSetContractMetadata) | [MsgSetContractMetadataResponse](#archway.rewards.v1beta1.MsgSetContractMetadataResponse) | SetContractMetadata creates or updates an existing contract metadata. Method is authorized to the contract owner (admin if no metadata exists). | |
| `WithdrawRewards` | [MsgWithdrawRewards](#archway.rewards.v1beta1.MsgWithdrawRewards) | [MsgWithdrawRewardsResponse](#archway.rewards.v1beta1.MsgWithdrawRewardsResponse) | WithdrawRewards performs collected rewards distribution. Rewards might be credited from multiple contracts (rewards_address must be set in the corresponding contract metadata). | |
| `SetFlatFee` | [MsgSetFlatFee](#archway.rewards.v1beta1.MsgSetFlatFee) | [MsgSetFlatFeeResponse](#archway.rewards.v1beta1.MsgSetFlatFeeResponse) | SetFlatFee sets or updates or removes the flat fee to interact with the contract Method is authorized to the contract owner. | |

 <!-- end services -->



<a name="archway/tracking/v1beta1/tracking.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## archway/tracking/v1beta1/tracking.proto



<a name="archway.tracking.v1beta1.BlockTracking"></a>

### BlockTracking
BlockTracking is the tracking information for a block.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `txs` | [TxTracking](#archway.tracking.v1beta1.TxTracking) | repeated | txs defines the list of transactions tracked in the block. |






<a name="archway.tracking.v1beta1.ContractOperationInfo"></a>

### ContractOperationInfo
ContractOperationInfo keeps a single contract operation gas consumption data.
Object is being created by the IngestGasRecord call from the wasmd.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  | id defines the unique operation ID. |
| `tx_id` | [uint64](#uint64) |  | tx_id defines a transaction ID operation relates to (TxInfo.id). |
| `contract_address` | [string](#string) |  | contract_address defines the contract address operation relates to. |
| `operation_type` | [ContractOperation](#archway.tracking.v1beta1.ContractOperation) |  | operation_type defines the gas consumption type. |
| `vm_gas` | [uint64](#uint64) |  | vm_gas is the gas consumption reported by the WASM VM. Value is adjusted by this module (CalculateUpdatedGas func). |
| `sdk_gas` | [uint64](#uint64) |  | sdk_gas is the gas consumption reported by the SDK gas meter and the WASM GasRegister (cost of Execute/Query/etc). Value is adjusted by this module (CalculateUpdatedGas func). |






<a name="archway.tracking.v1beta1.TxInfo"></a>

### TxInfo
TxInfo keeps a transaction gas tracking data.
Object is being created at the module EndBlocker.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  | id defines the unique transaction ID. |
| `height` | [int64](#int64) |  | height defines the block height of the transaction. |
| `total_gas` | [uint64](#uint64) |  | total_gas defines total gas consumption by the transaction. It is the sum of gas consumed by all contract operations (VM + SDK gas). |






<a name="archway.tracking.v1beta1.TxTracking"></a>

### TxTracking
TxTracking is the tracking information for a single transaction.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `info` | [TxInfo](#archway.tracking.v1beta1.TxInfo) |  | info defines the transaction details. |
| `contract_operations` | [ContractOperationInfo](#archway.tracking.v1beta1.ContractOperationInfo) | repeated | contract_operations defines the list of contract operations consumed by the transaction. |





 <!-- end messages -->


<a name="archway.tracking.v1beta1.ContractOperation"></a>

### ContractOperation
ContractOperation denotes which operation consumed gas.

| Name | Number | Description |
| ---- | ------ | ----------- |
| CONTRACT_OPERATION_UNSPECIFIED | 0 | Invalid or unknown operation |
| CONTRACT_OPERATION_INSTANTIATION | 1 | Instantiate operation |
| CONTRACT_OPERATION_EXECUTION | 2 | Execute operation |
| CONTRACT_OPERATION_QUERY | 3 | Query |
| CONTRACT_OPERATION_MIGRATE | 4 | Migrate operation |
| CONTRACT_OPERATION_IBC | 5 | IBC operations |
| CONTRACT_OPERATION_SUDO | 6 | Sudo operation |
| CONTRACT_OPERATION_REPLY | 7 | Reply callback operation |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="archway/tracking/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## archway/tracking/v1beta1/genesis.proto



<a name="archway.tracking.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the initial state of the tracking module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_info_last_id` | [uint64](#uint64) |  | tx_info_last_id defines the last unique ID for a TxInfo objs. |
| `tx_infos` | [TxInfo](#archway.tracking.v1beta1.TxInfo) | repeated | tx_infos defines a list of all the tracked transactions. |
| `contract_op_info_last_id` | [uint64](#uint64) |  | contract_op_info_last_id defines the last unique ID for ContractOperationInfo objs. |
| `contract_op_infos` | [ContractOperationInfo](#archway.tracking.v1beta1.ContractOperationInfo) | repeated | contract_op_infos defines a list of all the tracked contract operations. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="archway/tracking/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## archway/tracking/v1beta1/query.proto



<a name="archway.tracking.v1beta1.QueryBlockGasTrackingRequest"></a>

### QueryBlockGasTrackingRequest
QueryBlockGasTrackingRequest is the request for Query.BlockGasTracking.






<a name="archway.tracking.v1beta1.QueryBlockGasTrackingResponse"></a>

### QueryBlockGasTrackingResponse
QueryBlockGasTrackingResponse is the response for Query.BlockGasTracking.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `block` | [BlockTracking](#archway.tracking.v1beta1.BlockTracking) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="archway.tracking.v1beta1.Query"></a>

### Query
Query service for the tracking module.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `BlockGasTracking` | [QueryBlockGasTrackingRequest](#archway.tracking.v1beta1.QueryBlockGasTrackingRequest) | [QueryBlockGasTrackingResponse](#archway.tracking.v1beta1.QueryBlockGasTrackingResponse) | BlockGasTracking returns block gas tracking for the current block | GET|/archway/tracking/v1/block_gas_tracking|

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

