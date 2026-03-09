# OpenID Connect (OIDC)
Using the lessons learned in the previous labs, build out a complete serverless application.

### What the Application does
- Forces a login based on Google ID. (/index.html redirects to /login.html) 
- The OIDC JWT is sent as a POST callback to /v1/verify_token.
- The callback is handled by a lambda function:
  - Generates a UUID
  - Stores the UUIC and OIDC provided email in DynamoDB
  - Redirects to /dashboard.html
- /dashboard.html displays: OIDC JWT contents, email, and UUID values.

### Components
| CF* | Function | Purpose | Notes |  
| :---: | :---: | :--- | :--- | 
| ❌ | Github | Source (App &IoC) | Clone locally for customization
| ❌ | Google | OIDC provider | Generate client secret; set scopes
| ❌ | CloudFormation | IoC | Need to set custom values
| ✅ | S3 Bucket | Static web content & Lambda packages | Globally unique; user defined name 
| ✅ | Lambda | OIDC callback and session creation | 
| ✅ | DynamoDB | Storage of session UUID | 
| ✅ | API GatewayV2 | Control access to Lambda functions 
| ✅ | Route53 | Provide friendly URL | Optional: requires domain oownership 
| ✅ | CloudFront | CDN for static pages and controls access to ApiGatewayV2 | 

CF*: IoC deployment based on CloudFormation

## Setup (Github, Google, CloudFormation)
### Github
Clone this repo to AWS cloud shell or your local machine.
Optional: Fork this repo to allow for automated workflows and making your work public.
### Google
Follow the Google provided steps to create OAuth 2.0 Client IDs: [LINK](https://developers.google.com/identity/openid-connect/openid-connect)  

> [!IMPORTANT]
> You will need to come back and adjust these settings once the CloudFront URLs are known.

__URIs:__
- The authorized Javascript origin will limit where your client ID can be used.
- The redirect URL will limit where the callback can be redirected to.
- You can have multiple values: (dev, testing, and production)
- Ports matter:  http://localhost and http://localhost:8080 are not the same.

> [!IMPORTANT]
> Edit the lambda function deploy/lambda/verifyToken.py and website/login.html to use your client id.

You can click on the more info button (upper right) to see your client id and secret.
![console capture](images/gcp-console.png)


### CloudFormation
Three (3) CloudFormation templates have been defined.  Storage, Backend, and Distribution.  
A shell script (deploy.sh) has been provided to deploy each of these stacks.  
deploy.sh takes one argument.  A prefix name to be used in naming resources.

> [!TIP]
> Make your prefix name globally unique and lowercase.  This is a S3 limitation.  "it718" is not going to fly.

When a stack deployment completes, one or more URLs will be shown.  You can use these URLs to connect to your S3, Lambda, API, etc.  
Copy the CloudFront URL.  Example: "https://d1mvssppd7zkjp.cloudfront.net"  Go back to the Google Development console and add this as "Authorized JavaScript origins" and to "Authorized redirect URIs" add the url with "/v1/verifyToken" append.  Example: "https://d1mvssppd7zkjp.cloudfront.net/v1/verifyToken"

### S3
The bucket holds the ./website content and the ./deploy/lambda zip package

### Lambda
verifyToken.py handles the OIDC callback, generates a uuid which is stored in DynamoDB, and returns the Google generated JWT.

### DynamoDB
Storage of session uuid, repeated calls are handled as overwrites.

### API GatewayV2
Defines one route /v1/verifyToken.  POST requests that invokes the lambda function.

### Route53 (optional)
Provides custom (friendly) URL to CloudFront.  If you own a domain this is a easier and more predicitable way to setup Google as a OIDC provider.

### CloudFront
Used to cache and serve static files (e.g., HTML, CSS, JavaScript) from an S3 bucket to an origin close to your users for low latency.  It also controls access to the API Gateway: for dynamic requests (e.g., POST, GET, PUT) to your backend services or AWS Lambda.

## Lab Report
Submit to Canvas your login URL.  This is pass/fail based on my login to your site with my Google id.
