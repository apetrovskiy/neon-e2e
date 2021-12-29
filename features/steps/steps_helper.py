import allure
from src.constants.eth_constants import ETHER
from src.helpers.web3.web3_helper import get_web3

w3 = get_web3()


@allure.step("getting balance")
async def set_balance(address: str) -> int:
    balance = await w3.eth.get_balance(address)
    # balance = initial_balance
    return balance


@allure.step("converting balance into Ethers")
def get_ethers_amount(balance: int):
    ethers_amount = w3.fromWei(int(balance), ETHER)
    print(f"user B balance = {balance}")
    print(f"user B balance in ETH = {ethers_amount}")
    return ethers_amount
