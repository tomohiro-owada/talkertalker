# TalkerTalker

superwhisperで翻訳されたテキストを自動で読み上げるmacOS用GUIアプリケーション（Go製）

## 特徴

- **タガログ語・日本語対応**
  - タガログ語: 女性音声（Blessica）
  - 日本語: 男性音声（Keita）
- **2つのアプリを同時起動可能**
  - タガログ語専用アプリ
  - 日本語専用アプリ
- **高品質な音声**: Microsoft Edge TTSを使用（無料）
- **シンプルなGUI**: Fyne製の軽量アプリ
- **クリップボード連携**: "Paste & Speak"ボタンでワンクリック読み上げ

## 必要要件

- macOS
- Python 3.x（edge-ttsコマンド用）
- インターネット接続（Edge TTS APIを使用するため）

## インストール

### 1. Python依存関係のインストール

```bash
pip install --upgrade edge-tts
# edge-tts 7.2.3以降が必要です
```

### 2. アプリケーションのビルド（オプション）

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

### 方法1: テキストボックスで入力

1. アプリを起動
2. テキストボックスにテキストを入力または貼り付け
3. "Speak"ボタンをクリック

### 方法2: クリップボードから即座に読み上げ

1. superwhisperで音声入力→翻訳
2. 翻訳結果をコピー（自動でクリップボードにコピーされる）
3. TalkerTalkerアプリで"Paste & Speak"ボタンをクリック
4. 自動で読み上げ開始！

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

## 対応音声

| 言語 | コード | 音声名 | 性別 |
|------|--------|--------|------|
| タガログ語 | tl | fil-PH-BlessicaNeural | 女性 |
| 日本語 | ja | ja-JP-KeitaNeural | 男性 |

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

## 今後の拡張予定

- [ ] グローバルホットキー機能
- [ ] メニューバーアプリ化
- [ ] 音声速度調整
- [ ] 音声保存機能

## ライセンス

MIT License

## 使用技術への感謝

- [Fyne](https://fyne.io/) - Cross-platform GUI toolkit
- [Edge TTS](https://github.com/rany2/edge-tts) - Microsoft Edge Text-to-Speech
- [clipboard](https://github.com/atotto/clipboard) - Cross-platform clipboard library
