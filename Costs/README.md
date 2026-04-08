Command to generate data (change the detes to matxh your needs)
```
aws ce get-cost-and-usage   --time-period Start=2026-04-07,End=2026-04-8 --granularity DAILY   --metrics "UnblendedCost" "UsageQuantity"   --group-by Type=DIMENSION,Key=SERVICE Type=DIMENSION,Key=USAGE_TYPE   --output json > weekly2-cost.json
```

[Render cost report](https://github.com/kengraf/UNH-IT718/blob/main/Costs/cost-report.html)
