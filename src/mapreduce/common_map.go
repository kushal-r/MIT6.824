package mapreduce

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"os"
)

func doMap(
	jobName string, // the name of the MapReduce job
	mapTask int, // map task number
	inFile string,
	nReduce int, // the number of reduce task that will be run ("R" in the paper)
	mapF func(filename string, contents string) []KeyValue,
) {
	content, err := ioutil.ReadFile(inFile)
	if err != nil {
		fmt.Println("Read file failed")
		return
	}

	kvs := mapF(inFile, string(content))
	files := make([]*os.File, nReduce)
	encs := make([]*json.Encoder, nReduce)

	for i := 0; i < nReduce; i++ {
		f, err := os.Create(reduceName(jobName, mapTask, i))
		if err != nil {
			fmt.Print("Create file failed")
		} else {
			files[i] = f
			encs[i] = json.NewEncoder(f)
		}
	}

	for _, kv := range kvs {
		r := ihash(kv.Key) % nReduce
		if encs[r] != nil {
			encs[r].Encode(&kv)
		}
	}

	for _, f := range files {
		f.Close()
	}
}

//
func ihash(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32() & 0x7fffffff)
}
