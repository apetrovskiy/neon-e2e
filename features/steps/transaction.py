import allure
import asyncio
from behave import given, step, then, when
from behave.api.async_step import async_run_until_complete
from behave.api.async_step import use_or_create_async_context, AsyncContext
from behave.runner import Context
from features.steps.steps_helper import ASYNC_CONTEXT, TEST_DATA, get_ethers_amount
from features.steps.test_data import TestData
from src.constants.eth_constants import ETHER
from src.helpers.account_factory import AccountFactory
from src.helpers.web3.web3_helper import get_web3

w3 = get_web3()


@given(u'there is user Alice in Ethereum network ' +
       u'with the initial balance {initial_balance}Ξ')
@async_run_until_complete
async def step_user_alice_initial_balance(context: Context,
                                          initial_balance: str):
    async_context = use_or_create_async_context(context, ASYNC_CONTEXT)
    async_context.data[TEST_DATA].user_alice = await AccountFactory().create(
        initial_balance)
    print(type(async_context.data[TEST_DATA].user_alice))
    '''
    # print(f"user A: {async_context.data[TEST_DATA].user_alice._address}")
    # balance = w3.eth.get_balance(async_context.data[TEST_DATA].user_alice.address)
    balance = initial_balance

    ethers_amount = w3.fromWei(int(balance), ETHER)
    print(f"user A balance = {balance}")
    print(f"user A balance in ETH = {ethers_amount}")
    async_context.data[TEST_DATA].initial_balance_alice = balance
    '''
    async_context.data[TEST_DATA].initial_balance_alice = initial_balance
    ethers_amount = get_ethers_amount(initial_balance)

    assert initial_balance == str(
        ethers_amount
    ), f"Initial balance {initial_balance} does not equal Ethers amount {ethers_amount}"


@given(u'there is user Bob in Ethereum network ' +
       u'with the initial balance {initial_balance}Ξ')
@async_run_until_complete
async def step_user_bob_initial_balance(context: Context,
                                        initial_balance: str):
    async_context = context.async_context1
    async_context.data[TEST_DATA].user_bob = await AccountFactory().create(
        initial_balance)
    '''
    # balance = w3.eth.get_balance(context.data[TEST_DATA].user_bob.address)
    balance = initial_balance

    ethers_amount = w3.fromWei(int(balance), ETHER)
    print(f"user B balance = {balance}")
    print(f"user B balance in ETH = {ethers_amount}")
    context.data[TEST_DATA].initial_balance_bob = balance
    '''
    async_context.data[TEST_DATA].initial_balance_bob = initial_balance
    ethers_amount = get_ethers_amount(initial_balance)

    assert initial_balance == str(
        ethers_amount
    ), f"Initial balance {initial_balance} does not equal Ethers amount {ethers_amount}"


@when(u'user Alice sends {eth_number}Ξ to user Bob')
@async_run_until_complete
async def step_transaction(context: Context, eth_number: str):
    async_context = context.async_context1

    print(f"Attempting to send {eth_number}Ξ from \
        {async_context.data[TEST_DATA].user_alice.address} to \
          {async_context.data[TEST_DATA].user_bob.address}")

    print(f"value = {w3.toWei(eth_number, ETHER)}")
    print(f"gas price = {w3.eth.gas_price}")
    print(
        f"private key = {async_context.data[TEST_DATA].user_alice.privateKey}")
    print(f"coinbase = {async_context.data[TEST_DATA].user_alice.address}")
    # print(
    #     f"nonce = {w3.eth.get_transaction_count(async_context.data[TEST_DATA].user_alice.address)}"
    # )
    nonce = w3.eth.get_transaction_count(
        async_context.data[TEST_DATA].user_alice.address)
    print(f"nonce = {nonce}")

    # original version
    '''
    txn = dict(nonce=w3.eth.get_transaction_count(
        async_context.data[TEST_DATA].user_alice.address),
               to=str(async_context.data[TEST_DATA].user_bob.address),
               value=w3.toWei(eth_number, ETHER),
               gas=4294967295,
               gasPrice=w3.eth.gas_price,
               data=b'')
    '''

    txn = dict(
        nonce=nonce,
        # w3.eth.get_transaction_count(
        # w3.eth.coinbase), # no supported
        # async_context.data[TEST_DATA].user_alice.address),
        to=async_context.data[TEST_DATA].user_bob.address,
        #  value=w3.toWei(eth_number, ETHER),
        # value=eth_number,
        # value=1,
        # value=w3.toWei(50, 'gwei'),
        value=w3.toWei(eth_number, 'ether'),
        # maxFeePerGas=2000000000,
        # maxPriorityFeePerGas=1000000000,
        gas=4294967295,
        maxFeePerGas=2000000000,
        maxPriorityFeePerGas=1000000000,
        # 100000,
        # 4294967295,
        # gasPrice=w3.toWei(50, 'gwei'),
        # w3.eth.gas_price,
        # type=2,
        chainId=w3.eth.chain_id,
        data=b'')
    print(f"transaction: {txn}")

    try:
        print("Trying to sign transaction...")
        print(f"{async_context.data[TEST_DATA].user_alice.privateKey}")

        signed_txn = w3.eth.account.sign_transaction(
            txn,
            private_key=async_context.data[TEST_DATA].user_alice.privateKey)

        print(f"signed txn: {signed_txn}")

        # Deploy transaction
        create_receipt = w3.eth.send_raw_transaction(signed_txn.rawTransaction)

        print(f"Transaction successful with hash: \
            {create_receipt.transactionHash}")

    except Exception as e:
        print(e)
        assert 1 == 2, "Failed to send a raw transaction"

    print('when is finished')


@then(u'the recipient has balance increased by {eth_number}Ξ')
@async_run_until_complete
async def step_user_bob_result(context: Context, eth_number: str):
    async_context = context.async_context1
    balance = w3.eth.get_balance(
        async_context.data[TEST_DATA].user_bob.address)
    expected_balance = async_context.data[
        TEST_DATA].initial_balance_bob + w3.toWei(eth_number, ETHER)

    print(
        f"initial user B balance {async_context.data[TEST_DATA].initial_balance_bob}"
    )
    print(f"amount {w3.toWei(eth_number, ETHER)}")
    print(f"expected user B balance = {expected_balance}")
    print(f"resulting user B balance {balance}")

    assert balance == expected_balance, \
        f"The expected balance {expected_balance} does not equal the actual balance {balance}"


@then(u'the sender has balance decreased more than by {eth_number}Ξ')
@async_run_until_complete
async def step_user_alice_result(context: Context, eth_number):
    async_context = context.async_context1
    balance = w3.eth.get_balance(
        async_context.data[TEST_DATA].user_alice.address)
    expected_balance = async_context.data[
        TEST_DATA].initial_balance_alice - w3.toWei(eth_number, ETHER)

    print(
        f"initial user A balance {async_context.data[TEST_DATA].initial_balance_alice}"
    )
    print(f"amount {w3.toWei(eth_number, ETHER)}")
    print(f"expected user A balance = {expected_balance}")
    print(f"resulting user A balance {balance}")

    assert balance <= expected_balance, \
        f"The expected balance {expected_balance} is not less than the actual balance {balance}"


@then(u'the sender has balance decreased by {eth_number}Ξ')
@async_run_until_complete
async def step_user_alice_no_changes(context: Context, eth_number):
    async_context = context.async_context1
    balance = w3.eth.get_balance(
        async_context.data[TEST_DATA].user_alice.address)

    print(
        f"initial user A balance {async_context.data[TEST_DATA].initial_balance_alice}"
    )
    print(f"amount {w3.toWei(eth_number, ETHER)}")
    print(f"resulting user A balance {balance}")

    expected_balance = async_context.data[
        TEST_DATA].initial_balance_alice - get_web3.toWei(eth_number, ETHER)
    assert balance == expected_balance, \
        f"The expected balance {expected_balance} does not equal actual balance {balance}"
