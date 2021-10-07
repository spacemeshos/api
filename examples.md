# Usage Examples
Examples below assume you are running a locan full node with the default provided reference config file.

## Working with Binary Data
Binary data submitted via grpcurl to the api must be a base64 string encoded in the following way: `base64(hex.decode(hex_string))`.
So for example for address represted as a hex string of `0x8C99365feE31f1cEDB056A35D7f561584679ea3c`, the base64 encoding is `jJk2X+4x8c7bBWo11/VhWEZ56jw='`

## Account Balance
Returns the current and project balance of an account identified by an address.

```bash
grpcurl -d '{"filter" : {"accountId":{"address":"jJk2X+4x8c7bBWo11/VhWEZ56jw="}, "account_data_flags":4}}' -plaintext localhost:9092 spacemesh.v1.GlobalStateService.AccountDataQuery
{
  "totalResults": 1,
  "accountItem": [
    {
      "accountWrapper": {
        "accountId": {
          "address": "jJk2X+4x8c7bBWo11/VhWEZ56jw="
        },
        "stateCurrent": {
          "balance": {
            "value": "152276432652715"
          }
        },
        "stateProjected": {
          "balance": {
            "value": "152276432652715"
          }
        }
      }
    }
  ]
}
```


## Account Rewards
Returns the smeshing rewards of a rewards account identified by its address.

```bash
grpcurl -d '{"filter" : {"accountId":{"address":"jJk2X+4x8c7bBWo11/VhWEZ56jw="}, "account_data_flags":4}}' -plaintext localhost:9092 spacemesh.v1.GlobalStateService.AccountDataQuery
{

}
```


