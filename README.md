# Reports
http://docs.neon-labs.org/neon-e2e/

# Installation
## Settings
Modify the .env file or use the environment variables exactly like those that are in the .env file

## Python

### Prerequisites
Ubuntu 21.04
```
brew install python
sudo apt install pipenv
sudo apt install python3-behave
sudo apt install -y python3-bitarray
brew install allure

pip install lru-dict
```
Mac OSX 11
```
brew install python
brew install pipenv
brew install allure
```

### Initialization
```
cd neon-e2e
pipenv sync
pipenv shell
```

### Run
```
behave
allure serve
```

## Node.js

### Prerequisites
TBD

### Initialization
```
npm i
```

### Run
```
npm test
```


## REMIX EXAMPLE PROJECT

Remix example project is present when Remix loads very first time or there are no files existing in the File Explorer. 
It contains 3 directories:

1. 'contracts': Holds three contracts with different complexity level, denoted with number prefix in file name.
2. 'scripts': Holds two scripts to deploy a contract. It is explained below.
3. 'tests': Contains one test file for 'Ballot' contract with unit tests in Solidity.

### SCRIPTS

The 'scripts' folder contains example async/await scripts for deploying the 'Storage' contract.
For the deployment of any other contract, 'contractName' and 'constructorArgs' should be updated (along with other code if required). 
Scripts have full access to the web3.js and ethers.js libraries.

To run a script, right click on file name in the file explorer and click 'Run'. Remember, Solidity file must already be compiled.

Output from script will appear in remix terminal.
