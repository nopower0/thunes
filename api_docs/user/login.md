# User Login API

### Introduction
URI: `/user/login`

Description:  
This API is used to associate a user to the token


### Request

Parameters

|Name|Type|Description|Required?|Format|Sample|
|----|----|-----------|---------|------|------|
| username | string | | Y | - | - |
| password | string | | Y | - | - |

Sample
```json
{
    "username": "robert",
    "password": "robert"
}
```
### Response

Parameters

|Name|Type|Description|Must|Format|Sample|
|----|----|-----------|---------|------|------|
| uid | int | user ID | Y | - | 1 |
| username | string | | Y | - | - |

Sample
```json
{
    "code": "A0000",
    "msg": "Success",
    "data": {
        "uid": 1,
        "username": "robert"
    }
}
```
