require('dotenv').config();

let Web3 = require('web3');
// ganache local blockchain
let url_ganache = 'http://127.0.0.1:7545';

// infura
let url_infura = 'https://rinkeby.infura.io/v3/6f20cda0c408428a92a4c8680552f5a1';

// using geth
let url_geth ='http://127.0.0.1:8545';

let url_evm = process.env.RPC_URL;

// let web3 = new Web3(url_ganache);
// let web3 = new Web3(url_infura);
// let web3 = new Web3(url_geth);
// let web3 = new Web3(url_evm);

// console.log(web3);