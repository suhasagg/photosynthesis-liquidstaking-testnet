#!/bin/bash

for i in {1..16};do
    OUTPUT=$(/home/keanu-xbox/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv13/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /home/keanu-xbox/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv13/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 tx wasm execute archway14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9sy85n2u '{"increase_allowance": {"spender": "archway1qygx0pxuttycdddzz5lre5rlxcxjemthwmlh63", "amount": "1"}}' --from pval2 --chain-id localnet --keyring-backend=test --gas 205000 --fees 20000uarch -y)

    sleep 2
    echo "$OUTPUT"
    echo "Script execution completed."

    txhash=$(echo "$OUTPUT" | grep -oP 'txhash: \K.*')
    echo "$txhash"
done

for i in {1..16};do
    OUTPUT=$(/home/keanu-xbox/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv13/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /home/keanu-xbox/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv13/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 tx wasm execute archway1suhgf5svhu4usrurvxzlgn54ksxmn8gljarjtxqnapv8kjnp4nrsmjcxz4 '{"increase_allowance": {"spender": "archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m", "amount": "1"}}' --from pval5 --chain-id localnet --keyring-backend=test --gas 205000 --fees 20000uarch -y)

    sleep 2
    echo "$OUTPUT"
    echo "Script execution completed."

    txhash=$(echo "$OUTPUT" | grep -oP 'txhash: \K.*')
    echo "$txhash"
done