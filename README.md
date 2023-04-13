# sommelier
Securely generate keys and addresses for Sommelier Finance network.

## Air Gapped Seed Phrase Generator and Tendermint Address Tool
To be used on an air-gapped machine.

This is a command-line tool to generate a random seed phrase, shard it into multiple pieces and combine them to derive a Tendermint address. It is designed to be used on an air-gapped computer for added security. This document explains the usage of the tool and the various options available.

### Prerequisites
Go version 1.20 or later must be installed on your computer.

### Installation
Clone this repository:

`git clone https://github.com/[user]/[repo].git`

`cd [repo]`

#### Install dependencies:
`go get ./...`

Build the binary file:

_Note: it is important that the `-trimpath` flag be added to ensure a reproducible build_

`go build -trimpath main.go generate_key.go`

Run the binary file:

`./main`

### Usage
The tool accepts the following command-line flags:

* -c: Command to execute [new, derive, combineShares]
* -n: Number of shares to combine (required for combineShares command)
* -shares: Total number of shares to generate (default is 3)
* -threshold: Number of shares required to reconstruct the seed phrase (default is 2)

#### Generating a new seed phrase
To generate a new seed phrase, use the new command:

`./main -c new -shares <number of total shares> -threshold <threshold required to reconstruct phrase>`

This is because the program employs [Shamir's Secret Sharing algorithm](https://en.wikipedia.org/wiki/Shamir%27s_secret_sharing).

By default, if the shares and threshold flags are not provided, it becomes a `2 of 3` sharded phrase.

##### Example:
`./main -c new -shares 5 -threshold 3`

This generates a new seed phrase and shards it into 5 pieces. To reconstruct the seed phrase, at least 3 pieces are required.

**Example Output:**
```shell
./main -c new -shares 5 -threshold 3
generated mnemonic:  slam sock rib build diesel repair arrange always describe reform surprise crop nasty oval genuine magic divorce borrow session old manual approve trap mail
seed shard:  0 :  d4344518a3dcd99820d82c292a6fa6ec9039c3c96d7dc7785a734d0fb6866a1d98f8e181691b60761eded3c791c4487e5c0d40014124b261b4c6b31c305e3dbb9b01b1f73ac04060a2885e15faa487c476da5cd12f09eae3852c6a405ce0c96500e50871737015c0644ec6867438b0676f9e940ec79607bf2d3fa33c8e5820c89467d87d4bc34120522f1f79b7fa34f5c68ef700f238a57065b0f101
seed shard:  1 :  bc3792bf1cef2f53ed6bdfc79f022fee0be9abe41e6b275b184607f3519e4992cd3c2611fc3fcbe8129d85543ffca1ad88e7bf7e668c0e190368b24098b750d4be7e38b8df9644ae53634c54039bef7f1819703aac606074a8d7b9bfcd197946114419c765b5ca9e64e1c0ab8443a751bc5e3ada2ce0a86d982a3cecf86920a0271f6bcb37a020e638ac9b0a46664f4a0e95487ffb8dea15c55c5329
seed shard:  2 :  54998d4feb4f5e31c81247b1a10391d37b2243045fdfe208c4a12f4b40991a6b367eca3b3fb8b5719cd2d0d0a80603c839eba9b14f5c570200ff21dc481551e8aab18679f3c33e25f6fc822652b507f2b78f03a5db1b708a48bbdca97448ae1571f37957d164213f64e5874bb2f68018d966ad6bb348ac9ba630362750f420485e04125b5612af99eddb384be3c2fabdc23cbfb09cb32be0e1174a20
```

#### Deriving a Tendermint address
To derive a Tendermint address from a seed phrase, use the derive command:

`./main -c derive -phrase <"seed phrase (wrapped in quotes)"> -path <"HD PATH (wrapped in quotes)"> -hrp <address prefix>`

This generates a new seed phrase and derives a Tendermint address from it.

The default HD path is `m/44'/118'/0'/0/0`, but you can specify a different path if needed.

**Example Output deriving a Sommelier Network Address:**
```shell
./main -c derive -phrase "slam sock rib build diesel repair arrange always describe reform surprise crop nasty oval genuine magic divorce borrow session old manual approve trap mail" -path "m/44'/118'/0'/0/1" -hrp "somm"
somm1azsdxclsdsjuzevtk6552jlthmkmfxktw50we8
```

#### Combining seed phrase shares
To combine seed phrase shares into a full seed phrase, use the combineShares command:

`./main -c combineShares -n 3 "share1" "share2" "share3"`

This combines three seed phrase shares into a full seed phrase. The shares must be separated by spaces.

You also will need to add `-n <number of shares>` to specify how many shards you are recovering from.

**Example Output combining 2 of 3 seed phrase shares:**
```shell
./main -c combineShares -n 2 d4344518a3dcd99820d82c292a6fa6ec9039c3c96d7dc7785a734d0fb6866a1d98f8e181691b60761eded3c791c4487e5c0d40014124b261b4c6b31c305e3dbb9b01b1f73ac04060a2885e15faa487c476da5cd12f09eae3852c6a405ce0c96500e50871737015c0644ec6867438b0676f9e940ec79607bf2d3fa33c8e5820c89467d87d4bc34120522f1f79b7fa34f5c68ef700f238a57065b0f101 54998d4feb4f5e31c81247b1a10391d37b2243045fdfe208c4a12f4b40991a6b367eca3b3fb8b5719cd2d0d0a80603c839eba9b14f5c570200ff21dc481551e8aab18679f3c33e25f6fc822652b507f2b78f03a5db1b708a48bbdca97448ae1571f37957d164213f64e5874bb2f68018d966ad6bb348ac9ba630362750f420485e04125b5612af99eddb384be3c2fabdc23cbfb09cb32be0e1174a20
original seed:  slam sock rib build diesel repair arrange always describe reform surprise crop nasty oval genuine magic divorce borrow session old manual approve trap mail
```
## Conclusion
This tool provides a simple and secure way to generate and use seed phrases on an air-gapped computer.

Below is a procedure document as to how this tool should be used securely.
**WIP**