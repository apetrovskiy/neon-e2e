from behave import given, then, when
from web3 import Web3


@given(u'there is user Alice in Ethereum network')
def step_impl1(context):
    # raise NotImplementedError(u'STEP: Given there is user Alice in Ethereum network')
    pass


@given(u'there is user Bob in Ethereum network')
def step_impl2(context):
    # raise NotImplementedError(u'STEP: Given there is user Bob in Ethereum network')
    pass


@when(u'user Alice sends {eth_number}Ξ to user Bob')
def step_impl3(context, eth_number):
    print(f"Eth number = {eth_number}")
    # raise NotImplementedError(u'STEP: When user Alice sends Eth to user Bob')
    pass


@then(u'the recipient has balance increased')
def step_impl4(context):
    # raise NotImplementedError(
        # u'STEP: Then the recipient has balance increased')
    pass


@then(u'the sender has balance decreased')
def step_impl5(context):
    # raise NotImplementedError(u'STEP: Then the sender has balance decreased')
    pass
