# JSON Convert

使 JSON 讀入後可按指定命名樣式輸出欄位格式。(尚未完成)

## 使用方法

在結構加入 `naming` 標籤，即可在輸出時按指定樣式輸出。

### 1. 加入 `naming` 標籤: 

- snake_case
- camelCase
- CamelCase

```golang
type Config struct {
	Version string  `naming:"snake_case"`
	Rules   []*Rule `naming:"snake_case"`
}
```

### 2. 使用 `convert.Naming` 轉換格式

```golang
package convert

func TestConvertNaming(t *testing.T) {
	var cfg cors.Config

	f, _ := os.Open("../cors-rules.json")
	json.NewDecoder(f).Decode(&cfg)

	newStruct := Naming(&cfg)

	jsonStr, _ := json.MarshalIndent(newStruct, "", "    ")
	fmt.Println(string(jsonStr))
}
```

```json
{
    "version": "1.0",
    "rules": [
        {   // 目前碰到巢狀自訂結構還沒想到辦法解
            "Resource": {
                "Path": "/api/data/documents",
                "StartsWith": false,
                "Exact": false
            },
            "AllowOrigins": [
                "http://this.example.com",
                "http://that.example.com"
            ],
            "AllowMethods": [
                "GET"
            ],
            "AllowCredentials": true,
            "ExposeHeaders": null
        }
    ]
}
```