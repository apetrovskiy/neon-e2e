Feature: deploy contract

  Scenario Outline: simple contract deployment
    Given there is a contract '<contract file>'
    When compiling the contract
    Then there is no errors

    Examples:
      | contract file |
      | 1_Storage.sol |
      # | 2_Owner.sol   |
      | 3_Ballot.sol  |
      | proof.sol     |
