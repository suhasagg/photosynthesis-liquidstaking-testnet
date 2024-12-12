Testnet run Instructions 

cd /home/keanu-xbox/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv13/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/

make start-docker build=sp

//Deploy photosynthesis liquid staking smart contract 

build/archwayd --home /home/keanu-xbox/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv13/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 tx wasm store /home/keanu-xbox/smartcontractlogs/liquidstakingv2/liquidstaking/artifacts/cosmwasm_liquid_staking.wasm --from pval5 --keyring-backend test --fees 500000uarch --gas 2000000 --chain-id localnet -y
build/archwayd tx --home /home/keanu-xbox/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv13/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 wasm instantiate 1 '{
     "liquid_staking_interval": 100,
     "arch_liquid_stake_interval": 150,
     "redemption_rate_query_interval": 100,
     "rewards_withdrawal_interval": 70,
     "redemption_interval_threshold": 1000
   }' --label "Liquid Staking Contract" --from pval5 --keyring-backend=test --chain-id localnet --fees 30000uarch --gas auto --gas-adjustment 1.5 --no-admin -y
   
//Deploy Dapp1    
./Dappsimulation.sh

//Register Dapp metadata in photosynthesis liquid staking smart contract

//Sample Values
//CONTRACT_ADDRESS_TO_SET="archway1ghd753shjuwexxywmgs4xz7x2q732vcnkm6h2pyv9s6ah3hylvrqvlzdpl"  # Address of the contract to set metadata for
//REWARDS_ADDRESS="archway1m062ly7dt42nke40hsmqsnk0xtag8rqhhruf8x"  # Rewards address associated with the contract
//LIQUIDITY_PROVIDER_ADDRESS="archway1he6g65qj44gtgnve7dzre6m55dz7ql0srzn46h"
//REDEMPTION_ADDRESS="archway1dpm8hstetl3y8agyznwaak0ruhap0dmznd79lu"  # Liquidity provider address
//MINIMUM_REWARD_AMOUNT="1000"  # Minimum reward amount in uarch 
//MAXIMUM_REWARD_AMOUNT="100000000000" # Maximum reward amount in uarch  

/home/keanu-xbox/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv13/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/deps/archway-main/contrib/localnet/opt/contract-metadata.sh

//Deploy Dapp2

./Dappsimulationadserver.sh

//Register Dapp metadata in photosynthesis liquid staking smart contract 

/home/keanu-xbox/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv13/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/deps/archway-main/contrib/localnet/opt/contract-metadata.sh

//Deploy Dapp3

./Dappsimulationcookiesync.sh 

//Register Dapp metadata in photosynthesis liquid staking smart contract 

/home/keanu-xbox/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv13/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/deps/archway-main/contrib/localnet/opt/contract-metadata.sh


//This script claims periodic rewards for different Dapps registered above from Archway chain and ingest rewards in photosynthesis liquid staking smart contract


* * * * * /home/keanu-xbox/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv13/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/deps/archway-main/contrib/localnet/opt/rewardswithdrawal.sh >> /home/keanu-xbox/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv13/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/rewardslogs.log



* * * * * /home/keanu-xbox/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv13/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/deps/archway-main/contrib/localnet/opt/add_cron.sh >> /home/keanu-xbox/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv13/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/cronlogs.log


//This script starts process of liquid staking and periodic distribution of liquidity.


*/6 * * * * /home/keanu-xbox/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv13/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/deps/archway-main/contrib/localnet/opt/liquidstake.sh >> /home/keanu-xbox/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv13/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/liquidstakelogs && /home/keanu-xbox/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv13/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/deps/archway-main/contrib/localnet/opt/DistributeLiquidity.sh >> /home/keanu-xbox/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv13/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/distributeliquiditylogs


//This task retrieves and logs the current redemption rates for the tokens.

* * * * * /home/keanu-xbox/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv13/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/strided --home /home/keanu-xbox/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv13/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/stride1 q stakeibc list-host-zone >> /home/keanu-xbox/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv13/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/redemptionrate


//Generate transactions for Dapps deployed on archway chain to generate rewards  

* * * * * /home/keanu-xbox/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv13/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/deps/archway-main/contrib/localnet/opt/transactions.sh >> /root/transactions.txt



Elasticsearch index creation 

node create_index.js


Elasticsearch Data Layer scripts 


node listener.js


node ingestlogs.js


node ingestredemptiondatatoelasticsearch.js


node ingestunbondingdatainelasticsearch.js


After running liquid staking for some time, run redeem stake to benchmark revenue (Set up explained fully in demo document)

/home/keanu-xbox/Pictures/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv13/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/deps/archway-main/contrib/localnet/opt/redeemstake.sh



![395124315-537af863-4a26-4e4e-bdc6-2bc79bf84cd7](https://github.com/user-attachments/assets/3af51e9d-a574-4023-8243-8a80d2dd00b9)



![395185977-c773ecfc-16aa-4530-bd52-0539884359fa](https://github.com/user-attachments/assets/e97e8c56-4889-47b2-9ccc-dc0e47ebe296)



![395135121-2fd3013f-88a8-4cd9-90cc-a0551d6e6700](https://github.com/user-attachments/assets/ebef78c9-ff27-4523-9f65-e34209b2a736)


# Cosmos SDK Port of smart contract is also available at - 

https://github.com/suhasagg/photosynthesis-cosmos-sdk/tree/main/deps/archway-main/x/photosynthesis

