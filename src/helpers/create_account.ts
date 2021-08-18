require('./connect_probe.ts');

let Web3 = require('web3');
// let web3 = new Web3();

// let url_neon = process.env.NODE_ENV !== 'production'
// let web3 = new Web3(url_neon);

let url_evm = process.env.RPC_URL;

let web3 = new Web3(url_evm);

console.log(web3);
// console.log(url_evm);

let result = web3.eth.accounts.create();
console.log(result);
