# API仕様  
本リポジトリのAPIは、過去に作成した記事管理サービスを想定している。

## 各エンドポイントのリクエストとレスポンス
### /read/articles
- method：GET
#### PathParam   
> - genre : string
> - month : int
> - year  : int 

例
> - genre : "ドラマ"
> - month : int or null
> - year  : int or null

###### Body
パターン1
```cassandraql
{
    "genre": "ドラマ" ,
    "month": 6 ,
    "year": 2020
}
```
パターン2
```
{
    "genre": "ドラマ" ,
    "month": null ,
    "year": null
}
```
パターン3
```
{
    "genre": "ドラマ" 
}
```
---
#### Response(1つ目のもの)  
> - ArticleID : string  
> - Title     : string  
> - ImagePath : string  
> - Tag       : string  
```cassandraql
{
    "articles": [
        {
            "id": "14",
            "title": "東京ラブストーリー",
            "imagePath": "sample",
            "tags": [
                "ドラマ",
                "名作",
                "東京ラブストーリー",
                "織田裕二"
            ]
        }
}
```

### /query/tag/articles  
- method：GET
#### PathParam   
> - tag : string  

例
> - tag : "ジャンル"  

---
#### Response 
> - ArticleID : string  
> - Title     : string  
> - ImagePath : string  
> - Tag       : string  
```cassandraql
{
    "articles": [
        {
            "id": "26",
            "title": "田原 俊彦デビュー",
            "imagePath": "",
            "tags": [
                "音楽",
                "アイドル",
                "田原 俊彦",
                "デビュー"
            ]
        }
    ]
}
```
---

### /read/article 
- method：GET
###### Header
```cassandraql
UserID:11111
```
#### PathParam   
> - articleID : string    

例  
> - articleID: "1"

---
#### Response
> - Title     : string  
> - ImagePath : string  
> - Context   : string  
> - Nice      : int  
> - NiceStatus: boolean(bool)  
> - ListStatus: boolean(bool) 
>- Tag        :string
> - userName  : string
> - userImage : string
> - contents  : string
```cassandraql
{
    "articleDetail": {
        "title": "鬼滅の刃",
        "imagePath": "",
        "context": "「鬼滅の刃」は、週刊少年ジャンプ(集英社)にて2016年11号から連載がスタートし2020年24号に終了した日本の漫画である。作者は吾峠呼世晴。舞台は日本の大正時代。炭を売る少年・竈門炭次郎はいつものように炭を売るために町へ降り、家に帰ると家族全員鬼に皆殺しにされていた。さらに、唯一の生き残りである妹・禰豆子(禰はしめすへん)は鬼に変貌していた。絶望の淵に立たされた炭治郎は禰豆子を人間に戻すべく、”鬼狩り”をする非政府組織・鬼殺隊に醜態するのであった。\n\n本作はテレビアニメ化を機に人気を呼び、単行本は累計発行部数が1億部を突破した。さらに、テレビアニメが大きな反響を及んだため2020年10月16日にテレビシリーズの続きである「劇場版　鬼滅の刃　無限列車編」が公開された。\n",
        "nice": 114512,
        "userStatus": {
            "Nice": false,
            "List": true
        },
        "tags": [
            "アニメ・漫画",
            "鬼滅の刃",
            "週刊少年ジャンプ",
            "吾峠呼世晴"
        ],
        "comments": [
            {
                "userName": "taketo",
                "userImage": "",
                "contents": "面白い"
            }
        ]
    }
}
```

### /mutation/add/like
- method：POST
#### Request   
> - articleID : string
###### Header
```cassandraql
UserID:b1018085@fun.ac.jp
```
###### Body
```cassandraql
{
    "articleID":"2"
}
```
---
#### Response(リクエスト送信元ユーザーがまだ記事に対していいねしていない場合)
前提：init.sqlでarticle_idのniceは100と設定されている。過去にリクエストを送ったことがない
> - nice : int
```cassandraql
{
    "nice": 101
}
```
以後増えて減ってを繰り返す
