import { Before, BeforeAll, HookScenarioResult, setDefaultTimeout } from 'cucumber';
import { logger } from 'src/utils/logger';
import { decorate } from 'allure-decorators';
import { allure, MochaAllure } from 'allure-mocha/runtime';

BeforeAll(function () {
  const time = 30000;
  setDefaultTimeout(time);
  decorate<MochaAllure>(allure);
});

Before(function (scenario: HookScenarioResult) {
  const {
    pickle: { name }
  } = scenario;

  decorate<MochaAllure>(allure);

  logger.notice(`🧪 Starting Scenario: ${name}`);
});
