# 各種コマンド

## マイグレーションファイル作成
```
migrate create -ext sql -dir migrations -seq {マイグレーションファイル名};
```

## マイグレーション
```
migrate -source file://./migrations -database mysql://root:secret@tcp\(127.0.0.1:3307\)/dev up
```