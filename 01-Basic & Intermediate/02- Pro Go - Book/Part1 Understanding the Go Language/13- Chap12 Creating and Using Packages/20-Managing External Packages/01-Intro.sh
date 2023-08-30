# The go get command adds dependencies to the go.mod file, 
# but these are not removed automatically if the
# external package is no longer required.


# To update the go.mod file to reflect the change, run the command
# in the package folder:
go mod tidy`