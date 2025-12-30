# GCP												
1.	Create an account on GCP, if needed. [LINK](https://console.cloud.google.com/)
2.	Here is the URL you will need to access in order to request a $50 Google Cloud coupon. You will be asked to provide your school email address and name. An email will be sent to you to confirm these details before a coupon is sent to you. [Student Coupon Retrieval Link](h[ttps://nam12.safelinks.protection.outlook.com/?url=https%3A%2F%2Fgcp.secure.force.com%2FGCPEDU%3Fcid%3Dq7L0A2AHlh%252BVwJyYftvWjikKyDKmniI5F6MnNM3TNDyKhNM3NLxLqL4vdGZ%252BfQM0%2F&data=05%7C02%7Cken.graf%40unh.edu%7Ca43680c3100045c235a308dd2388b921%7Cd6241893512d46dc8d2bbe47e25f5666%7C0%7C0%7C638705792348872738%7CUnknown%7CTWFpbGZsb3d8eyJFbXB0eU1hcGkiOnRydWUsIlYiOiIwLjAuMDAwMCIsIlAiOiJXaW4zMiIsIkFOIjoiTWFpbCIsIldUIjoyfQ%3D%3D%7C0%7C%7C%7C&sdata=Eqi%2BOyrbiGqDhktdNFUbVR9im5RxQaDr%2FOtHavnrzAc%3D&reserved=0](https://nam12.safelinks.protection.outlook.com/?url=https%3A%2F%2Fgcp.secure.force.com%2FGCPEDU%3Fcid%3D94o8bP7yNl0cKqVcfzjMu3O5NyWMY5ZTevQW32Uluj%252B1UYUN7E3jN10qQ5QlcIxj%2F&data=05%7C02%7Cken.graf%40unh.edu%7C072e9eeabf8c461e413108de32c06c72%7Cd6241893512d46dc8d2bbe47e25f5666%7C0%7C0%7C639003999200054432%7CUnknown%7CTWFpbGZsb3d8eyJFbXB0eU1hcGkiOnRydWUsIlYiOiIwLjAuMDAwMCIsIlAiOiJXaW4zMiIsIkFOIjoiTWFpbCIsIldUIjoyfQ%3D%3D%7C0%7C%7C%7C&sdata=JpQG6m36UVSgtWlUm0iHojpfMWmS03LRbM6x4Kqn%2Fws%3D&reserved=0))
3.	It is HIGHLY RECOMMENDED that you enable MFA.  Provide a screenshot in your lab report showing the console, it should look similar to this:<br/>   ![Console](Lab1-GCP-console.png)
4.	Open a cloudshell and issue the following commands, replace environment variable with your values. 
```
BILLING_ACCOUNT_ID="YOUR_BILLING_ACCOUNT_ID"
BUDGET_NAME="DailyBudgetAlert"
```
```
gcloud billing budgets create --billing-account=$BILLING_ACCOUNT_ID \
 --display-name=$BUDGET_NAME --budget-amount=0.01USD --threshold-rule=percent=1.0 \
 --threshold-rule=percent=1.0,basis=forecasted-spend
```
# Lab report
In addition to the console screenshot, provide a screenshot of the command shell interaction in your lab report.  
![Budget](Lab1-GCP-budget.png)
