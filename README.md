# tryEcho

[Quick Start](https://echo.labstack.com/docs/quick-start)

[Github](https://github.com/orechen422dazo/tryEcho)

```shell
go get github.com/labstack/echo/v4
```

example:
```go
package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

// URL
// http://localhost:1323/
func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
```

Start server:
```shell
go run main.go
```

# お買い物リスト API ドキュメント
API仕様書です。

## 概要
このAPIは、お買い物リストの管理を行うためのものです。商品の追加、取得、更新、削除が可能です。

ベースURL: `http://localhost:1323`

## エンドポイント

### 1. 全ての商品を取得

- **URL**: `/items`
- **メソッド**: GET
- **成功レスポンス**:
    - コード: 200
    - 内容: 商品オブジェクトの配列
  ```json
  [
    {
      "id": 1,
      "name": "りんご",
      "price": 100,
      "purchased": false
    },
    {
      "id": 2,
      "name": "バナナ",
      "price": 80,
      "purchased": true
    }
  ]
  ```

### 2. 特定の商品を取得

- **URL**: `/items/:id`
- **メソッド**: GET
- **URL パラメータ**: id=[integer]
- **成功レスポンス**:
    - コード: 200
    - 内容: 単一の商品オブジェクト
  ```json
  {
    "id": 1,
    "name": "りんご",
    "price": 100,
    "purchased": false
  }
  ```
- **エラーレスポンス**:
    - コード: 404
    - 内容: `{ "error": "Item not found" }`

### 3. 新しい商品を追加

- **URL**: `/items`
- **メソッド**: POST
- **データパラメータ**:
  ```json
  {
    "name": "商品名",
    "price": 価格,
    "purchased": 購入済みかどうか(true/false)
  }
  ```
- **成功レスポンス**:
    - コード: 201
    - 内容: 作成された商品オブジェクト（IDを含む）
- **エラーレスポンス**:
    - コード: 400
    - 内容: `{ "error": "Invalid input" }`

### 4. 商品を更新

- **URL**: `/items/:id`
- **メソッド**: PUT
- **URL パラメータ**: id=[integer]
- **データパラメータ**:
  ```json
  {
    "name": "更新後の商品名",
    "price": 更新後の価格,
    "purchased": 更新後の購入状態(true/false)
  }
  ```
- **成功レスポンス**:
    - コード: 200
    - 内容: 更新された商品オブジェクト
- **エラーレスポンス**:
    - コード: 404
    - 内容: `{ "error": "Item not found" }`

### 5. 商品を削除

- **URL**: `/items/:id`
- **メソッド**: DELETE
- **URL パラメータ**: id=[integer]
- **成功レスポンス**:
    - コード: 200
    - 内容: `{ "message": "Item deleted successfully" }`
- **エラーレスポンス**:
    - コード: 404
    - 内容: `{ "error": "Item not found" }`

## 使用例

### cURLを使用した例

1. 全ての商品を取得:
   ```
   curl http://localhost:1323/items
   ```

2. 新しい商品を追加:
   ```
   curl -X POST -H "Content-Type: application/json" -d '{"name":"りんご", "price":100, "purchased":false}' http://localhost:1323/items
   ```

3. 特定の商品を取得 (IDを1と仮定):
   ```
   curl http://localhost:1323/items/1
   ```

4. 商品を更新:
   ```
   curl -X PUT -H "Content-Type: application/json" -d '{"name":"りんご", "price":120, "purchased":true}' http://localhost:1323/items/1
   ```

5. 商品を削除:
   ```
   curl -X DELETE http://localhost:1323/items/1
   ```

## 注意事項
- このAPIは認証を実装していません。実際の運用では適切な認証メカニズムを実装してください。
- エラーハンドリングは基本的なものにとどまっています。実際の使用では、より詳細なエラーメッセージを提供することをお勧めします。
- データの永続化は実装されていません。サーバーの再起動時にデータは失われます。