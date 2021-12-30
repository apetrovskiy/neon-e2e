from behave.runner import Context
from features.steps.steps_helper import TEST_DATA
from features.steps.test_data import TestData

data = TestData(None, None, 0, 0)


def before_scenario(scenario, context: Context):
    context.data = {}
    context.data[TEST_DATA] = data
