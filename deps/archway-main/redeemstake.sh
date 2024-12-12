#!/bin/bash

# Directories
BASE_DIR="/home/keanu-xbox/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv13/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow"
LOG_DIR="$BASE_DIR/dockernet/logs"
STATE_DIR="$BASE_DIR/dockernet/state"
BIN_DIR="$BASE_DIR/build"

ARCHWAYD="$BIN_DIR/archwayd --home $STATE_DIR/photo1"
STRIDED="$BIN_DIR/strided --home $STATE_DIR/stride1"

# Variables
LIQUIDITY_FILE="$LOG_DIR/liquiditydistributionforDapps"
ENABLER_FILE="$LOG_DIR/enableredeemstake"
REDEEM_EPOCH_FILE="$LOG_DIR/redeemepoch"
TOTAL_REDEMPTION_AMOUNT=0

# Ensure ~/bin is in PATH
if ! echo "$PATH" | grep -q "$HOME/bin"; then
  export PATH="$HOME/bin:$PATH"
  echo 'export PATH="$HOME/bin:$PATH"' >> ~/.bashrc
  source ~/.bashrc
fi

# Execute a command and convert YAML to JSON
execute_cmd() {
  local CMD=$1
  OUTPUT=$(eval "$CMD")
  echo "$OUTPUT" | yq eval -j -
}

# Query the total redemption amount
query_total_redemption_amount() {
  CMD="$ARCHWAYD q photosynthesis total-redemption-amount --node http://localhost:26457 --chain-id localnet -o json"
  echo "Executing: $CMD"  # Print the command being executed for visibility

  # Execute the command and capture the output
  JSON_OUTPUT=$(eval "$CMD")

  # Print the raw output for debugging
  echo "Raw JSON Output: $JSON_OUTPUT" | jq -R -c '{"message": .}'

  # Extract the redemption amount using jq
  TOTAL_REDEMPTION_AMOUNT=$(echo "$JSON_OUTPUT" | jq -r '.amount')

  # Print the final redemption amount
  if [ -n "$TOTAL_REDEMPTION_AMOUNT" ] && [ "$TOTAL_REDEMPTION_AMOUNT" != "null" ]; then
    echo "Total Redemption Amount: $TOTAL_REDEMPTION_AMOUNT" | jq -R -c '{"message": .}'
  else
    echo "Error: Failed to retrieve redemption amount." | jq -R -c '{"message": .}'
    exit 1
  fi
}

# Execute a transaction with retries
execute_transaction() {
  local CMD=$1
  local MAX_RETRIES=5
  local RETRY_COUNT=0

  while [ $RETRY_COUNT -lt $MAX_RETRIES ]; do
    OUTPUT=$(eval "$CMD")
    echo "$OUTPUT" | jq -R -c '{"message": .}'

    if [[ "$OUTPUT" != *"account sequence mismatch"* ]]; then
      echo "Transaction successful!" | jq -R -c '{"message": .}'
      return 0
    else
      echo "Sequence mismatch. Retrying in 5 seconds..." | jq -R -c '{"message": .}'
      ((RETRY_COUNT++))
      sleep 5
    fi
  done

  echo "Transaction failed after $MAX_RETRIES attempts." | jq -R -c '{"message": .}'
  return 1
}

# Redeem stake
redeem_stake() {
  CMD="$STRIDED tx stakeibc redeem-stake $TOTAL_REDEMPTION_AMOUNT localnet archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m --from admin --keyring-backend test --chain-id STRIDE -y"
  execute_transaction "$CMD"
}

# Claim undelegated tokens for claimable records
claim_undelegated_tokens() {
  local json="$1"

  echo "${json}" | jq -r '.epoch_unbonding_record[] | @base64' | while read epoch; do
    epoch_data=$(echo "$epoch" | base64 --decode)

    echo "$epoch_data" | jq -r '.host_zone_unbondings[] | @base64' | while read row; do
      row_data=$(echo "$row" | base64 --decode)
      status=$(echo "$row_data" | jq -r '.status')

      if [ "$status" = "CLAIMABLE" ]; then
        records=$(echo "$row_data" | jq -c '.user_redemption_records[]')

        echo "$records" | jq -r '.' | while read record; do
          part1=$(echo "$record" | jq -r '.hostzone')
          part2=$(echo "$record" | jq -r '.epoch')
          part3=$(echo "$record" | jq -r '.stride_address')

          echo "Claiming undelegated tokens for Hostzone: $part1, Epoch: $part2" | jq -R -c '{"message": .}'

          CMD="$STRIDED tx stakeibc claim-undelegated-tokens $part1 $part2 $part3 --from admin --keyring-backend test --chain-id STRIDE -y"
          execute_transaction "$CMD"
        done
      fi
    done
  done
}

# Query and claim epoch unbonding records
process_epoch_unbonding_records() {
  local CMD="$STRIDED q records list-epoch-unbonding-record"
  JSON_OUTPUT=$(execute_cmd "$CMD")
  echo "Raw JSON Output: $JSON_OUTPUT" | jq -R -c '{"message": .}'
  echo "Processing epoch unbonding records" | jq -R -c '{"message": .}'
  claim_undelegated_tokens "$JSON_OUTPUT"
}

# Main workflow logic
main() {
  echo "Querying total redemption amount" | jq -R -c '{"message": .}'
  query_total_redemption_amount

  echo "Redeeming total stake" | jq -R -c '{"message": .}'
  redeem_stake

  sleep 500
  echo "Processing epoch unbonding records" | jq -R -c '{"message": .}'
  process_epoch_unbonding_records

  echo "Redemption complete. Final Redemption Amount: $TOTAL_REDEMPTION_AMOUNT" | jq -R -c '{"message": .}'
  echo "Final Redemption Amount: $TOTAL_REDEMPTION_AMOUNT" >> "$LIQUIDITY_FILE"

  if [ -f "$ENABLER_FILE" ]; then
    last_line=$(tail -n 1 "$ENABLER_FILE")
    echo "Last line of enabler file: $last_line" | jq -R -c '{"message": .}'
    echo "$last_line" >> "$REDEEM_EPOCH_FILE"
    rm "$ENABLER_FILE"
  else
    echo "Enabler file not found." | jq -R -c '{"message": .}'
  fi
}

# Start the script
main
