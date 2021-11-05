from behave import given, step, then, when
from behave.runner import Context


@step("0007")
@given(u'there is user Alice in Ethereum network with no initial balance')
def step_user_no_initial_balance(context: Context):
    print("no initial balance")
    raise NotImplementedError(
        u'STEP: Given there is user Alice in Ethereum network with no initial balance'
    )


@step("0008")
@when(u'the user requests the ERC20 faucet for {eth_number}Ξ')
def step_user_requests_the_erc20_faucet(context: Context, eth_number: str):
    print("erc20")
    raise NotImplementedError(
        u'STEP: When the user requests the ERC20 faucet for 1Ξ')


@step("0009")
@then(u'the balance equals {eth_number}Ξ')
def step_balance_equals(context: Context, eth_number: str):
    print("equals")
    raise NotImplementedError(u'STEP: Then the balance equals 1Ξ')
