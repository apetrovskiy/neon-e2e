Feature: sample e2e

  # Scenario: a new account's activity
  #   Given there is user Alice in Ethereum network
  #   And there is user Bob in Ethereum network
  #   When user Alice sends 1Ξ to user Bob
  #   Then the recipient has balance increased
  #   And the sender has balance decreased

  Scenario Outline: a user to user successful transaction
    Given there is user Alice in Ethereum network
    And there is user Bob in Ethereum network
    When user Alice sends <amount>Ξ to user Bob
    Then the recipient has balance increased
    And the sender has balance decreased

    Examples:
      | initial balance | amount |
      | 10              | 1      |
      | 10              | 2      |
      | 10              | 9      |
