# Database Session

> [!NOTE]
> This lab only needs to be completed using AWS.

> [!IMPORTANT]
> Instructions are provided for AWS, Azure, and GCP.  Pick only one for your lab report.

You will deploy an AWS Lambda function written in python (code provided).
The function takes two http get requests:
- /new?email=testmail.com
- /get?email=testmail.com&uuid=uuid-value

/new generates, stores, and returns a uuid as a cookie.  Multiple calls to /new will update the uuid value.  
/get validates the UUID against the stored value.  

The ?mail query parameter value is *NOT* checked for validity.  We will leave that for the next lab.
