import { defineFeature, loadFeature, DefineStepFunction } from 'jest-cucumber';
// import { SeriesSolver } from '../../src/series-solver';

const feature = loadFeature('specs/features/more-scenario-outlines.feature');

defineFeature(feature, (test) => {
  // let solver: SeriesSolver;
  let solution: string;
  let terms: string[];
  let operator: string;

  beforeEach(() => {
    // solver = new SeriesSolver();
    console.log('BEFORE EACH');
  });

  const whenISolveTheSeries = (when: DefineStepFunction) => {
    when(/^I solve the series$/, () => {
      // solution = solver.solve(terms, operator);
      console.log('WHEN');
    });
  };

  const thenIShouldGetXAsTheAnswer = (then: DefineStepFunction) => {
    then(/^I should get (.*) as the answer$/, (expectedSolution) => {
      console.log('THEN');
      expect(solution).toBe(expectedSolution);
    });
  };

  test.skip('Solving series', ({ given, when, then }) => {
    given(
      /^I have a series (.*) (.*) (.*) (.*) (.*) (.*) (.*)$/,
      (firstTerm, firstOperator, secondTerm, secondOperator, thirdTerm, thirdOperator, forthTerm) => {
        console.log('GIVEN');
        expect(firstOperator).toEqual(secondOperator);
        expect(firstOperator).toEqual(thirdOperator);

        operator = firstOperator;
        terms = [firstTerm, secondTerm, thirdTerm, forthTerm];
      });

    whenISolveTheSeries(when);

    thenIShouldGetXAsTheAnswer(then);
  });

  // test('Adding series', ({ given, when, then }) => {
  //   given(
  //     /^I add the following series:$/,
  //     (table: [{ Series: string, Operator: string, Solution: string }]) => {
  //       const row = table[0];
  //       terms = row.Series.split(` ${row.Operator} `);
  //       operator = row.Operator;
  //       // solver.add(terms, operator, row.Solution);
  //     });

  //   whenISolveTheSeries(when);

  //   thenIShouldGetXAsTheAnswer(then);
  // });
});
