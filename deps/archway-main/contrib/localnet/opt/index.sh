#!/bin/bash

# Elasticsearch configuration
ES_HOST="localhost"
ES_PORT="9200"
INDEX_NAME="liquid-stake-cycles-data"

# Delete the index if it exists (optional)
echo "Deleting existing index (if any)..."
curl -X DELETE "http://$ES_HOST:$ES_PORT/$INDEX_NAME" -H 'Content-Type: application/json'

# Create the new index with proper mapping
echo "Creating index with mappings..."
curl -X PUT "http://$ES_HOST:$ES_PORT/$INDEX_NAME" -H 'Content-Type: application/json' -d '
{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 1
  },
  "mappings": {
    "properties": {
      "timestamp": {
        "type": "date"
      },
      "query": {
        "type": "keyword"
      },
      "data": {
        "properties": {
          "owner": {
            "type": "keyword"
          },
          "liquid_staking_interval": {
            "type": "integer"
          },
          "arch_liquid_stake_interval": {
            "type": "integer"
          },
          "redemption_rate_query_interval": {
            "type": "integer"
          },
          "rewards_withdrawal_interval": {
            "type": "integer"
          },
          "redemption_interval_threshold": {
            "type": "integer"
          },
          "total_liquid_stake": {
            "type": "double"
          },
          "stuarch_amount": {
            "type": "double"
          },
          "deposit_records": {
            "type": "nested",
            "properties": {
              "contract_address": {
                "type": "keyword"
              },
              "amount": {
                "type": "double"
              },
              "status": {
                "type": "keyword"
              }
            }
          },
          "stake_ratios": {
            "type": "nested",
            "properties": {
              "contract_address": {
                "type": "keyword"
              },
              "ratio": {
                "type": "double"
              }
            }
          },
          "redeem_token_ratios": {
            "type": "nested",
            "properties": {
              "contract_address": {
                "type": "keyword"
              },
              "ratio": {
                "type": "double"
              }
            }
          },
          "contract_metadata": {
            "properties": {
              "rewards_address": {
                "type": "keyword"
              },
              "liquidity_provider_address": {
                "type": "keyword"
              },
              "minimum_reward_amount": {
                "type": "double"
              },
              "maximum_reward_amount": {
                "type": "double"
              }
            }
          },
          "reward": {
            "type": "double"
          }
        }
      }
    }
  }
}
'

echo "Index created successfully!"
