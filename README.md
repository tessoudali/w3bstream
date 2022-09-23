# w3bstream

## Arch

![w3bstream](__doc__/modules_and_dataflow.png)

## Features

1. wasm applet management
2. wasm runtime instance deployment
3. interact with wasm (a word count demo)

## How to run

### dependencies:

- docker: to start a postgres
- httpie: a simple curl command
- tinygo: to build wasm code

### init database

```sh
make run_depends # start postgres and mqtt
make migrate     # create or update schema
```

### create admin account

> if admin already created, skip this step

```sh
make create_admin
> username: admin
> password: c6a0a469cf1b506c251ccf966ce315e1
> please remember it
```

### login(get token)

command

```sh
echo '{"username":"admin","password":"{password}"}' | http put :8888/srv-applet-mgr/v0/login 
```

output like

```json
{
  "accountID": "4d6a0bda-bfe2-4d6c-b146-a7205a7bafae",
  "expireAt": "2022-09-23T07:20:08.099601+08:00",
  "issuer": "srv-applet-mgr",
  "token": "{token}"
}
```

### create your project

command

```sh
echo '{"name":"project_name","version":"0.0.1"}' | http post :8888/srv-applet-mgr/v0/project -A bearer -a {token}
```

output like

```json
{
  "accountID": "4d6a0bda-bfe2-4d6c-b146-a7205a7bafae",
  "createdAt": "2022-09-23T07:26:52.013626+08:00",
  "name": "project_name",
  "projectID": "254fd639-ae90-479c-9788-e2890c56e2c4",
  "updatedAt": "2022-09-23T07:26:52.013626+08:00",
  "version": "0.0.1"
}
```

### build wasm demo

```sh
make wasm_demo ## build to pkg/modules/vm/testdata/ use to deploy wasm applet
``` 

### create and deploy applet


command

```sh
http --form post :8888/srv-applet-mgr/v0/applet file@{path_to_wasm_file} info='{"projectID":"{project_id}","appletName":"{applet_name}"}' -A bearer -a {token}

http post :8888/srv-applet-mgr/v0/deploy/applet/{applet_id} -A bearer -a {token}

http put :8888/srv-applet-mgr/v0/deploy/{instance_id}/START -A bearer -a {token}
```

output like

```json
{
  "appletID": "924ac719-875a-4983-896f-5292273100b3",
  "config": null,
  "createdAt": "2022-09-23T07:37:08.101494+08:00",
  "name": "applet_name",
  "projectID": "254fd639-ae90-479c-9788-e2890c56e2c4",
  "updatedAt": "2022-09-23T07:37:08.101494+08:00"
}
```

### publish event to server

```sh
curl --location --request POST 'localhost:8888/srv-applet-mgr/v0/event/{projectID}/{appletID}/start' \
--header 'publisher: publisherID' \
--header 'Content-Type: text/plain' \
--data-raw 'testdata'
```
