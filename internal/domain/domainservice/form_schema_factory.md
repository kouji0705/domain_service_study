## `form_schema_factory.go` がやること（全体像）

`FormSchemaFactory` は、条件に応じて `FormSchema` を**生成する Factory** です。

「Composer（合成する人）」という名前は、内部で共通・種類別・影響度別・分岐の各パーツを `Combine` して1つにまとめている、という実装の見え方から付けていました。ただ、外から見るとやっていることは「入力を渡したら FormSchema を作って返す」なので、より一般的な **Factory** の方が意図が伝わりやすいと判断し、改名しました。

## 返すもの / 入力するもの

- 入力
  - `troubleType`: トラブル種類（例: `pc`, `network`）
  - `impactLevel`: 影響度（例: `individual`, `team`, `company_wide`）
  - `answers`: 既にある回答（分岐の前提になる）
- 出力
  - `model.FormSchema`: この画面で聞くべき質問項目の集合

## 全体の処理手順（細かい話は後で）

`NewFormSchema` は概ね次の流れです。

1. 入力値がこのサンプルで扱えるか検査する
2. トラブル種類に応じた「基本の質問群」を用意する
3. 影響度に応じた「追加の質問群」および必須条件の上書きを行う
4. すでにある回答に応じて「分岐質問」を追加する
5. 合成で矛盾（重複など）があればエラーにする

## 詳細：モジュール構造

生成はモジュール単位で行われます。

- `casetype.CommonModule`：全報告書に共通な質問項目
- `casetype.(PCModule|NetworkModule)`：トラブル種類ごとの基本質問 + 回答依存の分岐質問
- `casetype.ImpactModule`：影響度に応じた追加質問 + 必須条件（`required`）の上書き

またトラブル種類側は `typeFormSchemaModule` という抽象で、「差し替え点」として表現しています。

## 詳細：`NewFormSchema` がやっている生成の順番

実装の順番に沿って書くと、概ね以下です。

1. `common.NewFormSchema()` を取得
2. 種類モジュールの `NewBaseFormSchema()` を `schema.Combine(...)` で合成
3. `impact.NewFormSchema(impactLevel)` を `schema.Combine(...)` で合成
4. 種類モジュールの `NewAnswerDependentFormSchema(answers)` を合成
5. 最後に `impact.ApplyRequiredOverrides(schema, impactLevel)` で `required` を調整

## エラーの出どころ（後半で見る）

- 未知の `troubleType / impactLevel`
  - `casetype.UnknownTroubleTypeError` / `casetype.UnknownImpactLevelError`
- 合成矛盾（質問ID重複など）
  - `model.FormSchema.Combine` による `DuplicateQuestionIDError`
- 分岐のための前提回答が無い/条件を満たさない
  - `NewAnswerDependentFormSchema` が空の `FormSchema{}` を返す設計（エラーではなく「追加しない」）
