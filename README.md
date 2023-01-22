# Mutualfund price change notification

## Prerequisites

* go 1.19

## How to run

1. clone this repo
2. install dependency
   ```go mod download```
3. set the required env**
4. run the main
   ```go run cmd/aya```

## Required ENV

* DISCORD_WEBHOOK_URL=<discord webhook url>
* MF_PRODUCT_IDS=<comma separated mutualfund product id>
