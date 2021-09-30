/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable @typescript-eslint/no-non-null-assertion */
/* eslint-disable @typescript-eslint/no-unsafe-assignment */
/* eslint-disable @typescript-eslint/no-unsafe-call */
/* eslint-disable @typescript-eslint/no-unsafe-member-access */
/* eslint-disable @typescript-eslint/no-unsafe-return */
/* eslint-disable @typescript-eslint/explicit-module-boundary-types */
import { unit } from 'src/constants/unit';
import { Account } from 'web3-core';

import { Web3Helper } from './web3-helper';

export class TxHelper {
  public static async signTransaction(sender: Account, recipient: Account, amount: number) {
    return await Web3Helper.getWeb3().eth.accounts.signTransaction(
      {
        from: sender.address,
        gas: '4294967295',
        to: recipient.address,
        value: Web3Helper.getWeb3().utils.toWei(amount.toString(), unit.ether)
      },
      sender.privateKey
    );
  }

  public static async deployTransaction(transaction: any) {
    return await Web3Helper.getWeb3().eth.sendSignedTransaction(transaction.rawTransaction!);
  }
}
