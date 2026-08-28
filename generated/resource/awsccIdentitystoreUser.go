package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIdentitystoreUser = `{
  "block": {
    "attributes": {
      "addresses": {
        "computed": true,
        "description": "A list of addresses associated with the user.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "country": {
              "computed": true,
              "description": "The country of the address.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "formatted": {
              "computed": true,
              "description": "A formatted version of the address for display.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "locality": {
              "computed": true,
              "description": "A string of the address locality.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "postal_code": {
              "computed": true,
              "description": "The postal code of the address.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "primary": {
              "computed": true,
              "description": "Whether this is the primary address.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "region": {
              "computed": true,
              "description": "The region of the address.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "street_address": {
              "computed": true,
              "description": "The street of the address.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "computed": true,
              "description": "The type of address.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the user.",
        "description_kind": "plain",
        "type": "string"
      },
      "birthdate": {
        "computed": true,
        "description": "The user's birthdate in YYYY-MM-DD format.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The date and time the user was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_by": {
        "computed": true,
        "description": "The identifier of the user or system that created the user.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "A string containing the name of the user for display.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "emails": {
        "computed": true,
        "description": "A list of email addresses associated with the user.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "primary": {
              "computed": true,
              "description": "Whether this is the primary email address.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "type": {
              "computed": true,
              "description": "The type of email address.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The email address.",
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
      "identity_store_id": {
        "description": "The globally unique identifier for the identity store.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "locale": {
        "computed": true,
        "description": "The geographical region or location of the user.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the user.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "family_name": {
              "computed": true,
              "description": "The family name of the user.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "formatted": {
              "computed": true,
              "description": "A string containing a formatted version of the name for display.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "given_name": {
              "computed": true,
              "description": "The given name of the user.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "honorific_prefix": {
              "computed": true,
              "description": "The honorific prefix of the user.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "honorific_suffix": {
              "computed": true,
              "description": "The honorific suffix of the user.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "middle_name": {
              "computed": true,
              "description": "The middle name of the user.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "nick_name": {
        "computed": true,
        "description": "An alternate name for the user.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "phone_numbers": {
        "computed": true,
        "description": "A list of phone numbers associated with the user.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "primary": {
              "computed": true,
              "description": "Whether this is the primary phone number.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "type": {
              "computed": true,
              "description": "The type of phone number.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The phone number.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "photos": {
        "computed": true,
        "description": "A list of photos associated with the user.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "display": {
              "computed": true,
              "description": "A display name for the photo.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "primary": {
              "computed": true,
              "description": "Whether this is the primary photo.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "type": {
              "computed": true,
              "description": "The type of photo.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The photo data or URL.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "preferred_language": {
        "computed": true,
        "description": "The preferred language of the user.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "profile_url": {
        "computed": true,
        "description": "A URL associated with the user.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "roles": {
        "computed": true,
        "description": "A list of roles associated with the user.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "primary": {
              "computed": true,
              "description": "Whether this is the primary role.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "type": {
              "computed": true,
              "description": "The type of role.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The role name.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "timezone": {
        "computed": true,
        "description": "The time zone for the user.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "title": {
        "computed": true,
        "description": "The title of the user.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description": "The date and time the user was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "updated_by": {
        "computed": true,
        "description": "The identifier of the user or system that last updated the user.",
        "description_kind": "plain",
        "type": "string"
      },
      "user_id": {
        "computed": true,
        "description": "The identifier for a user in the identity store.",
        "description_kind": "plain",
        "type": "string"
      },
      "user_name": {
        "computed": true,
        "description": "A unique string used to identify the user.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_status": {
        "computed": true,
        "description": "The current status of the user account.",
        "description_kind": "plain",
        "type": "string"
      },
      "user_type": {
        "computed": true,
        "description": "A string indicating the type of user.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "website": {
        "computed": true,
        "description": "The user's personal website or blog URL.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Creates a user within the specified identity store.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIdentitystoreUserSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIdentitystoreUser), &result)
	return &result
}
