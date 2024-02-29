Changing the case of file and folder names in a specified directory 

#### Flags
```
-json
    json output
```

```
-path string 
    path to scan and change case folder (default "./dir")
```

```
-upper 
    to upper case
```
#### Example

```
./ccid -json -upper -path /foo
```

#### Build
```
env GOOS=windows go build ccid
env GOOS=linux go build ccid
```

