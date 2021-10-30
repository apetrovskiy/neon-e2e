from behave import given, step, then, when
from behave.runner import Context


@step("0010")
@given(u'there is a contract \'{contract_file}\'')
def step_there_is_a_contract(context: Context, contract_file: str):
    print(contract_file)
    raise NotImplementedError(
        u'STEP: Given there is a contract \'{contract_file}\'')


@step("0011")
@given(u'the contract is compiled')
def step_the_contract_is_compiled(context: Context):
    print(context)
    raise NotImplementedError(u'STEP: Given the contract is compiled')


@step("0012")
@when(u'the contract is deployed')
def step_the_contract_is_deployed(context: Context):
    print(context)
    raise NotImplementedError(u'STEP: When the contract is deployed')


@step("0014")
@then(u'there are no errors')
def step_there_are_no_errors(context: Context):
    print(context)
    raise NotImplementedError(u'STEP: Then there are no errors')
