from behave import given, step, then, when
from behave.runner import Context
from features.steps.steps_helper import TEST_DATA, get_balance
from features.steps.test_data import TestData
from src.helpers.account_factory import AccountFactory
from src.helpers.faucet.faucet_requester import request_faucet
from src.helpers.web3.web3_helper import get_web3

w3 = get_web3()


@given(u'there is user Alice in Ethereum network with no initial balance')
async def step_user_alice_no_initial_balance(context: Context):
    print("no initial balance")
    context.data[TEST_DATA].user_alice = AccountFactory().create(0)


@given(u'there is user Bob in Ethereum network with no initial balance')
async def step_user_bob_no_initial_balance(context):
    print("no initial balance")
    context.data[TEST_DATA].user_bob = AccountFactory().create(0)


@given(u'user Alice requests the Ether faucet for {eth_number}Ξ')
@when(u'user Alice requests the Ether faucet for {eth_number}Ξ')
async def step_user_alice_requests_the_token_faucet(context: Context,
                                                    eth_number: str):
    await request_faucet(context.data[TEST_DATA].user_alice.address,
                         eth_number)


@given(u'user Bob requests the Ether faucet for {eth_number}Ξ')
async def step_user_bob_requests_the_token_faucet(context: Context,
                                                  eth_number: str):
    await request_faucet(context.data[TEST_DATA].user_bob.address, eth_number)


@then(u'user Alice\'s balance equals {eth_number}Ξ')
async def step_balance_equals(context: Context, eth_number: str):
    # balance = await w3.eth.get_balance(context.data[TEST_DATA].user_alice.address)
    balance = await get_balance(context.data[TEST_DATA].user_alice.address)
    assert balance == eth_number, f"Actual balance of {balance} does not equal Eth number {eth_number}"


@then(u'user Bob\'s balance equals {eth_number}Ξ')
async def step_impl(context: Context, eth_number: str):
    balance = await get_balance(context.data[TEST_DATA].user_bob.address)
    assert balance == eth_number, f"Actual balance of {balance} does not equal Eth number {eth_number}"
