/* eslint-disable @typescript-eslint/restrict-plus-operands */
/* eslint-disable @typescript-eslint/restrict-template-expressions */
/* eslint-disable @typescript-eslint/no-unsafe-assignment */
import axios from 'axios';

import { Config } from '../../../config/default';
import { logger } from '../../utils/logger';
import { Urls } from '../common/urls';
import { FaucetRequest } from './faucet-request';

export const requestFaucet = async (wallet: string, amount: number): Promise<void> => {
  const data: FaucetRequest = { amount: amount, wallet: wallet };
  const url: string = Config.baseUrl + Urls.request_erc20_tokens;
  logger.notice(`URL: ${url}`);
  logger.notice(`Wallet = ${data.wallet}, amount = ${data.amount}`);
  const result = await axios.post(url, data);
  logger.notice(result);
};
