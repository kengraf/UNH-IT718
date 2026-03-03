# AWS												
1.	Create a database to store UUIDs
2.	Deploy lambda function with proper permissions
3.	Test functionality

### Create Database
```
DB_ARN=`aws dynamodb create-table \
    --table-name Lab6-session \
    --attribute-definitions AttributeName=email,AttributeType=S \
    --key-schema AttributeName=email,KeyType=HASH \
    --billing-mode PAY_PER_REQUEST --query "TableDescription.TableArn" --output text`
sed -i "s|YOUR_DB_ARN|${DB_ARN}|g" Lab6-dynamodb-policy.json
```
### Create an IAM Role for Lambda
```
aws iam create-role \
    --role-name Lab6-LambdaDynamoDBRole \
    --assume-role-policy-document file://Lab6-trust-policy.json
ROLE_ARN=` aws iam get-role --role-name Lab6-LambdaDynamoDBRole --query "Role.Arn" --output text`
```
### Deploy cloud function
```
zip function.zip Lab6-session.py
aws lambda create-function --function-name Lab6-session --runtime python3.13 \
    --role $ROLE_ARN --handler Lab6-session.lambda_handler --zip-file fileb://function.zip \
    --query "FunctionName" --output text
aws lambda create-function-url-config --function-name Lab6-session --auth-type NONE
aws lambda add-permission --function-name Lab6-session --action lambda:InvokeFunctionUrl \
    --principal "*" --function-url-auth-type NONE --statement-id FunctionURLPublicAccess
```
### Attach DynamoDB Permissions to the Role
```
aws iam put-role-policy \
    --role-name Lab6-LambdaDynamoDBRole \
    --policy-name Lab6-DynamoDBAccessPolicy \
    --policy-document file://Lab6-dynamodb-policy.json
```
### Attach Additional Permissions for Lambda Execution
```
aws iam attach-role-policy \
    --role-name Lab6-LambdaDynamoDBRole \
    --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole
```
### Attach the Role to the Lambda Function
```
aws lambda update-function-configuration \
    --function-name Lab6-session \
    --role $ROLE_ARN  --query "Role" --output text
```
### Retrieve URL for lab report
```
aws lambda get-function-url-config --function-name Lab6-session --query "FunctionUrl" --output text
```
### Lab report commands to validate deployment
The sample screenshot below shows adding a new user (twice to generate a new and old uuid).  Then testing both the new and old uuid.
Provide the screenshot and the function URL as a link in your lab report.

![CLI screen capture](Lab6-AWS-cli.png)
