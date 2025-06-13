package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCustomerprofilesObjectType = `{
  "block": {
    "attributes": {
      "allow_profile_creation": {
        "computed": true,
        "description": "Indicates whether a profile should be created when data is received.",
        "description_kind": "plain",
        "type": "bool"
      },
      "created_at": {
        "computed": true,
        "description": "The time of this integration got created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Description of the profile object type.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
        "computed": true,
        "description": "The unique name of the domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "encryption_key": {
        "computed": true,
        "description": "The default encryption key",
        "description_kind": "plain",
        "type": "string"
      },
      "expiration_days": {
        "computed": true,
        "description": "The default number of days until the data within the domain expires.",
        "description_kind": "plain",
        "type": "number"
      },
      "fields": {
        "computed": true,
        "description": "A list of the name and ObjectType field.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "object_type_field": {
              "computed": true,
              "description": "Represents a field in a ProfileObjectType.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "content_type": {
                    "computed": true,
                    "description": "The content type of the field. Used for determining equality when searching.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "source": {
                    "computed": true,
                    "description": "A field of a ProfileObject. For example: _source.FirstName, where \"_source\" is a ProfileObjectType of a Zendesk user and \"FirstName\" is a field in that ObjectType.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "target": {
                    "computed": true,
                    "description": "The location of the data in the standard ProfileObject model. For example: _profile.Address.PostalCode.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "list"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "keys": {
        "computed": true,
        "description": "A list of unique keys that can be used to map data to the profile.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "object_type_key_list": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "field_names": {
                    "computed": true,
                    "description": "The reference for the key name of the fields map. ",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "standard_identifiers": {
                    "computed": true,
                    "description": "The types of keys that a ProfileObject can have. Each ProfileObject can have only 1 UNIQUE key but multiple PROFILE keys. PROFILE means that this key can be used to tie an object to a PROFILE. UNIQUE means that it can be used to uniquely identify an object. If a key a is marked as SECONDARY, it will be used to search for profiles after all other PROFILE keys have been searched. A LOOKUP_ONLY key is only used to match a profile but is not persisted to be used for searching of the profile. A NEW_ONLY key is only used if the profile does not already exist before the object is ingested, otherwise it is only used for matching objects to profiles.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "list"
        }
      },
      "last_updated_at": {
        "computed": true,
        "description": "The time of this integration got last updated at.",
        "description_kind": "plain",
        "type": "string"
      },
      "max_available_profile_object_count": {
        "computed": true,
        "description": "The maximum available number of profile objects",
        "description_kind": "plain",
        "type": "number"
      },
      "max_profile_object_count": {
        "computed": true,
        "description": "The maximum number of profile objects for this object type",
        "description_kind": "plain",
        "type": "number"
      },
      "object_type_name": {
        "computed": true,
        "description": "The name of the profile object type.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_last_updated_timestamp_format": {
        "computed": true,
        "description": "The format of your sourceLastUpdatedTimestamp that was previously set up.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags (keys and values) associated with the integration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "template_id": {
        "computed": true,
        "description": "A unique identifier for the object template.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::CustomerProfiles::ObjectType",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCustomerprofilesObjectTypeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCustomerprofilesObjectType), &result)
	return &result
}
