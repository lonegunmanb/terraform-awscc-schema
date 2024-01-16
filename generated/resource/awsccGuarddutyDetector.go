package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGuarddutyDetector = `{
  "block": {
    "attributes": {
      "data_sources": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kubernetes": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "audit_logs": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "enable": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "malware_protection": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "scan_ec_2_instance_with_findings": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "ebs_volumes": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
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
            "s3_logs": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enable": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
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
      "enable": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "features": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "additional_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "status": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "status": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "finding_publishing_frequency": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::GuardDuty::Detector",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccGuarddutyDetectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGuarddutyDetector), &result)
	return &result
}
