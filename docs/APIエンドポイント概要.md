### app handlers
- **POST /api/app/users**
  - ユーザー登録を行うエンドポイント。
  - 招待コードを用いて登録した場合は、招待クーポンを付与する。

- **POST /api/app/payment-methods**
  - 決済トークンの登録を行うエンドポイント。
  - 認証が必要。

- **GET /api/app/rides**
  - ユーザーが完了済みのライド一覧を取得するエンドポイント。
  - 認証が必要。

- **POST /api/app/rides**
  - ユーザーが配車を要求するエンドポイント。
  - クーポンを所有している場合、自動で利用する。
  - 認証が必要。

- **POST /api/app/rides/estimated-fare**
  - ライドの運賃を見積もるエンドポイント。
  - 認証が必要。

- **POST /api/app/rides/ride_id/evaluation**
  - ユーザーがライドを評価するエンドポイント。
  - 認証が必要。

- **GET /api/app/notification**
  - ユーザー向け通知エンドポイント。
  - 認証が必要。

- **GET /api/app/nearby-chairs**
  - ユーザーの近くにいる椅子を取得するエンドポイント。
  - 認証が必要。

### owner handlers
- **POST /api/owner/owners**
  - 椅子のオーナーが会員登録を行うエンドポイント。

- **GET /api/owner/sales**
  - 椅子のオーナーが指定期間の全体・椅子ごと・モデルごとの売上情報を取得するエンドポイント。
  - 認証が必要。

- **GET /api/owner/chairs**
  - 椅子のオーナーが管理している椅子の一覧を取得するエンドポイント。
  - 認証が必要。

### chair handlers
- **POST /api/chair/chairs**
  - 椅子の登録を行うエンドポイント。

- **POST /api/chair/activity**
  - 椅子が配車受付を開始・停止するエンドポイント。
  - 認証が必要。

- **POST /api/chair/coordinate**
  - 椅子が自身の位置情報を送信するエンドポイント。
  - 認証が必要。

- **GET /api/chair/notification**
  - 椅子向け通知エンドポイント。
  - 認証が必要。

- **POST /api/chair/rides/ride_id/status**
  - 椅子がライドのステータスを更新するエンドポイント。
  - 認証が必要。

### internal handlers
- **GET /api/internal/matching**
  - ライドのマッチングを行うエンドポイント。
  - 内部からのみアクセス可能。


```mermaid
graph TD
    subgraph User
        A["ユーザー登録"] -->|POST /api/app/users| B["ユーザー"]
        B -->|POST /api/app/payment-methods| C["決済トークン登録"]
        B -->|GET /api/app/rides| D["完了済みライド一覧取得"]
        B -->|POST /api/app/rides| E["配車要求"]
        B -->|POST /api/app/rides/estimated-fare| F["運賃見積もり"]
        B -->|POST /api/app/rides/ride_id/evaluation| G["ライド評価"]
        B -->|GET /api/app/notification| H["通知取得"]
        B -->|GET /api/app/nearby-chairs| I["近くの椅子取得"]
    end
```

```mermaid
graph TD
    subgraph Owner
        J["オーナー登録"] -->|POST /api/owner/owners| K["オーナー"]
        K -->|GET /api/owner/sales| L["売上情報取得"]
        K -->|GET /api/owner/chairs| M["椅子一覧取得"]
    end
```

```mermaid
graph TD
    subgraph Chair
        N["椅子登録"] -->|POST /api/chair/chairs| O["椅子"]
        O -->|POST /api/chair/activity| P["配車受付開始・停止"]
        O -->|POST /api/chair/coordinate| Q["位置情報送信"]
        O -->|GET /api/chair/notification| R["通知取得"]
        O -->|POST /api/chair/rides/ride_id/status| S["ライドステータス更新"]
    end

    subgraph Internal
        T["マッチング"] -->|GET /api/internal/matching| U["内部処理"]
    end
```