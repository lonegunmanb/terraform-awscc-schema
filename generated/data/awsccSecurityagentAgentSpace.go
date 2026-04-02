package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecurityagentAgentSpace = `{
  "block": {
    "attributes": {
      "agent_space_id": {
        "computed": true,
        "description": "Unique identifier of the agent space",
        "description_kind": "plain",
        "type": "string"
      },
      "aws_resources": {
        "computed": true,
        "description": "AWS resource configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "iam_roles": {
              "computed": true,
              "description": "IAM role ARNs",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "lambda_function_arns": {
              "computed": true,
              "description": "Lambda function ARNs used to retrieve tester credentials for pentests",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "log_groups": {
              "computed": true,
              "description": "CloudWatch log group ARNs",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "s3_buckets": {
              "computed": true,
              "description": "S3 bucket ARNs",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "secret_arns": {
              "computed": true,
              "description": "SecretsManager secret ARNs used to store tester credentials for pentests",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "vpcs": {
              "computed": true,
              "description": "VPC configurations",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "security_group_arns": {
                    "computed": true,
                    "description": "List of security group ARNs in the customer VPC",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "subnet_arns": {
                    "computed": true,
                    "description": "List of subnet ARNs in the customer VPC",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "vpc_arn": {
                    "computed": true,
                    "description": "ARN of the customer VPC",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "code_review_settings": {
        "computed": true,
        "description": "Details of code review settings",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "controls_scanning": {
              "computed": true,
              "description": "Whether Controls are utilized for code review analysis",
              "description_kind": "plain",
              "type": "bool"
            },
            "general_purpose_scanning": {
              "computed": true,
              "description": "Whether general purpose analysis is performed for code review",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "created_at": {
        "computed": true,
        "description": "Timestamp when the agent space was created",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Description of the agent space",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "integrated_resources": {
        "computed": true,
        "description": "Integrated Resources configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "integration": {
              "computed": true,
              "description": "Unique identifier of the Provider Integration",
              "description_kind": "plain",
              "type": "string"
            },
            "provider_resources": {
              "computed": true,
              "description": "List of selected Resources from the Integration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "git_hub_capabilities": {
                    "computed": true,
                    "description": "GitHub repository capabilities",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "leave_comments": {
                          "computed": true,
                          "description": "Enables Code Review in the repository",
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "remediate_code": {
                          "computed": true,
                          "description": "Enables creation of pull requests with automated fixes",
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "git_hub_repository": {
                    "computed": true,
                    "description": "GitHub repository details",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
                          "computed": true,
                          "description": "GitHub repository name",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "owner": {
                          "computed": true,
                          "description": "GitHub repository owner (user or organization)",
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
            }
          },
          "nesting_mode": "list"
        }
      },
      "kms_key_id": {
        "computed": true,
        "description": "Identifier of the KMS key used to encrypt data. Can be a key ID, key ARN, alias name, or alias ARN. If not specified, an AWS managed key is used.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Name of the agent space",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags for the agent space",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "target_domain_ids": {
        "computed": true,
        "description": "List of target domain identifiers registered with the agent space",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "updated_at": {
        "computed": true,
        "description": "Timestamp when the agent space was last updated",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SecurityAgent::AgentSpace",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSecurityagentAgentSpaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecurityagentAgentSpace), &result)
	return &result
}
