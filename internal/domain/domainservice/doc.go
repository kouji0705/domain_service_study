// Package domainservice は、複数のドメインモデルをまたいで協調する処理を置く。
//
// FormSchema（質問項目の定義）の合成や TroubleReport の生成など、1つの集約だけでは完結しない
// ドメイン知識をここに集める。アプリケーション層は model の生成関数を直接呼ばず、
// このパッケージを経由する。
package domainservice
