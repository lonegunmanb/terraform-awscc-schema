package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCodeartifactPackageGroup = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the package group.",
        "description_kind": "plain",
        "type": "string"
      },
      "contact_info": {
        "computed": true,
        "description": "The contact info of the package group.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The text description of the package group.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
        "computed": true,
        "description": "The name of the domain that contains the package group.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_owner": {
        "computed": true,
        "description": "The 12-digit account ID of the AWS account that owns the domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "origin_configuration": {
        "computed": true,
        "description": "The package origin configuration of the package group.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "restrictions": {
              "computed": true,
              "description": "The origin configuration that is applied to the package group.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "external_upstream": {
                    "computed": true,
                    "description": "The external upstream restriction determines if new package versions can be ingested or retained from external connections.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "repositories": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "restriction_mode": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "internal_upstream": {
                    "computed": true,
                    "description": "The internal upstream restriction determines if new package versions can be ingested or retained from upstream repositories.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "repositories": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "restriction_mode": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "publish": {
                    "computed": true,
                    "description": "The publish restriction determines if new package versions can be published.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "repositories": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "restriction_mode": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "pattern": {
        "computed": true,
        "description": "The package group pattern that is used to gather packages.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to the package group.",
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
              "description": "The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::CodeArtifact::PackageGroup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCodeartifactPackageGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCodeartifactPackageGroup), &result)
	return &result
}
