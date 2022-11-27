# Decentralized-print

Enabling autonomous and open-for-all printing ecosystem using Ethereum smart contract. \
\
Final project for 2110595 ADV. TOPICS IN COMPUTER ENGINEERING (BLOCKCHAIN).

# Project Info

This project aims to serve as a protocol to connect between printer owners and those with documents to be printed. \
The parties involved are as follows:

  - `Users` : (or customers) want to print some documents.
  - `Providers` :  (or printer owners) want to offer a printing service, in exchange for some $$$.

The project is structured as follows:

  - `contract` : contains all the files related to smart contracts (implemented in Solidity using Hardhat). Transaction contract represents a printing order, Printer contract represents a printer and CentralServer acts as a printer registry, keeping track of the open printers.
  - `frontend` : contains the code for the frontend UI. Although users may interact with contracts directly (or however they wish), this web UI was created for convenience and demostration purposes.
  - `printer-node` : contains the code for integrating a printer with our project. This was implemented in Go.

More details can be found in `final_presentation.pdf`

# Set-up & Installation

## Setup Central server

```bash
cd central_server_contract
go run main.go
```

then enter all information that prompt in the terminal and it will deploy central server and output the public key of the instance.

## Frontend

### Pre-install

- Node.js
- Yarn

### Start the service

```bash
cd frontend
yarn
yarn start
```

service will be at localhost:3000

## Printer-node

### Pre-install

- Go 1.19

### Start the service

```bash
cd printer-node
go mod tidy
go run main.go
```

then enter all information that prompt in the terminal after that it will service will running until you pressed CTRL+C it will stopped the service

# Team members
1. 6230252121 Tarm Kalavantavanich
2. 6231301421 Kanokpich Chaiyawan
3. 6231350121 Puwadol Vichitbandha
4. 6232002021 Khajornphong Phimmuang
