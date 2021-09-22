import { Given } from 'cucumber';
import { ChainId } from '@uniswap/sdk';

Given('probe', async () => {
  console.log(ChainId);
});
