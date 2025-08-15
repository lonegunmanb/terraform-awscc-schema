package data

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
        "type": "string"
      },
      "hierarchy_group_arn": {
        "computed": true,
        "description": "The identifier of the hierarchy group for the user.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
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
              "type": "string"
            },
            "first_name": {
              "computed": true,
              "description": "The first name. This is required if you are using Amazon Connect or SAML for identity management.",
              "description_kind": "plain",
              "type": "string"
            },
            "last_name": {
              "computed": true,
              "description": "The last name. This is required if you are using Amazon Connect or SAML for identity management.",
              "description_kind": "plain",
              "type": "string"
            },
            "mobile": {
              "computed": true,
              "description": "The mobile phone number.",
              "description_kind": "plain",
              "type": "string"
            },
            "secondary_email": {
              "computed": true,
              "description": "The secondary email address. If you provide a secondary email, the user receives email notifications -- other than password reset notifications -- to this email address instead of to their primary email address.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "instance_arn": {
        "computed": true,
        "description": "The identifier of the Amazon Connect instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "password": {
        "computed": true,
        "description": "The password for the user account. A password is required if you are using Amazon Connect for identity management. Otherwise, it is an error to include a password.",
        "description_kind": "plain",
        "type": "string"
      },
      "phone_config": {
        "computed": true,
        "description": "The phone settings for the user.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "after_contact_work_time_limit": {
              "computed": true,
              "description": "The After Call Work (ACW) timeout setting, in seconds.",
              "description_kind": "plain",
              "type": "number"
            },
            "auto_accept": {
              "computed": true,
              "description": "The Auto accept setting.",
              "description_kind": "plain",
              "type": "bool"
            },
            "desk_phone_number": {
              "computed": true,
              "description": "The phone number for the user's desk phone.",
              "description_kind": "plain",
              "type": "string"
            },
            "persistent_connection": {
              "computed": true,
              "description": "The Persistent Connection setting.",
              "description_kind": "plain",
              "type": "bool"
            },
            "phone_type": {
              "computed": true,
              "description": "The phone type.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "routing_profile_arn": {
        "computed": true,
        "description": "The identifier of the routing profile for the user.",
        "description_kind": "plain",
        "type": "string"
      },
      "security_profile_arns": {
        "computed": true,
        "description": "One or more security profile arns for the user",
        "description_kind": "plain",
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
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
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
              "computed": true,
              "description": "The name of user's proficiency. You must use name of predefined attribute present in the Amazon Connect instance.",
              "description_kind": "plain",
              "type": "string"
            },
            "attribute_value": {
              "computed": true,
              "description": "The value of user's proficiency. You must use value of predefined attribute present in the Amazon Connect instance.",
              "description_kind": "plain",
              "type": "string"
            },
            "level": {
              "computed": true,
              "description": "The level of the proficiency. The valid values are 1, 2, 3, 4 and 5.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "list"
        }
      },
      "username": {
        "computed": true,
        "description": "The user name for the account.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Connect::User",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccConnectUserSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectUser), &result)
	return &result
}
