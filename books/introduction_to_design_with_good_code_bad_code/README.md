# 良いコード/悪いコードで学ぶ設計入門

## 1. 悪しき構造の弊害を知覚する

悪しき構造として, `意味不明な命名`、`巨大なネスト`、`データクラス`がある

### 意味不明な命名

意味不明な命名として、`技術駆動命名`や`連番命名`があり、意図が読み取れないコードになってしまう

- 技術駆動命名
  - 型情報やメモリ制御などのプログラミング用語やコンピューター用語に基づいた名前
- 連番命名

### ネスト

何重にもネストしたコードはコードの見通しが悪くなり、if文のブロックの範囲を読み解くのが難しくなってしまう

### データクラス

データしか持たないクラスにすると以下のような弊害が発生する可能性がある

- 重複コード
  - 修正漏れ
  - 可読性の低下
- 未初期化状態(生焼けオブジェクト)
  - 不正値の混入

## 2. 設計の初歩

基本的な設計の考え方として以下の4点がある

- 省略せずに意図が伝わる名前
- 変数名を使い回さない
- 意味のあるまとまりでメソッド化
- 関係し合うデータとロジックをクラスにまとめる

## 3. クラス設計

クラス自身が単体で正常に動作するよう設計する必要があり、その構成要素として以下の2点がある(場合によってはその限りではない)

- インスタンス変数
- インスタンス変数を不正状態から防御し、正常に操作するメソッド

成熟したクラスへ成長させるために

- コンストラクタで確実に正常値を設定する
  - インスタンス変数を初期化し、バリデーションを実装
  - ガード節(早期return)を用いてバリデーションすると、後続のロジックがシンプルになったり、ネストが増えずに見通しが良くなる
- ロジックをデータ保持側に寄せる
- インスタンス変数やメソッド引数、ローカル変数を不変にする
- 変更したい場合は新しいインスタンスを作成する
- 値渡しの間違いを型で防止する
  - プリミティブ型だと、誤って意図が異なる値を渡ってしまう。そこで、型を定義して渡すことで誤って値が渡ってもコンパイルエラーで弾くことができる
- 現実の営みにはないメソッドは実装しない

## 4. 不変の活用

- 可変がもたらす影響
  - インスタンスの使い回しによる意図しない挙動
  - 関数の副作用による意図しない挙動
    - 主作用: 関数が引数を受け取り、値を返すこと
    - 副作用: 主作用以外にインスタンス変数やグローバル変数、引数などの状態変更すること
- 変数や引数は不変
  - 再代入を防ぐ
  - 影響範囲が限定的になる
  - 変更したい場合は新しいインスタンスを作成する
- 可変にする場合は、正しく状態変更するメソッドを設計する

## 5. 低凝集

- staticメソッドの誤用
  - staticメソッドは、データとロジックが別々になり、低凝集な構造になる
  - インスタンス変数を使ってないインスタンスメソッドや共通処理クラスも低凝集な構造
  - 正しいstaticメソッドの使用例
    - ファクトリメソッド
    - 横断的関心事
      - フォーマット変換
      - ログ出力
      - エラー検出
      - デバッグ
      - 例外処理
      - キャッシュ
      - 同期処理
      - 分散処理 etc...
- 初期化ロジックの分散
  - コンストラクタが様々な目的に使用され、初期化ロジックが分散されがちになる
  - そこで、コンストラクタをprivateにし、目的別のファクトリメソッドを用意する
- 結果を返すために引数を使う
  - 出力になるデータを引数として扱わない(出力引数)
  - 出力引数は低凝集だけでなく、引数が変更されることが外部からわからない
  - そこで、データとデータを操作するロジックを同じクラスに凝集する
- 引数が多すぎる
  - 引数の数が多いということは、処理する内容が多いということ
  - プリミティブ型ではなく、意味のある単位ごとにクラスを定義して、データに対する処理をまとめる
- メソッドチェーン
  - メソッドチェーンでのインスタンス変数の参照や書き換えするような書き方すると同じようなコードが複数箇所に実装される恐れがある
  - `Tell, Don't Ask`(尋ねるな、命じろ)
    - オブジェクトの内部状態を尋ねたり、状態に応じて呼び出し側が判断するのではない
    - 呼び出し側はただメソッドを呼び出すだけ
    - 呼び出された側で状態の判断や制御を行うよう設計
  - インスタンス変数をprivateにして、メソッド経由でインスタンス変数を操作、参照する

## 6. 条件分岐

## 7. コレクション

## 8. 密結合

## 9. 設計の健全性をそこなうさまざまな悪魔たち

## 10. 名前設計

## 11. コメント

## 12. メソッド(関数)

## 13. モデリング

## 14. リファクタリング

## 15. 設計の意義と設計への向き合い方

## 16. 設計を妨げる開発プロセスとの戦い

## 17. 設計技術の理解の深め方
