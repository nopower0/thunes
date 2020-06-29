# Get User Wallet API

### Introduction
URI: `/wallet/transfer`

**Login Is Required**

Description:  
This API is used to transfer SGD to another user


### Request

Parameters

|Name|Type|Description|Required?|Format|Sample|
|----|----|-----------|---------|------|------|
| to | int | the ID of the user who you want to transfer SGD to | Y | - | 2 |
| amount | int | the SGD amount you want to transfer | Y | - | 3 |

Sample
```json
{
    "to": 3,
    "amount": 3
}
```
### Response

Parameters

|Name|Type|Description|Must|Format|Sample|
|----|----|-----------|---------|------|------|
| wallet | WalletObject| wallet information after transaction | Y | - | - |
|    |    |           |         |      |      |
|**WalletObject**| | | | | | |
| uid | int | user ID | Y | - | 1 |
| sgd | int | SGD balance | Y | - | 8 |

Sample
```json
{
    "code": "A0000",
    "msg": "Success",
    "data": {
        "wallet": {
            "uid": 1,
            "sgd": 5
        }
    }
}
```
