import { ethers } from 'ethers';
import { Pool } from '@uniswap/v3-sdk';
import { Address } from 'cluster';
import { Given } from 'cucumber';
import { Config } from 'config/default';
import { expect } from 'chai';

const ProviderShouldNotBeNull = 'Provider should not be null';

// let provider: ethers.providers.JsonRpcBatchProvider;

const getProvider = async (url: string) => {
  return new ethers.providers.JsonRpcProvider(Config.url);
};

const BlockNumberShouldBeGreaterThan0 = 'Block number should be greater than 0';
Given('there is connection to Neonswap', async () => {
  const provider = await getProvider(Config.url);
  console.log(`Provider = ${provider}`);
  expect(provider, ProviderShouldNotBeNull).to.not.be.null;
  console.log(`${provider.connection.url}`);
  console.log(`${provider.connection.headers}`);
  console.log(`Block number = ${provider.blockNumber}`);
  const currentBlockNumber = await provider.getBlockNumber();
  console.log(`Current block number = ${currentBlockNumber}`);
  expect(currentBlockNumber, BlockNumberShouldBeGreaterThan0).to.be.greaterThan(0);
});

Given('next attempts', async () => {
  const provider = await getProvider(Config.url);

  const poolAddress = '0x8ad599c3A0ff1De082011EFDDc58f1908eb6e6D8';

  const poolImmutablesAbi = [
    'function factory() external view returns (address)',
    'function token0() external view returns (address)',
    'function token1() external view returns (address)',
    'function fee() external view returns (uint24)',
    'function tickSpacing() external view returns (int24)',
    'function maxLiquidityPerTick() external view returns (uint128)'
  ];

  const poolContract = new ethers.Contract(poolAddress, poolImmutablesAbi, provider);

  async function getPoolImmutables() {
    const PoolImmutables: Immutables = {
      factory: await poolContract.factory(),
      token0: await poolContract.token0(),
      token1: await poolContract.token1(),
      fee: await poolContract.fee(),
      tickSpacing: await poolContract.tickSpacing(),
      maxLiquidityPerTick: await poolContract.maxLiquidityPerTick()
    };
    return PoolImmutables;
  }

  getPoolImmutables().then((result) => {
    console.log(result);
  });
});

interface Immutables {
  factory: Address;
  token0: Address;
  token1: Address;
  fee: number;
  tickSpacing: number;
  maxLiquidityPerTick: number;
}
