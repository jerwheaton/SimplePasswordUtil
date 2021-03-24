## Simple Password Utility
Contains tools for ensuring high quality passphrases.

### Common Password
`check` command is used to check if a password is in a list of plaintext passwords. Provided in `/sample` is a list of 1000000 common passwords which can be used for testing.

##### Flags
- `-b`: Use bloom filter, has a relatively low rate of false positives for the given sample file, but is actually slower than standard check. Use if memory use is a concern.
```sh
➜ make
➜ ./build/password check sample/passwords.txt password1
Matched:  true
➜ ./build/password check sample/passwords.txt '%q5vo!&c5hEUki1@byNRLvM$'
Matched:  false
```

### Rate Password
Uses a simplified string entropy formula and weighs it against the number of unique characters in the string. Outputs a float value greater than zero, which is also greater than `1` if the quality meets a preset threshold.
```sh
➜ ./build/password rate 'password1'
Rated:  0.439
➜ ./build/password rate '%q5vo!&c5hEUki1@byNRLvM$'
Rated:  1.427
```
