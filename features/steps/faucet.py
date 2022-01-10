import asyncio
from behave import given, step, then, when
from behave.api.async_step import async_run_until_complete
from behave.api.async_step import use_or_create_async_context, AsyncContext
from behave.runner import Context
from features.steps.steps_helper import ASYNC_CONTEXT, TEST_DATA, get_balance, get_ethers_amount
from features.steps.test_data import TestData
from src.helpers.account_factory import AccountFactory
from src.helpers.faucet.faucet_requester import request_faucet
from src.helpers.web3.web3_helper import get_web3

w3 = get_web3()


@given(u'there is user Alice in Ethereum network with no initial balance')
@async_run_until_complete
async def step_user_alice_no_initial_balance(context: Context):
    async_context = use_or_create_async_context(context, ASYNC_CONTEXT)
    async_context.data[TEST_DATA].user_alice = await AccountFactory().create(0)


@given(u'there is user Bob in Ethereum network with no initial balance')
@async_run_until_complete
async def step_user_bob_no_initial_balance(context):
    async_context = context.async_context1
    async_context.data[TEST_DATA].user_bob = await AccountFactory().create(0)


@given(u'user Alice requests the Ether faucet for {amount}Ξ')
@when(u'user Alice requests the Ether faucet for {amount}Ξ')
@async_run_until_complete
async def step_user_alice_requests_the_token_faucet(context: Context,
                                                    amount: int):
    async_context = context.async_context1
    await request_faucet(async_context.data[TEST_DATA].user_alice.address,
                         amount)
    balance = w3.eth.get_balance(
        async_context.data[TEST_DATA].user_alice.address)
    async_context.data[TEST_DATA].initial_balance_alice = amount


@given(u'user Bob requests the Ether faucet for {amount}Ξ')
@async_run_until_complete
async def step_user_bob_requests_the_token_faucet(context: Context,
                                                  amount: str):
    async_context = context.async_context1
    await request_faucet(async_context.data[TEST_DATA].user_bob.address,
                         amount)
    balance = w3.eth.get_balance(
        async_context.data[TEST_DATA].user_bob.address)
    async_context.data[TEST_DATA].initial_balance_bob = amount


@given(u'user Alice\'s balance equals {amount}Ξ')
@then(u'user Alice\'s balance equals {amount}Ξ')
@async_run_until_complete
async def step_balance_user_a_equals(context: Context, amount: int):
    async_context = context.async_context1
    compare_expected_and_actual_balance(
        async_context.data[TEST_DATA].user_alice.address, amount)


@given(u'user Bob\'s balance equals {amount}Ξ')
@then(u'user Bob\'s balance equals {amount}Ξ')
@async_run_until_complete
async def step_balance_user_b_equals(context: Context, amount: int):
    async_context = context.async_context1
    compare_expected_and_actual_balance(
        async_context.data[TEST_DATA].user_bob.address, amount)


async def compare_expected_and_actual_balance(address: str, amount: int):
    balance = await get_balance(address)
    balance_in_ethers = get_ethers_amount(balance)
    assert balance_in_ethers == int(
        amount
    ), f"Actual balance of {balance_in_ethers} does not equal Eth number {amount}"
