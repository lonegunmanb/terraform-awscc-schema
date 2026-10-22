package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockRetrievalInstruction = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "When the retrieval instruction was created. On a numbered version, when that version was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "customer_encryption_key_arn": {
        "computed": true,
        "description": "The ARN of the customer managed KMS key protecting this content, when one was supplied at create.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the retrieval instruction. On a numbered version, the description of that version.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instruction": {
        "computed": true,
        "description": "The reviewed content this instruction serves ? what AgenticRetrieveStream resolves while status is ACTIVE. Absent until the first create review passes. Always present on a numbered version.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "answer_format": {
              "computed": true,
              "description": "How the answer is structured, including how sources are cited.",
              "description_kind": "plain",
              "type": "string"
            },
            "answer_policy": {
              "computed": true,
              "description": "Whether and how to answer: how strongly an answer must be supported, and what to do with questions outside scope.",
              "description_kind": "plain",
              "type": "string"
            },
            "persona": {
              "computed": true,
              "description": "Who the assistant is: identity, tone and language.",
              "description_kind": "plain",
              "type": "string"
            },
            "retrieval_guidance": {
              "computed": true,
              "description": "How to interpret and decompose the question, and what to prioritise when searching.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "last_transition_status": {
        "computed": true,
        "description": "The outcome or in-progress state of the most recent create or update. CREATING or UPDATING means the operation is still running. Independent of status. Always has a value from the create 202 onward. On a numbered version this is always CREATE_SUCCESSFUL: a version is created once, from reviewed content, and is immutable.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the retrieval instruction.",
        "description_kind": "plain",
        "type": "string"
      },
      "retrieval_instruction_id": {
        "computed": true,
        "description": "The identifier of the retrieval instruction.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Whether the instruction can be used. Only ACTIVE is resolvable at retrieval time. Monotone: INACTIVE to ACTIVE once, never back. Numbered versions are always ACTIVE.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "updated_at": {
        "computed": true,
        "description": "When the retrieval instruction was last updated. Equal to createdAt on a numbered version, which is immutable.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Bedrock::RetrievalInstruction",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBedrockRetrievalInstructionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockRetrievalInstruction), &result)
	return &result
}
