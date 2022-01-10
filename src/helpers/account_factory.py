import allure
from web3 import Account, Web3, HTTPProvider
from config import config
from src.helpers.faucet.faucet_requester import request_faucet


class AccountFactory():
    @allure.step("creating an account with id")
    async def create_with_specific_id(self, id: str, amount: int) -> Account:
        url = config.PROXY_URL
        w3 = Web3(HTTPProvider(url))
        print(w3)
        if id is None or id == '':
            account: Account = w3.eth.account.create(amount)
        else:
            account: Account = w3.eth.account.create(id, amount)
        if amount > 0:
            await request_faucet(account.address, amount)
        print(account)
        return account

    @allure.step("creating an account")
    async def create(self, amount: int) -> Account:
        return await self.create_with_specific_id('', amount)
