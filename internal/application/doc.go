// Package application は、ユースケースを置くアプリケーション層。
//
// HTTP ハンドラや CLI など外部入力を受け取り、ドメイン型へ変換して
// Domain Service を呼び出す。ビジネスルール自体は domain パッケージが持つ。
package application
