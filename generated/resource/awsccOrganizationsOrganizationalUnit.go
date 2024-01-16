package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOrganizationsOrganizationalUnit = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of this OU.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "The unique identifier (ID) associated with this OU.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The friendly name of this OU.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parent_id": {
        "description": "The unique identifier (ID) of the parent root or OU that you want to create the new OU in.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of tags that you want to attach to the newly created OU.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key identifier, or name, of the tag.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The string value that's associated with the key of the tag. You can set the value of a tag to an empty string, but you can't set the value of a tag to null.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "You can use organizational units (OUs) to group accounts together to administer as a single unit. This greatly simplifies the management of your accounts. For example, you can attach a policy-based control to an OU, and all accounts within the OU automatically inherit the policy. You can create multiple OUs within a single organization, and you can create OUs within other OUs. Each OU can contain multiple accounts, and you can move accounts from one OU to another. However, OU names must be unique within a parent OU or root.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccOrganizationsOrganizationalUnitSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOrganizationsOrganizationalUnit), &result)
	return &result
}
