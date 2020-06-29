# Token Request API

### Introduction
URI: `/token/request`

Description:  
This API is used to get a token for other API requests.
If `username` or `password` is specified, there will be a verification.


### Request

Parameters

|Name|Type|Description|Required?|Format|Sample|
|----|----|-----------|---------|------|------|
| username | string | | N | - | - |
| password | string | | N | - | - |

Sample
```json
{}
```
### Response

Parameters

|Name|Type|Description|Must|Format|Sample|
|----|----|-----------|---------|------|------|
| token | string | token used to send other requests | Y | - | 5bd01800-74b7-4bba-ac6a-249c5703eef6 |
| expire_at | int | the timestamp (in second) when the token will expire | Y | - | 1593489671 |

Sample
```json
{
    "code": "A0000",
    "msg": "Success",
    "data": {
        "token": "5bd01800-74b7-4bba-ac6a-249c5703eef6",
        "expire_at": 1593489671
    }
}
```
