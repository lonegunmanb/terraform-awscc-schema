package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3OutpostsBucket = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the specified bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "bucket_name": {
        "computed": true,
        "description": "A name for the bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "lifecycle_configuration": {
        "computed": true,
        "description": "Rules that define how Amazon S3Outposts manages objects during their lifetime.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "rules": {
              "computed": true,
              "description": "A list of lifecycle rules for individual objects in an Amazon S3Outposts bucket.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "abort_incomplete_multipart_upload": {
                    "computed": true,
                    "description": "Specifies a lifecycle rule that stops incomplete multipart uploads to an Amazon S3Outposts bucket.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "days_after_initiation": {
                          "computed": true,
                          "description": "Specifies the number of days after which Amazon S3Outposts aborts an incomplete multipart upload.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "expiration_date": {
                    "computed": true,
                    "description": "Indicates when objects are deleted from Amazon S3Outposts. The date value must be in ISO 8601 format. The time is always midnight UTC.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "expiration_in_days": {
                    "computed": true,
                    "description": "Indicates the number of days after creation when objects are deleted from Amazon S3Outposts.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "filter": {
                    "computed": true,
                    "description": "The container for the filter of the lifecycle rule.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "and_operator": {
                          "computed": true,
                          "description": "The container for the AND condition for the lifecycle rule. A combination of Prefix and 1 or more Tags OR a minimum of 2 or more tags.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "prefix": {
                                "computed": true,
                                "description": "Prefix identifies one or more objects to which the rule applies.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "tags": {
                                "computed": true,
                                "description": "All of these tags must exist in the object's tag set in order for the rule to apply.",
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
                                  "nesting_mode": "set"
                                }
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "prefix": {
                          "computed": true,
                          "description": "Object key prefix that identifies one or more objects to which this rule applies.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "tag": {
                          "computed": true,
                          "description": "Specifies a tag used to identify a subset of objects for an Amazon S3Outposts bucket.",
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
                            "nesting_mode": "single"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "id": {
                    "computed": true,
                    "description": "Unique identifier for the lifecycle rule. The value can't be longer than 255 characters.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "status": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "outpost_id": {
        "computed": true,
        "description": "The id of the customer outpost on which the bucket resides.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An arbitrary set of tags (key-value pairs) for this S3Outposts bucket.",
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
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::S3Outposts::Bucket",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccS3OutpostsBucketSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3OutpostsBucket), &result)
	return &result
}
