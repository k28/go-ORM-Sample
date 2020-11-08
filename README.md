# 目的

DockerでGoとDatabaseのコンテナを動作させて動作確認する

# コマンドラインで使うコマンド

| コマンド                    | 動作                                   |
| :-------------------        | :---                                   |
| docker-compose up -d        | コンテナを起動                         |
| docker-compose down -v      | コンテナ停止 + 内容を破棄              |
| docker-compose exec db bash | DBが動作しているコンテナにログインする |


