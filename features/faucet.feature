Feature: faucet

  Scenario Outline: the first funding request
    Given there is user Alice in Ethereum network with no initial balance
    When the user requests the ERC20 faucet for <initial balance>Ξ
    Then the balance equals <initial balance>Ξ

    Examples:
      | initial balance |
      | 1               |
      | 10              |

  Scenario Outline: the second funding request
    Given there is user Alice in Ethereum network with the initial balance <initial balance>Ξ
    When the user requests the ERC20 faucet for <amount>Ξ
    Then the balance equals <total>Ξ

    Examples:
      | initial balance | amount | total |
      | 2               | 3      | 5     |
      | 10              | 10     | 20    |
