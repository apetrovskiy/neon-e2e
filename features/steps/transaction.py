from config import config
from behave import given, then, when
from src.helpers.account_factory import AccountFactory
from web3 import Web3


def get_web3():
    return Web3(Web3.HTTPProvider(config.HTTP_URL))


user_alice: object
user_bob: object
global initial_balance_alice
global initial_balance_bob


@given(
    u'there is user Alice in Ethereum network with the initial balance {initial_balance}Ξ'
)
def step_user_alice_initial_balance(context, initial_balance):
    print(f"user A the initial balance expected: {initial_balance}")
    user_alice = AccountFactory().create()
    balance = get_web3().eth.get_balance(user_alice.address)
    ethers_amount = get_web3().fromWei(balance, "ether")
    print(ethers_amount)
    initial_balance_alice = balance
    print(initial_balance_alice)
    print(type(initial_balance))
    print(type(ethers_amount))
    assert initial_balance == str(ethers_amount)


@given(
    u'there is user Bob in Ethereum network with the initial balance {initial_balance}Ξ'
)
def step_user_bob_initial_balance(context, initial_balance):
    print(f"user B the initial balance expected: {initial_balance}")
    user_bob = AccountFactory().create()
    balance = get_web3().eth.get_balance(user_bob.address)
    ethers_amount = get_web3().fromWei(balance, "ether")
    print(ethers_amount)
    initial_balance_bob = balance
    print(initial_balance_bob)
    print(type(initial_balance))
    print(type(ethers_amount))
    assert initial_balance == str(ethers_amount)


@when(u'user Alice sends {eth_number}Ξ to user Bob')
def step_transaction(context, eth_number):
    print(f"Eth number = {eth_number}")
    # raise NotImplementedError(u'STEP: When user Alice sends Eth to user Bob')
    pass


@then(u'the recipient has balance increased by {eth_number}Ξ')
def step_user_bob_result(context, eth_number):
    # raise NotImplementedError(
    # u'STEP: Then the recipient has balance increased')
    pass


@then(u'the sender has balance decreased more than by {eth_number}Ξ')
def step_user_alice_result(context, eth_number):
    # raise NotImplementedError(u'STEP: Then the sender has balance decreased')
    pass


@then(u'the sender has balance decreased by {eth_number}Ξ')
def step_user_alice_no_changes(context, eth_number):
    # raise NotImplementedError(u'STEP: Then the sender has balance decreased')
    pass
