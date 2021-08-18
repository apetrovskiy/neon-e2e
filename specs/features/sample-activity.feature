Feature: sample e2e

    Scenario: a new account's activity
        Given there is user Alice in Ethereum
        And there is user Bob in Ethereum
        When user A sends Eth to user B
        Then the recipient has balance increased
        And the sender has balance decreased
