go test -coverprofile=coverage.out ./...
awk '!/_test.go/' coverage.out > temp && mv temp coverage.out
awk '!/_mock.go/' coverage.out > temp && mv temp coverage.out
awk '!/internal\/assert/' coverage.out > temp && mv temp coverage.out
go tool cover -func=coverage.out