import { Given } from 'cucumber';
import { ChainId, Fetcher, Pair, Token, TokenAmount, WETH } from '@uniswap/sdk';
import { Config } from 'config/default';

Given('probe01', async () => {
  console.log(ChainId);
  console.log(ChainIdPlus);
  const chainId = ChainIdPlus.DEVNET; // parseInt(Config.networkId);
  const tokenAddress = Config.tokenAddressA;
  const decimals = 18;

  const DAI1 = new Token(chainId, tokenAddress, decimals);
  console.log(DAI1);
  const DAI2 = await Fetcher.fetchTokenData(chainId, tokenAddress);
  console.log(DAI2);
  const DAI3 = new Token(chainId, tokenAddress, decimals, 'DAI', 'Dai Stablecoin');
  console.log(DAI3);
  const DAI4 = await Fetcher.fetchTokenData(chainId, tokenAddress, undefined, 'DAI', 'Dai Stablecoin');
  console.log(DAI4);
});

Given('probe02', async () => {
  console.log('==== ChainId ====');
  console.log(ChainId);
  console.log('==== ChainIdPlus ====');
  console.log(ChainIdPlus);
  const chainId = ChainIdPlus.TESTNET;
  const tokenAddress = Config.tokenAddressA;
  const decimals = await getDecimals(chainId, tokenAddress);
  const DAI1 = new Token(chainId, tokenAddress, decimals, 'SYM', 'Neo');
  console.log('==== DAI1 ====');
  console.log(DAI1);

  const pair = await getPair(DAI1);
  console.log('==== pair ====');
  console.log(pair);
});

const getDecimals = async (chainId: ChainId, tokenAddress: string): Promise<number> => {
  return 18;
};

const ChainIdPlus = {
  ...ChainId,
  TESTNET: 111,
  DEVNET: 110
};

async function getPair(DAI: Token): Promise<Pair> {
  const pairAddress = Pair.getAddress(DAI, WETH[DAI.chainId]);

  const reserves: [any, any] | never[] = [
    /* */
  ];
  const [reserve0, reserve1] = reserves;

  const tokens = [DAI, WETH[DAI.chainId]];
  const [token0, token1] = tokens[0].sortsBefore(tokens[1]) ? tokens : [tokens[1], tokens[0]];

  const pair = new Pair(new TokenAmount(token0, reserve0), new TokenAmount(token1, reserve1));
  return pair;
}
