# パーフェクトRuby

rubyは全てがオブジェクトである。そのため、数値や真偽値、nilなどのプリミティブ型のような方もオブジェクトであるため、メソッドを読み出すことが可能。

```ruby
p 1.to_s     # => "1"
p true.to_s  # => "true"
p false.to_s # => "false"
p nil.to_s   # => ""
```

## クラス

### クラス定義

`class`キーワードを用いて以下のようにクラスを定義できる。

```ruby
class Hoge
end
```

`Hoge.new`でクラスのインスタンスを生成することができる。また、`initialize`が定義されていればそれを呼び出す。

### インスタンスメソッド

インスタンスメソッドの呼び出しはレシーバが省略して呼び出すことができる。省力した場合は、暗黙的に`self`がレシーバになり、selfはクラス自身。
メソッド名の末尾に`?`や`!`を使用できる。Rubyの慣習的に、末尾に`?`につくメソッドは、オブジェクトの状態を真偽値で返すメソッドを表し、`!`が末尾につくメソッドは破壊的メソッドを表す傾向がある。

### インスタンス変数へのアクセス

オブジェクトが持っているインスタンス変数へアクセスするには、以下のようなgetter, setterを定義する必要がある。
しかし、`attr_accessor :fizz`のようにattr_accessorにインスタンス変数の名前を渡すと自動で定義される。

```ruby
class Hoge
  def initialize(fizz)
    @fizz = fizz
  end

  # getter
  def fizz
    @fizz
  end

  # setter
  def fizz=(value)
    @fizz = value
  end
end
```

```ruby
TODO: attr_accessorの実装
```

### クラスメソッド

以下のようにクラスメソッドを定義することでクラスに対してメソッドを呼び出すことができる。

```
class Hoge
  def self.fizz
    :fizz
  end
end
Hoge.fizz # => :fizz
```

または、

```
class Hoge
  class << self
    def fizz
      :fizz
    end
  end
end
Hoge.fizz # => :fizz
```

## 2. Rubyの基礎
## 3. 制御構造/メソッド/組み込み関数
## 4. クラスとモジュール
### クラス
### モジュール
### オブジェクト
## 5. 主な組み込みクラス/モジュール
### Numeric
### String
### Regexp
### Comparable
### Enumerable
### Time
### IO/File
### Dir
### Thread
### Fiber
### Process
### Struct
### Marshal
### ObjectSpace
## 6. Rubyのクラスオブジェクト
### Classクラスからクラスを作る
### クラスオブジェクト
### 特異クラス
### メソッド探索方法
### Module#prepend
## 7. 動的なプログラミング
### オープンなクラス
### Refinements
### BasicObject#method_missing
### eval
## 8. Procオブジェクト
## 9. Methodクラス
## 10. Rubyでのリフレクションプログラミング
