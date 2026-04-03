Command to generate data (change the detes to matxh your needs)
```
aws ce get-cost-and-usage   --time-period Start=2025-12-01,End=2025-12-02   --granularity DAILY   --metrics "UnblendedCost" "UsageQuantity"   --group-by Type=DIMENSION,Key=SERVICE Type=DIMENSION,Key=USAGE_TYPE   --output json > weekly2-cost.json
```
