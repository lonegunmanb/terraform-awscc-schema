package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRamPermissionAssociation = `{
  "block": {
    "attributes": {
      "association_status": {
        "computed": true,
        "description": "The current status of the association between the permission and the resource share.",
        "description_kind": "plain",
        "type": "string"
      },
      "feature_set": {
        "computed": true,
        "description": "The feature set of the resource share.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "is_default": {
        "computed": true,
        "description": "Indicates whether the associated resource share is using the default version of the permission.",
        "description_kind": "plain",
        "type": "bool"
      },
      "last_updated_time": {
        "computed": true,
        "description": "The date and time when the association between the permission and the resource share was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "permission_arn": {
        "description": "Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the AWS RAM permission to associate with the resource share.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "permission_version": {
        "computed": true,
        "description": "The version of the permission currently associated with the resource share.",
        "description_kind": "plain",
        "type": "string"
      },
      "replace": {
        "computed": true,
        "description": "Specifies whether to replace the existing permission on the resource share. Use ` + "`" + `true` + "`" + ` to replace the current permission. Use ` + "`" + `false` + "`" + ` to add the permission when no permission is currently associated. The default value is ` + "`" + `false` + "`" + `. Updating an existing association also requires ` + "`" + `true` + "`" + `, because AWS RAM applies the change by re-associating the permission.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "resource_share_arn": {
        "description": "Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the resource share.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_type": {
        "computed": true,
        "description": "The resource type to which the permission applies.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Associates a specified AWS RAM permission with a resource share. You can only associate one permission with each resource type in a resource share.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRamPermissionAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRamPermissionAssociation), &result)
	return &result
}
