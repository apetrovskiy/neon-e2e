from behave import given, then, when


@given(u'I have previously created a password')
def step_impl(context):
    raise NotImplementedError(
        u'STEP: Given I have previously created a password')


@when(u'I enter my password correctly')
def step_impl(context):
    raise NotImplementedError(u'STEP: When I enter my password correctly')


@then(u'I should be granted access')
def step_impl(context):
    raise NotImplementedError(u'STEP: Then I should be granted access')
