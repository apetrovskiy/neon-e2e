Feature: transaction

  Scenario Outline: a user to user successful transaction
    Given there is user Alice in Ethereum network with the initial balance <initial balance>Ξ
    And there is user Bob in Ethereum network with the initial balance <initial balance>Ξ
    When user Alice sends <amount>Ξ to user Bob
    Then the recipient has balance increased by <amount>Ξ
    And the sender has balance decreased more than by <amount>Ξ

    Examples:
      | initial balance | amount |
      | 10              | 1      |
      | 10              | 2      |
      | 10              | 9      |

  Scenario Outline: a user to user failed transaction
    Given there is user Alice in Ethereum network with the initial balance <initial balance>Ξ
    And there is user Bob in Ethereum network with the initial balance <initial balance>Ξ
    When user Alice sends <amount>Ξ to user Bob
    Then the recipient has balance increased by 0Ξ
    And the sender has balance decreased by 0Ξ

    Examples:
      | initial balance | amount |
      # | 10              | 0      |
      # | 10              | 0.1    |
      | 10              | 10     |
      | 10              | 11     |
