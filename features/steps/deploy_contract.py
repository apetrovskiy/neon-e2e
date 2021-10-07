from behave import given, step, then, when
# from behave.api.async_step import async_run_until_complete
# import asyncio


@step("0010")
# @async_run_until_complete
@given(u'there is a contract \'{contract_file}\'')
# async
def step_there_is_a_contract(context, contract_file: str):
    print(contract_file)
    # await asyncio.sleep(5)
    # raise NotImplementedError(
    #     u'STEP: Given there is a contract \'{contract_file}\'')


@step("0011")
# @async_run_until_complete
@given(u'the contract is compiled')
# async
def step_the_contract_is_compiled(context):
    print(context)
    # await asyncio.sleep(5)
    # raise NotImplementedError(u'STEP: Given the contract is compiled')


@step("0012")
# @async_run_until_complete
@when(u'the contract is deployed')
# async
def step_the_contract_is_deployed(context):
    print(context)
    # raise NotImplementedError(u'STEP: When the contract is deployed')


@step("0014")
# @async_run_until_complete
@then(u'there are no errors')
# async
def step_there_are_no_errors(context):
    print(context)
    # raise NotImplementedError(u'STEP: Then there are no errors')
