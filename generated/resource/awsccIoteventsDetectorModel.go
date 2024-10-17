package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIoteventsDetectorModel = `{
  "block": {
    "attributes": {
      "detector_model_definition": {
        "description": "Information that defines how a detector operates.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "initial_state_name": {
              "description": "The state that is entered at the creation of each detector (instance).",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "states": {
              "description": "Information about the states of the detector.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "on_enter": {
                    "computed": true,
                    "description": "When entering this state, perform these ` + "`" + `` + "`" + `actions` + "`" + `` + "`" + ` if the ` + "`" + `` + "`" + `condition` + "`" + `` + "`" + ` is TRUE.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "events": {
                          "computed": true,
                          "description": "Specifies the actions that are performed when the state is entered and the ` + "`" + `` + "`" + `condition` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `TRUE` + "`" + `` + "`" + `.",
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
                                            "description": "The name of the timer to clear.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "dynamo_d_bv_2": {
                                      "computed": true,
                                      "description": "Writes to the DynamoDB table that you created. The default action payload contains all attribute-value pairs that have the information about the detector model instance and the event that triggered the action. You can customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html). A separate column of the DynamoDB table receives one attribute-value pair in the payload that you specify. For more information, see [Actions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-event-actions.html) in *Developer Guide*.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n By default, ITE generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `` + "`" + `contentExpression` + "`" + `` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `` + "`" + `'\u003cstring\u003e'` + "`" + `` + "`" + `), variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `` + "`" + `STRING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            },
                                            "optional": true
                                          },
                                          "table_name": {
                                            "computed": true,
                                            "description": "The name of the DynamoDB table.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "dynamo_db": {
                                      "computed": true,
                                      "description": "Writes to the DynamoDB table that you created. The default action payload contains all attribute-value pairs that have the information about the detector model instance and the event that triggered the action. You can customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html). One column of the DynamoDB table receives all attribute-value pairs in the payload that you specify. For more information, see [Actions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-event-actions.html) in *Developer Guide*.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "hash_key_field": {
                                            "computed": true,
                                            "description": "The name of the hash key (also called the partition key). The ` + "`" + `` + "`" + `hashKeyField` + "`" + `` + "`" + ` value must match the partition key of the target DynamoDB table.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "hash_key_type": {
                                            "computed": true,
                                            "description": "The data type for the hash key (also called the partition key). You can specify the following values:\n  +   ` + "`" + `` + "`" + `'STRING'` + "`" + `` + "`" + ` - The hash key is a string.\n  +   ` + "`" + `` + "`" + `'NUMBER'` + "`" + `` + "`" + ` - The hash key is a number.\n  \n If you don't specify ` + "`" + `` + "`" + `hashKeyType` + "`" + `` + "`" + `, the default value is ` + "`" + `` + "`" + `'STRING'` + "`" + `` + "`" + `.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "hash_key_value": {
                                            "computed": true,
                                            "description": "The value of the hash key (also called the partition key).",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "operation": {
                                            "computed": true,
                                            "description": "The type of operation to perform. You can specify the following values: \n  +   ` + "`" + `` + "`" + `'INSERT'` + "`" + `` + "`" + ` - Insert data as a new item into the DynamoDB table. This item uses the specified hash key as a partition key. If you specified a range key, the item uses the range key as a sort key.\n  +   ` + "`" + `` + "`" + `'UPDATE'` + "`" + `` + "`" + ` - Update an existing item of the DynamoDB table with new data. This item's partition key must match the specified hash key. If you specified a range key, the range key must match the item's sort key.\n  +   ` + "`" + `` + "`" + `'DELETE'` + "`" + `` + "`" + ` - Delete an existing item of the DynamoDB table. This item's partition key must match the specified hash key. If you specified a range key, the range key must match the item's sort key.\n  \n If you don't specify this parameter, ITE triggers the ` + "`" + `` + "`" + `'INSERT'` + "`" + `` + "`" + ` operation.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n By default, ITE generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `` + "`" + `contentExpression` + "`" + `` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `` + "`" + `'\u003cstring\u003e'` + "`" + `` + "`" + `), variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `` + "`" + `STRING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            },
                                            "optional": true
                                          },
                                          "payload_field": {
                                            "computed": true,
                                            "description": "The name of the DynamoDB column that receives the action payload.\n If you don't specify this parameter, the name of the DynamoDB column is ` + "`" + `` + "`" + `payload` + "`" + `` + "`" + `.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "range_key_field": {
                                            "computed": true,
                                            "description": "The name of the range key (also called the sort key). The ` + "`" + `` + "`" + `rangeKeyField` + "`" + `` + "`" + ` value must match the sort key of the target DynamoDB table.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "range_key_type": {
                                            "computed": true,
                                            "description": "The data type for the range key (also called the sort key), You can specify the following values:\n  +   ` + "`" + `` + "`" + `'STRING'` + "`" + `` + "`" + ` - The range key is a string.\n  +   ` + "`" + `` + "`" + `'NUMBER'` + "`" + `` + "`" + ` - The range key is number.\n  \n If you don't specify ` + "`" + `` + "`" + `rangeKeyField` + "`" + `` + "`" + `, the default value is ` + "`" + `` + "`" + `'STRING'` + "`" + `` + "`" + `.",
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
                                            "computed": true,
                                            "description": "The name of the DynamoDB table. The ` + "`" + `` + "`" + `tableName` + "`" + `` + "`" + ` value must match the table name of the target DynamoDB table.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
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
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "You can configure the action payload when you send a message to an Amazon Data Firehose delivery stream.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `` + "`" + `'\u003cstring\u003e'` + "`" + `` + "`" + `), variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `` + "`" + `STRING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
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
                                      "description": "Sends ITE input, which passes information about the detector model instance and the event that triggered the action.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "input_name": {
                                            "computed": true,
                                            "description": "The name of the ITE input where the data is sent.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "You can configure the action payload when you send a message to an ITE input.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `` + "`" + `'\u003cstring\u003e'` + "`" + `` + "`" + `), variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `` + "`" + `STRING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
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
                                      "description": "Sends information about the detector model instance and the event that triggered the action to an asset property in ITSW .",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "asset_id": {
                                            "computed": true,
                                            "description": "The ID of the asset that has the specified property.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "entry_id": {
                                            "computed": true,
                                            "description": "A unique identifier for this entry. You can use the entry ID to track which data entry causes an error in case of failure. The default is a new unique identifier.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "property_alias": {
                                            "computed": true,
                                            "description": "The alias of the asset property.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "property_id": {
                                            "computed": true,
                                            "description": "The ID of the asset property.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "property_value": {
                                            "computed": true,
                                            "description": "The value to send to the asset property. This value contains timestamp, quality, and value (TQV) information.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "quality": {
                                                  "computed": true,
                                                  "description": "The quality of the asset property value. The value must be ` + "`" + `` + "`" + `'GOOD'` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `'BAD'` + "`" + `` + "`" + `, or ` + "`" + `` + "`" + `'UNCERTAIN'` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "timestamp": {
                                                  "computed": true,
                                                  "description": "The timestamp associated with the asset property value. The default is the current event time.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "offset_in_nanos": {
                                                        "computed": true,
                                                        "description": "The nanosecond offset converted from ` + "`" + `` + "`" + `timeInSeconds` + "`" + `` + "`" + `. The valid range is between 0-999999999.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "time_in_seconds": {
                                                        "computed": true,
                                                        "description": "The timestamp, in seconds, in the Unix epoch format. The valid range is between 1-31556889864403199.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  },
                                                  "optional": true
                                                },
                                                "value": {
                                                  "computed": true,
                                                  "description": "The value to send to an asset property.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "boolean_value": {
                                                        "computed": true,
                                                        "description": "The asset property value is a Boolean value that must be ` + "`" + `` + "`" + `'TRUE'` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `'FALSE'` + "`" + `` + "`" + `. You must use an expression, and the evaluated result should be a Boolean value.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "double_value": {
                                                        "computed": true,
                                                        "description": "The asset property value is a double. You must use an expression, and the evaluated result should be a double.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "integer_value": {
                                                        "computed": true,
                                                        "description": "The asset property value is an integer. You must use an expression, and the evaluated result should be an integer.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "string_value": {
                                                        "computed": true,
                                                        "description": "The asset property value is a string. You must use an expression, and the evaluated result should be a string.",
                                                        "description_kind": "plain",
                                                        "optional": true,
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
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "iot_topic_publish": {
                                      "computed": true,
                                      "description": "Publishes an MQTT message with the given topic to the IoT message broker.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "mqtt_topic": {
                                            "computed": true,
                                            "description": "The MQTT topic of the message. You can use a string expression that includes variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `) and input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `) as the topic string.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "You can configure the action payload when you publish a message to an IoTCore topic.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `` + "`" + `'\u003cstring\u003e'` + "`" + `` + "`" + `), variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `` + "`" + `STRING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
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
                                      "description": "Calls a Lambda function, passing in information about the detector model instance and the event that triggered the action.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "function_arn": {
                                            "computed": true,
                                            "description": "The ARN of the Lambda function that is executed.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "You can configure the action payload when you send a message to a Lambda function.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `` + "`" + `'\u003cstring\u003e'` + "`" + `` + "`" + `), variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `` + "`" + `STRING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
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
                                    "reset_timer": {
                                      "computed": true,
                                      "description": "Information needed to reset the timer.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "timer_name": {
                                            "computed": true,
                                            "description": "The name of the timer to reset.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "set_timer": {
                                      "computed": true,
                                      "description": "Information needed to set the timer.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "duration_expression": {
                                            "computed": true,
                                            "description": "The duration of the timer, in seconds. You can use a string expression that includes numbers, variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), and input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `) as the duration. The range of the duration is 1-31622400 seconds. To ensure accuracy, the minimum duration is 60 seconds. The evaluated result of the duration is rounded down to the nearest whole number.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "seconds": {
                                            "computed": true,
                                            "description": "The number of seconds until the timer expires. The minimum value is 60 seconds to ensure accuracy. The maximum value is 31622400 seconds.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          },
                                          "timer_name": {
                                            "computed": true,
                                            "description": "The name of the timer.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "set_variable": {
                                      "computed": true,
                                      "description": "Sets a variable to a specified value.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "value": {
                                            "computed": true,
                                            "description": "The new value of the variable.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "variable_name": {
                                            "computed": true,
                                            "description": "The name of the variable.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "sns": {
                                      "computed": true,
                                      "description": "Sends an Amazon SNS message.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "payload": {
                                            "computed": true,
                                            "description": "You can configure the action payload when you send a message as an Amazon SNS push notification.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `` + "`" + `'\u003cstring\u003e'` + "`" + `` + "`" + `), variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `` + "`" + `STRING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            },
                                            "optional": true
                                          },
                                          "target_arn": {
                                            "computed": true,
                                            "description": "The ARN of the Amazon SNS target where the message is sent.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "sqs": {
                                      "computed": true,
                                      "description": "Sends an Amazon SNS message.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "payload": {
                                            "computed": true,
                                            "description": "You can configure the action payload when you send a message to an Amazon SQS queue.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `` + "`" + `'\u003cstring\u003e'` + "`" + `` + "`" + `), variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `` + "`" + `STRING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            },
                                            "optional": true
                                          },
                                          "queue_url": {
                                            "computed": true,
                                            "description": "The URL of the SQS queue where the data is written.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "use_base_64": {
                                            "computed": true,
                                            "description": "Set this to TRUE if you want the data to be base-64 encoded before it is written to the queue. Otherwise, set this to FALSE.",
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
                              },
                              "condition": {
                                "computed": true,
                                "description": "Optional. The Boolean expression that, when TRUE, causes the ` + "`" + `` + "`" + `actions` + "`" + `` + "`" + ` to be performed. If not present, the actions are performed (=TRUE). If the expression result is not a Boolean value, the actions are not performed (=FALSE).",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "event_name": {
                                "computed": true,
                                "description": "The name of the event.",
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
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "on_exit": {
                    "computed": true,
                    "description": "When exiting this state, perform these ` + "`" + `` + "`" + `actions` + "`" + `` + "`" + ` if the specified ` + "`" + `` + "`" + `condition` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `TRUE` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "events": {
                          "computed": true,
                          "description": "Specifies the ` + "`" + `` + "`" + `actions` + "`" + `` + "`" + ` that are performed when the state is exited and the ` + "`" + `` + "`" + `condition` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `TRUE` + "`" + `` + "`" + `.",
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
                                            "description": "The name of the timer to clear.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "dynamo_d_bv_2": {
                                      "computed": true,
                                      "description": "Writes to the DynamoDB table that you created. The default action payload contains all attribute-value pairs that have the information about the detector model instance and the event that triggered the action. You can customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html). A separate column of the DynamoDB table receives one attribute-value pair in the payload that you specify. For more information, see [Actions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-event-actions.html) in *Developer Guide*.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n By default, ITE generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `` + "`" + `contentExpression` + "`" + `` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `` + "`" + `'\u003cstring\u003e'` + "`" + `` + "`" + `), variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `` + "`" + `STRING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            },
                                            "optional": true
                                          },
                                          "table_name": {
                                            "computed": true,
                                            "description": "The name of the DynamoDB table.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "dynamo_db": {
                                      "computed": true,
                                      "description": "Writes to the DynamoDB table that you created. The default action payload contains all attribute-value pairs that have the information about the detector model instance and the event that triggered the action. You can customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html). One column of the DynamoDB table receives all attribute-value pairs in the payload that you specify. For more information, see [Actions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-event-actions.html) in *Developer Guide*.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "hash_key_field": {
                                            "computed": true,
                                            "description": "The name of the hash key (also called the partition key). The ` + "`" + `` + "`" + `hashKeyField` + "`" + `` + "`" + ` value must match the partition key of the target DynamoDB table.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "hash_key_type": {
                                            "computed": true,
                                            "description": "The data type for the hash key (also called the partition key). You can specify the following values:\n  +   ` + "`" + `` + "`" + `'STRING'` + "`" + `` + "`" + ` - The hash key is a string.\n  +   ` + "`" + `` + "`" + `'NUMBER'` + "`" + `` + "`" + ` - The hash key is a number.\n  \n If you don't specify ` + "`" + `` + "`" + `hashKeyType` + "`" + `` + "`" + `, the default value is ` + "`" + `` + "`" + `'STRING'` + "`" + `` + "`" + `.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "hash_key_value": {
                                            "computed": true,
                                            "description": "The value of the hash key (also called the partition key).",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "operation": {
                                            "computed": true,
                                            "description": "The type of operation to perform. You can specify the following values: \n  +   ` + "`" + `` + "`" + `'INSERT'` + "`" + `` + "`" + ` - Insert data as a new item into the DynamoDB table. This item uses the specified hash key as a partition key. If you specified a range key, the item uses the range key as a sort key.\n  +   ` + "`" + `` + "`" + `'UPDATE'` + "`" + `` + "`" + ` - Update an existing item of the DynamoDB table with new data. This item's partition key must match the specified hash key. If you specified a range key, the range key must match the item's sort key.\n  +   ` + "`" + `` + "`" + `'DELETE'` + "`" + `` + "`" + ` - Delete an existing item of the DynamoDB table. This item's partition key must match the specified hash key. If you specified a range key, the range key must match the item's sort key.\n  \n If you don't specify this parameter, ITE triggers the ` + "`" + `` + "`" + `'INSERT'` + "`" + `` + "`" + ` operation.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n By default, ITE generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `` + "`" + `contentExpression` + "`" + `` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `` + "`" + `'\u003cstring\u003e'` + "`" + `` + "`" + `), variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `` + "`" + `STRING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            },
                                            "optional": true
                                          },
                                          "payload_field": {
                                            "computed": true,
                                            "description": "The name of the DynamoDB column that receives the action payload.\n If you don't specify this parameter, the name of the DynamoDB column is ` + "`" + `` + "`" + `payload` + "`" + `` + "`" + `.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "range_key_field": {
                                            "computed": true,
                                            "description": "The name of the range key (also called the sort key). The ` + "`" + `` + "`" + `rangeKeyField` + "`" + `` + "`" + ` value must match the sort key of the target DynamoDB table.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "range_key_type": {
                                            "computed": true,
                                            "description": "The data type for the range key (also called the sort key), You can specify the following values:\n  +   ` + "`" + `` + "`" + `'STRING'` + "`" + `` + "`" + ` - The range key is a string.\n  +   ` + "`" + `` + "`" + `'NUMBER'` + "`" + `` + "`" + ` - The range key is number.\n  \n If you don't specify ` + "`" + `` + "`" + `rangeKeyField` + "`" + `` + "`" + `, the default value is ` + "`" + `` + "`" + `'STRING'` + "`" + `` + "`" + `.",
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
                                            "computed": true,
                                            "description": "The name of the DynamoDB table. The ` + "`" + `` + "`" + `tableName` + "`" + `` + "`" + ` value must match the table name of the target DynamoDB table.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
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
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "You can configure the action payload when you send a message to an Amazon Data Firehose delivery stream.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `` + "`" + `'\u003cstring\u003e'` + "`" + `` + "`" + `), variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `` + "`" + `STRING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
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
                                      "description": "Sends ITE input, which passes information about the detector model instance and the event that triggered the action.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "input_name": {
                                            "computed": true,
                                            "description": "The name of the ITE input where the data is sent.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "You can configure the action payload when you send a message to an ITE input.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `` + "`" + `'\u003cstring\u003e'` + "`" + `` + "`" + `), variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `` + "`" + `STRING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
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
                                      "description": "Sends information about the detector model instance and the event that triggered the action to an asset property in ITSW .",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "asset_id": {
                                            "computed": true,
                                            "description": "The ID of the asset that has the specified property.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "entry_id": {
                                            "computed": true,
                                            "description": "A unique identifier for this entry. You can use the entry ID to track which data entry causes an error in case of failure. The default is a new unique identifier.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "property_alias": {
                                            "computed": true,
                                            "description": "The alias of the asset property.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "property_id": {
                                            "computed": true,
                                            "description": "The ID of the asset property.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "property_value": {
                                            "computed": true,
                                            "description": "The value to send to the asset property. This value contains timestamp, quality, and value (TQV) information.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "quality": {
                                                  "computed": true,
                                                  "description": "The quality of the asset property value. The value must be ` + "`" + `` + "`" + `'GOOD'` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `'BAD'` + "`" + `` + "`" + `, or ` + "`" + `` + "`" + `'UNCERTAIN'` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "timestamp": {
                                                  "computed": true,
                                                  "description": "The timestamp associated with the asset property value. The default is the current event time.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "offset_in_nanos": {
                                                        "computed": true,
                                                        "description": "The nanosecond offset converted from ` + "`" + `` + "`" + `timeInSeconds` + "`" + `` + "`" + `. The valid range is between 0-999999999.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "time_in_seconds": {
                                                        "computed": true,
                                                        "description": "The timestamp, in seconds, in the Unix epoch format. The valid range is between 1-31556889864403199.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  },
                                                  "optional": true
                                                },
                                                "value": {
                                                  "computed": true,
                                                  "description": "The value to send to an asset property.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "boolean_value": {
                                                        "computed": true,
                                                        "description": "The asset property value is a Boolean value that must be ` + "`" + `` + "`" + `'TRUE'` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `'FALSE'` + "`" + `` + "`" + `. You must use an expression, and the evaluated result should be a Boolean value.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "double_value": {
                                                        "computed": true,
                                                        "description": "The asset property value is a double. You must use an expression, and the evaluated result should be a double.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "integer_value": {
                                                        "computed": true,
                                                        "description": "The asset property value is an integer. You must use an expression, and the evaluated result should be an integer.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "string_value": {
                                                        "computed": true,
                                                        "description": "The asset property value is a string. You must use an expression, and the evaluated result should be a string.",
                                                        "description_kind": "plain",
                                                        "optional": true,
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
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "iot_topic_publish": {
                                      "computed": true,
                                      "description": "Publishes an MQTT message with the given topic to the IoT message broker.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "mqtt_topic": {
                                            "computed": true,
                                            "description": "The MQTT topic of the message. You can use a string expression that includes variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `) and input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `) as the topic string.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "You can configure the action payload when you publish a message to an IoTCore topic.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `` + "`" + `'\u003cstring\u003e'` + "`" + `` + "`" + `), variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `` + "`" + `STRING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
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
                                      "description": "Calls a Lambda function, passing in information about the detector model instance and the event that triggered the action.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "function_arn": {
                                            "computed": true,
                                            "description": "The ARN of the Lambda function that is executed.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "You can configure the action payload when you send a message to a Lambda function.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `` + "`" + `'\u003cstring\u003e'` + "`" + `` + "`" + `), variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `` + "`" + `STRING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
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
                                    "reset_timer": {
                                      "computed": true,
                                      "description": "Information needed to reset the timer.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "timer_name": {
                                            "computed": true,
                                            "description": "The name of the timer to reset.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "set_timer": {
                                      "computed": true,
                                      "description": "Information needed to set the timer.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "duration_expression": {
                                            "computed": true,
                                            "description": "The duration of the timer, in seconds. You can use a string expression that includes numbers, variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), and input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `) as the duration. The range of the duration is 1-31622400 seconds. To ensure accuracy, the minimum duration is 60 seconds. The evaluated result of the duration is rounded down to the nearest whole number.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "seconds": {
                                            "computed": true,
                                            "description": "The number of seconds until the timer expires. The minimum value is 60 seconds to ensure accuracy. The maximum value is 31622400 seconds.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          },
                                          "timer_name": {
                                            "computed": true,
                                            "description": "The name of the timer.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "set_variable": {
                                      "computed": true,
                                      "description": "Sets a variable to a specified value.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "value": {
                                            "computed": true,
                                            "description": "The new value of the variable.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "variable_name": {
                                            "computed": true,
                                            "description": "The name of the variable.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "sns": {
                                      "computed": true,
                                      "description": "Sends an Amazon SNS message.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "payload": {
                                            "computed": true,
                                            "description": "You can configure the action payload when you send a message as an Amazon SNS push notification.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `` + "`" + `'\u003cstring\u003e'` + "`" + `` + "`" + `), variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `` + "`" + `STRING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            },
                                            "optional": true
                                          },
                                          "target_arn": {
                                            "computed": true,
                                            "description": "The ARN of the Amazon SNS target where the message is sent.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "sqs": {
                                      "computed": true,
                                      "description": "Sends an Amazon SNS message.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "payload": {
                                            "computed": true,
                                            "description": "You can configure the action payload when you send a message to an Amazon SQS queue.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `` + "`" + `'\u003cstring\u003e'` + "`" + `` + "`" + `), variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `` + "`" + `STRING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            },
                                            "optional": true
                                          },
                                          "queue_url": {
                                            "computed": true,
                                            "description": "The URL of the SQS queue where the data is written.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "use_base_64": {
                                            "computed": true,
                                            "description": "Set this to TRUE if you want the data to be base-64 encoded before it is written to the queue. Otherwise, set this to FALSE.",
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
                              },
                              "condition": {
                                "computed": true,
                                "description": "Optional. The Boolean expression that, when TRUE, causes the ` + "`" + `` + "`" + `actions` + "`" + `` + "`" + ` to be performed. If not present, the actions are performed (=TRUE). If the expression result is not a Boolean value, the actions are not performed (=FALSE).",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "event_name": {
                                "computed": true,
                                "description": "The name of the event.",
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
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "on_input": {
                    "computed": true,
                    "description": "When an input is received and the ` + "`" + `` + "`" + `condition` + "`" + `` + "`" + ` is TRUE, perform the specified ` + "`" + `` + "`" + `actions` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "events": {
                          "computed": true,
                          "description": "Specifies the actions performed when the ` + "`" + `` + "`" + `condition` + "`" + `` + "`" + ` evaluates to TRUE.",
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
                                            "description": "The name of the timer to clear.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "dynamo_d_bv_2": {
                                      "computed": true,
                                      "description": "Writes to the DynamoDB table that you created. The default action payload contains all attribute-value pairs that have the information about the detector model instance and the event that triggered the action. You can customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html). A separate column of the DynamoDB table receives one attribute-value pair in the payload that you specify. For more information, see [Actions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-event-actions.html) in *Developer Guide*.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n By default, ITE generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `` + "`" + `contentExpression` + "`" + `` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `` + "`" + `'\u003cstring\u003e'` + "`" + `` + "`" + `), variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `` + "`" + `STRING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            },
                                            "optional": true
                                          },
                                          "table_name": {
                                            "computed": true,
                                            "description": "The name of the DynamoDB table.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "dynamo_db": {
                                      "computed": true,
                                      "description": "Writes to the DynamoDB table that you created. The default action payload contains all attribute-value pairs that have the information about the detector model instance and the event that triggered the action. You can customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html). One column of the DynamoDB table receives all attribute-value pairs in the payload that you specify. For more information, see [Actions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-event-actions.html) in *Developer Guide*.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "hash_key_field": {
                                            "computed": true,
                                            "description": "The name of the hash key (also called the partition key). The ` + "`" + `` + "`" + `hashKeyField` + "`" + `` + "`" + ` value must match the partition key of the target DynamoDB table.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "hash_key_type": {
                                            "computed": true,
                                            "description": "The data type for the hash key (also called the partition key). You can specify the following values:\n  +   ` + "`" + `` + "`" + `'STRING'` + "`" + `` + "`" + ` - The hash key is a string.\n  +   ` + "`" + `` + "`" + `'NUMBER'` + "`" + `` + "`" + ` - The hash key is a number.\n  \n If you don't specify ` + "`" + `` + "`" + `hashKeyType` + "`" + `` + "`" + `, the default value is ` + "`" + `` + "`" + `'STRING'` + "`" + `` + "`" + `.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "hash_key_value": {
                                            "computed": true,
                                            "description": "The value of the hash key (also called the partition key).",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "operation": {
                                            "computed": true,
                                            "description": "The type of operation to perform. You can specify the following values: \n  +   ` + "`" + `` + "`" + `'INSERT'` + "`" + `` + "`" + ` - Insert data as a new item into the DynamoDB table. This item uses the specified hash key as a partition key. If you specified a range key, the item uses the range key as a sort key.\n  +   ` + "`" + `` + "`" + `'UPDATE'` + "`" + `` + "`" + ` - Update an existing item of the DynamoDB table with new data. This item's partition key must match the specified hash key. If you specified a range key, the range key must match the item's sort key.\n  +   ` + "`" + `` + "`" + `'DELETE'` + "`" + `` + "`" + ` - Delete an existing item of the DynamoDB table. This item's partition key must match the specified hash key. If you specified a range key, the range key must match the item's sort key.\n  \n If you don't specify this parameter, ITE triggers the ` + "`" + `` + "`" + `'INSERT'` + "`" + `` + "`" + ` operation.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n By default, ITE generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `` + "`" + `contentExpression` + "`" + `` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `` + "`" + `'\u003cstring\u003e'` + "`" + `` + "`" + `), variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `` + "`" + `STRING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            },
                                            "optional": true
                                          },
                                          "payload_field": {
                                            "computed": true,
                                            "description": "The name of the DynamoDB column that receives the action payload.\n If you don't specify this parameter, the name of the DynamoDB column is ` + "`" + `` + "`" + `payload` + "`" + `` + "`" + `.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "range_key_field": {
                                            "computed": true,
                                            "description": "The name of the range key (also called the sort key). The ` + "`" + `` + "`" + `rangeKeyField` + "`" + `` + "`" + ` value must match the sort key of the target DynamoDB table.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "range_key_type": {
                                            "computed": true,
                                            "description": "The data type for the range key (also called the sort key), You can specify the following values:\n  +   ` + "`" + `` + "`" + `'STRING'` + "`" + `` + "`" + ` - The range key is a string.\n  +   ` + "`" + `` + "`" + `'NUMBER'` + "`" + `` + "`" + ` - The range key is number.\n  \n If you don't specify ` + "`" + `` + "`" + `rangeKeyField` + "`" + `` + "`" + `, the default value is ` + "`" + `` + "`" + `'STRING'` + "`" + `` + "`" + `.",
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
                                            "computed": true,
                                            "description": "The name of the DynamoDB table. The ` + "`" + `` + "`" + `tableName` + "`" + `` + "`" + ` value must match the table name of the target DynamoDB table.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
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
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "You can configure the action payload when you send a message to an Amazon Data Firehose delivery stream.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `` + "`" + `'\u003cstring\u003e'` + "`" + `` + "`" + `), variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `` + "`" + `STRING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
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
                                      "description": "Sends ITE input, which passes information about the detector model instance and the event that triggered the action.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "input_name": {
                                            "computed": true,
                                            "description": "The name of the ITE input where the data is sent.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "You can configure the action payload when you send a message to an ITE input.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `` + "`" + `'\u003cstring\u003e'` + "`" + `` + "`" + `), variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `` + "`" + `STRING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
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
                                      "description": "Sends information about the detector model instance and the event that triggered the action to an asset property in ITSW .",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "asset_id": {
                                            "computed": true,
                                            "description": "The ID of the asset that has the specified property.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "entry_id": {
                                            "computed": true,
                                            "description": "A unique identifier for this entry. You can use the entry ID to track which data entry causes an error in case of failure. The default is a new unique identifier.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "property_alias": {
                                            "computed": true,
                                            "description": "The alias of the asset property.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "property_id": {
                                            "computed": true,
                                            "description": "The ID of the asset property.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "property_value": {
                                            "computed": true,
                                            "description": "The value to send to the asset property. This value contains timestamp, quality, and value (TQV) information.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "quality": {
                                                  "computed": true,
                                                  "description": "The quality of the asset property value. The value must be ` + "`" + `` + "`" + `'GOOD'` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `'BAD'` + "`" + `` + "`" + `, or ` + "`" + `` + "`" + `'UNCERTAIN'` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "timestamp": {
                                                  "computed": true,
                                                  "description": "The timestamp associated with the asset property value. The default is the current event time.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "offset_in_nanos": {
                                                        "computed": true,
                                                        "description": "The nanosecond offset converted from ` + "`" + `` + "`" + `timeInSeconds` + "`" + `` + "`" + `. The valid range is between 0-999999999.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "time_in_seconds": {
                                                        "computed": true,
                                                        "description": "The timestamp, in seconds, in the Unix epoch format. The valid range is between 1-31556889864403199.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  },
                                                  "optional": true
                                                },
                                                "value": {
                                                  "computed": true,
                                                  "description": "The value to send to an asset property.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "boolean_value": {
                                                        "computed": true,
                                                        "description": "The asset property value is a Boolean value that must be ` + "`" + `` + "`" + `'TRUE'` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `'FALSE'` + "`" + `` + "`" + `. You must use an expression, and the evaluated result should be a Boolean value.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "double_value": {
                                                        "computed": true,
                                                        "description": "The asset property value is a double. You must use an expression, and the evaluated result should be a double.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "integer_value": {
                                                        "computed": true,
                                                        "description": "The asset property value is an integer. You must use an expression, and the evaluated result should be an integer.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "string_value": {
                                                        "computed": true,
                                                        "description": "The asset property value is a string. You must use an expression, and the evaluated result should be a string.",
                                                        "description_kind": "plain",
                                                        "optional": true,
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
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "iot_topic_publish": {
                                      "computed": true,
                                      "description": "Publishes an MQTT message with the given topic to the IoT message broker.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "mqtt_topic": {
                                            "computed": true,
                                            "description": "The MQTT topic of the message. You can use a string expression that includes variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `) and input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `) as the topic string.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "You can configure the action payload when you publish a message to an IoTCore topic.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `` + "`" + `'\u003cstring\u003e'` + "`" + `` + "`" + `), variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `` + "`" + `STRING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
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
                                      "description": "Calls a Lambda function, passing in information about the detector model instance and the event that triggered the action.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "function_arn": {
                                            "computed": true,
                                            "description": "The ARN of the Lambda function that is executed.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "You can configure the action payload when you send a message to a Lambda function.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `` + "`" + `'\u003cstring\u003e'` + "`" + `` + "`" + `), variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `` + "`" + `STRING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
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
                                    "reset_timer": {
                                      "computed": true,
                                      "description": "Information needed to reset the timer.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "timer_name": {
                                            "computed": true,
                                            "description": "The name of the timer to reset.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "set_timer": {
                                      "computed": true,
                                      "description": "Information needed to set the timer.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "duration_expression": {
                                            "computed": true,
                                            "description": "The duration of the timer, in seconds. You can use a string expression that includes numbers, variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), and input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `) as the duration. The range of the duration is 1-31622400 seconds. To ensure accuracy, the minimum duration is 60 seconds. The evaluated result of the duration is rounded down to the nearest whole number.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "seconds": {
                                            "computed": true,
                                            "description": "The number of seconds until the timer expires. The minimum value is 60 seconds to ensure accuracy. The maximum value is 31622400 seconds.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          },
                                          "timer_name": {
                                            "computed": true,
                                            "description": "The name of the timer.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "set_variable": {
                                      "computed": true,
                                      "description": "Sets a variable to a specified value.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "value": {
                                            "computed": true,
                                            "description": "The new value of the variable.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "variable_name": {
                                            "computed": true,
                                            "description": "The name of the variable.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "sns": {
                                      "computed": true,
                                      "description": "Sends an Amazon SNS message.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "payload": {
                                            "computed": true,
                                            "description": "You can configure the action payload when you send a message as an Amazon SNS push notification.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `` + "`" + `'\u003cstring\u003e'` + "`" + `` + "`" + `), variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `` + "`" + `STRING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            },
                                            "optional": true
                                          },
                                          "target_arn": {
                                            "computed": true,
                                            "description": "The ARN of the Amazon SNS target where the message is sent.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "sqs": {
                                      "computed": true,
                                      "description": "Sends an Amazon SNS message.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "payload": {
                                            "computed": true,
                                            "description": "You can configure the action payload when you send a message to an Amazon SQS queue.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `` + "`" + `'\u003cstring\u003e'` + "`" + `` + "`" + `), variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `` + "`" + `STRING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            },
                                            "optional": true
                                          },
                                          "queue_url": {
                                            "computed": true,
                                            "description": "The URL of the SQS queue where the data is written.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "use_base_64": {
                                            "computed": true,
                                            "description": "Set this to TRUE if you want the data to be base-64 encoded before it is written to the queue. Otherwise, set this to FALSE.",
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
                              },
                              "condition": {
                                "computed": true,
                                "description": "Optional. The Boolean expression that, when TRUE, causes the ` + "`" + `` + "`" + `actions` + "`" + `` + "`" + ` to be performed. If not present, the actions are performed (=TRUE). If the expression result is not a Boolean value, the actions are not performed (=FALSE).",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "event_name": {
                                "computed": true,
                                "description": "The name of the event.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "transition_events": {
                          "computed": true,
                          "description": "Specifies the actions performed, and the next state entered, when a ` + "`" + `` + "`" + `condition` + "`" + `` + "`" + ` evaluates to TRUE.",
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
                                            "description": "The name of the timer to clear.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "dynamo_d_bv_2": {
                                      "computed": true,
                                      "description": "Writes to the DynamoDB table that you created. The default action payload contains all attribute-value pairs that have the information about the detector model instance and the event that triggered the action. You can customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html). A separate column of the DynamoDB table receives one attribute-value pair in the payload that you specify. For more information, see [Actions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-event-actions.html) in *Developer Guide*.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n By default, ITE generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `` + "`" + `contentExpression` + "`" + `` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `` + "`" + `'\u003cstring\u003e'` + "`" + `` + "`" + `), variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `` + "`" + `STRING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            },
                                            "optional": true
                                          },
                                          "table_name": {
                                            "computed": true,
                                            "description": "The name of the DynamoDB table.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "dynamo_db": {
                                      "computed": true,
                                      "description": "Writes to the DynamoDB table that you created. The default action payload contains all attribute-value pairs that have the information about the detector model instance and the event that triggered the action. You can customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html). One column of the DynamoDB table receives all attribute-value pairs in the payload that you specify. For more information, see [Actions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-event-actions.html) in *Developer Guide*.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "hash_key_field": {
                                            "computed": true,
                                            "description": "The name of the hash key (also called the partition key). The ` + "`" + `` + "`" + `hashKeyField` + "`" + `` + "`" + ` value must match the partition key of the target DynamoDB table.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "hash_key_type": {
                                            "computed": true,
                                            "description": "The data type for the hash key (also called the partition key). You can specify the following values:\n  +   ` + "`" + `` + "`" + `'STRING'` + "`" + `` + "`" + ` - The hash key is a string.\n  +   ` + "`" + `` + "`" + `'NUMBER'` + "`" + `` + "`" + ` - The hash key is a number.\n  \n If you don't specify ` + "`" + `` + "`" + `hashKeyType` + "`" + `` + "`" + `, the default value is ` + "`" + `` + "`" + `'STRING'` + "`" + `` + "`" + `.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "hash_key_value": {
                                            "computed": true,
                                            "description": "The value of the hash key (also called the partition key).",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "operation": {
                                            "computed": true,
                                            "description": "The type of operation to perform. You can specify the following values: \n  +   ` + "`" + `` + "`" + `'INSERT'` + "`" + `` + "`" + ` - Insert data as a new item into the DynamoDB table. This item uses the specified hash key as a partition key. If you specified a range key, the item uses the range key as a sort key.\n  +   ` + "`" + `` + "`" + `'UPDATE'` + "`" + `` + "`" + ` - Update an existing item of the DynamoDB table with new data. This item's partition key must match the specified hash key. If you specified a range key, the range key must match the item's sort key.\n  +   ` + "`" + `` + "`" + `'DELETE'` + "`" + `` + "`" + ` - Delete an existing item of the DynamoDB table. This item's partition key must match the specified hash key. If you specified a range key, the range key must match the item's sort key.\n  \n If you don't specify this parameter, ITE triggers the ` + "`" + `` + "`" + `'INSERT'` + "`" + `` + "`" + ` operation.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "Information needed to configure the payload.\n By default, ITE generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use ` + "`" + `` + "`" + `contentExpression` + "`" + `` + "`" + `.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `` + "`" + `'\u003cstring\u003e'` + "`" + `` + "`" + `), variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `` + "`" + `STRING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            },
                                            "optional": true
                                          },
                                          "payload_field": {
                                            "computed": true,
                                            "description": "The name of the DynamoDB column that receives the action payload.\n If you don't specify this parameter, the name of the DynamoDB column is ` + "`" + `` + "`" + `payload` + "`" + `` + "`" + `.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "range_key_field": {
                                            "computed": true,
                                            "description": "The name of the range key (also called the sort key). The ` + "`" + `` + "`" + `rangeKeyField` + "`" + `` + "`" + ` value must match the sort key of the target DynamoDB table.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "range_key_type": {
                                            "computed": true,
                                            "description": "The data type for the range key (also called the sort key), You can specify the following values:\n  +   ` + "`" + `` + "`" + `'STRING'` + "`" + `` + "`" + ` - The range key is a string.\n  +   ` + "`" + `` + "`" + `'NUMBER'` + "`" + `` + "`" + ` - The range key is number.\n  \n If you don't specify ` + "`" + `` + "`" + `rangeKeyField` + "`" + `` + "`" + `, the default value is ` + "`" + `` + "`" + `'STRING'` + "`" + `` + "`" + `.",
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
                                            "computed": true,
                                            "description": "The name of the DynamoDB table. The ` + "`" + `` + "`" + `tableName` + "`" + `` + "`" + ` value must match the table name of the target DynamoDB table.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
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
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "You can configure the action payload when you send a message to an Amazon Data Firehose delivery stream.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `` + "`" + `'\u003cstring\u003e'` + "`" + `` + "`" + `), variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `` + "`" + `STRING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
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
                                      "description": "Sends ITE input, which passes information about the detector model instance and the event that triggered the action.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "input_name": {
                                            "computed": true,
                                            "description": "The name of the ITE input where the data is sent.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "You can configure the action payload when you send a message to an ITE input.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `` + "`" + `'\u003cstring\u003e'` + "`" + `` + "`" + `), variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `` + "`" + `STRING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
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
                                      "description": "Sends information about the detector model instance and the event that triggered the action to an asset property in ITSW .",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "asset_id": {
                                            "computed": true,
                                            "description": "The ID of the asset that has the specified property.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "entry_id": {
                                            "computed": true,
                                            "description": "A unique identifier for this entry. You can use the entry ID to track which data entry causes an error in case of failure. The default is a new unique identifier.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "property_alias": {
                                            "computed": true,
                                            "description": "The alias of the asset property.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "property_id": {
                                            "computed": true,
                                            "description": "The ID of the asset property.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "property_value": {
                                            "computed": true,
                                            "description": "The value to send to the asset property. This value contains timestamp, quality, and value (TQV) information.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "quality": {
                                                  "computed": true,
                                                  "description": "The quality of the asset property value. The value must be ` + "`" + `` + "`" + `'GOOD'` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `'BAD'` + "`" + `` + "`" + `, or ` + "`" + `` + "`" + `'UNCERTAIN'` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "timestamp": {
                                                  "computed": true,
                                                  "description": "The timestamp associated with the asset property value. The default is the current event time.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "offset_in_nanos": {
                                                        "computed": true,
                                                        "description": "The nanosecond offset converted from ` + "`" + `` + "`" + `timeInSeconds` + "`" + `` + "`" + `. The valid range is between 0-999999999.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "time_in_seconds": {
                                                        "computed": true,
                                                        "description": "The timestamp, in seconds, in the Unix epoch format. The valid range is between 1-31556889864403199.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      }
                                                    },
                                                    "nesting_mode": "single"
                                                  },
                                                  "optional": true
                                                },
                                                "value": {
                                                  "computed": true,
                                                  "description": "The value to send to an asset property.",
                                                  "description_kind": "plain",
                                                  "nested_type": {
                                                    "attributes": {
                                                      "boolean_value": {
                                                        "computed": true,
                                                        "description": "The asset property value is a Boolean value that must be ` + "`" + `` + "`" + `'TRUE'` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `'FALSE'` + "`" + `` + "`" + `. You must use an expression, and the evaluated result should be a Boolean value.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "double_value": {
                                                        "computed": true,
                                                        "description": "The asset property value is a double. You must use an expression, and the evaluated result should be a double.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "integer_value": {
                                                        "computed": true,
                                                        "description": "The asset property value is an integer. You must use an expression, and the evaluated result should be an integer.",
                                                        "description_kind": "plain",
                                                        "optional": true,
                                                        "type": "string"
                                                      },
                                                      "string_value": {
                                                        "computed": true,
                                                        "description": "The asset property value is a string. You must use an expression, and the evaluated result should be a string.",
                                                        "description_kind": "plain",
                                                        "optional": true,
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
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "iot_topic_publish": {
                                      "computed": true,
                                      "description": "Publishes an MQTT message with the given topic to the IoT message broker.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "mqtt_topic": {
                                            "computed": true,
                                            "description": "The MQTT topic of the message. You can use a string expression that includes variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `) and input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `) as the topic string.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "You can configure the action payload when you publish a message to an IoTCore topic.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `` + "`" + `'\u003cstring\u003e'` + "`" + `` + "`" + `), variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `` + "`" + `STRING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
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
                                      "description": "Calls a Lambda function, passing in information about the detector model instance and the event that triggered the action.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "function_arn": {
                                            "computed": true,
                                            "description": "The ARN of the Lambda function that is executed.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "payload": {
                                            "computed": true,
                                            "description": "You can configure the action payload when you send a message to a Lambda function.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `` + "`" + `'\u003cstring\u003e'` + "`" + `` + "`" + `), variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `` + "`" + `STRING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
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
                                    "reset_timer": {
                                      "computed": true,
                                      "description": "Information needed to reset the timer.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "timer_name": {
                                            "computed": true,
                                            "description": "The name of the timer to reset.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "set_timer": {
                                      "computed": true,
                                      "description": "Information needed to set the timer.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "duration_expression": {
                                            "computed": true,
                                            "description": "The duration of the timer, in seconds. You can use a string expression that includes numbers, variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), and input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `) as the duration. The range of the duration is 1-31622400 seconds. To ensure accuracy, the minimum duration is 60 seconds. The evaluated result of the duration is rounded down to the nearest whole number.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "seconds": {
                                            "computed": true,
                                            "description": "The number of seconds until the timer expires. The minimum value is 60 seconds to ensure accuracy. The maximum value is 31622400 seconds.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
                                          },
                                          "timer_name": {
                                            "computed": true,
                                            "description": "The name of the timer.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "set_variable": {
                                      "computed": true,
                                      "description": "Sets a variable to a specified value.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "value": {
                                            "computed": true,
                                            "description": "The new value of the variable.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "variable_name": {
                                            "computed": true,
                                            "description": "The name of the variable.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "sns": {
                                      "computed": true,
                                      "description": "Sends an Amazon SNS message.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "payload": {
                                            "computed": true,
                                            "description": "You can configure the action payload when you send a message as an Amazon SNS push notification.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `` + "`" + `'\u003cstring\u003e'` + "`" + `` + "`" + `), variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `` + "`" + `STRING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            },
                                            "optional": true
                                          },
                                          "target_arn": {
                                            "computed": true,
                                            "description": "The ARN of the Amazon SNS target where the message is sent.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "sqs": {
                                      "computed": true,
                                      "description": "Sends an Amazon SNS message.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "payload": {
                                            "computed": true,
                                            "description": "You can configure the action payload when you send a message to an Amazon SQS queue.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_expression": {
                                                  "computed": true,
                                                  "description": "The content of the payload. You can use a string expression that includes quoted strings (` + "`" + `` + "`" + `'\u003cstring\u003e'` + "`" + `` + "`" + `), variables (` + "`" + `` + "`" + `$variable.\u003cvariable-name\u003e` + "`" + `` + "`" + `), input values (` + "`" + `` + "`" + `$input.\u003cinput-name\u003e.\u003cpath-to-datum\u003e` + "`" + `` + "`" + `), string concatenations, and quoted strings that contain ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + ` as the content. The recommended maximum size of a content expression is 1 KB.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "type": {
                                                  "computed": true,
                                                  "description": "The value of the payload type can be either ` + "`" + `` + "`" + `STRING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            },
                                            "optional": true
                                          },
                                          "queue_url": {
                                            "computed": true,
                                            "description": "The URL of the SQS queue where the data is written.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "use_base_64": {
                                            "computed": true,
                                            "description": "Set this to TRUE if you want the data to be base-64 encoded before it is written to the queue. Otherwise, set this to FALSE.",
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
                              },
                              "condition": {
                                "computed": true,
                                "description": "Required. A Boolean expression that when TRUE causes the actions to be performed and the ` + "`" + `` + "`" + `nextState` + "`" + `` + "`" + ` to be entered.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "event_name": {
                                "computed": true,
                                "description": "The name of the transition event.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "next_state": {
                                "computed": true,
                                "description": "The next state to enter.",
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
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "state_name": {
                    "description": "The name of the state.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "detector_model_description": {
        "computed": true,
        "description": "A brief description of the detector model.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "detector_model_name": {
        "computed": true,
        "description": "The name of the detector model.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "evaluation_method": {
        "computed": true,
        "description": "Information about the order in which events are evaluated and how actions are executed.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "key": {
        "computed": true,
        "description": "The value used to identify a detector instance. When a device or system sends input, a new detector instance with a unique key value is created. ITE can continue to route input to its corresponding detector instance based on this identifying information. \n This parameter uses a JSON-path expression to select the attribute-value pair in the message payload that is used for identification. To route the message to the correct detector instance, the device must send a message payload that contains the same attribute-value.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role_arn": {
        "description": "The ARN of the role that grants permission to ITE to perform its operations.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.\n For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag's key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag's value.",
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
    "description": "The AWS::IoTEvents::DetectorModel resource creates a detector model. You create a *detector model* (a model of your equipment or process) using *states*. For each state, you define conditional (Boolean) logic that evaluates the incoming inputs to detect significant events. When an event is detected, it can change the state or trigger custom-built or predefined actions using other AWS services. You can define additional events that trigger actions when entering or exiting a state and, optionally, when a condition is met. For more information, see [How to Use](https://docs.aws.amazon.com/iotevents/latest/developerguide/how-to-use-iotevents.html) in the *Developer Guide*.\n  When you successfully update a detector model (using the ITE console, ITE API or CLI commands, or CFN) all detector instances created by the model are reset to their initial states. (The detector's ` + "`" + `` + "`" + `state` + "`" + `` + "`" + `, and the values of any variables and timers are reset.)\n When you successfully update a detector model (using the ITE console, ITE API or CLI commands, or CFN) the version number of the detector model is incremented. (A detector model with version number 1 before the update has version number 2 after the update succeeds.)\n If you attempt to update a detector model using CFN and the update does not succeed, the system may, in some cases, restore the original detector model. When this occurs, the detector model's version is incremented twice (for example, from version 1 to version 3) and the detector instances are reset.\n Also, be aware that if you attempt to update several detector models at once using CFN, some updates may succeed and others fail. In this case, the effects on each detector model's detector instances and version number depend on whether the update succeeded or failed, with the results as stated.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIoteventsDetectorModelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIoteventsDetectorModel), &result)
	return &result
}
