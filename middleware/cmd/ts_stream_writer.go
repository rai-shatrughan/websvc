package main

import (
	// go modules
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
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
	topic = "ts"
	// groupID          = "iotts-consumer-group"
	chanBufferSize   = 1000
	kafkaReaderCount = 5
	kvWriterCount    = 10
	tsParserCount    = 10
	msgChan          = make(chan kafka.Message, chanBufferSize)
	tskvChan         = make(chan tskv, chanBufferSize*3)
	kvCounterChan    = make(chan int8, chanBufferSize*3)
	etc              = mw.KV{}
	tsdb             = mw.TSDB{}
	readers          []*kafka.Reader
	brokers          = [...]string{"172.18.0.41:9092", "172.18.0.42:9092", "172.18.0.43:9092"}
	logger           *zap.Logger
	err              error
)

type tskv struct {
	key   string
	value string
}

func main() {
	setup()
	groupConsumer()
}

func setup() {
	logger, err = zap.NewProduction()
	if err != nil {
		log.Fatalf("Can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	etc.New()
	for i := 0; i < kafkaReaderCount; i++ {
		logger.Info("Initializing reader",
			zap.Int("reader", i),
		)
		readers = append(readers, reader(i))
	}

	tsdb.New()
}

func groupConsumer() {
	done := make(chan bool)

	for i := 0; i < kafkaReaderCount; i++ {
		go readKafkaMsg(readers[i], msgChan)
	}

	for i := 0; i < tsParserCount; i++ {
		go parseKafkaMsg(msgChan, tskvChan)
	}

	for i := 0; i < kvWriterCount; i++ {
		go writeKV(tskvChan, kvCounterChan)
	}

	go msgCounter(kvCounterChan)

	done <- true

}

func msgCounter(kvCounterChan <-chan int8) {
	var msgCounter int64
	for {
		<-kvCounterChan
		msgCounter++
		// logger.Info("Total messages written to etcd", zap.Int64("count", msgCounter))
		log.Println("Total messages written to etcd", msgCounter)
	}

}

func writeKV(tskvChan <-chan tskv, kvCounterChan chan<- int8) {
	for {
		for kv := range tskvChan {
			go func(kv tskv) {
				etc.Put(kv.key, kv.value)
				kvCounterChan <- 1
			}(kv)
		}
	}

}

func parseKafkaMsg(msgChan <-chan kafka.Message, tskvChan chan<- tskv) {
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
			// dt := *tsa[i].Timestamp

			datapoint.Timestamp = time.Now()
			datapoint.Value = *tsa[i].Value

			for i := 0; i < len(labels); i++ {
				log.Printf("%s ", labels[i].Name)
				log.Printf("%s ", labels[i].Value)
			}

			tsdb.Write(labels, datapoint)
			log.Printf("%s ", datapoint.Timestamp)
			log.Printf("%f ", datapoint.Value)

		}
		logger.Info("Time elapsed for parser",
			zap.String("duration", time.Since(start).String()),
		)
	}
}

func reader(partition int) *kafka.Reader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   brokers[:],
		Topic:     topic,
		Partition: partition,
		// GroupID:  groupID,
		MinBytes: 0,
		MaxBytes: 10e6, // 10MB
	})

	// startTime := time.Now().Add(-time.Minute * 1)
	// r.SetOffsetAt(context.Background(), startTime)
	r.SetOffset(kafka.LastOffset)

	setupCloseReaderHandler(r)

	return r
}

func readKafkaMsg(r *kafka.Reader, msgChan chan<- kafka.Message) {
	for {
		logger.Debug("Reading message from group",
			zap.String("groupId", r.Config().GroupID),
		)
		start := time.Now()
		m, err := r.ReadMessage(context.TODO())

		if err != nil {
			logger.Fatal("Error in Reading from kafka")
		}
		elapsed := time.Since(start)
		logger.Info("Time elapsed for Kafka Reader",
			zap.String("duration", elapsed.String()),
		)
		// TODO: process message
		logger.Info("message at offset : ",
			zap.Int64("offsetValue", m.Offset),
		)
		msgChan <- m
	}
}

func closeReader(r *kafka.Reader) {
	logger.Info("Closing reader...")
	if err := r.Close(); err != nil {
		logger.Fatal("Failed to close reader...")
	}

}

func setupCloseReaderHandler(r *kafka.Reader) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		logger.Info("\r- Ctrl+C pressed in Terminal")
		closeReader(r)
		os.Exit(0)
	}()
}
