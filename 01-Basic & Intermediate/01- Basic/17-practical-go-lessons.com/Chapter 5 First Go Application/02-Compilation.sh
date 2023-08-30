cd Documents/code/dateAndTime
go build main.go

ls -lh
# Output:
# total 4160
# -rwxr-xr-x  1 maximilienandile  staff   2.0M Aug 16 11:27 main
# -rw-r--r--  1 maximilienandile  staff    94B Aug 16 11:00 main.go


# We use the command ls. (for windows user, you can use the command dir). You can see that we have two files :

#     main that weights 2.0M (MegaBytes), and that is executable.

#     main.go that weights only 94 Bytes. (this is our source file).

# Now it’s time to launch our application :

./main
# Output: 
# 2019-08-16 11:45:44.435637 +0200 CEST m=+0.000263533
