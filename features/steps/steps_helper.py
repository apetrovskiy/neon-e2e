import allure
from src.constants.eth_constants import ETHER
from src.helpers.web3.web3_helper import get_web3

w3 = get_web3()

TEST_DATA = "test_data"
ASYNC_CONTEXT = "async_context1"


@allure.step("getting balance")
async def get_balance(address: str) -> int:
    balance = w3.eth.get_balance(address)
    return balance


@allure.step("converting balance into Ethers")
def get_ethers_amount(balance: int):
    ethers_amount = w3.fromWei(int(balance), ETHER)
    print(f"user's' balance = {balance}")
    print(f"user's'balance in ETH = {ethers_amount}")
    return ethers_amount
