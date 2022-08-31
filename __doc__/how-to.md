# W3bstream

## code structure

```
.
├── __doc__
├── build
├── cmd
│   ├── pub_client        # mock message pub client
│   └── srv-applet-mgr    # applet management backend
├── pkg
│   ├── enums             
│   ├── errors
│   │   └── status        # http status
│   ├── models            # database models
│   └── modules
│       ├── applet        # applet
│       ├── applet_deploy # applet deploy 
│       ├── model         # applet model initialization
│       ├── resource      # applet assert storage
│       ├── testdata   
│       └── vm            # wasm vm instance management
└── testutil
```

## run

> dependencies: mqtt and postgres

### dependencies (if need)

```sh
make run_depends
```

> modify `cmd/srv-applet-mgr/config/local.yaml` to use your config

### migrate database

```sh
make migrate
```

### run server

```sh
make run_server
```

