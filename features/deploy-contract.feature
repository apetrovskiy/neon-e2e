Feature: deploy contract

  # Scenario Outline: a user to user successful transaction
  Scenario: simple contract deployment
    Given there is a contract '1_Storage.sol'
    When compiling the contract
    Then there is no errors

    # Examples:
    #   | initial balance | amount |
    #   | 10              | 1      |
    #   | 10              | 2      |
    #   | 10              | 9      |
