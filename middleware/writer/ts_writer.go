package main

import (
	// go modules
	"encoding/json"
	"strings"
	"time"
	"os"
	"os/signal"
	"context"

	// git modules
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
	kafkaReaderCount = 5
	tsParserCount    = 10
	kvWriterCount    = 10
	kafkaMsgChan     = make(chan kafka.Message, chanBufferSize)
	tskvChan         = make(chan tskv, chanBufferSize*3)
	kvCounterChan    = make(chan int8, chanBufferSize*3)
	etc              = mw.KV{}
	kf               = mw.KafkaReaders{}
	brokers          []string
	logger           = mw.Logger{}
	ctx, cancel = context.WithCancel(context.Background())
)

type tskv struct {
	key   string
	value string
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
}

func main() {
	groupConsumer()
}

func groupConsumer() {
	done := make(chan bool)
	for i := 0; i < kafkaReaderCount; i++ {
		go readKafkaMsg(kf.Readers[i], kafkaMsgChan)
	}

	for i := 0; i < tsParserCount; i++ {
		go parseKafkaMsg(kafkaMsgChan, tskvChan)
	}

	for i := 0; i < kvWriterCount; i++ {
		go writeKV(tskvChan, kvCounterChan)
	}

	go msgCounter(kvCounterChan)

	quitHandler()
	<-done
}

func msgCounter(kvCounterChan <-chan int8) {
	var kvCounter int64
	for {
		select {
		case <-ctx.Done():
			logger.Info("stopping counter")
			return
		case <-kvCounterChan:
			kvCounter++
			logger.Info("Total messages written to etcd", zap.Int64("count", kvCounter))
		}
	}

}

func writeKV(tskvChan <-chan tskv, kvCounterChan chan<- int8) {
	for {
		select {
		case <-ctx.Done():
			logger.Info("stopping parser")
			return
		default:
			kv := <-tskvChan
			etc.Put(kv.key, kv.value)
			kvCounterChan <- 1
		}
	}

}

func parseKafkaMsg(msgChan <-chan kafka.Message, tskvChan chan<- tskv) {
	for {
		select {
		case <-ctx.Done():
			logger.Info("stopping parser")
			return
		default:		
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
			}
			logger.Info("Time elapsed for parser",
				zap.String("duration", time.Since(start).String()),
			)
		}
	}
}

func readKafkaMsg(reader *kafka.Reader, kafkaMsgChan chan<- kafka.Message){
	for {
		select {
		case <-ctx.Done():
			logger.Info("stopping reader")
			return
		default:
			kafkaMsgChan <- kf.Read(reader)
		}
		
	}
}

func quitHandler() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		<-sig
		logger.Info("\r- Ctrl+C pressed - stopping writer now")
		cancel()
	}()
}