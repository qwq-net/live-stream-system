# Live Stream System

Go言語によるリアルタイム動画ストリーミング・サーバーの試験的な実装です。

## 技術スタック

- **Language**: Go (Golang)
- **Infrastructure**: Docker, Docker Compose
- **Command Management**: [Task](https://taskfile.dev/) (Taskfile.yml)

## セットアップと起動

このプロジェクトでは、コマンドの管理に `Task` を使用しています。

### 1. 依存関係のインストール

GoとTaskがインストールされていることを確認してください。

- [Go Installation Guide](https://go.dev/doc/install)
- [Task Installation Guide](https://taskfile.dev/installation/)

### 2. コードのフォーマット

```bash
task fmt
```

### 3. アプリケーションの実行

ローカル環境で直接実行する場合：

```bash
task run
```

Docker Composeを使用して実行する場合：

```bash
task docker:up
```

## プロジェクト構成

```text
.
├── Taskfile.yml         # タスク定義（コマンドショートカット）
├── cmd/
│   └── video-server/    # メインエントリポイント
├── internal/
│   ├── config/          # 設定管理
│   ├── server/          # サーバーコアロジック
│   └── stream/          # ストリーミング処理
├── pkg/                 # 外部から再利用可能なライブラリ（現在は配置のみ）
├── Dockerfile           # アプリケーションのDockerイメージ定義
└── compose.yaml         # Docker Compose構成ファイル
```

## 利用可能なタスク一覧

`task --list` で詳細を確認できます。

| コマンド            | 内容                   |
| :------------------ | :--------------------- |
| `task fmt`          | Goコードのフォーマット |
| `task test`         | テストの実行           |
| `task build`        | 実行バイナリのビルド   |
| `task run`          | ローカルでの実行       |
| `task docker:build` | Dockerイメージのビルド |
| `task docker:up`    | Docker Composeでの起動 |
| `task docker:down`  | Docker Composeでの停止 |
