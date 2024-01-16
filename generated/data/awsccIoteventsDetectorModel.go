package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIoteventsDetectorModel = `{
  "block": {
    "attributes": {
      "detector_model_definition": {
        "computed": true,
        "description": "Information that defines how a detector operates.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "initial_state_name": {
              "computed": true,
              "description": "The state that is entered at the creation of each detector (instance).",
              "description_kind": "plain",
              "type": "string"
            },
            "states": {
              "computed": true,
              "description": "Information about the states of the detector.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "on_enter": {
                    "computed": true,
                    "description": "When entering this state, perform these ` + "`" + `actions` + "`" + ` if the ` + "`" + `condition` + "`" + ` is ` + "`" + `TRUE` + "`" + `.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "events": {
                          "computed": true,
                          "description": "Specifies the ` + "`" + `actions` + "`" + ` that are performed when the state is entered and the ` + "`" + `condition` + "`" + ` is ` + "`" + `TRUE` + "`" + `.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "actions": {
                                "computed": true,
                                "description": "The actions to be performed.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "clear_timer": {
                                      "computed": true,
                                      "description": "Information needed to clear the timer.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "timer_name": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "dynamo_d_bv_2": {
                                      "computed": true,
                                      "description": "Defines an action to write to the Amazon DynamoDB table that you created. The default action payload contains all attribute-value pairs that have the information about the detector model instance and the event that triggered the action. You can also customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html). A separate column of the DynamoDB table receives one attribute-value pair in the payload that you specify.\n\nYou can use expressions for parameters that are strings. For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *AWS IoT Events Developer Guide*.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "table_name": {
                                            "computed": true,
                                            "description": "The name of the DynamoDB table.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "dynamo_db": {
                                      "computed": true,
                                      "description": "Writes to the DynamoDB table that you created. The default action payload contains all attribute-value pairs that have the information about the detector model instance and the event that triggered the action. You can also customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html). One column of the DynamoDB table receives all attribute-value pairs in the payload that you specify. For more information, see [Actions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-event-actions.html) in *AWS IoT Events Developer Guide*.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "hash_key_field": {
                                            "computed": true,
                                            "description": "The name of the hash key (also called the partition key).",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "hash_key_type": {
                                            "computed": true,
                                            "description": "The data type for the hash key (also called the partition key). You can specify the following values:\n\n* ` + "`" + `STRING` + "`" + ` - The hash key is a string.\n\n* ` + "`" + `NUMBER` + "`" + ` - The hash key is a number.\n\nIf you don't specify ` + "`" + `hashKeyType` + "`" + `, the default value is ` + "`" + `STRING` + "`" + `.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "hash_key_value": {
                                            "computed": true,
                                            "description": "The value of the hash key (also called the partition key).",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "operation": {
                                            "computed": true,
                                            "description": "The type of operation to perform. You can specify the following values:\n\n* ` + "`" + `INSERT` + "`" + ` - Insert data as a new item into the DynamoDB table. This item uses the specified hash key as a partition key. If you specified a range key, the item uses the range key as a sort key.\n\n* ` + "`" + `UPDATE` + "`" + ` - Update an existing item of the DynamoDB table with new data. This item's partition key must match the specified hash key. If you specified a range key, the range key must match the item's sort key.\n\n* ` + "`" + `DELETE` + "`" + ` - Delete an existing item of the DynamoDB table. This item's partition key must match the specified hash key. If you specified a range key, the range key must match the item's sort key.\n\nIf you don't specify this parameter, AWS IoT Events triggers the ` + "`" + `INSERT` + "`" + ` operation.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "payload_field": {
                                            "computed": true,
                                            "description": "The name of the DynamoDB column that receives the action payload.\n\nIf you don't specify this parameter, the name of the DynamoDB column is ` + "`" + `payload` + "`" + `.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "range_key_field": {
                                            "computed": true,
                                            "description": "The name of the range key (also called the sort key).",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "range_key_type": {
                                            "computed": true,
                                            "description": "The data type for the range key (also called the sort key), You can specify the following values:\n\n* ` + "`" + `STRING` + "`" + ` - The range key is a string.\n\n* ` + "`" + `NUMBER` + "`" + ` - The range key is number.\n\nIf you don't specify ` + "`" + `rangeKeyField` + "`" + `, the default value is ` + "`" + `STRING` + "`" + `.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "range_key_value": {
                                            "computed": true,
                                            "description": "The value of the range key (also called the sort key).",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "table_name": {
                                            "computed": true,
                                            "description": "The name of the DynamoDB table.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "firehose": {
                                      "computed": true,
                                      "description": "Sends information about the detector model instance and the event that triggered the action to an Amazon Kinesis Data Firehose delivery stream.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "delivery_stream_name": {
                                            "computed": true,
                                            "description": "The name of the Kinesis Data Firehose delivery stream where the data is written.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "separator": {
                                            "computed": true,
                                            "description": "A character separator that is used to separate records written to the Kinesis Data Firehose delivery stream. Valid values are: '\\n' (newline), '\\t' (tab), '\\r\\n' (Windows newline), ',' (comma).",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "iot_events": {
                                      "computed": true,
                                      "description": "Sends an AWS IoT Events input, passing in information about the detector model instance and the event that triggered the action.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "input_name": {
                                            "computed": true,
                                            "description": "The name of the AWS IoT Events input where the data is sent.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
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
                                    "iot_site_wise": {
                                      "computed": true,
                                      "description": "Sends information about the detector model instance and the event that triggered the action to a specified asset property in AWS IoT SiteWise.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "asset_id": {
                                            "computed": true,
                                            "description": "The ID of the asset that has the specified property. You can specify an expression.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "entry_id": {
                                            "computed": true,
                                            "description": "A unique identifier for this entry. You can use the entry ID to track which data entry causes an error in case of failure. The default is a new unique identifier. You can also specify an expression.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "property_alias": {
                                            "computed": true,
                                            "description": "The alias of the asset property. You can also specify an expression.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "property_id": {
                                            "computed": true,
                                            "description": "The ID of the asset property. You can specify an expression.",
                                            "description_kind": "plain",
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
                                                        "type": "string"
                                                      },
                                                      "time_in_seconds": {
                                                        "computed": true,
                                                        "description": "The nanosecond offset converted from ` + "`" + `timeInSeconds` + "`" + `. The valid range is between ` + "`" + `0-999999999` + "`" + `. You can also specify an expression.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "value": {
                                                  "computed": true,
                                                  "description": "A structure that contains an asset property value. For more information, see [Variant](https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_Variant.html) in the *AWS IoT SiteWise API Reference*.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "boolean_value": {
                                                        "computed": true,
                                                        "description": "The asset property value is a Boolean value that must be ` + "`" + `TRUE` + "`" + ` or ` + "`" + `FALSE` + "`" + `. You can also specify an expression. If you use an expression, the evaluated result should be a Boolean value.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "double_value": {
                                                        "computed": true,
                                                        "description": "The asset property value is a double. You can also specify an expression. If you use an expression, the evaluated result should be a double.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "integer_value": {
                                                        "computed": true,
                                                        "description": "The asset property value is an integer. You can also specify an expression. If you use an expression, the evaluated result should be an integer.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "string_value": {
                                                        "computed": true,
                                                        "description": "The asset property value is a string. You can also specify an expression. If you use an expression, the evaluated result should be a string.",
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
                                    "iot_topic_publish": {
                                      "computed": true,
                                      "description": "Information required to publish the MQTT message through the AWS IoT message broker.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "mqtt_topic": {
                                            "computed": true,
                                            "description": "The MQTT topic of the message. You can use a string expression that includes variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `) and input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `) as the topic string.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
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
                                    "lambda": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "function_arn": {
                                            "computed": true,
                                            "description": "The ARN of the Lambda function that is executed.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
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
                                    "reset_timer": {
                                      "computed": true,
                                      "description": "Information required to reset the timer. The timer is reset to the previously evaluated result of the duration. The duration expression isn't reevaluated when you reset the timer.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "timer_name": {
                                            "computed": true,
                                            "description": "The name of the timer to reset.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "set_timer": {
                                      "computed": true,
                                      "description": "Information needed to set the timer.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "duration_expression": {
                                            "computed": true,
                                            "description": "The duration of the timer, in seconds. You can use a string expression that includes numbers, variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), and input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `) as the duration. The range of the duration is ` + "`" + `1-31622400` + "`" + ` seconds. To ensure accuracy, the minimum duration is ` + "`" + `60` + "`" + ` seconds. The evaluated result of the duration is rounded down to the nearest whole number.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "seconds": {
                                            "computed": true,
                                            "description": "The number of seconds until the timer expires. The minimum value is ` + "`" + `60` + "`" + ` seconds to ensure accuracy. The maximum value is ` + "`" + `31622400` + "`" + ` seconds.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          },
                                          "timer_name": {
                                            "computed": true,
                                            "description": "The name of the timer.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "set_variable": {
                                      "computed": true,
                                      "description": "Information about the variable and its new value.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "value": {
                                            "computed": true,
                                            "description": "The new value of the variable.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "variable_name": {
                                            "computed": true,
                                            "description": "The name of the variable.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "sns": {
                                      "computed": true,
                                      "description": "Information required to publish the Amazon SNS message.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "target_arn": {
                                            "computed": true,
                                            "description": "The ARN of the Amazon SNS target where the message is sent.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "sqs": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "queue_url": {
                                            "computed": true,
                                            "description": "The URL of the SQS queue where the data is written.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "use_base_64": {
                                            "computed": true,
                                            "description": "Set this to ` + "`" + `TRUE` + "`" + ` if you want the data to be base-64 encoded before it is written to the queue. Otherwise, set this to ` + "`" + `FALSE` + "`" + `.",
                                            "description_kind": "plain",
                                            "type": "bool"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    }
                                  },
                                  "nesting_mode": "list"
                                }
                              },
                              "condition": {
                                "computed": true,
                                "description": "The Boolean expression that, when ` + "`" + `TRUE` + "`" + `, causes the ` + "`" + `actions` + "`" + ` to be performed. If not present, the ` + "`" + `actions` + "`" + ` are performed (=` + "`" + `TRUE` + "`" + `). If the expression result is not a ` + "`" + `Boolean` + "`" + ` value, the ` + "`" + `actions` + "`" + ` are not performed (=` + "`" + `FALSE` + "`" + `).",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "event_name": {
                                "computed": true,
                                "description": "The name of the event.",
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
                  "on_exit": {
                    "computed": true,
                    "description": "When exiting this state, perform these ` + "`" + `actions` + "`" + ` if the specified ` + "`" + `condition` + "`" + ` is ` + "`" + `TRUE` + "`" + `.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "events": {
                          "computed": true,
                          "description": "Specifies the ` + "`" + `actions` + "`" + ` that are performed when the state is exited and the ` + "`" + `condition` + "`" + ` is ` + "`" + `TRUE` + "`" + `.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "actions": {
                                "computed": true,
                                "description": "The actions to be performed.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "clear_timer": {
                                      "computed": true,
                                      "description": "Information needed to clear the timer.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "timer_name": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "dynamo_d_bv_2": {
                                      "computed": true,
                                      "description": "Defines an action to write to the Amazon DynamoDB table that you created. The default action payload contains all attribute-value pairs that have the information about the detector model instance and the event that triggered the action. You can also customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html). A separate column of the DynamoDB table receives one attribute-value pair in the payload that you specify.\n\nYou can use expressions for parameters that are strings. For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *AWS IoT Events Developer Guide*.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "table_name": {
                                            "computed": true,
                                            "description": "The name of the DynamoDB table.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "dynamo_db": {
                                      "computed": true,
                                      "description": "Writes to the DynamoDB table that you created. The default action payload contains all attribute-value pairs that have the information about the detector model instance and the event that triggered the action. You can also customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html). One column of the DynamoDB table receives all attribute-value pairs in the payload that you specify. For more information, see [Actions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-event-actions.html) in *AWS IoT Events Developer Guide*.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "hash_key_field": {
                                            "computed": true,
                                            "description": "The name of the hash key (also called the partition key).",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "hash_key_type": {
                                            "computed": true,
                                            "description": "The data type for the hash key (also called the partition key). You can specify the following values:\n\n* ` + "`" + `STRING` + "`" + ` - The hash key is a string.\n\n* ` + "`" + `NUMBER` + "`" + ` - The hash key is a number.\n\nIf you don't specify ` + "`" + `hashKeyType` + "`" + `, the default value is ` + "`" + `STRING` + "`" + `.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "hash_key_value": {
                                            "computed": true,
                                            "description": "The value of the hash key (also called the partition key).",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "operation": {
                                            "computed": true,
                                            "description": "The type of operation to perform. You can specify the following values:\n\n* ` + "`" + `INSERT` + "`" + ` - Insert data as a new item into the DynamoDB table. This item uses the specified hash key as a partition key. If you specified a range key, the item uses the range key as a sort key.\n\n* ` + "`" + `UPDATE` + "`" + ` - Update an existing item of the DynamoDB table with new data. This item's partition key must match the specified hash key. If you specified a range key, the range key must match the item's sort key.\n\n* ` + "`" + `DELETE` + "`" + ` - Delete an existing item of the DynamoDB table. This item's partition key must match the specified hash key. If you specified a range key, the range key must match the item's sort key.\n\nIf you don't specify this parameter, AWS IoT Events triggers the ` + "`" + `INSERT` + "`" + ` operation.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "payload_field": {
                                            "computed": true,
                                            "description": "The name of the DynamoDB column that receives the action payload.\n\nIf you don't specify this parameter, the name of the DynamoDB column is ` + "`" + `payload` + "`" + `.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "range_key_field": {
                                            "computed": true,
                                            "description": "The name of the range key (also called the sort key).",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "range_key_type": {
                                            "computed": true,
                                            "description": "The data type for the range key (also called the sort key), You can specify the following values:\n\n* ` + "`" + `STRING` + "`" + ` - The range key is a string.\n\n* ` + "`" + `NUMBER` + "`" + ` - The range key is number.\n\nIf you don't specify ` + "`" + `rangeKeyField` + "`" + `, the default value is ` + "`" + `STRING` + "`" + `.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "range_key_value": {
                                            "computed": true,
                                            "description": "The value of the range key (also called the sort key).",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "table_name": {
                                            "computed": true,
                                            "description": "The name of the DynamoDB table.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "firehose": {
                                      "computed": true,
                                      "description": "Sends information about the detector model instance and the event that triggered the action to an Amazon Kinesis Data Firehose delivery stream.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "delivery_stream_name": {
                                            "computed": true,
                                            "description": "The name of the Kinesis Data Firehose delivery stream where the data is written.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "separator": {
                                            "computed": true,
                                            "description": "A character separator that is used to separate records written to the Kinesis Data Firehose delivery stream. Valid values are: '\\n' (newline), '\\t' (tab), '\\r\\n' (Windows newline), ',' (comma).",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "iot_events": {
                                      "computed": true,
                                      "description": "Sends an AWS IoT Events input, passing in information about the detector model instance and the event that triggered the action.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "input_name": {
                                            "computed": true,
                                            "description": "The name of the AWS IoT Events input where the data is sent.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
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
                                    "iot_site_wise": {
                                      "computed": true,
                                      "description": "Sends information about the detector model instance and the event that triggered the action to a specified asset property in AWS IoT SiteWise.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "asset_id": {
                                            "computed": true,
                                            "description": "The ID of the asset that has the specified property. You can specify an expression.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "entry_id": {
                                            "computed": true,
                                            "description": "A unique identifier for this entry. You can use the entry ID to track which data entry causes an error in case of failure. The default is a new unique identifier. You can also specify an expression.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "property_alias": {
                                            "computed": true,
                                            "description": "The alias of the asset property. You can also specify an expression.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "property_id": {
                                            "computed": true,
                                            "description": "The ID of the asset property. You can specify an expression.",
                                            "description_kind": "plain",
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
                                                        "type": "string"
                                                      },
                                                      "time_in_seconds": {
                                                        "computed": true,
                                                        "description": "The nanosecond offset converted from ` + "`" + `timeInSeconds` + "`" + `. The valid range is between ` + "`" + `0-999999999` + "`" + `. You can also specify an expression.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "value": {
                                                  "computed": true,
                                                  "description": "A structure that contains an asset property value. For more information, see [Variant](https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_Variant.html) in the *AWS IoT SiteWise API Reference*.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "boolean_value": {
                                                        "computed": true,
                                                        "description": "The asset property value is a Boolean value that must be ` + "`" + `TRUE` + "`" + ` or ` + "`" + `FALSE` + "`" + `. You can also specify an expression. If you use an expression, the evaluated result should be a Boolean value.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "double_value": {
                                                        "computed": true,
                                                        "description": "The asset property value is a double. You can also specify an expression. If you use an expression, the evaluated result should be a double.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "integer_value": {
                                                        "computed": true,
                                                        "description": "The asset property value is an integer. You can also specify an expression. If you use an expression, the evaluated result should be an integer.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "string_value": {
                                                        "computed": true,
                                                        "description": "The asset property value is a string. You can also specify an expression. If you use an expression, the evaluated result should be a string.",
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
                                    "iot_topic_publish": {
                                      "computed": true,
                                      "description": "Information required to publish the MQTT message through the AWS IoT message broker.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "mqtt_topic": {
                                            "computed": true,
                                            "description": "The MQTT topic of the message. You can use a string expression that includes variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `) and input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `) as the topic string.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
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
                                    "lambda": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "function_arn": {
                                            "computed": true,
                                            "description": "The ARN of the Lambda function that is executed.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
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
                                    "reset_timer": {
                                      "computed": true,
                                      "description": "Information required to reset the timer. The timer is reset to the previously evaluated result of the duration. The duration expression isn't reevaluated when you reset the timer.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "timer_name": {
                                            "computed": true,
                                            "description": "The name of the timer to reset.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "set_timer": {
                                      "computed": true,
                                      "description": "Information needed to set the timer.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "duration_expression": {
                                            "computed": true,
                                            "description": "The duration of the timer, in seconds. You can use a string expression that includes numbers, variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), and input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `) as the duration. The range of the duration is ` + "`" + `1-31622400` + "`" + ` seconds. To ensure accuracy, the minimum duration is ` + "`" + `60` + "`" + ` seconds. The evaluated result of the duration is rounded down to the nearest whole number.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "seconds": {
                                            "computed": true,
                                            "description": "The number of seconds until the timer expires. The minimum value is ` + "`" + `60` + "`" + ` seconds to ensure accuracy. The maximum value is ` + "`" + `31622400` + "`" + ` seconds.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          },
                                          "timer_name": {
                                            "computed": true,
                                            "description": "The name of the timer.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "set_variable": {
                                      "computed": true,
                                      "description": "Information about the variable and its new value.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "value": {
                                            "computed": true,
                                            "description": "The new value of the variable.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "variable_name": {
                                            "computed": true,
                                            "description": "The name of the variable.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "sns": {
                                      "computed": true,
                                      "description": "Information required to publish the Amazon SNS message.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "target_arn": {
                                            "computed": true,
                                            "description": "The ARN of the Amazon SNS target where the message is sent.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "sqs": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "queue_url": {
                                            "computed": true,
                                            "description": "The URL of the SQS queue where the data is written.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "use_base_64": {
                                            "computed": true,
                                            "description": "Set this to ` + "`" + `TRUE` + "`" + ` if you want the data to be base-64 encoded before it is written to the queue. Otherwise, set this to ` + "`" + `FALSE` + "`" + `.",
                                            "description_kind": "plain",
                                            "type": "bool"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    }
                                  },
                                  "nesting_mode": "list"
                                }
                              },
                              "condition": {
                                "computed": true,
                                "description": "The Boolean expression that, when ` + "`" + `TRUE` + "`" + `, causes the ` + "`" + `actions` + "`" + ` to be performed. If not present, the ` + "`" + `actions` + "`" + ` are performed (=` + "`" + `TRUE` + "`" + `). If the expression result is not a ` + "`" + `Boolean` + "`" + ` value, the ` + "`" + `actions` + "`" + ` are not performed (=` + "`" + `FALSE` + "`" + `).",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "event_name": {
                                "computed": true,
                                "description": "The name of the event.",
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
                  "on_input": {
                    "computed": true,
                    "description": "When an input is received and the ` + "`" + `condition` + "`" + ` is ` + "`" + `TRUE` + "`" + `, perform the specified ` + "`" + `actions` + "`" + `.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "events": {
                          "computed": true,
                          "description": "Specifies the ` + "`" + `actions` + "`" + ` performed when the ` + "`" + `condition` + "`" + ` evaluates to ` + "`" + `TRUE` + "`" + `.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "actions": {
                                "computed": true,
                                "description": "The actions to be performed.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "clear_timer": {
                                      "computed": true,
                                      "description": "Information needed to clear the timer.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "timer_name": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "dynamo_d_bv_2": {
                                      "computed": true,
                                      "description": "Defines an action to write to the Amazon DynamoDB table that you created. The default action payload contains all attribute-value pairs that have the information about the detector model instance and the event that triggered the action. You can also customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html). A separate column of the DynamoDB table receives one attribute-value pair in the payload that you specify.\n\nYou can use expressions for parameters that are strings. For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *AWS IoT Events Developer Guide*.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "table_name": {
                                            "computed": true,
                                            "description": "The name of the DynamoDB table.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "dynamo_db": {
                                      "computed": true,
                                      "description": "Writes to the DynamoDB table that you created. The default action payload contains all attribute-value pairs that have the information about the detector model instance and the event that triggered the action. You can also customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html). One column of the DynamoDB table receives all attribute-value pairs in the payload that you specify. For more information, see [Actions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-event-actions.html) in *AWS IoT Events Developer Guide*.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "hash_key_field": {
                                            "computed": true,
                                            "description": "The name of the hash key (also called the partition key).",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "hash_key_type": {
                                            "computed": true,
                                            "description": "The data type for the hash key (also called the partition key). You can specify the following values:\n\n* ` + "`" + `STRING` + "`" + ` - The hash key is a string.\n\n* ` + "`" + `NUMBER` + "`" + ` - The hash key is a number.\n\nIf you don't specify ` + "`" + `hashKeyType` + "`" + `, the default value is ` + "`" + `STRING` + "`" + `.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "hash_key_value": {
                                            "computed": true,
                                            "description": "The value of the hash key (also called the partition key).",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "operation": {
                                            "computed": true,
                                            "description": "The type of operation to perform. You can specify the following values:\n\n* ` + "`" + `INSERT` + "`" + ` - Insert data as a new item into the DynamoDB table. This item uses the specified hash key as a partition key. If you specified a range key, the item uses the range key as a sort key.\n\n* ` + "`" + `UPDATE` + "`" + ` - Update an existing item of the DynamoDB table with new data. This item's partition key must match the specified hash key. If you specified a range key, the range key must match the item's sort key.\n\n* ` + "`" + `DELETE` + "`" + ` - Delete an existing item of the DynamoDB table. This item's partition key must match the specified hash key. If you specified a range key, the range key must match the item's sort key.\n\nIf you don't specify this parameter, AWS IoT Events triggers the ` + "`" + `INSERT` + "`" + ` operation.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "payload_field": {
                                            "computed": true,
                                            "description": "The name of the DynamoDB column that receives the action payload.\n\nIf you don't specify this parameter, the name of the DynamoDB column is ` + "`" + `payload` + "`" + `.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "range_key_field": {
                                            "computed": true,
                                            "description": "The name of the range key (also called the sort key).",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "range_key_type": {
                                            "computed": true,
                                            "description": "The data type for the range key (also called the sort key), You can specify the following values:\n\n* ` + "`" + `STRING` + "`" + ` - The range key is a string.\n\n* ` + "`" + `NUMBER` + "`" + ` - The range key is number.\n\nIf you don't specify ` + "`" + `rangeKeyField` + "`" + `, the default value is ` + "`" + `STRING` + "`" + `.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "range_key_value": {
                                            "computed": true,
                                            "description": "The value of the range key (also called the sort key).",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "table_name": {
                                            "computed": true,
                                            "description": "The name of the DynamoDB table.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "firehose": {
                                      "computed": true,
                                      "description": "Sends information about the detector model instance and the event that triggered the action to an Amazon Kinesis Data Firehose delivery stream.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "delivery_stream_name": {
                                            "computed": true,
                                            "description": "The name of the Kinesis Data Firehose delivery stream where the data is written.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "separator": {
                                            "computed": true,
                                            "description": "A character separator that is used to separate records written to the Kinesis Data Firehose delivery stream. Valid values are: '\\n' (newline), '\\t' (tab), '\\r\\n' (Windows newline), ',' (comma).",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "iot_events": {
                                      "computed": true,
                                      "description": "Sends an AWS IoT Events input, passing in information about the detector model instance and the event that triggered the action.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "input_name": {
                                            "computed": true,
                                            "description": "The name of the AWS IoT Events input where the data is sent.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
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
                                    "iot_site_wise": {
                                      "computed": true,
                                      "description": "Sends information about the detector model instance and the event that triggered the action to a specified asset property in AWS IoT SiteWise.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "asset_id": {
                                            "computed": true,
                                            "description": "The ID of the asset that has the specified property. You can specify an expression.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "entry_id": {
                                            "computed": true,
                                            "description": "A unique identifier for this entry. You can use the entry ID to track which data entry causes an error in case of failure. The default is a new unique identifier. You can also specify an expression.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "property_alias": {
                                            "computed": true,
                                            "description": "The alias of the asset property. You can also specify an expression.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "property_id": {
                                            "computed": true,
                                            "description": "The ID of the asset property. You can specify an expression.",
                                            "description_kind": "plain",
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
                                                        "type": "string"
                                                      },
                                                      "time_in_seconds": {
                                                        "computed": true,
                                                        "description": "The nanosecond offset converted from ` + "`" + `timeInSeconds` + "`" + `. The valid range is between ` + "`" + `0-999999999` + "`" + `. You can also specify an expression.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "value": {
                                                  "computed": true,
                                                  "description": "A structure that contains an asset property value. For more information, see [Variant](https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_Variant.html) in the *AWS IoT SiteWise API Reference*.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "boolean_value": {
                                                        "computed": true,
                                                        "description": "The asset property value is a Boolean value that must be ` + "`" + `TRUE` + "`" + ` or ` + "`" + `FALSE` + "`" + `. You can also specify an expression. If you use an expression, the evaluated result should be a Boolean value.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "double_value": {
                                                        "computed": true,
                                                        "description": "The asset property value is a double. You can also specify an expression. If you use an expression, the evaluated result should be a double.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "integer_value": {
                                                        "computed": true,
                                                        "description": "The asset property value is an integer. You can also specify an expression. If you use an expression, the evaluated result should be an integer.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "string_value": {
                                                        "computed": true,
                                                        "description": "The asset property value is a string. You can also specify an expression. If you use an expression, the evaluated result should be a string.",
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
                                    "iot_topic_publish": {
                                      "computed": true,
                                      "description": "Information required to publish the MQTT message through the AWS IoT message broker.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "mqtt_topic": {
                                            "computed": true,
                                            "description": "The MQTT topic of the message. You can use a string expression that includes variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `) and input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `) as the topic string.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
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
                                    "lambda": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "function_arn": {
                                            "computed": true,
                                            "description": "The ARN of the Lambda function that is executed.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
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
                                    "reset_timer": {
                                      "computed": true,
                                      "description": "Information required to reset the timer. The timer is reset to the previously evaluated result of the duration. The duration expression isn't reevaluated when you reset the timer.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "timer_name": {
                                            "computed": true,
                                            "description": "The name of the timer to reset.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "set_timer": {
                                      "computed": true,
                                      "description": "Information needed to set the timer.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "duration_expression": {
                                            "computed": true,
                                            "description": "The duration of the timer, in seconds. You can use a string expression that includes numbers, variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), and input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `) as the duration. The range of the duration is ` + "`" + `1-31622400` + "`" + ` seconds. To ensure accuracy, the minimum duration is ` + "`" + `60` + "`" + ` seconds. The evaluated result of the duration is rounded down to the nearest whole number.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "seconds": {
                                            "computed": true,
                                            "description": "The number of seconds until the timer expires. The minimum value is ` + "`" + `60` + "`" + ` seconds to ensure accuracy. The maximum value is ` + "`" + `31622400` + "`" + ` seconds.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          },
                                          "timer_name": {
                                            "computed": true,
                                            "description": "The name of the timer.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "set_variable": {
                                      "computed": true,
                                      "description": "Information about the variable and its new value.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "value": {
                                            "computed": true,
                                            "description": "The new value of the variable.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "variable_name": {
                                            "computed": true,
                                            "description": "The name of the variable.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "sns": {
                                      "computed": true,
                                      "description": "Information required to publish the Amazon SNS message.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "target_arn": {
                                            "computed": true,
                                            "description": "The ARN of the Amazon SNS target where the message is sent.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "sqs": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "queue_url": {
                                            "computed": true,
                                            "description": "The URL of the SQS queue where the data is written.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "use_base_64": {
                                            "computed": true,
                                            "description": "Set this to ` + "`" + `TRUE` + "`" + ` if you want the data to be base-64 encoded before it is written to the queue. Otherwise, set this to ` + "`" + `FALSE` + "`" + `.",
                                            "description_kind": "plain",
                                            "type": "bool"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    }
                                  },
                                  "nesting_mode": "list"
                                }
                              },
                              "condition": {
                                "computed": true,
                                "description": "The Boolean expression that, when ` + "`" + `TRUE` + "`" + `, causes the ` + "`" + `actions` + "`" + ` to be performed. If not present, the ` + "`" + `actions` + "`" + ` are performed (=` + "`" + `TRUE` + "`" + `). If the expression result is not a ` + "`" + `Boolean` + "`" + ` value, the ` + "`" + `actions` + "`" + ` are not performed (=` + "`" + `FALSE` + "`" + `).",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "event_name": {
                                "computed": true,
                                "description": "The name of the event.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "transition_events": {
                          "computed": true,
                          "description": "Specifies the ` + "`" + `actions` + "`" + ` performed, and the next ` + "`" + `state` + "`" + ` entered, when a ` + "`" + `condition` + "`" + ` evaluates to ` + "`" + `TRUE` + "`" + `.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "actions": {
                                "computed": true,
                                "description": "The actions to be performed.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "clear_timer": {
                                      "computed": true,
                                      "description": "Information needed to clear the timer.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "timer_name": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "dynamo_d_bv_2": {
                                      "computed": true,
                                      "description": "Defines an action to write to the Amazon DynamoDB table that you created. The default action payload contains all attribute-value pairs that have the information about the detector model instance and the event that triggered the action. You can also customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html). A separate column of the DynamoDB table receives one attribute-value pair in the payload that you specify.\n\nYou can use expressions for parameters that are strings. For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *AWS IoT Events Developer Guide*.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "table_name": {
                                            "computed": true,
                                            "description": "The name of the DynamoDB table.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "dynamo_db": {
                                      "computed": true,
                                      "description": "Writes to the DynamoDB table that you created. The default action payload contains all attribute-value pairs that have the information about the detector model instance and the event that triggered the action. You can also customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html). One column of the DynamoDB table receives all attribute-value pairs in the payload that you specify. For more information, see [Actions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-event-actions.html) in *AWS IoT Events Developer Guide*.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "hash_key_field": {
                                            "computed": true,
                                            "description": "The name of the hash key (also called the partition key).",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "hash_key_type": {
                                            "computed": true,
                                            "description": "The data type for the hash key (also called the partition key). You can specify the following values:\n\n* ` + "`" + `STRING` + "`" + ` - The hash key is a string.\n\n* ` + "`" + `NUMBER` + "`" + ` - The hash key is a number.\n\nIf you don't specify ` + "`" + `hashKeyType` + "`" + `, the default value is ` + "`" + `STRING` + "`" + `.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "hash_key_value": {
                                            "computed": true,
                                            "description": "The value of the hash key (also called the partition key).",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "operation": {
                                            "computed": true,
                                            "description": "The type of operation to perform. You can specify the following values:\n\n* ` + "`" + `INSERT` + "`" + ` - Insert data as a new item into the DynamoDB table. This item uses the specified hash key as a partition key. If you specified a range key, the item uses the range key as a sort key.\n\n* ` + "`" + `UPDATE` + "`" + ` - Update an existing item of the DynamoDB table with new data. This item's partition key must match the specified hash key. If you specified a range key, the range key must match the item's sort key.\n\n* ` + "`" + `DELETE` + "`" + ` - Delete an existing item of the DynamoDB table. This item's partition key must match the specified hash key. If you specified a range key, the range key must match the item's sort key.\n\nIf you don't specify this parameter, AWS IoT Events triggers the ` + "`" + `INSERT` + "`" + ` operation.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "payload_field": {
                                            "computed": true,
                                            "description": "The name of the DynamoDB column that receives the action payload.\n\nIf you don't specify this parameter, the name of the DynamoDB column is ` + "`" + `payload` + "`" + `.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "range_key_field": {
                                            "computed": true,
                                            "description": "The name of the range key (also called the sort key).",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "range_key_type": {
                                            "computed": true,
                                            "description": "The data type for the range key (also called the sort key), You can specify the following values:\n\n* ` + "`" + `STRING` + "`" + ` - The range key is a string.\n\n* ` + "`" + `NUMBER` + "`" + ` - The range key is number.\n\nIf you don't specify ` + "`" + `rangeKeyField` + "`" + `, the default value is ` + "`" + `STRING` + "`" + `.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "range_key_value": {
                                            "computed": true,
                                            "description": "The value of the range key (also called the sort key).",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "table_name": {
                                            "computed": true,
                                            "description": "The name of the DynamoDB table.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "firehose": {
                                      "computed": true,
                                      "description": "Sends information about the detector model instance and the event that triggered the action to an Amazon Kinesis Data Firehose delivery stream.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "delivery_stream_name": {
                                            "computed": true,
                                            "description": "The name of the Kinesis Data Firehose delivery stream where the data is written.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "separator": {
                                            "computed": true,
                                            "description": "A character separator that is used to separate records written to the Kinesis Data Firehose delivery stream. Valid values are: '\\n' (newline), '\\t' (tab), '\\r\\n' (Windows newline), ',' (comma).",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "iot_events": {
                                      "computed": true,
                                      "description": "Sends an AWS IoT Events input, passing in information about the detector model instance and the event that triggered the action.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "input_name": {
                                            "computed": true,
                                            "description": "The name of the AWS IoT Events input where the data is sent.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
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
                                    "iot_site_wise": {
                                      "computed": true,
                                      "description": "Sends information about the detector model instance and the event that triggered the action to a specified asset property in AWS IoT SiteWise.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "asset_id": {
                                            "computed": true,
                                            "description": "The ID of the asset that has the specified property. You can specify an expression.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "entry_id": {
                                            "computed": true,
                                            "description": "A unique identifier for this entry. You can use the entry ID to track which data entry causes an error in case of failure. The default is a new unique identifier. You can also specify an expression.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "property_alias": {
                                            "computed": true,
                                            "description": "The alias of the asset property. You can also specify an expression.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "property_id": {
                                            "computed": true,
                                            "description": "The ID of the asset property. You can specify an expression.",
                                            "description_kind": "plain",
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
                                                        "type": "string"
                                                      },
                                                      "time_in_seconds": {
                                                        "computed": true,
                                                        "description": "The nanosecond offset converted from ` + "`" + `timeInSeconds` + "`" + `. The valid range is between ` + "`" + `0-999999999` + "`" + `. You can also specify an expression.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  }
                                                },
                                                "value": {
                                                  "computed": true,
                                                  "description": "A structure that contains an asset property value. For more information, see [Variant](https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_Variant.html) in the *AWS IoT SiteWise API Reference*.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "boolean_value": {
                                                        "computed": true,
                                                        "description": "The asset property value is a Boolean value that must be ` + "`" + `TRUE` + "`" + ` or ` + "`" + `FALSE` + "`" + `. You can also specify an expression. If you use an expression, the evaluated result should be a Boolean value.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "double_value": {
                                                        "computed": true,
                                                        "description": "The asset property value is a double. You can also specify an expression. If you use an expression, the evaluated result should be a double.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "integer_value": {
                                                        "computed": true,
                                                        "description": "The asset property value is an integer. You can also specify an expression. If you use an expression, the evaluated result should be an integer.",
                                                        "description_kind": "plain",
                                                        "type": "string"
                                                      },
                                                      "string_value": {
                                                        "computed": true,
                                                        "description": "The asset property value is a string. You can also specify an expression. If you use an expression, the evaluated result should be a string.",
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
                                    "iot_topic_publish": {
                                      "computed": true,
                                      "description": "Information required to publish the MQTT message through the AWS IoT message broker.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "mqtt_topic": {
                                            "computed": true,
                                            "description": "The MQTT topic of the message. You can use a string expression that includes variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `) and input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `) as the topic string.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
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
                                    "lambda": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "function_arn": {
                                            "computed": true,
                                            "description": "The ARN of the Lambda function that is executed.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
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
                                    "reset_timer": {
                                      "computed": true,
                                      "description": "Information required to reset the timer. The timer is reset to the previously evaluated result of the duration. The duration expression isn't reevaluated when you reset the timer.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "timer_name": {
                                            "computed": true,
                                            "description": "The name of the timer to reset.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "set_timer": {
                                      "computed": true,
                                      "description": "Information needed to set the timer.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "duration_expression": {
                                            "computed": true,
                                            "description": "The duration of the timer, in seconds. You can use a string expression that includes numbers, variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), and input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `) as the duration. The range of the duration is ` + "`" + `1-31622400` + "`" + ` seconds. To ensure accuracy, the minimum duration is ` + "`" + `60` + "`" + ` seconds. The evaluated result of the duration is rounded down to the nearest whole number.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "seconds": {
                                            "computed": true,
                                            "description": "The number of seconds until the timer expires. The minimum value is ` + "`" + `60` + "`" + ` seconds to ensure accuracy. The maximum value is ` + "`" + `31622400` + "`" + ` seconds.",
                                            "description_kind": "plain",
                                            "type": "number"
                                          },
                                          "timer_name": {
                                            "computed": true,
                                            "description": "The name of the timer.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "set_variable": {
                                      "computed": true,
                                      "description": "Information about the variable and its new value.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "value": {
                                            "computed": true,
                                            "description": "The new value of the variable.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "variable_name": {
                                            "computed": true,
                                            "description": "The name of the variable.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "sns": {
                                      "computed": true,
                                      "description": "Information required to publish the Amazon SNS message.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "target_arn": {
                                            "computed": true,
                                            "description": "The ARN of the Amazon SNS target where the message is sent.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    },
                                    "sqs": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n\nBy default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `contentExpression` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `'\u003cstring\u003e'` + "`" + `), variables (` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `), input values (` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `${}` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `STRING` + "`" + ` or ` + "`" + `JSON` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            }
                                          },
                                          "queue_url": {
                                            "computed": true,
                                            "description": "The URL of the SQS queue where the data is written.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "use_base_64": {
                                            "computed": true,
                                            "description": "Set this to ` + "`" + `TRUE` + "`" + ` if you want the data to be base-64 encoded before it is written to the queue. Otherwise, set this to ` + "`" + `FALSE` + "`" + `.",
                                            "description_kind": "plain",
                                            "type": "bool"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      }
                                    }
                                  },
                                  "nesting_mode": "list"
                                }
                              },
                              "condition": {
                                "computed": true,
                                "description": "A Boolean expression that when ` + "`" + `TRUE` + "`" + ` causes the ` + "`" + `actions` + "`" + ` to be performed and the ` + "`" + `nextState` + "`" + ` to be entered.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "event_name": {
                                "computed": true,
                                "description": "The name of the event.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "next_state": {
                                "computed": true,
                                "description": "The next state to enter.",
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
                  "state_name": {
                    "computed": true,
                    "description": "The name of the state.",
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
      "detector_model_description": {
        "computed": true,
        "description": "A brief description of the detector model.",
        "description_kind": "plain",
        "type": "string"
      },
      "detector_model_name": {
        "computed": true,
        "description": "The name of the detector model.",
        "description_kind": "plain",
        "type": "string"
      },
      "evaluation_method": {
        "computed": true,
        "description": "Information about the order in which events are evaluated and how actions are executed.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "key": {
        "computed": true,
        "description": "The value used to identify a detector instance. When a device or system sends input, a new detector instance with a unique key value is created. AWS IoT Events can continue to route input to its corresponding detector instance based on this identifying information.\n\nThis parameter uses a JSON-path expression to select the attribute-value pair in the message payload that is used for identification. To route the message to the correct detector instance, the device must send a message payload that contains the same attribute-value.",
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "The ARN of the role that grants permission to AWS IoT Events to perform its operations.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.\n\nFor more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "Key of the Tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Value of the Tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::IoTEvents::DetectorModel",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIoteventsDetectorModelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIoteventsDetectorModel), &result)
	return &result
}
