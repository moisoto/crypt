# Perform Unit Testing and save a coverage profile
go test -v -cover -coverprofile=coverage.tmp

# Get me an HTML with detailed info on code touched by Unit Testing
go tool cover -html=coverage.tmp

# Remove Temp File
rm coverage.tmp
