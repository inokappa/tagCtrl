# tagCtrl

## これなに

EC2 タグの取得, 追加, 削除を行うワンバイナリツールです.

## 使い方

### インストール

https://github.com/inokappa/tagCtrl/releases から環境に応じたバイナリをダウンロードしてください.

```
wget https://github.com/inokappa/tagCtrl/releases/download/v0.0.1/tagCtrl_darwin_amd64 -O ~/bin/ec2ctrl
chmod +x ~/bin/tagCtrl
```

### ヘルプ

```sh
$ tagCtrl -h
Usage of tagCtrl:
  -add
        タグをインスタンスに付与.
  -del
        タグをインスタンスから削除.
  -endpoint string
        AWS API のエンドポイントを指定.
  -instances string
        Instance ID 又は Instance Tag 名を指定.
  -list
        インスタンスのタグ一覧を取得.
  -profile string
        Profile 名を指定.
  -region string
        Region 名を指定. (default "ap-northeast-1")
  -tags string
        Tag Key(Key=) 及び Tag Value(Value=) を指定.
  -version
        バージョンを出力.
```

### タグ一覧の取得

指定したインスタンス ID のタグ一覧を表示する.

```sh
$ tagCtrl -endpoint=http://127.0.0.1:5000 -instances=i-8840e11cee96b5c0d,i-55e152af866937e65
+---------------------+------+---------------+
|     INSTANCEID      | KEY  |     VALUE     |
+---------------------+------+---------------+
| i-8840e11cee96b5c0d | json | {"foo":"bar"} |
+                     +------+---------------+
|                     | Name | ec2-1         |
+---------------------+------+---------------+
| i-55e152af866937e65 | json | {"foo":"bar"} |
+                     +------+---------------+
|                     | Name | ec2-2         |
+---------------------+------+---------------+
```

### タグの追加

`-add` 及び `-tags` オプションを使って EC2 にタグを付与する.

```sh
$ tagCtrl -endpoint=http://127.0.0.1:5000 -instances=i-8840e11cee96b5c0d,i-55e152af866937e65 -tags="Key=foo,Value=bar Key=baz,Value=qux" -add
$ tagCtrl -endpoint=http://127.0.0.1:5000 -instances=i-8840e11cee96b5c0d,i-55e152af866937e65
+---------------------+------+---------------+
|     INSTANCEID      | KEY  |     VALUE     |
+---------------------+------+---------------+
| i-8840e11cee96b5c0d | json | {"foo":"bar"} |
+                     +------+---------------+
|                     | foo  | bar           |
+                     +------+---------------+
|                     | baz  | qux           |
+                     +------+---------------+
|                     | Name | ec2-1         |
+---------------------+------+---------------+
| i-55e152af866937e65 | json | {"foo":"bar"} |
+                     +------+---------------+
|                     | foo  | bar           |
+                     +------+---------------+
|                     | baz  | qux           |
+                     +------+---------------+
|                     | Name | ec2-2         |
+---------------------+------+---------------+
```

以下のように指定することで, JSON 文字列もタグの値として付与することが可能.

```sh
$ tagCtrl -endpoint=http://127.0.0.1:5000 -instances=i-8840e11cee96b5c0d,i-55e152af866937e65 -tags="Key=json:sample,Value={\"baz\":\"qux\"}" -add
$ tagCtrl -endpoint=http://127.0.0.1:5000 -instances=i-8840e11cee96b5c0d,i-55e152af866937e65
+---------------------+-------------+---------------+
|     INSTANCEID      |     KEY     |     VALUE     |
+---------------------+-------------+---------------+
| i-8840e11cee96b5c0d | json        | {"foo":"bar"} |
+                     +-------------+---------------+
|                     | Name        | ec2-1         |
+                     +-------------+---------------+
|                     | json:sample | {"baz":"qux"} |
+---------------------+-------------+---------------+
| i-55e152af866937e65 | json        | {"foo":"bar"} |
+                     +-------------+---------------+
|                     | Name        | ec2-2         |
+                     +-------------+---------------+
|                     | json:sample | {"baz":"qux"} |
+---------------------+-------------+---------------+
```

### タグの削除

`-del` 及び `-tags` オプションを使って EC2 にタグを削除する.

```sh
$ tagCtrl -endpoint=http://127.0.0.1:5000 -instances=i-8840e11cee96b5c0d,i-55e152af866937e65 -tags="Key=foo" -del
$ tagCtrl -endpoint=http://127.0.0.1:5000 -instances=i-8840e11cee96b5c0d,i-55e152af866937e65 -tags="Key=baz,Value=qux" -del
$ tagCtrl -endpoint=http://127.0.0.1:5000 -instances=i-8840e11cee96b5c0d,i-55e152af866937e65
+---------------------+------+---------------+
|     INSTANCEID      | KEY  |     VALUE     |
+---------------------+------+---------------+
| i-8840e11cee96b5c0d | json | {"foo":"bar"} |
+                     +------+---------------+
|                     | Name | ec2-1         |
+---------------------+------+---------------+
| i-55e152af866937e65 | json | {"foo":"bar"} |
+                     +------+---------------+
|                     | Name | ec2-2         |
+---------------------+------+---------------+
```
