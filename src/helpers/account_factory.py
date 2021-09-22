from config import config
from web3 import Web3, HTTPProvider


class AccountFactory():
    def create_with_specific_id(self, id: str):
        url = config.HTTP_URL
        w3 = Web3(HTTPProvider(url))
        print(w3)
        if id is None or id == '':
            return w3.eth.account.create()
        else:
            w3.eth.account.create(id)

    def create(self):
        return self.create_with_specific_id('')
