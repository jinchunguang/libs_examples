package main

import (
	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v3"
	"log"
	"strconv"
	"sync"
)

func main() {

	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		log.Panic(err)
	}
	hook, err := elogrus.NewElasticHook(client, "localhost", logrus.InfoLevel, "mylog")
	if err != nil {
		log.Panic(err)
	}
	logrus.AddHook(hook)

	ct := 10
	wg := &sync.WaitGroup{}
	wg.Add(ct)

	for i := 0; i < ct; i++ {
		index := strconv.Itoa(i)
		logrus.WithFields(logrus.Fields{
			"name": "joe" + index,
			"age":  42 + i,
		}).Error("Hello world!" + index)
		wg.Done()
	}
	wg.Wait()
}
