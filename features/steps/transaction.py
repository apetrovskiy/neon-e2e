import allure
from behave import given, step, then, when
from behave.runner import Context
from features.steps.steps_helper import TEST_DATA, get_ethers_amount
from features.steps.test_data import TestData
from src.constants.eth_constants import ETHER
# from behave.api.async_step import async_run_until_complete
# import asyncio
from src.helpers.account_factory import AccountFactory
from src.helpers.web3.web3_helper import get_web3

w3 = get_web3()


# @async_run_until_complete
@given(u'there is user Alice in Ethereum network ' +
       u'with the initial balance {initial_balance}Ξ')
async def step_user_alice_initial_balance(context: Context,
                                          initial_balance: str):
    context.data[TEST_DATA].user_alice = AccountFactory().create(
        initial_balance)
    print(type(context.data[TEST_DATA].user_alice))
    '''
    # print(f"user A: {context.data[TEST_DATA].user_alice._address}")
    # balance = w3.eth.get_balance(context.data[TEST_DATA].user_alice.address)
    balance = initial_balance

    ethers_amount = w3.fromWei(int(balance), ETHER)
    print(f"user A balance = {balance}")
    print(f"user A balance in ETH = {ethers_amount}")
    context.data[TEST_DATA].initial_balance_alice = balance
    '''
    # context.data[TEST_DATA].initial_balance_alice = set_balance(context.data[TEST_DATA].user_alice.address)
    context.data[TEST_DATA].initial_balance_alice = initial_balance
    ethers_amount = get_ethers_amount(initial_balance)

    assert initial_balance == str(ethers_amount)


# @async_run_until_complete
@given(u'there is user Bob in Ethereum network ' +
       u'with the initial balance {initial_balance}Ξ')
async def step_user_bob_initial_balance(context: Context,
                                        initial_balance: str):
    context.data[TEST_DATA].user_bob = AccountFactory().create(initial_balance)
    # print(f"user B: {context.data[TEST_DATA].user_bob._address}")
    '''
    # balance = w3.eth.get_balance(context.data[TEST_DATA].user_bob.address)
    balance = initial_balance

    ethers_amount = w3.fromWei(int(balance), ETHER)
    print(f"user B balance = {balance}")
    print(f"user B balance in ETH = {ethers_amount}")
    context.data[TEST_DATA].initial_balance_bob = balance
    '''
    # context.data[TEST_DATA].initial_balance_bob = set_balance(context.data[TEST_DATA].user_bob.address)
    context.data[TEST_DATA].initial_balance_bob = initial_balance
    ethers_amount = get_ethers_amount(initial_balance)

    assert initial_balance == str(
        ethers_amount
    ), f"Initial balance {initial_balance} does not equal Ethers amount {ethers_amount}"


# @async_run_until_complete
@when(u'user Alice sends {eth_number}Ξ to user Bob')
async def step_transaction(context: Context, eth_number: str):
    print(f"Attempting to send {eth_number}Ξ from \
        {context.data[TEST_DATA].user_alice.address} to \
          {context.data[TEST_DATA].user_bob.address}")

    print(f"value = {w3.toWei(eth_number, ETHER)}")
    print(f"gas price = {w3.eth.gas_price}")
    print(f"private key = {context.data[TEST_DATA].user_alice.privateKey}")
    print(f"coinbase = {context.data[TEST_DATA].user_alice.address}")
    print(
        f"nonce = {w3.eth.get_transaction_count(context.data[TEST_DATA].user_alice.address)}"
    )
    txn = dict(nonce=w3.eth.get_transaction_count(
        context.data[TEST_DATA].user_alice.address),
               to=str(context.data[TEST_DATA].user_bob.address),
               value=w3.toWei(eth_number, ETHER),
               gas=4294967295,
               gasPrice=w3.eth.gas_price,
               data=b'')
    print(f"transaction: {txn}")

    try:
        signed_txn: str = w3.eth.signTransaction(
            txn, str(context.data[TEST_DATA].user_alice.privateKey))

        # Deploy transaction
        create_receipt = w3.eth.send_raw_transaction(signed_txn.rawTransaction)

        print(f"Transaction successful with hash: \
            {create_receipt.transactionHash}")

    except Exception as e:
        print(e)

    print('when is finished')


# @async_run_until_complete
@then(u'the recipient has balance increased by {eth_number}Ξ')
async def step_user_bob_result(context: Context, eth_number: str):
    balance = w3.eth.get_balance(context.data[TEST_DATA].user_bob.address)

    print(
        f"initial user B balance {context.data[TEST_DATA].initial_balance_bob}"
    )
    print(f"amount {w3.toWei(eth_number, ETHER)}")
    print(f"resulting user B balance {balance}")

    expected_balance = context.data[TEST_DATA].initial_balance_bob + w3.toWei(
        eth_number, ETHER)
    assert balance == expected_balance


# @async_run_until_complete
@then(u'the sender has balance decreased more than by {eth_number}Ξ')
async def step_user_alice_result(context: Context, eth_number):
    balance = w3.eth.get_balance(context.data[TEST_DATA].user_alice.address)

    print(
        f"initial user A balance {context.data[TEST_DATA].initial_balance_alice}"
    )
    print(f"amount {w3.toWei(eth_number, ETHER)}")
    print(f"resulting user A balance {balance}")

    expected_balance = context.data[
        TEST_DATA].initial_balance_alice - w3.toWei(eth_number, ETHER)
    assert balance <= expected_balance


# @async_run_until_complete
@then(u'the sender has balance decreased by {eth_number}Ξ')
async def step_user_alice_no_changes(context: Context, eth_number):
    balance = w3.eth.get_balance(context.data[TEST_DATA].user_alice.address)

    print(
        f"initial user A balance {context.data[TEST_DATA].initial_balance_alice}"
    )
    print(f"amount {w3.toWei(eth_number, ETHER)}")
    print(f"resulting user A balance {balance}")

    expected_balance = context.data[
        TEST_DATA].initial_balance_alice - get_web3.toWei(eth_number, ETHER)
    assert balance == expected_balance
