# GCP												
1.	Create an account on GCP, if needed. [LINK](https://console.cloud.google.com/)
2.	Here is the URL you will need to access in order to request a $50 Google Cloud coupon. You will be asked to provide your school email address and name. An email will be sent to you to confirm these details before a coupon is sent to you. [Student Coupon Retrieval Link](https://vector.my.salesforce-sites.com/GCPEDU?cid=94o8bP7yNl0cKqVcfzjMu3O5NyWMY5ZTevQW32Uluj%2B1UYUN7E3jN10qQ5QlcIxj/)
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
