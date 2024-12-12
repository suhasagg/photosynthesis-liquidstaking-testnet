#!/bin/bash

set -eu 
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

source "${SCRIPT_DIR}/../config.sh"

for chain in ${HOST_CHAINS[@]:-}; do
    # Retrieve chain-specific variables
    chain_id=$(GET_VAR_VALUE "${chain}_CHAIN_ID")
    chain_name=$(printf "%s" "$chain" | awk '{ print tolower($0) }')
    account_name=$(GET_VAR_VALUE "RELAYER_${chain}_ACCT")
    mnemonic=$(GET_VAR_VALUE "RELAYER_${chain}_MNEMONIC")

    relayer_logs="${LOGS}/relayer-${chain_name}.log"
    relayer_home="$STATE/relayer-${chain_name}"
    relayer_config="$relayer_home/config"

    mkdir -p "$relayer_config"
    chmod -R 777 "$relayer_home"
    cp "${DOCKERNET_HOME}/config/relayer_config.yaml" "$relayer_config/config.yaml"

    printf "STRIDE <> %s - Adding relayer keys...\n" "$chain"
    # Restore the Stride key with the correct mnemonic
    rly keys restore stride "$RELAYER_STRIDE_ACCT" "$RELAYER_STRIDE_MNEMONIC1" --home "$relayer_home" >> "$relayer_logs" 2>&1
    # Restore the counterparty chain key with its mnemonic
    rly keys restore "$chain_name" "$account_name" "$mnemonic" --home "$relayer_home" >> "$relayer_logs" 2>&1
    echo "Done"

    printf "STRIDE <> %s - Creating client, connection, and transfer channel...\n" "$chain" | tee -a "$relayer_logs"
    rly transact link stride-"$chain_name" --home "$relayer_home" >> "$relayer_logs" 2>&1
    echo "Done"

    # Start the relayer in the background
    echo "Starting relayer..." | tee -a "$relayer_logs"
    nohup rly start stride-"$chain_name" --home "$relayer_home" >> "$relayer_logs" 2>&1 &
    echo "Relayer started for STRIDE <> %s\n" "$chain"
done
