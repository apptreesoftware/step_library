# Creating a new GO step package

1. Create a new Folder in step_library
2. Create a package.yaml
3. In Package.yaml, provide the name of your package
4. Create a main.go
5. go mod init <package_name>
6. Import the SDK into main.go by running `go get github.com/apptreesoftware/go-workflow` from within the new package directory

# Creating a new Javascript (node) step package

1. Create a new folder in step_library
2. Create a package.yaml using the following template
3. use the following template
```
name: json
version: 1.0

executables:
  linux:
    amd64: index-linux
  darwin:
    amd64: index-macos
  windows:
    amd64: index-windows.exe

steps:
  <step_name>@1.0:
    inputs:
      <input1>:
        required: true
        description: The text you want to parse as JSON
        sample: '{"key" : "value"}'
    outputs:
      <output1>:
        description: The record that was parsed from the JSON string

```
4. Create index.js
5. Run `npm init`
6. Run `npm install apptree-workflow-sdk`
7. Write your step functions
8. Register them with the SDK using `apptree.addStep()`

## To build your package for for publishing
1. Install `pkg` using npm. Run `npm install -g pkg`
2. `cd` into the package folder
3. Run `pkg -t node12-linux-x64,node12-macos-x64 .` or 'pkg -t node12-linux-x64,node12-macos-x64 index.js' 

## To publish your package
1. `cd` into the package folder
2. `apptree publish package -d . --host https://platform.apptreeio.com`

## VS Code launch.json configuration template

```
{
    "version": "0.2.0",
    "configurations": [
        {
            "type": "node",
            "request": "launch",
            "name": "Launch Program",
            "program": "${workspaceFolder}/<pkg_folder>/index.js",
            "cwd": "${workspaceFolder}/<pkg_folder>",
            "args": [
                "--serve" //Optional if you want to run in serve mode
                "--port" //Optional if you want to change the port
                "5001"
            ]
        }
    ]
}
```