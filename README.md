# Authentication service
A small part of the authentication server

## Technologies:
- Go
- JWT
- MongoDB

### POST /auth/sign-in

Request to generate and get access and refresh tokens.
*Initially, the user is already created in the database*

##### Example Input:
```
{
    "id": "64e6690b209ea12434cc3ba2",
} 
```

##### Example Response 1:
```
{
    "status": "OK",
    "access_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTI4MjI2NjUsImlhdCI6MTY5MjgyMTc2NSwidXNlcl9pZCI6IjY0ZTY2OGNhMjA5ZWExMjQzNGNjM2JhMSJ9.F8ufLqZ4jtR9MP1XxEITtwpuuM_JK4Puu5-2ROn0mx_eP_Xt5w1LvE0MEMlyaouwLPl9vM1d8sRH7EslASxFhA",
    "refresh_token": "REJIb3pOWkxRbHRRQkpLb2xiQmFyQmlMUENqdmlUa1Q="
} 
```

##### Example Response 2:
```
{
    "status": "Error",
    "error": "user not found"
} 
```

### POST /auth/refresh

Request to refresh couple of access and refresh tokens.
*The refresh token must not be expired*

##### Example Input:
```
{
    "id": "64e6690b209ea12434cc3ba2",
    "refresh_token": "REJIb3pOWkxRbHRRQkpLb2xiQmFyQmlMUENqdmlUa1Q="
} 
```

##### Example Response 1:
```
{
    "status": "OK",
    "access_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTI4MjM1ODgsImlhdCI6MTY5MjgyMjY4OCwidXNlcl9pZCI6IjY0ZTY2OGNhMjA5ZWExMjQzNGNjM2JhMSJ9.fOvb7SRCTsZswmW2BYQ7gSwgK3-_OEyGA7LoW_LQe2kRXbdwaDCk694SgF0BPh2plZ0tu7gCzerTpd3GDeXKHQ",
    "refresh_token": "eGpncnR6T0NVa0diVm5GTUlJYkNIU1hTa3lTTlBVSHk="
} 
```

##### Example Response 2:
```
{
    "status": "Error",
    "error": "token has expired"
} 
```

## Mongo Document User
![alt text](example_user.bmp)


## Run Project

Use ```make start``` to build project in a binary file and run it