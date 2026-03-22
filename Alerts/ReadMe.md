# Important Security Alerts

|  | Name | Reason |
| :---: | :--- | :--- | 
| ✅ | root-usage | Root usage should be extremely rare. | 
| ✅ | iam-change | Covers most privilege changes. | 
| ✅ | new-key | Common persistence mechanism.  |    
| ✅ | policy-change | Direct privilege escalation paths.  |  
| ✅ | log-tampering | Attackers try to disable logging first.   |  
| ✅ | console-attacks | Detect brute force / credential stuffing.   |    
| ✅ | public-s3 | Common data exposure vector.   |  

## Groundwork before we start creating rules

### Varible definitons (pick yours)
```
REGION='us-east-1'  # FYI IAM events goto CloudTail in us-east-1
S3BUCKET='kengraf-alerts'  # Needs to be globally unique and lowercase
ACCOUNT_ID='788715698479'
EMAIL='kengraf57@gmail.com'
```

### Create SNS topic
```
TOPIC_ARN=`aws sns create-topic --name alerts-topic --output text`
```

### Subscribe
> [!NOTE]
> Email confirmation, check spam folder.  
> This is an instance where you need a few seconds between commands.
```
aws sns subscribe \
  --topic-arn $TOPIC_ARN \
  --protocol email \
  --notification-endpoint $EMAIL

```

### Create a bucket if you don't have one ready
```
aws s3api create-bucket \
  --bucket $S3BUCKET \
  --region $REGION

```
### Set bucket policy to allow CloudTrail writes
```
cat <<EOF > cloudtrail-bucket-policy.json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "AWSCloudTrailAclCheck",
      "Effect": "Allow",
      "Principal": { "Service": "cloudtrail.amazonaws.com" },
      "Action": "s3:GetBucketAcl",
      "Resource": "arn:aws:s3:::${S3BUCKET}"
    },
    {
      "Sid": "AWSCloudTrailWrite",
      "Effect": "Allow",
      "Principal": { "Service": "cloudtrail.amazonaws.com" },
      "Action": "s3:PutObject",
      "Resource": "arn:aws:s3:::${S3BUCKET}/AWSLogs/${ACCOUNT_ID}/*",
      "Condition": {
        "StringEquals": {
          "s3:x-amz-acl": "bucket-owner-full-control"
        }
      }
    }
  ]
}
EOF

```
### Apply policy
```
aws s3api put-bucket-policy \
  --bucket $S3BUCKET \
  --policy file://cloudtrail-bucket-policy.json

```

### Create and start trail
```
aws cloudtrail create-trail \
  --name alerts-trail \
  --s3-bucket-name $S3BUCKET \
  --is-multi-region-trail
aws cloudtrail start-logging --name alerts-trail

```

### Allow eventbridge to publish to sns
```
aws sns set-topic-attributes \
  --topic-arn $TOPIC_ARN \
  --attribute-name Policy \
  --attribute-value '{
    "Version": "2012-10-17",
    "Statement": [{
      "Effect": "Allow",
      "Principal": { "Service": "events.amazonaws.com" },
      "Action": "sns:Publish",
      "Resource": "'$TOPIC_ARN'"
    }]
  }'
```
## Seven great opex/security rules
#### 1. Root account usage (critical)
```
NAME='root-usage'
aws events put-rule \
  --region $REGION \
  --name $NAME \
  --event-pattern '{
    "detail-type": ["AWS Console Sign In via CloudTrail"],
    "detail": {
      "userIdentity": {
        "type": ["Root"]
      }
    }
  }'
aws events put-targets --rule $NAME \
  --targets "Id"=$NAME,"Arn"="$TOPIC_ARN"
```

#### 2. IAM changes (core coverage)
```
NAME="iam-change"
aws events put-rule \
  --region $REGION \
  --name $NAME \
  --event-pattern '{
    "source": ["aws.iam"],
    "detail-type": ["AWS API Call via CloudTrail"],
    "detail": {
      "readOnly": [false]
    }
  }'
aws events put-targets --rule $NAME \
  --targets "Id"=$NAME,"Arn"="$TOPIC_ARN"
```
  
#### 3. ACCESS KEY CREATION (high-risk)
```
NAME="new-key"
aws events put-rule \
  --name $NAME \
  --region $REGION \
  --event-pattern '{
    "detail-type": ["AWS API Call via CloudTrail"],
    "detail": {
      "eventSource": ["iam.amazonaws.com"],
      "eventName": ["CreateAccessKey"]
    }
  }'
aws events put-targets --rule $NAME \
  --targets "Id"=$NAME,"Arn"="$TOPIC_ARN"
```
   
#### 4. POLICY / PRIVILEGE ESCALATION
```
NAME="policy-change"
aws events put-rule \
  --name $NAME \
  --region $REGION \
  --event-pattern '{
    "detail-type": ["AWS API Call via CloudTrail"],
    "detail": {
      "eventSource": ["iam.amazonaws.com"],
      "eventName": [
        "AttachRolePolicy",
        "PutUserPolicy",
        "PutRolePolicy",
        "AttachUserPolicy",
        "AddUserToGroup"
      ]
    }
  }'
aws events put-targets --rule $NAME \
  --targets "Id"=$NAME,"Arn"="$TOPIC_ARN"
```
   
#### 5. CLOUDTRAIL TAMPERING
```
NAME="log-tampering"
aws events put-rule \
  --name $NAME \
  --region $REGION \
  --event-pattern '{
    "source": ["aws.cloudtrail"],
    "detail-type": ["AWS API Call via CloudTrail"],
    "detail": {
      "eventName": [
        "StopLogging",
        "DeleteTrail",
        "UpdateTrail"
      ]
    }
  }'
aws events put-targets --rule $NAME \
  --targets "Id"=$NAME,"Arn"="$TOPIC_ARN"
```
   
#### 6. CONSOLE LOGIN FAILURES
```
NAME="console-attacks"
aws events put-rule \
  --name $NAME \
  --region $REGION \
  --event-pattern '{
    "detail-type": ["AWS Console Sign In via CloudTrail"],
    "detail": {
      "responseElements": {
        "ConsoleLogin": ["Failure"]
      }
    }
  }'
aws events put-targets --rule $NAME \
  --targets "Id"=$NAME,"Arn"="$TOPIC_ARN"
```
   
### 7. S3 PUBLIC ACCESS CHANGES
```
NAME="public-s3"
aws events put-rule \
  --name $NAME \
  --region $REGION \
  --event-pattern '{
    "source": ["aws.s3"],
    "detail-type": ["AWS API Call via CloudTrail"],
    "detail": {
      "eventName": [
        "PutBucketAcl",
        "PutBucketPolicy",
        "PutPublicAccessBlock"
      ]
    }
  }'
aws events put-targets --rule $NAME \
  --targets "Id"=$NAME,"Arn"="$TOPIC_ARN"
```  
### CLI commands to fire events
```
echo "1. ROOT ACCOUNT USAGE"
echo "⚠️ No CLI for this (must use root login)"

echo "🔐 2. IAM WRITE CHANGES"
USER=test-iam-$(date +%s)
aws iam create-user --user-name $USER

echo "🔐 3. ACCESS KEY CREATION"
USER_KEY=test-key-$(date +%s)
aws iam create-user --user-name $USER_KEY
KEY_ID=`aws iam create-access-key --user-name $USER_KEY`

echo "🔐 4. PRIVILEGE ESCALATION (policy attach)"
USER_PRIV=test-priv-$(date +%s)
aws iam create-user --user-name $USER_PRIV
aws iam attach-user-policy \
  --user-name $USER_PRIV \
  --policy-arn arn:aws:iam::aws:policy/ReadOnlyAccess

echo "🔐 5. CLOUDTRAIL TAMPERING 🚨"
echo "⚠️ Safe test: use a dummy trail, NOT your real one"
aws cloudtrail create-trail \
  --name test-trail \
  --s3-bucket-name $S3BUCKET
aws cloudtrail stop-logging --name test-trail

echo "🔐 6. CONSOLE LOGIN FAILURE"
echo "No CLI. Nust simulate bad login on console"

echo "🔐 7. S3 PUBLIC ACCESS CHANGES"
TEST_BUCKET=it718-test-public-$(date +%s)
aws s3api create-bucket \
  --bucket $BUCKET \
  --region $REGION
echo "The ACL will fail, but the attempt is logged and will fire event
aws s3api put-bucket-acl \
  --bucket $BUCKET \
  --acl public-read
```

### Cleanup test resources
```
aws iam delete-user --user-name $USER
aws iam delete-access-key --user-name $USER_KEY --access-key-id $KEY_ID
aws iam delete-user --user-name $USER_KEY
aws iam detach-user-policy \
  --user-name $USER_PRIV \
  --policy-arn arn:aws:iam::aws:policy/ReadOnlyAccess
aws iam delete-user --user-name $USER_PRIV
aws cloudtrail delete-trail --name test-trail
aws s3api delete-bucket --bucket $TEST_BUCKET
```
