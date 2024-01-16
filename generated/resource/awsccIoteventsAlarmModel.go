package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIoteventsAlarmModel = `{
  "block": {
    "attributes": {
      "alarm_capabilities": {
        "computed": true,
        "description": "Contains the configuration information of alarm state changes",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "acknowledge_flow": {
              "computed": true,
              "description": "Specifies whether to get notified for alarm state changes.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enabled": {
                    "computed": true,
                    "description": "The value must be TRUE or FALSE. If TRUE, you receive a notification when the alarm state changes. You must choose to acknowledge the notification before the alarm state can return to NORMAL. If FALSE, you won't receive notifications. The alarm automatically changes to the NORMAL state when the input property value returns to the specified range.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "initialization_configuration": {
              "computed": true,
              "description": "Specifies the default alarm state. The configuration applies to all alarms that were created based on this alarm model.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "disabled_on_initialization": {
                    "computed": true,
                    "description": "The value must be TRUE or FALSE. If FALSE, all alarm instances created based on the alarm model are activated. The default value is TRUE.",
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
      "alarm_event_actions": {
        "computed": true,
        "description": "Contains information about one or more alarm actions.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "alarm_actions": {
              "computed": true,
              "description": "Specifies one or more supported actions to receive notifications when the alarm state changes.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dynamo_d_bv_2": {
                    "computed": true,
                    "description": "Defines an action to write to the Amazon DynamoDB table that you created. The default action payload contains all attribute-value pairs that have the information about the alarm model instance and the event that triggered the action. You can also customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html). A separate column of the DynamoDB table receives one attribute-value pair in the payload that you specify.\n\nYou can use expressions for parameters that are strings. For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *AWS IoT Events Developer Guide*.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "payload": {
                          "computed": true,
                          "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the alarm model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "content_expression": {
                                "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "type": {
                                "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "table_name": {
                          "description": "The name of the DynamoDB table.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "dynamo_db": {
                    "computed": true,
                    "description": "Writes to the DynamoDB table that you created. The default action payload contains all attribute-value pairs that have the information about the alarm model instance and the event that triggered the action. You can also customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html). One column of the DynamoDB table receives all attribute-value pairs in the payload that you specify. For more information, see [Actions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-event-actions.html) in *AWS IoT Events Developer Guide*.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "hash_key_field": {
                          "description": "The name of the hash key (also called the partition key).",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "hash_key_type": {
                          "computed": true,
                          "description": "The data type for the hash key (also called the partition key). You can specify the following values:\n\n* ` + "`" + `STRING` + "`" + ` - The hash key is a string.\n\n* ` + "`" + `NUMBER` + "`" + ` - The hash key is a number.\n\nIf you don't specify ` + "`" + `hashKeyType` + "`" + `, the default value is ` + "`" + `STRING` + "`" + `.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "hash_key_value": {
                          "description": "The value of the hash key (also called the partition key).",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "operation": {
                          "computed": true,
                          "description": "The type of operation to perform. You can specify the following values:\n\n* ` + "`" + `INSERT` + "`" + ` - Insert data as a new item into the DynamoDB table. This item uses the specified hash key as a partition key. If you specified a range key, the item uses the range key as a sort key.\n\n* ` + "`" + `UPDATE` + "`" + ` - Update an existing item of the DynamoDB table with new data. This item's partition key must match the specified hash key. If you specified a range key, the range key must match the item's sort key.\n\n* ` + "`" + `DELETE` + "`" + ` - Delete an existing item of the DynamoDB table. This item's partition key must match the specified hash key. If you specified a range key, the range key must match the item's sort key.\n\nIf you don't specify this parameter, AWS IoT Events triggers the ` + "`" + `INSERT` + "`" + ` operation.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "payload": {
                          "computed": true,
                          "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the alarm model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "content_expression": {
                                "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "type": {
                                "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "payload_field": {
                          "computed": true,
                          "description": "The name of the DynamoDB column that receives the action payload.\n\nIf you don't specify this parameter, the name of the DynamoDB column is ` + "`" + `payload` + "`" + `.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "range_key_field": {
                          "computed": true,
                          "description": "The name of the range key (also called the sort key).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "range_key_type": {
                          "computed": true,
                          "description": "The data type for the range key (also called the sort key), You can specify the following values:\n\n* ` + "`" + `STRING` + "`" + ` - The range key is a string.\n\n* ` + "`" + `NUMBER` + "`" + ` - The range key is number.\n\nIf you don't specify ` + "`" + `rangeKeyField` + "`" + `, the default value is ` + "`" + `STRING` + "`" + `.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "range_key_value": {
                          "computed": true,
                          "description": "The value of the range key (also called the sort key).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "table_name": {
                          "description": "The name of the DynamoDB table.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "firehose": {
                    "computed": true,
                    "description": "Sends information about the alarm model instance and the event that triggered the action to an Amazon Kinesis Data Firehose delivery stream.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "delivery_stream_name": {
                          "description": "The name of the Kinesis Data Firehose delivery stream where the data is written.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "payload": {
                          "computed": true,
                          "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the alarm model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "content_expression": {
                                "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "type": {
                                "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "separator": {
                          "computed": true,
                          "description": "A character separator that is used to separate records written to the Kinesis Data Firehose delivery stream. Valid values are: '\\n' (newline), '\\t' (tab), '\\r\\n' (Windows newline), ',' (comma).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "iot_events": {
                    "computed": true,
                    "description": "Sends an AWS IoT Events input, passing in information about the alarm model instance and the event that triggered the action.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "input_name": {
                          "description": "The name of the AWS IoT Events input where the data is sent.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "payload": {
                          "computed": true,
                          "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the alarm model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "content_expression": {
                                "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "type": {
                                "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
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
                  "iot_site_wise": {
                    "computed": true,
                    "description": "Sends information about the alarm model instance and the event that triggered the action to a specified asset property in AWS IoT SiteWise.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "asset_id": {
                          "computed": true,
                          "description": "The ID of the asset that has the specified property. You can specify an expression.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "entry_id": {
                          "computed": true,
                          "description": "A unique identifier for this entry. You can use the entry ID to track which data entry causes an error in case of failure. The default is a new unique identifier. You can also specify an expression.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "property_alias": {
                          "computed": true,
                          "description": "The alias of the asset property. You can also specify an expression.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "property_id": {
                          "computed": true,
                          "description": "The ID of the asset property. You can specify an expression.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "property_value": {
                          "computed": true,
                          "description": "A structure that contains value information. For more information, see [AssetPropertyValue](https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_AssetPropertyValue.html) in the *AWS IoT SiteWise API Reference*.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "quality": {
                                "computed": true,
                                "description": "The quality of the asset property value. The value must be ` + "`" + `GOOD` + "`" + `, ` + "`" + `BAD` + "`" + `, or ` + "`" + `UNCERTAIN` + "`" + `. You can also specify an expression.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "timestamp": {
                                "computed": true,
                                "description": "A structure that contains timestamp information. For more information, see [TimeInNanos](https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_TimeInNanos.html) in the *AWS IoT SiteWise API Reference*.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "offset_in_nanos": {
                                      "computed": true,
                                      "description": "The timestamp, in seconds, in the Unix epoch format. The valid range is between ` + "`" + `1-31556889864403199` + "`" + `. You can also specify an expression.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "time_in_seconds": {
                                      "description": "The nanosecond offset converted from ` + "`" + `timeInSeconds` + "`" + `. The valid range is between ` + "`" + `0-999999999` + "`" + `. You can also specify an expression.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "value": {
                                "description": "A structure that contains an asset property value. For more information, see [Variant](https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_Variant.html) in the *AWS IoT SiteWise API Reference*.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "boolean_value": {
                                      "computed": true,
                                      "description": "The asset property value is a Boolean value that must be ` + "`" + `TRUE` + "`" + ` or ` + "`" + `FALSE` + "`" + `. You can also specify an expression. If you use an expression, the evaluated result should be a Boolean value.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "double_value": {
                                      "computed": true,
                                      "description": "The asset property value is a double. You can also specify an expression. If you use an expression, the evaluated result should be a double.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "integer_value": {
                                      "computed": true,
                                      "description": "The asset property value is an integer. You can also specify an expression. If you use an expression, the evaluated result should be an integer.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "string_value": {
                                      "computed": true,
                                      "description": "The asset property value is a string. You can also specify an expression. If you use an expression, the evaluated result should be a string.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
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
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "iot_topic_publish": {
                    "computed": true,
                    "description": "Information required to publish the MQTT message through the AWS IoT message broker.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "mqtt_topic": {
                          "description": "The MQTT topic of the message. You can use a string expression that includes variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `) and input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `) as the topic string.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "payload": {
                          "computed": true,
                          "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the alarm model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "content_expression": {
                                "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "type": {
                                "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
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
                  "lambda": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "function_arn": {
                          "description": "The ARN of the Lambda function that is executed.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "payload": {
                          "computed": true,
                          "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the alarm model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "content_expression": {
                                "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "type": {
                                "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
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
                  "sns": {
                    "computed": true,
                    "description": "Information required to publish the Amazon SNS message.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "payload": {
                          "computed": true,
                          "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the alarm model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "content_expression": {
                                "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "type": {
                                "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "target_arn": {
                          "description": "The ARN of the Amazon SNS target where the message is sent.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "sqs": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "payload": {
                          "computed": true,
                          "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the alarm model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "content_expression": {
                                "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "type": {
                                "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "queue_url": {
                          "description": "The URL of the SQS queue where the data is written.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "use_base_64": {
                          "computed": true,
                          "description": "Set this to ` + "`" + `TRUE` + "`" + ` if you want the data to be base-64 encoded before it is written to the queue. Otherwise, set this to ` + "`" + `FALSE` + "`" + `.",
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
                "nesting_mode": "list"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "alarm_model_description": {
        "computed": true,
        "description": "A brief description of the alarm model.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "alarm_model_name": {
        "computed": true,
        "description": "The name of the alarm model.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "alarm_rule": {
        "description": "Defines when your alarm is invoked.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "simple_rule": {
              "computed": true,
              "description": "A rule that compares an input property value to a threshold value with a comparison operator.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison_operator": {
                    "description": "The comparison operator.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "input_property": {
                    "description": "The value on the left side of the comparison operator. You can specify an AWS IoT Events input attribute as an input property.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "threshold": {
                    "description": "The value on the right side of the comparison operator. You can enter a number or specify an AWS IoT Events input attribute.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
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
      "key": {
        "computed": true,
        "description": "The value used to identify a alarm instance. When a device or system sends input, a new alarm instance with a unique key value is created. AWS IoT Events can continue to route input to its corresponding alarm instance based on this identifying information.\n\nThis parameter uses a JSON-path expression to select the attribute-value pair in the message payload that is used for identification. To route the message to the correct alarm instance, the device must send a message payload that contains the same attribute-value.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role_arn": {
        "description": "The ARN of the role that grants permission to AWS IoT Events to perform its operations.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "severity": {
        "computed": true,
        "description": "A non-negative integer that reflects the severity level of the alarm.\n\n",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.\n\nFor more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "Key of the Tag.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "Value of the Tag.",
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
    "description": "The AWS::IoTEvents::AlarmModel resource creates a alarm model. AWS IoT Events alarms help you monitor your data for changes. The data can be metrics that you measure for your equipment and processes. You can create alarms that send notifications when a threshold is breached. Alarms help you detect issues, streamline maintenance, and optimize performance of your equipment and processes.\n\nAlarms are instances of alarm models. The alarm model specifies what to detect, when to send notifications, who gets notified, and more. You can also specify one or more supported actions that occur when the alarm state changes. AWS IoT Events routes input attributes derived from your data to the appropriate alarms. If the data that you're monitoring is outside the specified range, the alarm is invoked. You can also acknowledge the alarms or set them to the snooze mode.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIoteventsAlarmModelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIoteventsAlarmModel), &result)
	return &result
}
