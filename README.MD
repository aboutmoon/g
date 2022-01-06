# GoFrame Demos

This repo implements some demos using GoFrame.
1. A simple websocket chat service.
1. Basic API example for user SignUp/SignIn.
1. Universal CURD service.

## Installation

### 1. You need a go development environment setup before everything starts taking off.
### 2. Use `git clone` cloing the repo to your local folder. 
```
git clone https://github.com/aboutmoon/g
```

### 3. Import `document/sql/create.sql` to your database. 

### 4. Create configuration file from `config.example.toml`. 
```
cp config/config.example.toml config/config.toml
```
Update `config.toml` according to your local configurations if necessary.

### 5. Run command `go run main.go`, and you'll see something as follows if success:
```
  SERVER  | DOMAIN  | ADDRESS | METHOD |        ROUTE        |                          HANDLER                          |                                                  MIDDLEWARE                                                    
----------|---------|---------|--------|---------------------|-----------------------------------------------------------|----------------------------------------------------------------------------------------------------------------
  default | default | :8199   | ALL    | /chat               | github.com/aboutmoon/g/app/api.(*apiChat).Index         | service.(*serviceMiddleware).Ctx-fm,service.(*serviceMiddleware).CORS-fm                                       
----------|---------|---------|--------|---------------------|-----------------------------------------------------------|----------------------------------------------------------------------------------------------------------------
  default | default | :8199   | ALL    | /chat/index         | github.com/aboutmoon/g/app/api.(*apiChat).Index         | service.(*serviceMiddleware).Ctx-fm,service.(*serviceMiddleware).CORS-fm                                       
----------|---------|---------|--------|---------------------|-----------------------------------------------------------|----------------------------------------------------------------------------------------------------------------
  default | default | :8199   | ALL    | /chat/setname       | github.com/aboutmoon/g/app/api.(*apiChat).SetName       | service.(*serviceMiddleware).Ctx-fm,service.(*serviceMiddleware).CORS-fm                                       
----------|---------|---------|--------|---------------------|-----------------------------------------------------------|----------------------------------------------------------------------------------------------------------------
  default | default | :8199   | ALL    | /chat/websocket     | github.com/aboutmoon/g/app/api.(*apiChat).WebSocket     | service.(*serviceMiddleware).Ctx-fm,service.(*serviceMiddleware).CORS-fm                                       
----------|---------|---------|--------|---------------------|-----------------------------------------------------------|----------------------------------------------------------------------------------------------------------------
  default | default | :8199   | ALL    | /swagger/*          | github.com/gogf/swagger.(*Swagger).Install.func1.1        | HOOK_BEFORE_SERVE                                                                                              
----------|---------|---------|--------|---------------------|-----------------------------------------------------------|----------------------------------------------------------------------------------------------------------------
  default | default | :8199   | ALL    | /user/checknickname | github.com/aboutmoon/g/app/api.(*apiUser).CheckNickName | service.(*serviceMiddleware).Ctx-fm,service.(*serviceMiddleware).CORS-fm                                       
----------|---------|---------|--------|---------------------|-----------------------------------------------------------|----------------------------------------------------------------------------------------------------------------
  default | default | :8199   | ALL    | /user/checkpassport | github.com/aboutmoon/g/app/api.(*apiUser).CheckPassport | service.(*serviceMiddleware).Ctx-fm,service.(*serviceMiddleware).CORS-fm                                       
----------|---------|---------|--------|---------------------|-----------------------------------------------------------|----------------------------------------------------------------------------------------------------------------
  default | default | :8199   | ALL    | /user/issignedin    | github.com/aboutmoon/g/app/api.(*apiUser).IsSignedIn    | service.(*serviceMiddleware).Ctx-fm,service.(*serviceMiddleware).CORS-fm                                       
----------|---------|---------|--------|---------------------|-----------------------------------------------------------|----------------------------------------------------------------------------------------------------------------
  default | default | :8199   | ALL    | /user/profile       | github.com/aboutmoon/g/app/api.(*apiUser).Profile-fm    | service.(*serviceMiddleware).Ctx-fm,service.(*serviceMiddleware).CORS-fm,service.(*serviceMiddleware).Auth-fm  
----------|---------|---------|--------|---------------------|-----------------------------------------------------------|----------------------------------------------------------------------------------------------------------------
  default | default | :8199   | ALL    | /user/signin        | github.com/aboutmoon/g/app/api.(*apiUser).SignIn        | service.(*serviceMiddleware).Ctx-fm,service.(*serviceMiddleware).CORS-fm                                       
----------|---------|---------|--------|---------------------|-----------------------------------------------------------|----------------------------------------------------------------------------------------------------------------
  default | default | :8199   | ALL    | /user/signout       | github.com/aboutmoon/g/app/api.(*apiUser).SignOut       | service.(*serviceMiddleware).Ctx-fm,service.(*serviceMiddleware).CORS-fm                                       
----------|---------|---------|--------|---------------------|-----------------------------------------------------------|----------------------------------------------------------------------------------------------------------------
  default | default | :8199   | ALL    | /user/signup        | github.com/aboutmoon/g/app/api.(*apiUser).SignUp        | service.(*serviceMiddleware).Ctx-fm,service.(*serviceMiddleware).CORS-fm                                       
----------|---------|---------|--------|---------------------|-----------------------------------------------------------|----------------------------------------------------------------------------------------------------------------
```

# GoFrame Sites
### GoFrame Repository
* [https://github.com/gogf/gf](https://github.com/gogf/gf)
* [https://gitee.com/johng/gf](https://gitee.com/johng/gf)

### GoFrame Home
* [https://goframe.org](https://goframe.org)
