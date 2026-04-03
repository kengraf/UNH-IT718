Command to generate data (change the detes to matxh your needs)
```
aws ce get-cost-and-usage   --time-period Start=$(date -d "8 days ago" +%Y-%m-%d),End=$(date +%Y-%m-%d) --granularity DAILY   --metrics "UnblendedCost" "UsageQuantity"   --group-by Type=DIMENSION,Key=SERVICE Type=DIMENSION,Key=USAGE_TYPE   --output json > weekly2-cost.json
```

[Render cost report](https://github.com/kengraf/UNH-IT718/blob/main/Costs/cost-report.html)
