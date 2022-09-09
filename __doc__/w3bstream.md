# w3bstream backend

## 目标

1. project注册


1. applet创建和注册
2. applet版本控制
3. applet wasm vm部署和生命周期控制
4. applet 数据库模型创建
5. 事件监听和分发
6. 可供wasm调用的native library
7. 可供开发者使用的webassembly library

## modules

### **Applet注册/创建**

#### applet包含的内容

- applet配置文件

一个例子

```yaml
name: demo
version: 0.0.1           # `applet_name`@`version`
description: Simple Wasm Applet
schema: schema.yaml      # applet数据库模型描述
wasm: demo.wasm          # applet wasm 
protocol: websocket|mqtt # host创建事件监听通道
```

- wasm二进制文件(vm runtime)

- 数据库模型描述文件(optional)

一个例子

```yaml
name: demo
version: 0.0.1
driver: postgres
tables:
  - name: locations
    table_name: f_locations
    comment: location records
    defs: # 模型约束
      - type: primary
        name:
        fields:
          - ID
      - type: index
        name:
        fields:
          - f_x
          - f_y
    fields: # 模型字段描述
      - name: f_x
        filed_name: X
        type: float64
        constraints:
        comment:
      - name: f_y
        filed_name: Y
        type: float64
        constraints:
        comment:
      - name: f_time
        filed_name: Time
        type: timestamp
        constraints:
        comment:
```

- applet可处理的事件的描述(abi)

一个例子

```yaml
name: run
inputs:
  - name: resource_id
    type: i32
    native_type: uint32
outputs:
  - name: status
    type: i32
    native_type: uint32
  - name: message
    type: i32
    native_type: string
    stash_type: log
  - name: location
    type: i32
    native_type: location
    stash_type: database
    schema: demo@0.0.1@f_location
```

#### applet静态资源的存储

### **Applet资产存储**

1. s3
2. local disk cache

### **Wasm VM**

### **事件中心**

#### 事件

```json
{
  "pub": "event source",
  "sub": "applet@1.1.1",
  "name": "run",
  "content": "..."
}
```

### **Applet数据模型**

### **Native Library**