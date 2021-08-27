from behave import given, then, when


@given(u'there is user Alice in Ethereum')
def step_impl(context):
    raise NotImplementedError(u'STEP: Given there is user Alice in Ethereum')


@given(u'there is user Bob in Ethereum')
def step_impl(context):
    raise NotImplementedError(u'STEP: Given there is user Bob in Ethereum')


@when(u'user A sends Eth to user B')
def step_impl(context):
    raise NotImplementedError(u'STEP: When user A sends Eth to user B')


@then(u'the recipient has balance increased')
def step_impl(context):
    raise NotImplementedError(
        u'STEP: Then the recipient has balance increased')


@then(u'the sender has balance decreased')
def step_impl(context):
    raise NotImplementedError(u'STEP: Then the sender has balance decreased')
