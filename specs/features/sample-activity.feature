Feature: sample e2e

    Scenario: a new account's activity
        Given there is a new user in Ethereum
        And there is another new user in Ethereum
        When user sends Eth
        Then the recipient has balance increased
        And the sender has balance decreased
