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

以下でプロジェクト内で有効な task コマンド一覧を確認できます。

```bash
task --list
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
