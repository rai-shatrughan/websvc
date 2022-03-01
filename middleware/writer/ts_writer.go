package main

import (
	// go modules
	"encoding/json"
	"strings"
	"time"

	// git modules
	"github.com/m3db/prometheus_remote_client_golang/promremote"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"

	// websvc modules
	mw "websvc/middleware"
	"websvc/timeseries/models"
)

var (
	topic            = "ts"
	groupID          = "ts-consumer-group"
	chanBufferSize   = 10000
	kafkaReaderCount = 10
	tsParserCount    = 10
	kvWriterCount    = 20
	m3WriterCount    = 20
	kafkaMsgChan     = make(chan kafka.Message, chanBufferSize)
	tskvChan         = make(chan tskv, chanBufferSize*3)
	m3kvChan         = make(chan m3kv, chanBufferSize*3)
	kvCounterChan    = make(chan int8, chanBufferSize*3)
	m3CounterChan    = make(chan int8, chanBufferSize*3)
	etc              = mw.KV{}
	tsdb             = mw.TSDB{}
	kf               = mw.KafkaReaders{}
	brokers          []string
	err              error
	logger           = mw.Logger{}
)

type tskv struct {
	key   string
	value string
}

type m3kv struct {
	lables    []promremote.Label
	datapoint promremote.Datapoint
}

func init() {
	conf := mw.Config{}
	conf.New()

	logger.New()

	brokers = conf.GetStringSlice("kafka.brokers")
	kf.GroupID = &groupID
	kf.Topic = &topic
	kf.Brokers = brokers
	kf.New(kafkaReaderCount)

	etc.New()
	tsdb.New()
}

func main() {
	groupConsumer()
}

func groupConsumer() {
	done := make(chan bool)

	for i := 0; i < kafkaReaderCount; i++ {
		go kf.ReadStream(kf.Readers[i], kafkaMsgChan)
	}

	for i := 0; i < tsParserCount; i++ {
		go parseKafkaMsg(kafkaMsgChan, tskvChan, m3kvChan)
	}

	for i := 0; i < kvWriterCount; i++ {
		go writeKV(tskvChan, kvCounterChan)
	}

	for i := 0; i < m3WriterCount; i++ {
		go writeM3(m3kvChan, m3CounterChan)
	}

	go msgCounter(kvCounterChan, m3CounterChan)

	done <- true

}

func msgCounter(kvCounterChan <-chan int8, m3CounterChan <-chan int8) {
	var kvCounter, m3Counter int64
	for {
		select {
		case <-kvCounterChan:
			kvCounter++
			logger.Info("Total messages written to etcd", zap.Int64("count", kvCounter))
		case <-m3CounterChan:
			m3Counter++
			logger.Info("Total messages written to M3DB", zap.Int64("count", m3Counter))
		}
	}

}

func writeM3(m3kvChan <-chan m3kv, m3CounterChan chan<- int8) {
	for {
		kv := <-m3kvChan
		tsdb.Write(kv.lables, kv.datapoint)
		m3CounterChan <- 1
	}

}

func writeKV(tskvChan <-chan tskv, kvCounterChan chan<- int8) {
	for {
		kv := <-tskvChan
		etc.Put(kv.key, kv.value)
		kvCounterChan <- 1
	}

}

func parseKafkaMsg(msgChan <-chan kafka.Message, tskvChan chan<- tskv, m3kvChan chan<- m3kv) {
	for {
		start := time.Now()
		msg := <-msgChan
		tsa := models.TimeseriesArray{}
		json.Unmarshal(msg.Value, &tsa)
		for i := range tsa {
			date := strings.Split(tsa[i].Timestamp.String(), "T")

			key := "/" + string(msg.Key) + "/" + date[0] + "/" + tsa[i].Timestamp.String()
			logger.Debug(key)

			ts, _ := json.Marshal(tsa[i])
			val := string(ts)
			logger.Debug(val)
			tskvChan <- tskv{key: key, value: val}

			var labels []promremote.Label
			var label promremote.Label
			label.Name = "property"
			label.Value = *tsa[i].Property

			labels = append(labels, label)

			label.Name = "unit"
			label.Value = *tsa[i].Unit

			labels = append(labels, label)

			label.Name = "__name__"
			label.Value = "asset"

			labels = append(labels, label)

			var datapoint promremote.Datapoint

			datapoint.Timestamp = time.Now()
			datapoint.Value = *tsa[i].Value

			m3kvChan <- m3kv{lables: labels, datapoint: datapoint}

		}
		logger.Info("Time elapsed for parser",
			zap.String("duration", time.Since(start).String()),
		)
	}
}
