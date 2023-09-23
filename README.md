# connect-go-todo-example

connect で todo アプリを作るぞ。Enablement Bootcamp の講義で教えてもらったこと

### 動かし方

- 必要なパッケージをインストール

```bash
$ go mod tidy
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
# シナリオテストをインストール
$ go install github.com/zoncoen/scenarigo/cmd/scenarigo@v0.15.1
# クライアントをインストール
$ brew tap ktr0731/evans
$ brew install evans
```

- サーバーを起動する

```bash
$ go run ./cmd/server/main.go
```

- プロトコルバッファーの定義を変更したら、コードに反映させる
```bash
$ buf lint
$ buf generate
```

- シナリオテストを実行する。プラグインのファイルが開けなくてエラーになる場合はシナリオテストのプラグインを更新して、再度実行する。

```bash
$ cd scenariotest
$ scenarigo run
```

- シナリオテストのプラグインを更新する

```bash
$ cd scenariotest
$ scenarigo plugin build
```

- クライアントを動かす。root 直下で以下のコマンドを実行する。

```bash
$ evans --proto proto/todo/v1/todo.proto repl --port 8080

  ______
 |  ____|
 | |__    __   __   __ _   _ __    ___
 |  __|   \ \ / /  / _. | | '_ \  / __|
 | |____   \ V /  | (_| | | | | | \__ \
 |______|   \_/    \__,_| |_| |_| |___/

 more expressive universal gRPC client


todo.v1.ToDoService@127.0.0.1:8080> show package
+---------+
| PACKAGE |
+---------+
| todo.v1 |
+---------+

todo.v1.ToDoService@127.0.0.1:8080> show service
+-------------+------------------+-------------------------+--------------------------+
|   SERVICE   |       RPC        |      REQUEST TYPE       |      RESPONSE TYPE       |
+-------------+------------------+-------------------------+--------------------------+
| ToDoService | CreateTask       | CreateTaskRequest       | CreateTaskResponse       |
| ToDoService | UpdateTaskStatus | UpdateTaskStatusRequest | UpdateTaskStatusResponse |
| ToDoService | DeleteTask       | DeleteTaskRequest       | DeleteTaskResponse       |
+-------------+------------------+-------------------------+--------------------------+

todo.v1.ToDoService@127.0.0.1:8080> show message
+--------------------------+
|         MESSAGE          |
+--------------------------+
| CreateTaskRequest        |
| CreateTaskResponse       |
| DeleteTaskRequest        |
| DeleteTaskResponse       |
| UpdateTaskStatusRequest  |
| UpdateTaskStatusResponse |
+--------------------------+

todo.v1.ToDoService@127.0.0.1:8080> service ToDoService

todo.v1.ToDoService@127.0.0.1:8080> call CreateTask
name (TYPE_STRING) => タスク1
✔ STATUS_TODO
{
  "id": "1",
  "name": "タスク1",
  "status": "STATUS_TODO"
}

```
