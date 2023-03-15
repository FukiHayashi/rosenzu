#　バス路線図描くためのデータが取れるAPI
## 概要

下記ができるAPIです。
- 地図上にバスの路線図を引くためのデータを取得出来ます。
- 路線の経路を引くためのデータも併せて取得できます。

## 実行方法

.envファイルを作成しpostgreSQLへの接続URLを記述。

```
DATABASE_URL = "host=yourDBHost port=yourDBPort dbname=yourDBName 
```

ビルドして実行
```
go build cmd/rosenzu
.rosenzu
```

データの取得
```
curl http://localhost:8000/line/:路線名
```

[APIの簡易説明1](https://note.com/ac23811s/n/n50524b939609)


[APIの簡易説明2](https://note.com/ac23811s/n/n6502b2c2ddb7)


## 路線データについて

Google mapストリートビューを見ながら緯度経度の座標を手動で取得して作っています。

## データの使い方

取得データで路線図を描くサンプルコードは
view/*.jsを参照下さい。


[データの使用方法簡易説明](https://note.com/ac23811s/n/nbc84b15e4920)
