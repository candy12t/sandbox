package main

import (
	"context"
	"log/slog"
	"os"
)

func main() {
	// メッセージを出力できる
	slog.Info("hello, world")

	// key-valueを追加できる
	user := "Mike"
	age := 20
	slog.Info("hello, world", "user", user, "age", age)

	// key-valueを追加する方法2
	slog.Info("hello, world",
		slog.String("user", user),
		slog.Int("age", age),
	)

	{
		// 出力フォーマットを変更できる
		// slog.Handlerを実装すれば自由にカスマイズできる
		// ex: https://pkg.go.dev/log/slog@go1.21.0#example-Handler-LevelHandler
		// built-inで用意されているのは、slog.JSONHandler, slog.TextHandler
		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		logger.Info("hello, world", "user", user, "age", age)
	}

	{
		// 2つの方法で特定のkey-valueを含めることができる
		{
			attrs := []slog.Attr{
				{
					Key:   "method",
					Value: slog.StringValue("GET"),
				},
				{
					Key:   "path",
					Value: slog.StringValue("/api/v1/users"),
				},
			}
			logger := slog.New(slog.NewJSONHandler(os.Stdout, nil).WithAttrs(attrs))
			logger.Info("hello, world", "user", user, "age", age)
		}
		{
			logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
			logger = logger.With(
				slog.String("method", "GET"),
				slog.String("path", "/api/v1/users"),
			)
			logger.Info("hello, world", "user", user, "age", age)
		}
	}

	// グルーピング
	{
		{
			logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
			logger.Info("hello, world",
				slog.Group("user1", slog.String("name", "mike")),
				slog.Group("user2", slog.String("name", "tom")),
			)
		}
		{
			logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
			logger1 := logger.WithGroup("user1")
			logger1.Info("hello, world", slog.String("name", "mike"))

			logger2 := logger.WithGroup("user2")
			logger2.Info("hello, world", slog.String("name", "tom"))
		}
	}

	{
		// オプション
		// - ソースコードの場所を含める
		// - ログレベルの指定 slog.LevelDebug, slog.LevelInfo, slog.LevelWarn, slog.LevelError
		// - 出力される値の書き換え
		options := slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if a.Key == "email" {
					return slog.String(a.Key, "****")
				}
				return a
			},
		}
		user := "Mike"
		age := 20
		email := "mike@example.com"
		logger := slog.New(slog.NewJSONHandler(os.Stdout, &options))
		logger.Debug("hello, world",
			slog.String("user", user),
			slog.Int("age", age),
			slog.String("email", email),
		)
	}

	{
		// LogValue() slog.Value を実装することでも出力する値を書き換えられる
		// structの特定のfieldはこの方法では無理ぽい -> slog.HandlerOptionsのReplaceAttrを使う
		email := Email("mike@example.com")
		slog.Info("hello, world", "email", email)

		user := User{
			Name:  "Mike",
			Age:   20,
			Email: "mike@example.com",
		}
		logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			ReplaceAttr: protectUserEmail,
		}))
		logger.Info("hello, world", "user", user)
	}

	{
		// 高速
		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		logger.LogAttrs(context.Background(), slog.LevelInfo, "hello world", slog.String("user", "mike"))
	}
}

func protectUserEmail(groups []string, a slog.Attr) slog.Attr {
	u, ok := a.Value.Any().(User)
	if !ok {
		return a
	}
	return slog.Any(a.Key, User{
		Name:  u.Name,
		Age:   u.Age,
		Email: "****",
	})
}

type Email string

func (e Email) LogValue() slog.Value {
	return slog.StringValue("****")
}

type User struct {
	Name  string
	Age   int
	Email Email
}

// func (u *User) LogValue() slog.Value {
// 	user := User{
// 		Name:  u.Name,
// 		Age:   u.Age,
// 		Email: "****",
// 	}
// 	return slog.AnyValue(user)
// }
