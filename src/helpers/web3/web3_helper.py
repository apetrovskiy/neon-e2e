from web3 import Web3
from config import config


def get_web3():
    return Web3(Web3.HTTPProvider(config.PROXY_URL))
