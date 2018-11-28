### Part I: Map/Reduce input and output
```shell
 cd ~/src/mapreduce
 go test -run Sequential
```

### Part II: Single-worker word count

```shell
cd ~/src/main
./ test-wc.sh
```

### Part III: Distributing MapReduce tasks
```shell
cd ~/src/mapreduce
go test -run TestParallel
```

### Part IV: Handling worker failures
```shell
cd ~/src/mapreduce
go test -run Failure
```