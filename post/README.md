# POST エンドポイント詳細

## エンドポイント
- URL: `http://localhost:1323/items`
- メソッド: POST

## リクエストボディ
リクエストボディはJSON形式で、以下の構造を持ちます：

```json
{
  "name": "商品名",
  "price": 商品の価格,
  "purchased": 購入済みかどうか（true/false）
}
```

### フィールドの説明
- `name` (文字列): 商品の名前
- `price` (数値): 商品の価格
- `purchased` (ブーリアン): 商品が購入済みかどうか（true：購入済み、false：未購入）

## リクエストの例

```json
{
  "name": "りんご",
  "price": 100,
  "purchased": false
}
```

## cURLを使用したリクエスト例

```bash
curl -X POST -H "Content-Type: application/json" -d '{"name":"りんご", "price":100, "purchased":false}' http://localhost:1323/items
```

## 期待されるレスポンス
成功した場合、HTTPステータスコード201（Created）と作成されたアイテムの詳細（IDを含む）が返されます。

```json
{
  "id": 1,
  "name": "りんご",
  "price": 100,
  "purchased": false
}
```

注意：
- `id`フィールドはサーバー側で自動的に割り当てられるため、リクエストボディに含める必要はありません。
- 価格は小数点を含む数値も指定可能です（例：99.99）。