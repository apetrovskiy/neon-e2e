Feature: sample e2e

    Scenario: a new account's activity
        Given there is user Alice in Ethereum network
        And there is user Bob in Ethereum network
        When user Alice sends 1 Eth to user Bob
        Then the recipient has balance increased
        And the sender has balance decreased
