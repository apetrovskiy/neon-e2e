# 📗 transaction-steps


11 steps defined.

## 📍 Given

#### - there is user Alice in Ethereum network with no initial balance

#### - there is user Alice in Ethereum network with the initial balance {int}Ξ

##### Parameters:

|  #  | Name | Type |
| --- | ---- | ---- |
| 1 | initialBalance | number |

#### - there is user Bob in Ethereum network with no initial balance

#### - there is user Bob in Ethereum network with the initial balance {int}Ξ

##### Parameters:

|  #  | Name | Type |
| --- | ---- | ---- |
| 1 | initialBalance | number |

#### - user Bob requests the Ether faucet for {int}Ξ

##### Parameters:

|  #  | Name | Type |
| --- | ---- | ---- |
| 1 | amount | number |

## 🎬 When

#### - user Alice requests the Ether faucet for {int}Ξ

##### Parameters:

|  #  | Name | Type |
| --- | ---- | ---- |
| 1 | amount | number |

#### - user Alice sends {int}Ξ to user Bob

##### Parameters:

|  #  | Name | Type |
| --- | ---- | ---- |
| 1 | ethNumber | number |

## ✅ Then

#### - user Alice's balance equals {int}Ξ

##### Parameters:

|  #  | Name | Type |
| --- | ---- | ---- |
| 1 | amount | number |

#### - the recipient has balance increased by {int}Ξ

##### Parameters:

|  #  | Name | Type |
| --- | ---- | ---- |
| 1 | ethNumber | number |

#### - the sender has balance decreased by {int}Ξ

##### Parameters:

|  #  | Name | Type |
| --- | ---- | ---- |
| 1 | ethNumber | number |

#### - the sender has balance decreased more than by {int}Ξ

##### Parameters:

|  #  | Name | Type |
| --- | ---- | ---- |
| 1 | ethNumber | number |
