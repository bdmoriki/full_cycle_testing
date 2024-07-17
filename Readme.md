go test -coverprofile=coverage.out
go tool cover -html=coverage.out  

go test -bench=.
go test -bench=. -run=^# 
go test -bench=. -run=^# -count=10
go test -bench=. -run=^# -count=100 -benchtime=3s
go test -bench=. -run=^# -benchmem

go test -fuzz=. 
go test -fuzz=. -run=ˆ= -fuzztime=5s