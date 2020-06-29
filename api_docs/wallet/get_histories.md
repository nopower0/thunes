# Get User Wallet API

### Introduction
URI: `/wallet/get_histories`

**Login Is Required**

Description:  
This API is used to transaction histories


### Request

Parameters

|Name|Type|Description|Required?|Format|Sample|
|----|----|-----------|---------|------|------|
| start | int | start transaction ID, used for pagination | N, default is 0 | - | 0 |
| length | int | transaction count | N, default is 10 | - | 10 |

Sample
```json
{
    "start": 0,
    "length": 10
}
```
### Response

Parameters

|Name|Type|Description|Must|Format|Sample|
|----|----|-----------|---------|------|------|
| histories | array of TransactionObject | transaction list ordered by transaction time desc | Y | - | - |
|    |    |           |         |      |      |
|**TransactionObject**| | | | | | |
| to | UserObject | transaction-to user information | Y | - | - |
| amount | int | transaction amount | Y | - | - |
| transaction_time | int | the timestamp when transaction happened | Y | - | - |
|    |    |           |         |      |      |
|**UserObject**| | | | | |
| uid | int | user ID | Y | - | - |
| username | string | | Y | - | - |

Sample
```json
{
    "code": "A0000",
    "msg": "Success",
    "data": {
        "histories": [
            {
                "to": {
                    "uid": 3,
                    "username": "jeremy"
                },
                "amount": 3,
                "transaction_time": 1593407997
            },
            {
                "to": {
                    "uid": 3,
                    "username": "jeremy"
                },
                "amount": 2,
                "transaction_time": 1592881867
            }
        ]
    }
}
```
