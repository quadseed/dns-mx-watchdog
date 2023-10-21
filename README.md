# dns-mx-watchdog
DNS MX Record Watchdog

### 概要
DNSのMXレコードを監視して一定期間ごとにslackへ通知します

### 環境変数
| 変数              | 内容                          |
|-----------------|-----------------------------|
| SLACK_BOT_TOKEN | `xoxb-` から始まるslack botのトークン |
| CHANNEL_ID      | 通知の送信先のslackのチャンネルID        |
| DNS_SERVER      | 参照するDNSサーバ (ポート番号まで指定する)    |
| DOMAIN          | 問い合わせたいドメイン                 |

### 設定例
```yaml
version: '3'
services:
  dns-watchdog:
    image: ghcr.io/quadseed/dns-mx-watchdog:latest
    restart: unless-stopped
    environment:
      TZ: Asia/Tokyo
      SLACK_BOT_TOKEN: xoxb-xxxxxxx
      CHANNEL_ID: XXXXXXXX
      DNS_SERVER: XXXXXXXX
      DOMAIN: XXXXXXXX
```

### 通知の仕様
#### 監視間隔
1時間に1回指定したDNSサーバに対して指定したドメインのMXレコードを問い合わせます
#### 正常に応答があった場合
正常な応答が24回続いた場合に日報として正常に稼働している旨の通知がslackに送信されます
#### 正常な応答がない場合
異常があった旨の通知が **チャンネルメンション付き** でslackに送信されます