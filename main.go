package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {
	kubeConfigPath := flag.String("kc", GetDefaultKubectlConfigPath(), "kubectl config file location.")
	groupBy := flag.String("gb", GetDefaultGroup(), fmt.Sprintf("Group output values by [%s]", strings.Join(ListAvailableGroupers(), ",") ))
	writerName := flag.String("wr", GetDefaultWriter(), fmt.Sprintf("Write to [%s]", strings.Join(ListAvailableWriters(), ",") ))
	flag.Parse()

	data,err := GetPods(*kubeConfigPath)
	if err != nil {
		log.Panicln(err)
	}

	data,err = Group(*groupBy, data)
	if err != nil {
		log.Panicln(err)
	}

	err = Write(*writerName, data)
	if err != nil {
		log.Panicln(err)
	}
}
