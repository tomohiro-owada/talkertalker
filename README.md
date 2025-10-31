# TalkerTalker

superwhisperで翻訳されたテキストを自動で読み上げるmacOS用GUIアプリケーション（Go製）

## 特徴

- **タガログ語・日本語対応**
  - タガログ語: 男性音声（Angelo）
  - 日本語: 女性音声（Nanami）
- **2つのアプリを同時起動可能**
  - タガログ語専用アプリ
  - 日本語専用アプリ
- **グローバルホットキー対応**
  - **Ctrl+Shift+G**: どこからでも日本語アプリをアクティブ化
  - **Ctrl+Shift+H**: どこからでもタガログ語アプリをアクティブ化
- **Auto Read機能**: テキストが貼り付けられると自動で読み上げ開始
- **自動クリア**: 読み上げ完了後、テキストを自動削除
- **高品質な音声**: Microsoft Edge TTSを使用（無料）
- **シンプルなGUI**: Fyne製の軽量アプリ

## 必要要件

- macOS
- Python 3.x（edge-ttsコマンド用）
- インターネット接続（Edge TTS APIを使用するため）
- **アクセシビリティ権限**（グローバルホットキー機能を使用する場合）

## インストール

### 1. Python依存関係のインストール

```bash
pip install --upgrade edge-tts
# edge-tts 7.2.3以降が必要です
```

### 2. アクセシビリティ権限の設定

グローバルホットキー（Ctrl+Shift+G / Ctrl+Shift+H）を使用するには、アプリにアクセシビリティ権限を付与する必要があります：

1. アプリを初回起動すると、権限のリクエストが表示されます
2. **システム設定** → **プライバシーとセキュリティ** → **アクセシビリティ** を開く
3. 「TalkerTalker Japanese」と「TalkerTalker Tagalog」を有効にする

権限を付与しないとホットキーは機能しませんが、アプリ自体は通常通り使用できます。

### 3. アプリケーションのビルド（オプション）

既にビルド済みのバイナリ（`talkertalker`）が含まれています。
再ビルドする場合は：

```bash
go build -o talkertalker .
```

## 使い方

### 方法1: Applicationsフォルダから起動（推奨）

既に2つの.appバンドルがApplicationsフォルダにインストールされています：

1. **TalkerTalker Japanese.app** - Launchpadまたはアプリケーションフォルダから起動
2. **TalkerTalker Tagalog.app** - Launchpadまたはアプリケーションフォルダから起動

両方同時に起動することもできます！

### 方法2: コマンドラインから起動

**タガログ語アプリを起動:**

```bash
./start-tagalog.sh
```

または

```bash
./talkertalker -lang tl
```

**日本語アプリを起動:**

```bash
./start-japanese.sh
```

または

```bash
./talkertalker -lang ja
```

**両方を同時に起動:**

```bash
./start-tagalog.sh &
./start-japanese.sh &
```

## 使用方法

### 方法1: Auto Read機能（推奨）

1. アプリを起動（Auto Readはデフォルトでオン）
2. superwhisperで音声入力→翻訳
3. 翻訳結果をTalkerTalkerのテキストボックスに貼り付け
4. **自動で読み上げ開始！**
5. 読み上げ完了後、テキストが自動でクリアされます

### 方法2: グローバルホットキーで素早くアクセス

1. 他のアプリを使用中でも **Ctrl+Shift+G**（日本語）または **Ctrl+Shift+H**（タガログ語）を押す
2. TalkerTalkerがアクティブになる
3. テキストを貼り付けると自動で読み上げ

### 方法3: 手動操作

1. アプリを起動
2. テキストボックスにテキストを入力または貼り付け
3. "Speak"ボタンをクリック（Auto Readがオフの場合）

## プロジェクト構造

```
talkertalker/
├── main.go                # メインアプリケーション
├── tts/
│   └── tts.go            # TTS機能
├── talkertalker          # ビルド済みバイナリ
├── start-tagalog.sh      # タガログ語起動スクリプト
├── start-japanese.sh     # 日本語起動スクリプト
├── go.mod                # Go依存関係
├── go.sum                # Go依存関係チェックサム
└── README.md             # このファイル

# Python版（旧バージョン）
├── tts_reader.py         # Python実装（非推奨）
└── requirements.txt      # Python依存関係
```

## 技術スタック

- **言語**: Go 1.25+
- **GUI**: Fyne v2
- **TTS**: Microsoft Edge TTS (edge-tts CLI経由)
- **音声再生**: macOS標準の`afplay`コマンド
- **クリップボード**: atotto/clipboard
- **グローバルホットキー**: robotn/gohook

## 対応音声

| 言語 | コード | 音声名 | 性別 |
|------|--------|--------|------|
| タガログ語 | tl | fil-PH-AngeloNeural | 男性 |
| 日本語 | ja | ja-JP-NanamiNeural | 女性 |

音声の変更は `tts/tts.go:16-19` で設定できます。

## トラブルシューティング

### edge-ttsコマンドが見つからない

```bash
pip install edge-tts
# または
pip3 install edge-tts
```

### 音声が再生されない

- インターネット接続を確認してください
- edge-ttsが正しくインストールされているか確認：

```bash
edge-tts --version
```

### ビルドエラー

```bash
go mod tidy
go build -o talkertalker .
```

### ホットキーが動作しない

1. システム設定でアクセシビリティ権限が付与されているか確認
2. アプリを一度終了して再起動
3. 権限を付与した後、macOS自体の再起動が必要な場合があります

## 今後の拡張予定

- [x] グローバルホットキー機能
- [x] Auto Read機能
- [x] 読み上げ後の自動クリア
- [ ] メニューバーアプリ化
- [ ] 音声速度調整
- [ ] 音声保存機能

## ライセンス

MIT License

## 使用技術への感謝

- [Fyne](https://fyne.io/) - Cross-platform GUI toolkit
- [Edge TTS](https://github.com/rany2/edge-tts) - Microsoft Edge Text-to-Speech
- [clipboard](https://github.com/atotto/clipboard) - Cross-platform clipboard library
- [gohook](https://github.com/robotn/gohook) - Global hotkey library for Go
