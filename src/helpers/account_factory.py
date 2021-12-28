import allure
from web3 import Web3, HTTPProvider
from config import config
from src.helpers.faucet.faucet_requester import request_faucet


class AccountFactory():
    @allure.step
    async def create_with_specific_id(self, id: str):
        url = config.PROXY_URL
        w3 = Web3(HTTPProvider(url))
        print(w3)
        if id is None or id == '':
            account = w3.eth.account.create()
        else:
            account = w3.eth.account.create(id)
        if config.USE_FAUCET:
            await request_faucet(account.address, config.REQUEST_AMOUNT * 10)

    @allure.step
    async def create(self):
        return await self.create_with_specific_id('')
