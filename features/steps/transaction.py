from behave import given, then, when
from web3 import Web3


@given(
    u'there is user Alice in Ethereum network with the initial balance {initial_balance}Ξ'
)
def step_impl1(context, initial_balance):
    # raise NotImplementedError(u'STEP: Given there is user Alice in Ethereum network')
    print(f"user A the initial balance expected: {initial_balance}")
    pass


@given(
    u'there is user Bob in Ethereum network with the initial balance {initial_balance}Ξ'
)
def step_impl2(context, initial_balance):
    # raise NotImplementedError(u'STEP: Given there is user Bob in Ethereum network')
    print(f"user B the initial balance expected: {initial_balance}")
    pass


@when(u'user Alice sends {eth_number}Ξ to user Bob')
def step_impl3(context, eth_number):
    print(f"Eth number = {eth_number}")
    # raise NotImplementedError(u'STEP: When user Alice sends Eth to user Bob')
    pass


@then(u'the recipient has balance increased by {eth_number}Ξ')
def step_impl4(context, eth_number):
    # raise NotImplementedError(
    # u'STEP: Then the recipient has balance increased')
    pass


@then(u'the sender has balance decreased more than by {eth_number}Ξ')
def step_impl5(context, eth_number):
    # raise NotImplementedError(u'STEP: Then the sender has balance decreased')
    pass


@then(u'the sender has balance decreased by {eth_number}Ξ')
def step_impl5(context, eth_number):
    # raise NotImplementedError(u'STEP: Then the sender has balance decreased')
    pass
