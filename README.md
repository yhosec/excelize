# Go Excelize
Tools to help with work related to Excel files.

## System Requirement
Build with `go1.22.1`
    
## Usage/Examples
### csv-to-excel
```bash
go run bin/csv-to-excel/main.go -prefix="RewardTiketPoints_202402" -separator="##"
```

### find-on-sheets
```bash
go run bin/find-on-sheets/main.go -file="RewardTiketPoints_202404.xlsx" -sheet="Sheet1" -key="B,G" -search="B,C"
```
