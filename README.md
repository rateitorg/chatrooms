# chatrooms

`chatrooms` will hold the handling of client <-> server connections.

## Installation

### 1. Clone the repository

```bash
git clone git@github.com:rateitorg/chatrooms.git
cd chatrooms
```

## Usage

### 1. Run `./server`

```bash
go run ./server
```

### 2. Open a WebSocket client

> [!NOTE]
> This example uses `wscat`.

```bash
wscat -c ws://127.0.0.1:8080/ws
```

### 3. Send `json`

```json
{"content":"Hello","sender":"user1","sentTime":"time"}
```

### Understanding `json` structure

> [!NOTE]
> Message struct is defined in `./entity/message.go`

| Key      | Value             |
| -------- | ----------------- |
| content  | message content   |
| sender   | message sender    |
| sentTime | message send time |

## Contributing

We welcome contributions! Please look at [Contributing.md](https://github.com/rateitorg/chatrooms/blob/main/templates/README.md) to learn how to contribute.

## Licence

The MIT Licence (MIT)

Copyright (c) 2023 - 2024 Jake Brunning, Gabriel Guimaraes, Callun Thiart. All Rights Reserved.
