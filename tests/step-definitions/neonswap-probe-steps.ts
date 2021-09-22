import { Given } from 'cucumber';
import { ChainId, Token } from '@uniswap/sdk';
import { Config } from 'config/default';

Given('probe', async () => {
  console.log(ChainId);
  const chainId = parseInt(Config.networkId);
  const tokenAddress = Config.tokenAddressA;
  const decimals = 18;

  const DAI = new Token(chainId, tokenAddress, decimals);
  console.log(DAI);
});
