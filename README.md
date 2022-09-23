# w3bstream

## Arch

![w3bstream](__doc__/modules_and_dataflow.png)

## Features

1. wasm applet management
2. wasm runtime instance deployment
3. interact with wasm (a word count demo)

## How to run

### Dependencies:

- docker: to start a postgres
- httpie: a simple curl command
- tinygo: to build wasm code

### Init protocols and database

```sh
make run_depends # start postgres and mqtt
make migrate     # create or update schema
```

### start a server
```sh
make run_server
```
keep the terminal alive, and open a new terminal for the other commands.

### create admin account

> if admin already created, skip this step

```sh
make create_admin
> username: admin
> password: {$password}
> please remember it
```

### login (fetch auth token)

command

```sh
echo '{"username":"admin","password":"{password}"}' | http put :8888/srv-applet-mgr/v0/login 
```

output like

```json
{
  "accountID": "{account_id}",
  "expireAt": "2022-09-23T07:20:08.099601+08:00",
  "issuer": "srv-applet-mgr",
  "token": "{token}"
}
```

### create your project

command

```sh
echo '{"name":"{applet_name}","version":"0.0.1"}' | http post :8888/srv-applet-mgr/v0/project -A bearer -a {token}
```

output like

```json
{
  "accountID": "{account_id}",
  "createdAt": "2022-09-23T07:26:52.013626+08:00",
  "name": "{applet_name}",
  "projectID": "{project_id}",
  "updatedAt": "2022-09-23T07:26:52.013626+08:00",
  "version": "0.0.1"
}
```

### build demo wasm scripts

```sh
make wasm_demo ## build to pkg/modules/vm/testdata/ use to deploy wasm applet
``` 

### create and deploy applet


upload wasm script

```sh
http --form post :8888/srv-applet-mgr/v0/applet file@{path_to_wasm_file} info='{"projectID":"{project_id}","appletName":"{applet_name}"}' -A bearer -a {token}
```

output like

```json
{
  "appletID": "{applet_id}",
  "config": null,
  "createdAt": "2022-09-23T07:37:08.101494+08:00",
  "name": "{project_name}",
  "projectID": "{project_id}",
  "updatedAt": "2022-09-23T07:37:08.101494+08:00"
}
```

deploy applet
```sh
http post :8888/srv-applet-mgr/v0/deploy/applet/{applet_id} -A bearer -a {token}
```

start applet
```sh
http put :8888/srv-applet-mgr/v0/deploy/{instance_id}/START -A bearer -a {token}
```

### publish event to server

```sh
curl --location --request POST 'localhost:8888/srv-applet-mgr/v0/event/{project_id}/{applet_id}/start' \
--header 'publisher: {publisher_id}' \
--header 'Content-Type: text/plain' \
--data-raw 'input a test sentence'
```
