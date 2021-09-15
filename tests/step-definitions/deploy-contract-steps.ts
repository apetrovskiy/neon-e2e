import { Given, Then, When } from 'cucumber';
import path from 'path';
import { CONTRACTS_DIR } from 'src/constants/directories';
import { compileContract } from 'src/helpers/contract-compiler';
import { readFileContent } from 'src/helpers/file-system';

let contractText: string;
let contractFilePath: string;

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

When('compiling the contract', async () => {
  const compiledContract = compileContract(contractFilePath);
  console.log('=================================================================');
  console.log(compiledContract);
});

Then('there is no errors', async () => {});
