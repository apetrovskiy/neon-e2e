from behave import given, step, then, when
from behave.runner import Context
from features.steps.steps_helper import set_balance
from features.steps.test_data import TestData
from src.helpers.account_factory import AccountFactory
from src.helpers.faucet.faucet_requester import request_faucet
from src.helpers.web3.web3_helper import get_web3

data = TestData(None, None, 0, 0)
w3 = get_web3()


@step("0007")
@given(u'there is user Alice in Ethereum network with no initial balance')
async def step_user_no_initial_balance(context: Context):
    print("no initial balance")
    data.user_alice = AccountFactory().create(0)


@step("0008")
@when(u'the user requests the Ether faucet for {eth_number}Ξ')
async def step_user_requests_the_token_faucet(context: Context,
                                              eth_number: str):
    await request_faucet(data.user_alice.address, eth_number)


@step("0009")
@then(u'the balance equals {eth_number}Ξ')
async def step_balance_equals(context: Context, eth_number: str):
    # balance = await w3.eth.get_balance(data.user_alice.address)
    balance = await set_balance(data.user_alice.address)
    assert balance == eth_number, f"Actual balance of {balance} does not equal Eth number {eth_number}"
