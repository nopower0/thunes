# API Basic

## Basic Information
* All APIs use `POST` method
* All APIs use `application/json` content type
* All Responses use JSON format
* All APIs except `token/request` must have `X-Token` header fulfilled with the token got from `token/request` API
* All APIs' responses are in the following format
```json
{
    "code": <Response Code>,
    "msg": <Response Message>,
    "data": <Response Data>
}
```

## Error Codes
|Code |Description|
|-----|-----------|
|A0000| Success |
|S0001| There is something wrong with server |
|T0001| Token is Invalid |
|U0001| There is no user associated with the token |
|U0002| Login failed due to invalid username or password |
|U0003| There is already a user associated with the token when calling login API |
|U0004| The transfer-to user does not exist when calling transfer API |
|W0001| The balance is not enough to complete this transfer transaction |
