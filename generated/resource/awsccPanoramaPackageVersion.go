package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPanoramaPackageVersion = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "is_latest_patch": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "mark_latest": {
        "computed": true,
        "description": "Whether to mark the new version as the latest version.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "owner_account": {
        "computed": true,
        "description": "An owner account.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "package_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "package_id": {
        "description": "A package ID.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "package_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "package_version": {
        "description": "A package version.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "patch_version": {
        "description": "A patch version.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "registered_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status_description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "updated_latest_patch_version": {
        "computed": true,
        "description": "If the version was marked latest, the new version to maker as latest.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Registers a package version.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccPanoramaPackageVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPanoramaPackageVersion), &result)
	return &result
}
