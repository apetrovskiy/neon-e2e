/* eslint-disable @typescript-eslint/no-unsafe-call */
/* eslint-disable @typescript-eslint/no-unsafe-assignment */
/* eslint-disable @typescript-eslint/no-unsafe-return */
/* eslint-disable @typescript-eslint/no-unsafe-member-access */
/* eslint-disable @typescript-eslint/restrict-template-expressions */
/* eslint-disable no-console */
import { CompileFailedError, CompileResult, compileSol } from 'solc-typed-ast';

export const compileContract = (filename: string): CompileResult | void => {
  let result: CompileResult;

  try {
    result = compileSol(filename, 'auto', []);
    console.log(result);
    return result;
  } catch (error) {
    if (error instanceof CompileFailedError) {
      console.error('Compile errors encountered:');

      for (const failure of error.failures) {
        console.error(`SolcJS ${failure.compilerVersion}:`);

        for (const error of failure.errors) {
          console.error(error);
        }
      }
    } else {
      console.error(error); //.message);
    }
  }
};

// console.log(result);
///////////////////////////////////////////////////

// var solc = require('solc');
// import solc from 'solc';

// var input = {
//   language: 'Solidity',
//   sources: {
//     'test.sol': {
//       content: 'contract C { function f() public { } }'
//     }
//   },
//   settings: {
//     outputSelection: {
//       '*': {
//         '*': ['*']
//       }
//     }
//   }
// };

// var output = JSON.parse(solc.compile(JSON.stringify(input)));

// // `output` here contains the JSON output as specified in the documentation
// for (var contractName in output.contracts['test.sol']) {
//   console.log(contractName + ': ' + output.contracts['test.sol'][contractName].evm.bytecode.object);
// }
