# sommelier
Securely generate keys and addresses for Sommelier Finance network.

## Air Gapped Seed Phrase Generator and Tendermint Address Tool
To be used on an air-gapped machine.

This is a command-line tool to generate a random seed phrase, shard it into multiple pieces and combine them to derive a Tendermint address. It is designed to be used on an air-gapped computer for added security. This document explains the usage of the tool and the various options available.

**Note: We can build the binary on a hot machine, then transfer it to the air gapped machine.
We can run a separate shasum check on the binary to make sure it's not tampered.**

## Procedure Document
Below is a procedure document as to how this tool should be used securely:

https://docs.google.com/document/d/1qtjWUvA-DNqzNYQe2OxFSzh6PNVG3wrF-zwTwLMzhfQ/edit?usp=sharing

^ Shared with team@unit410.com

### Prerequisites
Go version 1.20 or later must be installed on your computer.

### Installation (Build binary yourself)
Clone this repository:

`git clone https://github.com/mgallegos213/sommelier.git`

`cd sommelier`

#### Install dependencies:
`go get ./...`

Build the binary file:

_Note: As mentioned before, this binary should be built on a hot machine and transferred and verified on the air-gapped one where this will run. It is important that the `-trimpath` flag be added to ensure a reproducible build._

`go build -trimpath main.go`

Run the binary file:

`./main`

### Usage
The tool accepts the following command-line flags:

* -c: Command to execute [new, derive, combineShares]
* -unsafe: DEBUG ONLY - outputs seed shard and mnemonic values, used for project demo purposes

**Command: new**
* -shares: Total number of shares to generate (default is 3)
* -threshold: Number of shares required to reconstruct the seed phrase (default is 2)

**Command: derive**
* -phrase: seed phrase plaintext to derive address for
* -path: HD path to derive address for, defaults to `m/44'/118'/0'/0/0`
* -hrp: prefix used for bech32 address, `somm` or `cosmos` are examples

**Command: combineShares**
* -n: Number of shares to combine (required for combineShares command)
* encrypted shard paths as arguments

#### Generating a new seed phrase
To generate a new seed phrase, use the new command:

`./main -c new -shares <number of total shares> -threshold <threshold required to reconstruct phrase>`

This is because the program employs [Shamir's Secret Sharing algorithm](https://en.wikipedia.org/wiki/Shamir%27s_secret_sharing).

By default, if the shares and threshold flags are not provided, it becomes a `2 of 3` sharded phrase.

##### Example:
`./main -c new -shares 5 -threshold 3`

This generates a new seed phrase and shards it into 5 pieces. To reconstruct the seed phrase, at least 3 pieces are required.

**Example Output: (with -unsafe flag for visibility)**
```shell
./main -cmd new -unsafe -shares 5 -threshold 3
Generated mnemonic:  coach deposit camp dilemma crucial glance device push chef loan inform like scout town person expect second turtle drip setup away gold damage gravity
seed shard  0 :  a8ec9eb29838ecd56a0d946b1bf823617cd0a3e98daef650676e0373484612bf2511ebb786de7ac746359bdb25a949e1697da171c7719fa1bc99a23717d8d578efc6b93d17c5be1ddafc927bc25659de63273563c3a87e9d7b80256e361922783e9021384a50c801f2866e457959c39b33a0cf68f339acc7b82c4c7555262fedbfaaa74b0f5bc354a04c196c65225e1954bb7185efc772
seed shard  1 :  0699e8bb34f3c32d8d942ac6873a33edf3e9399acdac8943d17d656b9a72d8d363c115e925d8c173cd785ac63b079f17178e7e62bab1387b355d8428425a20f76c0110aa01f7bf00c7be898ba602e02adc1a26159d1b281e1cf83a958d62940a7693ae93ee165fd530a90a4bacb9261f63830a3899718c1c58b0fbba3bef0848696f79368dcc0ea7d285b0982f2bc4aef2665792aabba0
seed shard  2 :  bb2cc5958111a5a8f5fd26d08d74f7015e25b0602fbeaeb351bc5807647424d92b6c3f200822dd944ef7e2119d4250c5d8d2bd9efcade1f5c1bbc9d856b73d59ccdbd9122ed1697a2dfbc0471fd2615830978d76368766a22f4bf766b83991e7c9db7a7274d33837ec2b522695aa133c1c41833ef132185dad230b7efd734632810582f426ba168a6abdb6183b85d123c3f55b2ca23ef1
seed shard  3 :  56bf44bfe7ba5ab67a6c2c8d11d94e1142667f1e4211915801636c7065bf632493d488b914762535efe16c3ec159b43d3271462c93f2836f4d55605b4561bc98aa09ea7e548d0c5e6f80729843ee085c81135465c77c16412bee19e1c6ac07508ef261698cc3eb87b1a7014752a336f66cf8e84950b49b5b41fc86065bd0cd712cf797160cfee9402030e6123315219e746ae4845d946d
seed shard  4 :  3ecf65a4d0fd863ef91f85f730410efba3bfa28549b26a921425294dd6a37acadba62c75cd9f15028dc032e173c5174e38e975720041a6112f41c82e1c90945c8a861cb73507c0b0a82a0d40c696212ec8340097086452d7124c74a449b1766e720295846dfe3250cf39a904372d0d0d7651a6466f39bb30420f5ad3ec6c67706b5bf40b3fef47dcdb46224afa0541d2a0bd2e01f48620
Successfully stored encrypted key shards.
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
./main -c combineShares -n 2 -unsafe shard/shard_0.txt.gpg shard/shard_1.txt.gpg
Shard data:  [868948765aedccbf9ed103a236a4758cf44b54182b793092ad53dddd3725cb761e44c48d4e4240d7eedb2318279f77a5c2685470b43935ba5ff10b9e774ee523760e8f5031305196ca834eded2bcbc0e130a8c2eff84d9f8a1dc5cf6cce7939cc11adbcbc588b944905b48faca92afbde9b776aebaa8490301a186dd3e69a525402e30b1f12f8870c07c589ff859e4194b9e982a2c36a306cea2 f869b1ba33b8472b492feab5679a4fad4a7effd826d0f155e68d48ef29d0a59d0f5b687eb781c28d927789be0358fa57593aff3f1a6679cbc08e97dacedf237265bb55f6e348c0805296526f6d5db825f715a0ae27b534fc4edaeaf10586b1ffa19bdd827752a02d4a4a1b3ff65526e5186f392909e3208c1e0c0ec6c58c8be4eb040a4d52c734a2ef616d94978415d76d3f9df6b2b8cfb3d5d1]
Original seed:  abstract humble brain wasp female squirrel hello favorite kick wait actor parade nephew chat destroy silk penalty private sail bus door stand twelve wear
Saved original seed phrase to encrypted file:  shard/encrypted_phrase.txt.gpg
```
## Conclusion
This tool provides a simple and secure way to generate and use seed phrases on an air-gapped computer.
