# go-crypto-consigliere

[Lambda Execution] →    [Use Case]   → [Domain Logic] → [Exchange Adapter]
                            ↑
                    [Strategy Adapter]

## Debug

You must generate the executable golang debugger 

go build -gcflags="all=-N -l" -o debug_app cmd/main.go

Use './debug_app' to start your application