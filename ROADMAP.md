# ROADMAP.md

## 0. リポジトリ基盤
- [x] Go module 初期化
- [x] `cmd/proxy` エントリポイント
- [x] `mise.toml` 定義

## 1. 最小TCP Reverse Proxy（テーマ領域）
- [ ] Listenして接続受理
  - [x] 最小限実装
  - [x] `--listen` 引数を追加
  - [x] `conn.Read()` できるように修正
  - [x] `parseArgs()` のテストと実装を作成
  - [x] `conn.Close()` を `defer` で管理
  - [x] 接続処理を goroutine で実行
  - [x] 接続 timeout を設定
  - [ ] `WaitGroup` で in-flight を管理
- [ ] バックエンド選択（最初はラウンドロビン）
- [ ] upstreamへDial
- [ ] 双方向コピー（client->backend, backend->client）
- [ ] 切断時の後始末（close順序/リーク確認）

## 2. 接続管理・保護
- [ ] 同時接続数制限（全体）
- [ ] バックエンド単位の同時接続上限
- [ ] accept/dial/read/write timeout 方針を実装
- [ ] idle connection timeout
- [ ] graceful shutdown（新規停止＋既存ドレイン）

## 3. 負荷制御
- [ ] token bucketによるレート制限（IPまたは接続単位）
- [ ] バックプレッシャ時の振る舞い定義（待つ/落とす）
- [ ] queue長・drop数の計測
- [ ] 過負荷時のログ抑制（サンプリング）

## 4. ロードバランシング高度化（テーマ領域）
- [ ] ヘルスチェック（active/passive）
- [ ] 最小RTT優先
- [ ] EWMA重み付け
- [ ] slow start（復帰ノードの流量漸増）
- [ ] circuit breaker（backend単位）

## 5. 観測性（OTel準備）
- [ ] 構造化ログ導入（trace-like connection idを含める）
- [ ] メトリクス導入（active_conns, bytes, errors, latency）
- [ ] 接続トレースIDの生成と伝搬方針を定義
- [ ] ログ相関キーを統一
- [ ] pprof有効化（CPU/heap/goroutine）

## 6. 検証基盤（非テーマ領域）
- [ ] ローカル検証用backendエコーサーバを用意
- [ ] 負荷試験スクリプト（例: `ncat`/`vegeta`/自作）
- [ ] 障害注入スクリプト（遅延/ロス/切断）
- [ ] ベンチ結果保存フォーマットを決める
- [ ] 回帰確認チェックリストを作る

## 7. 運用性
- [ ] 設定ファイル（YAML/JSON）ローダ
- [ ] 設定のバリデーション
- [ ] SIGHUPまたは管理APIで設定再読込
- [ ] ドライラン（設定検証のみ）モード
- [ ] 管理用エンドポイント（health/metrics）

## 8. 発展アイデア
- [ ] PROXY protocol対応
- [ ] TLS終端/パススルー切替
- [ ] 再試行ポリシー（指数バックオフ）
- [ ] 優先度キュー（特定クライアント優先）
- [ ] eBPF連携でボトルネック分析

## 9. モチベ可視化タスク
- [ ] `docs/progress.md` を作り、達成項目を追記
- [ ] 各機能で「デモ手順」を1つずつ残す
- [ ] before/after ベンチ比較を最低3件残す
- [ ] 「壊して直した記録」を最低5件残す
- [ ] スクリーンショット/ログ断片を成果として保存

## メモ
- テーマ領域の本体コードはユーザーがスクラッチ実装。
- エージェントは主に非テーマ領域を実装し、テーマ領域は骨組み・検証観点を支援。
