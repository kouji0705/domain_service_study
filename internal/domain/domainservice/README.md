## `internal/domain/domainservice` について

このパッケージは、ドメインモデル（`model`）と事例知識（`casetype`）を「組み合わせて」成立させる処理をまとめたものです。

- **質問項目の定義（`model.FormSchema`）の合成**
  - 共通質問
  - トラブル種類（pc / network など）固有の質問
  - 影響度固有の追加質問・必須条件の上書き
  - 回答内容に依存する分岐質問
- **FormSchema から `model.TroubleReport` を生成**
  - 必須回答の検証
  - `model.NewTroubleReport` で最終構築

アプリケーション層（`internal/application`）は、`model.NewTroubleReport` などを直接呼ばず、このパッケージ経由で処理を行います。

## ドキュメント

- [trouble_report_service.md](./trouble_report_service.md) — 報告書生成の窓口
- [form_schema_composer.md](./form_schema_composer.md) — 質問項目定義の合成
