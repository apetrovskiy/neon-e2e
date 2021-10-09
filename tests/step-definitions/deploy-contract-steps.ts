import { Given, Then, When } from 'cucumber';
import path from 'path';
import { CompileResult } from 'solc-typed-ast';
import { CONTRACTS_DIR } from 'src/constants/directories';
import { compileContract } from 'src/helpers/contract-compiler';
import { readFileContent } from 'src/helpers/file-system';
import { Contract } from 'web3-eth-contract';

let contractText: string;
let contractFilePath: string;
let compiledContract: CompileResult;

Given("there is a contract '{}'", async (contractFileName: string) => {
  console.log('=================================================================');
  console.log(contractFileName);
  console.log('=================================================================');
  console.log(CONTRACTS_DIR);
  contractFilePath = path.resolve(CONTRACTS_DIR, contractFileName);
  console.log('=================================================================');
  console.log(contractFilePath);
  contractText = await readFileContent(contractFilePath);
  console.log('=================================================================');
  console.log(contractText);
});

Given('the contract is compiled', async () => {
  compiledContract = compileContract(contractFilePath);
  console.log('=============================== copilation result ==================================');
  console.log(compiledContract);
});

When('the contract is deployed', async () => {
  // (compiledContract as CompileResult)..deploy();
  const contract = new Contract(compiledContract.data);
  contract.methods.deploy(); //DeployOptions);
});

Then('there are no errors', async () => {});
