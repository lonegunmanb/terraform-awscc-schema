package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccKmsKey = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bypass_policy_lockout_safety_check": {
        "computed": true,
        "description": "Skips (\"bypasses\") the key policy lockout safety check. The default value is false.\n  Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.\n For more information, see [Default key policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policy-default.html#prevent-unmanageable-key) in the *Developer Guide*.\n  Use this parameter only when you intend to prevent the principal that is making the request from making a subsequent [PutKeyPolicy](https://docs.aws.amazon.com/kms/latest/APIReference/API_PutKeyPolicy.html) request on the KMS key.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "computed": true,
        "description": "A description of the KMS key. Use a description that helps you to distinguish this KMS key from others in the account, such as its intended use.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_key_rotation": {
        "computed": true,
        "description": "Enables automatic rotation of the key material for the specified KMS key. By default, automatic key rotation is not enabled.\n KMS supports automatic rotation only for symmetric encryption KMS keys (` + "`" + `` + "`" + `KeySpec` + "`" + `` + "`" + ` = ` + "`" + `` + "`" + `SYMMETRIC_DEFAULT` + "`" + `` + "`" + `). For asymmetric KMS keys, HMAC KMS keys, and KMS keys with Origin ` + "`" + `` + "`" + `EXTERNAL` + "`" + `` + "`" + `, omit the ` + "`" + `` + "`" + `EnableKeyRotation` + "`" + `` + "`" + ` property or set it to ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `.\n To enable automatic key rotation of the key material for a multi-Region KMS key, set ` + "`" + `` + "`" + `EnableKeyRotation` + "`" + `` + "`" + ` to ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` on the primary key (created by using ` + "`" + `` + "`" + `AWS::KMS::Key` + "`" + `` + "`" + `). KMS copies the rotation status to all replica keys. For details, see [Rotating multi-Region keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-manage.html#multi-region-rotate) in the *Developer Guide*.\n When you enable automatic rotation, KMS automatically creates new key material for the KMS key one year after the enable date and every year thereafter. KMS retains all key material until you delete the KMS key. Fo",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enabled": {
        "computed": true,
        "description": "Specifies whether the KMS key is enabled. Disabled KMS keys cannot be used in cryptographic operations.\n When ` + "`" + `` + "`" + `Enabled` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `, the *key state* of the KMS key is ` + "`" + `` + "`" + `Enabled` + "`" + `` + "`" + `. When ` + "`" + `` + "`" + `Enabled` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `, the key state of the KMS key is ` + "`" + `` + "`" + `Disabled` + "`" + `` + "`" + `. The default value is ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `.\n The actual key state of the KMS key might be affected by actions taken outside of CloudFormation, such as running the [EnableKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_EnableKey.html), [DisableKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_DisableKey.html), or [ScheduleKeyDeletion](https://docs.aws.amazon.com/kms/latest/APIReference/API_ScheduleKeyDeletion.html) operations.\n For information about the key states of a KMS key, see [Key state: Effect on your KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the *Developer Guide*.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "key_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "key_policy": {
        "computed": true,
        "description": "The key policy to attach to the KMS key.\n If you provide a key policy, it must meet the following criteria:\n  +  The key policy must allow the caller to make a subsequent [PutKeyPolicy](https://docs.aws.amazon.com/kms/latest/APIReference/API_PutKeyPolicy.html) request on the KMS key. This reduces the risk that the KMS key becomes unmanageable. For more information, see [Default key policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) in the *Developer Guide*. (To omit this condition, set ` + "`" + `` + "`" + `BypassPolicyLockoutSafetyCheck` + "`" + `` + "`" + ` to true.)\n  +  Each statement in the key policy must contain one or more principals. The principals in the key policy must exist and be visible to KMS. When you create a new AWS principal (for example, an IAM user or role), you might need to enforce a delay before including the new principal in a key policy because the new principal might not be immediately visible to KMS. For more information, see [",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key_spec": {
        "computed": true,
        "description": "Specifies the type of KMS key to create. The default value, ` + "`" + `` + "`" + `SYMMETRIC_DEFAULT` + "`" + `` + "`" + `, creates a KMS key with a 256-bit symmetric key for encryption and decryption. In China Regions, ` + "`" + `` + "`" + `SYMMETRIC_DEFAULT` + "`" + `` + "`" + ` creates a 128-bit symmetric key that uses SM4 encryption. You can't change the ` + "`" + `` + "`" + `KeySpec` + "`" + `` + "`" + ` value after the KMS key is created. For help choosing a key spec for your KMS key, see [Choosing a KMS key type](https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-choose.html) in the *Developer Guide*.\n The ` + "`" + `` + "`" + `KeySpec` + "`" + `` + "`" + ` property determines the type of key material in the KMS key and the algorithms that the KMS key supports. To further restrict the algorithms that can be used with the KMS key, use a condition key in its key policy or IAM policy. For more information, see [condition keys](https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms) in the *Developer Guide*.\n  If you change the value of the ` + "`" + `` + "`" + `KeySpec` + "`" + `` + "`" + ` property on an existing KMS key, the u",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key_usage": {
        "computed": true,
        "description": "Determines the [cryptographic operations](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations) for which you can use the KMS key. The default value is ` + "`" + `` + "`" + `ENCRYPT_DECRYPT` + "`" + `` + "`" + `. This property is required for asymmetric KMS keys and HMAC KMS keys. You can't change the ` + "`" + `` + "`" + `KeyUsage` + "`" + `` + "`" + ` value after the KMS key is created.\n  If you change the value of the ` + "`" + `` + "`" + `KeyUsage` + "`" + `` + "`" + ` property on an existing KMS key, the update request fails, regardless of the value of the [UpdateReplacePolicy attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatereplacepolicy.html). This prevents you from accidentally deleting a KMS key by changing an immutable property value.\n  Select only one valid value.\n  +  For symmetric encryption KMS keys, omit the property or specify ` + "`" + `` + "`" + `ENCRYPT_DECRYPT` + "`" + `` + "`" + `.\n  +  For asymmetric KMS keys with RSA key material, specify ` + "`" + `` + "`" + `ENCRYPT_DECRYPT` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `SIGN_VERIFY` + "`" + `` + "`" + `.\n  +  For asymmetric KMS keys with ECC key material, specify",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "multi_region": {
        "computed": true,
        "description": "Creates a multi-Region primary key that you can replicate in other AWS-Regions. You can't change the ` + "`" + `` + "`" + `MultiRegion` + "`" + `` + "`" + ` value after the KMS key is created.\n For a list of AWS-Regions in which multi-Region keys are supported, see [Multi-Region keys in](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html) in the **.\n  If you change the value of the ` + "`" + `` + "`" + `MultiRegion` + "`" + `` + "`" + ` property on an existing KMS key, the update request fails, regardless of the value of the [UpdateReplacePolicy attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatereplacepolicy.html). This prevents you from accidentally deleting a KMS key by changing an immutable property value.\n  For a multi-Region key, set to this property to ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `. For a single-Region key, omit this property or set it to ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `. The default value is ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `.\n *Multi-Region keys* are an KMS feature that lets you create multiple interoperable KMS keys in different AWS-Regions. Bec",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "origin": {
        "computed": true,
        "description": "The source of the key material for the KMS key. You cannot change the origin after you create the KMS key. The default is ` + "`" + `` + "`" + `AWS_KMS` + "`" + `` + "`" + `, which means that KMS creates the key material.\n To [create a KMS key with no key material](https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys-create-cmk.html) (for imported key material), set this value to ` + "`" + `` + "`" + `EXTERNAL` + "`" + `` + "`" + `. For more information about importing key material into KMS, see [Importing Key Material](https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html) in the *Developer Guide*.\n You can ignore ` + "`" + `` + "`" + `ENABLED` + "`" + `` + "`" + ` when Origin is ` + "`" + `` + "`" + `EXTERNAL` + "`" + `` + "`" + `. When a KMS key with Origin ` + "`" + `` + "`" + `EXTERNAL` + "`" + `` + "`" + ` is created, the key state is ` + "`" + `` + "`" + `PENDING_IMPORT` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `ENABLED` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `. After you import the key material, ` + "`" + `` + "`" + `ENABLED` + "`" + `` + "`" + ` updated to ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `. The KMS key can then be used for Cryptographic Operations. \n  CFN doesn't support creating an ` + "`" + `` + "`" + `Origin` + "`" + `` + "`" + ` parameter of the ` + "`" + `` + "`" + `AWS_CLOUDHSM` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `EXTERNAL_KEY_STORE` + "`" + `` + "`" + ` values.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pending_window_in_days": {
        "computed": true,
        "description": "Specifies the number of days in the waiting period before KMS deletes a KMS key that has been removed from a CloudFormation stack. Enter a value between 7 and 30 days. The default value is 30 days.\n When you remove a KMS key from a CloudFormation stack, KMS schedules the KMS key for deletion and starts the mandatory waiting period. The ` + "`" + `` + "`" + `PendingWindowInDays` + "`" + `` + "`" + ` property determines the length of waiting period. During the waiting period, the key state of KMS key is ` + "`" + `` + "`" + `Pending Deletion` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `Pending Replica Deletion` + "`" + `` + "`" + `, which prevents the KMS key from being used in cryptographic operations. When the waiting period expires, KMS permanently deletes the KMS key.\n KMS will not delete a [multi-Region primary key](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html) that has replica keys. If you remove a multi-Region primary key from a CloudFormation stack, its key state changes to ` + "`" + `` + "`" + `PendingReplicaDeletion` + "`" + `` + "`" + ` so it cannot be replicated or used in cryptographic ope",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tags": {
        "computed": true,
        "description": "Assigns one or more tags to the replica key.\n  Tagging or untagging a KMS key can allow or deny permission to the KMS key. For details, see [ABAC for](https://docs.aws.amazon.com/kms/latest/developerguide/abac.html) in the *Developer Guide*.\n  For information about tags in KMS, see [Tagging keys](https://docs.aws.amazon.com/kms/latest/developerguide/tagging-keys.html) in the *Developer Guide*. For information about tags in CloudFormation, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).",
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
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::KMS::Key` + "`" + `` + "`" + ` resource specifies an [KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#kms_keys) in KMSlong. You can use this resource to create symmetric encryption KMS keys, asymmetric KMS keys for encryption or signing, and symmetric HMAC KMS keys. You can use ` + "`" + `` + "`" + `AWS::KMS::Key` + "`" + `` + "`" + ` to create [multi-Region primary keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html#mrk-primary-key) of all supported types. To replicate a multi-Region key, use the ` + "`" + `` + "`" + `AWS::KMS::ReplicaKey` + "`" + `` + "`" + ` resource.\n  If you change the value of the ` + "`" + `` + "`" + `KeySpec` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `KeyUsage` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `Origin` + "`" + `` + "`" + `, or ` + "`" + `` + "`" + `MultiRegion` + "`" + `` + "`" + ` properties of an existing KMS key, the update request fails, regardless of the value of the [UpdateReplacePolicy attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatereplacepolicy.html). This prevents you from accidentally deleting a KMS key by changing any of its immutable property values.\n   KMS replaced th",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccKmsKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccKmsKey), &result)
	return &result
}
