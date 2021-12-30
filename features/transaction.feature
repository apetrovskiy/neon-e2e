Feature: transaction

  Scenario Outline: a user to user successful transaction
    # Given there is user Alice in Ethereum network with the initial balance <initial balance>Ξ
    Given there is user Alice in Ethereum network with no initial balance
    And user Alice requests the Ether faucet for <initial balance>Ξ
    # And there is user Bob in Ethereum network with the initial balance <initial balance>Ξ
    And there is user Bob in Ethereum network with no initial balance
    And user Bob requests the Ether faucet for <initial balance>Ξ
    When user Alice sends <amount>Ξ to user Bob
    Then the recipient has balance increased by <amount>Ξ
    # And the sender has balance decreased more than by <amount>Ξ
    And the sender has balance decreased by <amount>Ξ

    Examples:
      | initial balance | amount |
      | 10              | 1      |
      | 10              | 2      |
      | 10              | 9      |

  Scenario Outline: a user to user failed transaction
    # Given there is user Alice in Ethereum network with the initial balance <initial balance>Ξ
    Given there is user Alice in Ethereum network with no initial balance
    And user Alice requests the Ether faucet for <initial balance>Ξ
    # And there is user Bob in Ethereum network with the initial balance <initial balance>Ξ
    And there is user Bob in Ethereum network with no initial balance
    And user Bob requests the Ether faucet for <initial balance>Ξ
    When user Alice sends <amount>Ξ to user Bob
    Then the recipient has balance increased by <amount>Ξ
    And the sender has balance decreased by <amount>Ξ
    # earlier transaction would be reverted
    # Then user Alice's balance equals <initial balance>Ξ
    # And user Bob's balance equals <initial balance>Ξ

    Examples:
      | initial balance | amount |
      # | 10              | 0      |
      # | 10              | 0.1    |
      | 10              | 10     |
      | 10              | 11     |
