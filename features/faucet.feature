Feature: faucet

  Scenario Outline: the first funding request
    Given there is user Alice in Ethereum network with no initial balance
    When the user requests the Ether faucet for <initial balance>Ξ
    Then the balance equals <initial balance>Ξ

    Examples:
      | initial balance |
      | 1               |
      | 10              |

  Scenario Outline: the second funding request
    Given there is user Alice in Ethereum network with no initial balance
    When the user requests the Ether faucet for <first amount>Ξ
    And the user requests the Ether faucet for <second amount>Ξ
    Then the balance equals <total>Ξ

    Examples:
      | first amount | second amount | total |
      | 2            | 3             | 5     |
      | 10           | 10            | 20    |
