package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPollyLexicon = `{
  "block": {
    "attributes": {
      "alphabet": {
        "computed": true,
        "description": "The phonetic alphabet used in the lexicon. Valid values are ipa and x-sampa.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the lexicon.",
        "description_kind": "plain",
        "type": "string"
      },
      "content": {
        "computed": true,
        "description": "The content of the PLS lexicon in string format.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "language_code": {
        "computed": true,
        "description": "The language code that the lexicon applies to.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified": {
        "computed": true,
        "description": "The date the lexicon was last modified (ISO 8601 timestamp).",
        "description_kind": "plain",
        "type": "string"
      },
      "lexemes_count": {
        "computed": true,
        "description": "The number of lexemes in the lexicon.",
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "computed": true,
        "description": "The name of the lexicon. Must be an alphanumeric string up to 20 characters long.",
        "description_kind": "plain",
        "type": "string"
      },
      "size": {
        "computed": true,
        "description": "The total size of the lexicon, in characters.",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Data Source schema for AWS::Polly::Lexicon",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccPollyLexiconSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPollyLexicon), &result)
	return &result
}
