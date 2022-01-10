from behave.model import Scenario
from behave.runner import Context
from behave.api.async_step import async_run_until_complete
from behave.api.async_step import use_or_create_async_context, AsyncContext
from features.steps.steps_helper import ASYNC_CONTEXT, TEST_DATA
from features.steps.test_data import TestData

data = TestData(None, None, 0, 0)


@async_run_until_complete
async def before_scenario(context: Context, scenario: Scenario):
    async_context: AsyncContext = use_or_create_async_context(
        context, ASYNC_CONTEXT)
    async_context.data = {}
    async_context.data[TEST_DATA] = data
