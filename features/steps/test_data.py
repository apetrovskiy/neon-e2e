from dataclasses import dataclass
from web3 import Account


@dataclass
class TestData(object):
    user_alice: Account
    user_bob: Account
    initial_balance_alice: int
    initial_balance_bob: int
