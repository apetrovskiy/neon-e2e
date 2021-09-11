from behave import given, then, when
from web3 import Web3


@given(u'there is user Alice in Ethereum network')
def step_impl(context):
    # raise NotImplementedError(u'STEP: Given there is user Alice in Ethereum network')
    pass


@given(u'there is user Bob in Ethereum network')
def step_impl(context):
    # raise NotImplementedError(u'STEP: Given there is user Bob in Ethereum network')
    pass


@when(u'user Alice sends (.*) Eth to user Bob')
def step_impl(context, eth_number):
    print(f"Eth number = {eth_number}")
    # raise NotImplementedError(u'STEP: When user Alice sends Eth to user Bob')
    pass


@then(u'the recipient has balance increased')
def step_impl(context):
    # raise NotImplementedError(
        # u'STEP: Then the recipient has balance increased')
    pass


@then(u'the sender has balance decreased')
def step_impl(context):
    # raise NotImplementedError(u'STEP: Then the sender has balance decreased')
    pass
