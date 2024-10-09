# ベースイメージとしてDenoの公式イメージを使用
FROM denoland/deno:1.46.3

# ビルド引数としてPORTを定義（デフォルト値を設定）
ARG PORT
ENV PORT=${PORT}

# 作業ディレクトリを設定
WORKDIR /app

# アプリケーションの依存関係をコピー
# deps.tsは通常、プロジェクトの依存関係を管理するために使用されます。
COPY ./fresh/deps.ts .

# deps.ts に記載された依存関係をキャッシュします。
# これにより、後続のビルドや実行が高速化されます。
RUN deno cache deps.ts

# アプリケーションのソースコードをコピー
COPY ./fresh/ .

# アプリケーションをビルド
# Freshプロジェクトのビルドタスクを実行します。
# これは通常、deno.json または deno.jsonc に定義されたタスクです。
RUN deno task build

# 実行時の権限を設定
# メインアプリケーションファイルとその依存関係をキャッシュします。
# これにより、アプリケーションの起動が高速化されます
RUN deno cache main.ts

# ポートを公開
EXPOSE ${PORT}

# アプリケーションを実行
# コンテナ起動時に実行されるコマンドを指定します。
# "deno task start" は通常、Freshアプリケーションを起動するコマンドです。
CMD ["deno", "task", "start"]
