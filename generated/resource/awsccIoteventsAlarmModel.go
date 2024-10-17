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
        "description": "Contains the configuration information of alarm state changes.",
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
                    "description": "The value must be ` + "`" + `` + "`" + `TRUE` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `FALSE` + "`" + `` + "`" + `. If ` + "`" + `` + "`" + `TRUE` + "`" + `` + "`" + `, you receive a notification when the alarm state changes. You must choose to acknowledge the notification before the alarm state can return to ` + "`" + `` + "`" + `NORMAL` + "`" + `` + "`" + `. If ` + "`" + `` + "`" + `FALSE` + "`" + `` + "`" + `, you won't receive notifications. The alarm automatically changes to the ` + "`" + `` + "`" + `NORMAL` + "`" + `` + "`" + ` state when the input property value returns to the specified range.",
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
                    "description": "The value must be ` + "`" + `` + "`" + `TRUE` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `FALSE` + "`" + `` + "`" + `. If ` + "`" + `` + "`" + `FALSE` + "`" + `` + "`" + `, all alarm instances created based on the alarm model are activated. The default value is ` + "`" + `` + "`" + `TRUE` + "`" + `` + "`" + `.",
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
                    "description": "Defines an action to write to the Amazon DynamoDB table that you created. The default action payload contains all the information about the detector model instance and the event that triggered the action. You can customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html). A separate column of the DynamoDB table receives one attribute-value pair in the payload that you specify.\n You must use expressions for all parameters in ` + "`" + `` + "`" + `DynamoDBv2Action` + "`" + `` + "`" + `. The expressions accept literals, operators, functions, references, and substitution templates.\n  **Examples**\n +  For literal values, the expressions must contain single quotes. For example, the value for the ` + "`" + `` + "`" + `tableName` + "`" + `` + "`" + ` parameter can be ` + "`" + `` + "`" + `'GreenhouseTemperatureTable'` + "`" + `` + "`" + `.\n  +  For references, you must specify either variables or input values. For example, the value for the ` + "`" + `` + "`" + `tableName` + "`" + `` + "`" + ` parameter can be ` + "`" + `` + "`" + `$variable.ddbtableName` + "`" + `` + "`" + `.\n  +  For a substitution template, you must use ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + `, and the template must be in single quotes. A substitution template can also contain a combination of literals, operators, functions, references, and substitution templates.\n In the following example, the value for the ` + "`" + `` + "`" + `contentExpression` + "`" + `` + "`" + ` parameter in ` + "`" + `` + "`" + `Payload` + "`" + `` + "`" + ` uses a substitution template. \n  ` + "`" + `` + "`" + `'{\\\"sensorID\\\": \\\"${$input.GreenhouseInput.sensor_id}\\\", \\\"temperature\\\": \\\"${$input.GreenhouseInput.temperature * 9 / 5 + 32}\\\"}'` + "`" + `` + "`" + ` \n  +  For a string concatenation, you must use ` + "`" + `` + "`" + `+` + "`" + `` + "`" + `. A string concatenation can also contain a combination of literals, operators, functions, references, and substitution templates.\n In the following example, the value for the ` + "`" + `` + "`" + `tableName` + "`" + `` + "`" + ` parameter uses a string concatenation. \n  ` + "`" + `` + "`" + `'GreenhouseTemperatureTable ' + $input.GreenhouseInput.date` + "`" + `` + "`" + ` \n  \n For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *Developer Guide*.\n The value for the ` + "`" + `` + "`" + `type` + "`" + `` + "`" + ` parameter in ` + "`" + `` + "`" + `Payload` + "`" + `` + "`" + ` must be ` + "`" + `` + "`" + `JSON` + "`" + `` + "`" + `.",
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
                    "description": "Defines an action to write to the Amazon DynamoDB table that you created. The standard action payload contains all the information about the detector model instance and the event that triggered the action. You can customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html). One column of the DynamoDB table receives all attribute-value pairs in the payload that you specify.\n You must use expressions for all parameters in ` + "`" + `` + "`" + `DynamoDBAction` + "`" + `` + "`" + `. The expressions accept literals, operators, functions, references, and substitution templates.\n  **Examples**\n +  For literal values, the expressions must contain single quotes. For example, the value for the ` + "`" + `` + "`" + `hashKeyType` + "`" + `` + "`" + ` parameter can be ` + "`" + `` + "`" + `'STRING'` + "`" + `` + "`" + `.\n  +  For references, you must specify either variables or input values. For example, the value for the ` + "`" + `` + "`" + `hashKeyField` + "`" + `` + "`" + ` parameter can be ` + "`" + `` + "`" + `$input.GreenhouseInput.name` + "`" + `` + "`" + `.\n  +  For a substitution template, you must use ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + `, and the template must be in single quotes. A substitution template can also contain a combination of literals, operators, functions, references, and substitution templates.\n In the following example, the value for the ` + "`" + `` + "`" + `hashKeyValue` + "`" + `` + "`" + ` parameter uses a substitution template. \n  ` + "`" + `` + "`" + `'${$input.GreenhouseInput.temperature * 6 / 5 + 32} in Fahrenheit'` + "`" + `` + "`" + ` \n  +  For a string concatenation, you must use ` + "`" + `` + "`" + `+` + "`" + `` + "`" + `. A string concatenation can also contain a combination of literals, operators, functions, references, and substitution templates.\n In the following example, the value for the ` + "`" + `` + "`" + `tableName` + "`" + `` + "`" + ` parameter uses a string concatenation. \n  ` + "`" + `` + "`" + `'GreenhouseTemperatureTable ' + $input.GreenhouseInput.date` + "`" + `` + "`" + ` \n  \n For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *Developer Guide*.\n If the defined payload type is a string, ` + "`" + `` + "`" + `DynamoDBAction` + "`" + `` + "`" + ` writes non-JSON data to the DynamoDB table as binary data. The DynamoDB console displays the data as Base64-encoded text. The value for the ` + "`" + `` + "`" + `payloadField` + "`" + `` + "`" + ` parameter is ` + "`" + `` + "`" + `\u003cpayload-field\u003e_raw` + "`" + `` + "`" + `.",
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
                    "description": "Sends an ITE input, passing in information about the detector model instance and the event that triggered the action.",
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
                    "description": "Sends information about the detector model instance and the event that triggered the action to a specified asset property in ITSW.\n You must use expressions for all parameters in ` + "`" + `` + "`" + `IotSiteWiseAction` + "`" + `` + "`" + `. The expressions accept literals, operators, functions, references, and substitutions templates.\n  **Examples**\n +  For literal values, the expressions must contain single quotes. For example, the value for the ` + "`" + `` + "`" + `propertyAlias` + "`" + `` + "`" + ` parameter can be ` + "`" + `` + "`" + `'/company/windfarm/3/turbine/7/temperature'` + "`" + `` + "`" + `.\n  +  For references, you must specify either variables or input values. For example, the value for the ` + "`" + `` + "`" + `assetId` + "`" + `` + "`" + ` parameter can be ` + "`" + `` + "`" + `$input.TurbineInput.assetId1` + "`" + `` + "`" + `.\n  +  For a substitution template, you must use ` + "`" + `` + "`" + `${}` + "`" + `` + "`" + `, and the template must be in single quotes. A substitution template can also contain a combination of literals, operators, functions, references, and substitution templates.\n In the following example, the value for the ` + "`" + `` + "`" + `propertyAlias` + "`" + `` + "`" + ` parameter uses a substitution template. \n  ` + "`" + `` + "`" + `'company/windfarm/${$input.TemperatureInput.sensorData.windfarmID}/turbine/ ${$input.TemperatureInput.sensorData.turbineID}/temperature'` + "`" + `` + "`" + ` \n  \n You must specify either ` + "`" + `` + "`" + `propertyAlias` + "`" + `` + "`" + ` or both ` + "`" + `` + "`" + `assetId` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `propertyId` + "`" + `` + "`" + ` to identify the target asset property in ITSW.\n For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *Developer Guide*.",
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
                    "description": "Information required to publish the MQTT message through the IoT message broker.",
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
                  "sns": {
                    "computed": true,
                    "description": "Information required to publish the Amazon SNS message.",
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
                    "description": "Sends information about the detector model instance and the event that triggered the action to an Amazon SQS queue.",
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
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "alarm_model_description": {
        "computed": true,
        "description": "The description of the alarm model.",
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
                    "computed": true,
                    "description": "The comparison operator.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "input_property": {
                    "computed": true,
                    "description": "The value on the left side of the comparison operator. You can specify an ITE input attribute as an input property.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "threshold": {
                    "computed": true,
                    "description": "The value on the right side of the comparison operator. You can enter a number or specify an ITE input attribute.",
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
        "description": "An input attribute used as a key to create an alarm. ITE routes [inputs](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Input.html) associated with this key to the alarm.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role_arn": {
        "description": "The ARN of the IAM role that allows the alarm to perform actions and access AWS resources. For more information, see [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the *General Reference*.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "severity": {
        "computed": true,
        "description": "A non-negative integer that reflects the severity level of the alarm.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs that contain metadata for the alarm model. The tags help you manage the alarm model. For more information, see [Tagging your resources](https://docs.aws.amazon.com/iotevents/latest/developerguide/tagging-iotevents.html) in the *Developer Guide*.\n You can create up to 50 tags for one alarm model.",
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
    "description": "Represents an alarm model to monitor an ITE input attribute. You can use the alarm to get notified when the value is outside a specified range. For more information, see [Create an alarm model](https://docs.aws.amazon.com/iotevents/latest/developerguide/create-alarms.html) in the *Developer Guide*.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIoteventsAlarmModelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIoteventsAlarmModel), &result)
	return &result
}
