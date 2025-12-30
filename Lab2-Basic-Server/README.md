# Deploy a Basic Server Lab
## Objectives
For each of the three major cloud vendors (AWS, Azure, and GCP) complete the following actions:
1.	Use the command shell to deploy a server running nginx.
2.	Retrieve IP address from console.
3.	Verify running web page.
4.	In the console terminate the running instance.
### Comments about this lab
- There are many, many, ways to deploy a server to the cloud.  This is true for every cloud vendor.  As discussed in class it is easy to follow some web-based wizard and in a few clicks have a running server.
- This lab focuses on the API calls needed for a server deployment.
- In the real-world, deployments abstract away the cloud provider's API into templates, scripts, CI/CD processes.
- Deploying a server requires supporting infrastructure:  networking (VPC, subnets), security (keys, security groups) are common examples.
- The provided scripts assume the basic infrastructure needs to be created.  It is also assumed that you would like to keep the infrastructure created for later labs.
  
**BEFORE JUST BLINDLY CUTTING & PASTING:**  
- Should you alter the scripts to use existing infrastructure?
- Do you want to keep any created resources?  More meaningful names should be considered.
- Do you know what resources you will pay for?  Leaving unused compute running is a bad idea.
- The scripts in these labs do not clean up after themselves.

> [!NOTE]
> ** Forward thinking
> - Your project will require a running server, it can be but does not need to be a web server.  A later lab will show how static pages are best served by configuring cloud storage for that function.

## Lab Report
Provide one or more screenshots for each cloud vendor showing:
1.	The output from the command showing the VM IP addresses.
2.	wget/curl command retrieving the nginx default page,
3.	Successful SSH login.

