package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3StorageLensGroup = `{
  "block": {
    "attributes": {
      "filter": {
        "description": "Sets the Storage Lens Group filter.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "and": {
              "computed": true,
              "description": "The Storage Lens group will include objects that match all of the specified filter values.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "match_any_prefix": {
                    "computed": true,
                    "description": "Filter to match any of the specified prefixes.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "match_any_suffix": {
                    "computed": true,
                    "description": "Filter to match any of the specified suffixes.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "match_any_tag": {
                    "computed": true,
                    "description": "Filter to match any of the specified object tags.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "key": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "set"
                    },
                    "optional": true
                  },
                  "match_object_age": {
                    "computed": true,
                    "description": "Filter to match all of the specified values for the minimum and maximum object age.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "days_greater_than": {
                          "computed": true,
                          "description": "Minimum object age to which the rule applies.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "days_less_than": {
                          "computed": true,
                          "description": "Maximum object age to which the rule applies.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "match_object_size": {
                    "computed": true,
                    "description": "Filter to match all of the specified values for the minimum and maximum object size.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bytes_greater_than": {
                          "computed": true,
                          "description": "Minimum object size to which the rule applies.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "bytes_less_than": {
                          "computed": true,
                          "description": "Maximum object size to which the rule applies.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "match_any_prefix": {
              "computed": true,
              "description": "Filter to match any of the specified prefixes.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "match_any_suffix": {
              "computed": true,
              "description": "Filter to match any of the specified suffixes.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "match_any_tag": {
              "computed": true,
              "description": "Filter to match any of the specified object tags.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              },
              "optional": true
            },
            "match_object_age": {
              "computed": true,
              "description": "Filter to match all of the specified values for the minimum and maximum object age.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "days_greater_than": {
                    "computed": true,
                    "description": "Minimum object age to which the rule applies.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "days_less_than": {
                    "computed": true,
                    "description": "Maximum object age to which the rule applies.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "match_object_size": {
              "computed": true,
              "description": "Filter to match all of the specified values for the minimum and maximum object size.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bytes_greater_than": {
                    "computed": true,
                    "description": "Minimum object size to which the rule applies.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "bytes_less_than": {
                    "computed": true,
                    "description": "Maximum object size to which the rule applies.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "or": {
              "computed": true,
              "description": "The Storage Lens group will include objects that match any of the specified filter values.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "match_any_prefix": {
                    "computed": true,
                    "description": "Filter to match any of the specified prefixes.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "match_any_suffix": {
                    "computed": true,
                    "description": "Filter to match any of the specified suffixes.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "match_any_tag": {
                    "computed": true,
                    "description": "Filter to match any of the specified object tags.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "key": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "set"
                    },
                    "optional": true
                  },
                  "match_object_age": {
                    "computed": true,
                    "description": "Filter to match all of the specified values for the minimum and maximum object age.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "days_greater_than": {
                          "computed": true,
                          "description": "Minimum object age to which the rule applies.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "days_less_than": {
                          "computed": true,
                          "description": "Maximum object age to which the rule applies.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "match_object_size": {
                    "computed": true,
                    "description": "Filter to match all of the specified values for the minimum and maximum object size.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bytes_greater_than": {
                          "computed": true,
                          "description": "Minimum object size to which the rule applies.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "bytes_less_than": {
                          "computed": true,
                          "description": "Maximum object size to which the rule applies.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name that identifies the Amazon S3 Storage Lens Group.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "storage_lens_group_arn": {
        "computed": true,
        "description": "The ARN for the Amazon S3 Storage Lens Group.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A set of tags (key-value pairs) for this Amazon S3 Storage Lens Group.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "The AWS::S3::StorageLensGroup resource is an Amazon S3 resource type that you can use to create Storage Lens Group.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccS3StorageLensGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3StorageLensGroup), &result)
	return &result
}
