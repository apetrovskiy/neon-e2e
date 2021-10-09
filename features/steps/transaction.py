from config import config
from dataclasses import dataclass
from behave import given, step, then, when
# from behave.api.async_step import async_run_until_complete
# import asyncio
from src.helpers.account_factory import AccountFactory
from web3 import Web3

ETHER = "ether"


def get_web3():
    return Web3(Web3.HTTPProvider(config.PROXY_URL))


@dataclass
class Data(object):
    user_alice: object
    user_bob: object
    initial_balance_alice: int
    initial_balance_bob: int


data = Data(None, None, 0, 0)
w3 = get_web3()


@step("0001")
# @async_run_until_complete
@given(u'there is user Alice in Ethereum network ' +
       u'with the initial balance {initial_balance}Ξ')
# async
def step_user_alice_initial_balance(context, initial_balance: str):
    data.user_alice = AccountFactory().create()
    print(f"user A: {data.user_alice.address}")
    balance = w3.eth.get_balance(data.user_alice.address)
    ethers_amount = w3.fromWei(balance, ETHER)
    print(f"user A balance = {balance}")
    print(f"user A balance in ETH = {ethers_amount}")
    data.initial_balance_alice = balance
    assert initial_balance == str(ethers_amount)


@step("0002")
# @async_run_until_complete
@given(u'there is user Bob in Ethereum network ' +
       u'with the initial balance {initial_balance}Ξ')
# async
def step_user_bob_initial_balance(context, initial_balance: str):
    data.user_bob = AccountFactory().create()
    print(f"user B: {data.user_bob.address}")
    balance = w3.eth.get_balance(data.user_bob.address)
    ethers_amount = w3.fromWei(balance, ETHER)
    print(f"user B balance = {balance}")
    print(f"user B balance in ETH = {ethers_amount}")
    data.initial_balance_bob = balance
    assert initial_balance == str(ethers_amount)


@step("0003")
# @async_run_until_complete
@when(u'user Alice sends {eth_number}Ξ to user Bob')
# async
def step_transaction(context, eth_number: str):
    print(f"Attempting to send {eth_number}Ξ from \
        {data.user_alice.address} to {data.user_bob.address}")

    print(f"value = {w3.toWei(eth_number, ETHER)}")
    print(f"gas price = {w3.eth.gas_price}")
    print(f"private key = {data.user_alice.privateKey}")
    print(f"coinbase = {data.user_alice.address}")
    print(f"nonce = {w3.eth.get_transaction_count(data.user_alice.address)}")
    txn = dict(nonce=w3.eth.get_transaction_count(data.user_alice.address),
               to=str(data.user_bob.address),
               value=w3.toWei(eth_number, ETHER),
               gas=4294967295,
               gasPrice=w3.eth.gas_price,
               data=b'')
    print(f"transaction: {txn}")

    try:
        signed_txn: str = w3.eth.signTransaction(
            txn, str(data.user_alice.privateKey))

        # Deploy transaction
        create_receipt = w3.eth.send_raw_transaction(signed_txn.rawTransaction)

        print(f"Transaction successful with hash: \
            {create_receipt.transactionHash}")

    except Exception as e:
        print(e)

    print('when is finished')


@step("0004")
# @async_run_until_complete
@then(u'the recipient has balance increased by {eth_number}Ξ')
# async
def step_user_bob_result(context, eth_number: str):
    balance = w3.eth.get_balance(data.user_bob.address)

    print(f"initial user B balance {data.initial_balance_bob}")
    print(f"amount {w3.toWei(eth_number, ETHER)}")
    print(f"resulting user B balance {balance}")

    expected_balance = data.initial_balance_bob + w3.toWei(eth_number, ETHER)
    assert balance == expected_balance


@step("0005")
# @async_run_until_complete
@then(u'the sender has balance decreased more than by {eth_number}Ξ')
# async
def step_user_alice_result(context, eth_number):
    balance = w3.eth.get_balance(data.user_alice.address)

    print(f"initial user A balance {data.initial_balance_alice}")
    print(f"amount {w3.toWei(eth_number, ETHER)}")
    print(f"resulting user A balance {balance}")

    expected_balance = data.initial_balance_alice - w3.toWei(eth_number, ETHER)
    assert balance <= expected_balance


@step("0006")
# @async_run_until_complete
@then(u'the sender has balance decreased by {eth_number}Ξ')
# async
def step_user_alice_no_changes(context, eth_number):
    balance = w3.eth.get_balance(data.user_alice.address)

    print(f"initial user A balance {data.initial_balance_alice}")
    print(f"amount {w3.toWei(eth_number, ETHER)}")
    print(f"resulting user A balance {balance}")

    expected_balance = data.initial_balance_alice - get_web3.toWei(
        eth_number, ETHER)
    assert balance == expected_balance
