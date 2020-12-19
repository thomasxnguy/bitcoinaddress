# BitcoinAddress (HTTP)
[![GoDoc](https://godoc.org/github.com/thomasxnguy/bitcoinaddress?status.svg)](https://godoc.org/github.com/thomasxnguy/bitcoinaddress)
<a href="https://github.com/thomasxnguy/bitcoinaddress/LICENSE"><img src="https://img.shields.io/github/license/saniales/golang-crypto-trading-bot.svg?maxAge=2592000" alt="License"></a>
<a href="https://goreportcard.com/report/github.com/thomasxnguy/bitcoinaddress"><img src="https://goreportcard.com/badge/github.com/thomasxnguy/bitcoinaddress" alt="Goreportcard" /></a>

An HTTP server for bitcoin addresses generation.


| Method |    Endpoint      |                 Description                                         |
| :----: | :--------------: | ------------------------------------------------------------------  |
|  GET  | /address/gen       |Generate bitcoin segwit addresses for a user.  |      
|  GET  | /address/:user_id  | Get the bitcoin addresses of a user. This address is regenerated from the path's index (no key is actually stored server side).   |
|  POST | /p2sh              | Generate a n-out-of-m multisig p2sh address.  |

A use case for this boilerplate is to  build  a user's account management server, allowing users to receive payment in bitcoin.
A new address will be generated for each user in order to receive payment from their customers. Server will be responsible for signing transactions and manage user's fund.


## Prerequisites
- golang >= 1.15

## Get Started
1 - Clone this repository:

 ```go get -u github.com/thomasxnguy/bitcoinaddress```

2 - Run the application to see available commands:
 
 ```go run main.go --help```

3 - Run the application with command *serve*:
 
 ```go run main.go serve```

4 - API is running on port 3000 by default. Use the postman collection to start sending requests.

#### Run with Docker

To run the application in a docker container, type the following command in root directory 

```bash
docker build -t bitcoinaddress . && docker run -p 3000:3000 -it bitcoinaddress
```

## Security concerns

Seed (or mnemonic) is currently stored in the config.json. It has potential security vulnerabilities as anyone having access to the server would be able to read the seed, get access the master key and steal the fund of managed users. Multiple solutions can be considered.
1. Set the seed through a secure channel manually by an operator, either locally or through the network. The seed can only be set in memory during startup (We can imagine the server when re-started will be in "INIT" state waiting for configuration request by an operator in a protected endpoint).
2. Run the service in a protected memory zone such as an enclave to ensure that no one can access the part of the memory.

## Specs

| Method |    Endpoint      |     Request           |         Response        |  Note |
| :----: | :--------------: | :-------------------: | :--------------------:   | ---- |
|  GET  | /address/gen       |                      |{ user_id : UUID, <br> segwit_address : string, <br> native_segwit_address : string } | segwit_address also refers to nested segwit (with BIP49). native_segwit_address also refers to bech32 (with BIP84).        |
|  GET  | /address/:user_id  |                      |  {segwit_address : string, <br> native_segwit_address : string}  |   |
|  POST | /p2sh              | {req : int,  <br>public_keys: [pubkey1, pubkey2...] } | { p2sh_address : string} | |


## Examples (Curl)

HeathCheck
```bash
curl -X GET http://localhost:3000/health -H 'cache-control: no-cache'
```
Address generation

```bash
curl -X GET http://localhost:3000/address/gen -H 'Content-Type: application/json' \
-H 'cache-control: no-cache'
```

Get address
```bash
curl -X GET http://localhost:3000/address/:user_id -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache'
```

Generate P2SH address
```bash
curl -X POST \
  http://localhost:3000/p2sh \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache' \
  -d '{
    "req": 1,
    "public_keys": [
        "04a882d414e478039cd5b52a92ffb13dd5e6bd4515497439dffd691a0f12af9575fa349b5694ed3155b136f09e63975a1700c9f4d4df849323dac06cf3bd6458cd",
        "046ce31db9bdd543e72fe3039a1f1c047dab87037c36a669ff90e28da1848f640de68c2fe913d363a51154a0c62d7adea1b822d05035077418267b1a1379790187"
    ]
}'
```

### Improvements

- Add more unit tests and improve coverage
- Replace mock DB by a real DB
- Integration with a bitcoin indexer or node to get account value (amount)
- Implement sending transactions and implement address auto-increment
- Support other coin types

## Resources 

The folder 0_postman contains the postman collection script to run scenarios on the API

[BIP44](https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki)

[BIP49](https://github.com/bitcoin/bips/blob/master/bip-0049.mediawiki)

[BIP84](https://github.com/bitcoin/bips/blob/master/bip-0084.mediawiki)

[BIP16 (P2SH)](https://github.com/bitcoin/bips/blob/master/bip-0016.mediawiki#reference-implementation)