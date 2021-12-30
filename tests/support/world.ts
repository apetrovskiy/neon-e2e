import { setWorldConstructor, World } from 'cucumber';
import { Account } from 'web3-core';

class CustomWorld implements World {
  public timestamp = '';
  public productList = [];

  public userAlice = 'userAlice';
  public userBob = 'userBob';

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  [key: string]: any;
}

setWorldConstructor(CustomWorld);
