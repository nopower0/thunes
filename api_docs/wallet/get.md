# Get User Wallet API

### Introduction
URI: `/wallet/get`

**Login Is Required**

Description:  
This API is used to get wallet status for the current user.


### Request

Parameters

|Name|Type|Description|Required?|Format|Sample|
|----|----|-----------|---------|------|------|
|    |    |           |         |      |      |

Sample
```json
{}
```
### Response

Parameters

|Name|Type|Description|Must|Format|Sample|
|----|----|-----------|---------|------|------|
| wallet | WalletObject| wallet information | Y | - | - |
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
            "sgd": 8
        }
    }
}
```
