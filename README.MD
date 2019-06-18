# Creating a new package

1. Create a new Folder in step_library
2. Create a package.yaml
3. In Package.yaml, provide the name of your package
4. Create a main.go
5. go mod init <package_name>
6. Import the SDK into main.go by running `go get github.com/apptreesoftware/go-workflow` from within the new package directory