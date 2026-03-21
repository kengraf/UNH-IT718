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

```
### Create a bucket if you don't have one ready
```
aws s3api create-bucket \
  --bucket $S3BUCKET \
  --region $REGION \
  --create-bucket-configuration LocationConstraint=$REGION

```

```
RULE_NAME='root-usage'
aws events put-targets \
  --rule <RULE_NAME> \
  --targets "Id"="1","Arn"="<TOPIC_ARN>"
```
