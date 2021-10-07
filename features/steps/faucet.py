from behave import given, step, then, when
# from behave.api.async_step import async_run_until_complete
# import asyncio


@step("0007")
# @async_run_until_complete(timeout=3.0)
@given(u'there is user Alice in Ethereum network with no initial balance')
# async
def step_user_no_initial_balance(context):
    print("no initial balance")
    # await asyncio.sleep(5)
    # raise NotImplementedError(
    #     u'STEP: Given there is user Alice in Ethereum network with no initial balance'
    # )


@step("0008")
# @async_run_until_complete(timeout=3.0)
@when(u'the user requests the ERC20 faucet for {eth_number}Ξ')
# async
def step_user_requests_the_erc20_faucet(context, eth_number: str):
    print("erc20")
    # await asyncio.sleep(5)
    # raise NotImplementedError(
    #     u'STEP: When the user requests the ERC20 faucet for 1Ξ')


@step("0009")
# @async_run_until_complete
@then(u'the balance equals {eth_number}Ξ')
# async
def step_balance_equals(context, eth_number: str):
    print("equals")
    # await asyncio.sleep(5)
    # raise NotImplementedError(u'STEP: Then the balance equals 1Ξ')
