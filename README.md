# BitcoinAddress (HTTP)
[![GoDoc](https://godoc.org/github.com/thomasxnguy/bitcoinaddress?status.svg)](https://godoc.org/github.com/thomasxnguy/bitcoinaddress)
[![license](https://img.shields.io/github/license/thomasxnguy/bitcoinaddress.svg?maxAge=2592000)](https://github.com/thomasxnguy/bitcoinaddress/LICENSE)

An HTTP server for bitcoin addresses generation.


|    Method | Endpoint    | Description                                                       |
| :----:  | :-----------: | -------------------------------------------------------------------|
|  GET | /address/gen   | Generate a bitcoin segwit address for a user.        |
|  GET  | /address/:user_id   | Get the bitcoin address of a user. This address is regenerated from the path's index (no key is actually stored server side)      |
|  POST | /p2sh   | Generate a n-out-of-m multisig p2sh address.  |

A use case for this boilerplate code is to build a server for managing user's account wallet, allowing users to receive payment in bitcoin.


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
1. Set the seed through a secure channel manually by an operator. The seed can only be set in memory during startup (We can imagine the server when re-started will be in "INIT" state waiting for configuration request by an operator in a protected endpoint).
2. Run the service in a protected memory zone such as an enclave

## Examples

```bash
curl
```

### Todo

- [ ] Add integration and unit tests
- [ ] Store users address mapping and key generation index in DB (Use atomic increment)
- [ ] Get account current bitcoin amount
- [ ] Allow to send bitcoin to an address (sign transactions)

## Resources 

The folder 0_postman contains the postman collection script to run scenarios on the API