package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectUser = `{
  "block": {
    "attributes": {
      "directory_user_id": {
        "computed": true,
        "description": "The identifier of the user account in the directory used for identity management.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hierarchy_group_arn": {
        "computed": true,
        "description": "The identifier of the hierarchy group for the user.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "identity_info": {
        "computed": true,
        "description": "The information about the identity of the user.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "email": {
              "computed": true,
              "description": "The email address. If you are using SAML for identity management and include this parameter, an error is returned.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "first_name": {
              "computed": true,
              "description": "The first name. This is required if you are using Amazon Connect or SAML for identity management.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "last_name": {
              "computed": true,
              "description": "The last name. This is required if you are using Amazon Connect or SAML for identity management.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mobile": {
              "computed": true,
              "description": "The mobile phone number.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "secondary_email": {
              "computed": true,
              "description": "The secondary email address. If you provide a secondary email, the user receives email notifications -- other than password reset notifications -- to this email address instead of to their primary email address.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "instance_arn": {
        "description": "The identifier of the Amazon Connect instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "password": {
        "computed": true,
        "description": "The password for the user account. A password is required if you are using Amazon Connect for identity management. Otherwise, it is an error to include a password.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "phone_config": {
        "description": "The phone settings for the user.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "after_contact_work_time_limit": {
              "computed": true,
              "description": "The After Call Work (ACW) timeout setting, in seconds.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "auto_accept": {
              "computed": true,
              "description": "The Auto accept setting.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "desk_phone_number": {
              "computed": true,
              "description": "The phone number for the user's desk phone.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "phone_type": {
              "description": "The phone type.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "routing_profile_arn": {
        "description": "The identifier of the routing profile for the user.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "security_profile_arns": {
        "description": "One or more security profile arns for the user",
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "One or more tags.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "user_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the user.",
        "description_kind": "plain",
        "type": "string"
      },
      "user_proficiencies": {
        "computed": true,
        "description": "One or more predefined attributes assigned to a user, with a level that indicates how skilled they are.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "attribute_name": {
              "description": "The name of user's proficiency. You must use name of predefined attribute present in the Amazon Connect instance.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "attribute_value": {
              "description": "The value of user's proficiency. You must use value of predefined attribute present in the Amazon Connect instance.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "level": {
              "description": "The level of the proficiency. The valid values are 1, 2, 3, 4 and 5.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "username": {
        "description": "The user name for the account.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Connect::User",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccConnectUserSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectUser), &result)
	return &result
}
