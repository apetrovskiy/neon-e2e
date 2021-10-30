"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const allure_cucumberjs_1 = require("allure-cucumberjs");
const directories_1 = require("src/constants/directories");
class AllureReporter extends allure_cucumberjs_1.CucumberJSAllureFormatter {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any,@typescript-eslint/explicit-module-boundary-types
    constructor(options) {
        const allureRuntime = new allure_cucumberjs_1.AllureRuntime({ resultsDir: directories_1.ALLURE_REPORT_DIR });
        super(options, allureRuntime, {
            labels: {
                epic: [/@epic=(.*)/],
                issue: [/@issue=(.*)/]
            }
        });
        const properties = AllureReporter.buildEnvProperties();
        allureRuntime.writeEnvironmentInfo(properties);
    }
    static buildEnvProperties() {
        return Object.assign(Object.assign({}, AllureReporter.addEnvironment()), AllureReporter.addBranch());
    }
    static addEnvironment() {
        const targetHost = process.env.TARGET_HOST || 'local';
        return { Environment: targetHost };
    }
    static addBranch() {
        const branch = process.env.BRANCH;
        return branch ? { Branch: branch } : undefined;
    }
}
exports.default = AllureReporter;
//# sourceMappingURL=allure-reporter.js.map