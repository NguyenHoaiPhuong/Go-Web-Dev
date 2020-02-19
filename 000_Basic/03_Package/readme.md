Golang Package

Purpose:
- Import one package into another package
- Export a function in one package and be used in another package

Explanation:
- There are 2 packages in the program: main and stringutil
- Main package will use some functions and variable in the stringutil package
- Functions or variables inside the stringutil package that start with capital letter are able to be exported and used in main package

Go Build:
    go build reverse.go reverseTwo.go
 	won't produce an output file.

Go Install:
    will place the package inside the pkg directory of the workspace.