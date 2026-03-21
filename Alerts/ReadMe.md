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
REGION='us-east-2'
S3BUCKET='kengraf-alerts'  # Needs to be globally unique and lowercase
ACCOUNT_ID='788715698479' 
```
### Create a bucket if you don't have one ready
```
aws s3api create-bucket \
  --bucket $S3BUCKET \
  --region $REGION \
  --create-bucket-configuration LocationConstraint=$REGION

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
      "Resource": "arn:aws:s3:::$(S3_BUCKET)"
    },
    {
      "Sid": "AWSCloudTrailWrite",
      "Effect": "Allow",
      "Principal": { "Service": "cloudtrail.amazonaws.com" },
      "Action": "s3:PutObject",
      "Resource": "arn:aws:s3:::$(S3_BUCKET)/AWSLogs/$(ACCOUNTID)/*",
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

### Create trail
aws cloudtrail create-trail \
  --name alerts-trail \
  --s3-bucket-name $S3BUCKET \
  --is-multi-region-trail

aws cloudtrail put-event-selectors \
  --trail-name alerts-trail \
  --event-selectors '[
    {
      "ReadWriteType": "WriteOnly",
      "IncludeManagementEvents": true
    }
  ]'

aws cloudtrail start-logging --name iam-trail

# create sns topic
aws sns create-topic --name api-event-topic

# subscribe (offline confirmation)
aws sns subscribe \
  --topic-arn <TOPIC_ARN> \
  --protocol email \
  --notification-endpoint you@example.com

# create event-bridge rule
aws events put-rule \
  --name iam-write-rule \
  --event-pattern '{
    "source": ["aws.iam"],
    "detail-type": ["AWS API Call via CloudTrail"],
    "detail": {
      "readOnly": [false]
    }
  }'

# Allow eventbridge to publish to sns
aws sns set-topic-attributes \
  --topic-arn <TOPIC_ARN> \
  --attribute-name Policy \
  --attribute-value '{
    "Version": "2012-10-17",
    "Statement": [{
      "Effect": "Allow",
      "Principal": { "Service": "events.amazonaws.com" },
      "Action": "sns:Publish",
      "Resource": "<TOPIC_ARN>"
    }]
  }'

```
RULE_NAME='root-usage'
aws events put-targets \
  --rule <RULE_NAME> \
  --targets "Id"="1","Arn"="<TOPIC_ARN>"
```
