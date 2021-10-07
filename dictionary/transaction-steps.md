# 📗 transaction-steps


6 steps defined.

## 📍 Given

#### - there is user Alice in Ethereum network with the initial balance {int}Ξ

##### Parameters:

|  #  | Name | Type |
| --- | ---- | ---- |
| 1 | initialBalance | number |

#### - there is user Bob in Ethereum network with the initial balance {int}Ξ

##### Parameters:

|  #  | Name | Type |
| --- | ---- | ---- |
| 1 | initialBalance | number |

## 🎬 When

#### - user Alice sends {int}Ξ to user Bob

##### Parameters:

|  #  | Name | Type |
| --- | ---- | ---- |
| 1 | ethNumber | number |

## ✅ Then

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
