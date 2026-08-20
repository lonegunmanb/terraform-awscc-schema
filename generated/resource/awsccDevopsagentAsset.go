package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDevopsagentAsset = `{
  "block": {
    "attributes": {
      "agent_space_id": {
        "description": "The unique identifier of the parent Agent Space. The asset is created as a child of this agent space.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the asset. Nested under the parent Agent Space: arn:\u003cpartition\u003e:aidevops:\u003cregion\u003e:\u003caccount-id\u003e:agentspace/\u003cagentspace-id\u003e/asset/\u003casset-id\u003e.",
        "description_kind": "plain",
        "type": "string"
      },
      "asset_id": {
        "computed": true,
        "description": "The unique identifier of the asset (assigned by the service on Create).",
        "description_kind": "plain",
        "type": "string"
      },
      "asset_type": {
        "description": "The type of asset. The Asset API treats this as an open string; call ListAssetTypes for the current authoritative set of supported types. As of launch, customer-creatable types include skill, agents_md, and attachment.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the asset was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "files": {
        "computed": true,
        "description": "Inline file list. Mutually exclusive with Zip; enforced by the handler at Create/Update time. Write-only: not repopulated by Read.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "content_bytes": {
              "computed": true,
              "description": "Base64-encoded binary contents of the file. Mutually exclusive with ContentText (max 6 MiB).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "content_text": {
              "computed": true,
              "description": "UTF-8 text contents of the file. Mutually exclusive with ContentBytes (max 1.5 MiB).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "metadata": {
              "computed": true,
              "description": "Per-file metadata document. Values may be strings, numbers, booleans, or lists of any of those (validated server-side).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "path": {
              "computed": true,
              "description": "Path of this file within the asset bundle.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "metadata": {
        "computed": true,
        "description": "Asset metadata document. Required and optional keys depend on AssetType. Values may be strings, numbers, booleans, or lists of any of those - validated server-side; see the public Asset API docs for the per-type metadata schema.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the asset was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "computed": true,
        "description": "The current asset version. Server-managed; bumps on every successful Update (including no-op updates). This is the drift signal for change detection.",
        "description_kind": "plain",
        "type": "number"
      },
      "zip": {
        "computed": true,
        "description": "Base64-encoded zip bundle containing all files for the asset. Mutually exclusive with Files; enforced by the handler at Create/Update time. Write-only: not repopulated by Read. Server treats a zip as 'replace all files' (max 6 MiB).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::DevOpsAgent::Asset. An asset attached to an existing AWS DevOps Agent Space. Customer-creatable types include skill, agents_md, and attachment; call ListAssetTypes for the current authoritative set.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDevopsagentAssetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDevopsagentAsset), &result)
	return &result
}
