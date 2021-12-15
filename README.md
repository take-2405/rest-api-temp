# rest-api-temp


## リポジトリ概要
アドベントカレンダー用に作成  
https://qiita.com/advent-calendar/2021/hcb-2021

## 用途
迅速にrest apiを作らなきゃいけない場合のテンプレート   
注意
- テンプレートAPIの使用はdocs/api.mdに記載  
- ディレクトリ構成、DIはあまり意識できていない

### 実行方法
※実行にはdockerとgoの環境が必要(そのうちdocker images化するかも)  

1. リポジトリをクローン  
2. make compose-build  
3. make run  
4. make compose-clean  
4はDBを削除するコマンドのため作業終了時に使用してください。

### 使用ライブラリ  
- gin
- sql

### アーキテクチャ
mvc

### 作成者  
take-2405
