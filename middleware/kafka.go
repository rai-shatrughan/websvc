package middleware

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	kafka "github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

var (
	brokers        []string
	groupID            = "websvc-consumer-group"
	minBytes           = 1
	maxBytes       int = 10e6
	startOffset        = kafka.LastOffset
	logger             = Logger{}
)

func init() {
	conf := Config{}
	conf.New()
	brokers = conf.GetStringSlice("kafka.brokers")

	logger.New()
}

type kafkaBase struct {
	Brokers     []string
	Topic       *string
	GroupID     *string
	MinBytes    *int
	MaxBytes    *int
	StartOffset *int64	
}

//KafkaReaders wraps kafka.Reader
type KafkaReaders struct {
	Readers []*kafka.Reader
	kafkaBase
}

//KafkaWriter wraps kafka.Writer
type KafkaWriter struct {
	writer *kafka.Writer
	kafkaBase
}

func (kf *kafkaBase) fillDefaults() {
	if len(kf.Brokers) == 0 {
		kf.Brokers = make([]string, len(brokers))
	}

	if kf.Brokers[0] == "" {
		for i := 0; i < len(brokers); i++ {
			kf.Brokers = append(kf.Brokers, brokers[i])
		}
	}

	if kf.GroupID == nil {
		kf.GroupID = &groupID
	}

	if kf.MinBytes == nil {
		kf.MinBytes = &minBytes
	}

	if kf.MaxBytes == nil {
		kf.MaxBytes = &maxBytes
	}

	if kf.StartOffset == nil {
		kf.StartOffset = &startOffset
	}

}

//New initializes new instance of readers
func (kf *KafkaReaders) New(readerCount int) {
	kf.fillDefaults()
	for i := 0; i < readerCount; i++ {
		reader := kafka.NewReader(kafka.ReaderConfig{
			Brokers:     kf.Brokers[:],
			Topic:       *kf.Topic,
			GroupID:     *kf.GroupID,
			MinBytes:    *kf.MinBytes,
			MaxBytes:    *kf.MaxBytes,
			StartOffset: *kf.StartOffset,
		})
		kf.Readers = append(kf.Readers, reader)
		kf.setupCloseReaderHandler()
	}
}

//New initializes new instance of writer
func (kf *KafkaWriter) New() {
	kf.fillDefaults()
	if kf.writer == nil {
		kf.writer = &kafka.Writer{
			Addr:     kafka.TCP(kf.Brokers[:]...),
			Topic:    *kf.Topic,
			Balancer: &kafka.RoundRobin{},
			Async:    true,
		}
	}
	kf.setupCloseWriterHandler()
}

//Read returns kafka Message
func (kf *KafkaReaders) Read(reader *kafka.Reader) kafka.Message {

	logger.Info("Reading message from group",
		zap.String("groupId", reader.Config().GroupID),
	)

	var start time.Time
	var elapsed time.Duration

	start = time.Now()

	msg, err := reader.ReadMessage(context.Background())
	if err != nil {
		logger.Error("Error in Reading msg from Kafka ", zap.Error(err))
	}
	elapsed = time.Since(start)

	logger.Info("Kafka Read", zap.Duration("duration", elapsed))

	return msg

}

//Write pushes data to kafka
func (kf *KafkaWriter) Write(key, value []byte) {
	start := time.Now()
	logger.Info("Kafka Write", zap.String("start", start.String()))
	err := kf.writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   key,
			Value: value,
		},
	)
	elapsed := time.Since(start)
	logger.Info("Kafka Write", zap.Duration("duration", elapsed))

	if err != nil {
		logger.Error("Error Writing msg to Kafka ", zap.Error(err))
	}

}

//BatchWrite puts data into batches
func (kf *KafkaWriter) BatchWrite(key, val [][]byte) {

	msgs := make([]kafka.Message, len(key))
	for k := range key {
		msgs[k] = kafka.Message{Key: key[k], Value: val[k]}
	}

	start := time.Now()

	err := kf.writer.WriteMessages(context.Background(), msgs...)
	elapsed := time.Since(start)
	logger.Info("Kafka Write", zap.Duration("duration", elapsed))

	if err != nil {
		logger.Error("Error Writing msg to Kafka ", zap.Error(err))
	}

}

func (kf *KafkaWriter) closeWriter() {
	logger.Info("Closing writer")
	if err := kf.writer.Close(); err != nil {
		logger.Error("Error in Reading msg from Kafka ", zap.Error(err))
	}
}

func (kf *KafkaReaders) closeReaders() {
	for i := range kf.Readers {
		logger.Info("Closing reader : ", zap.Int("nos.", i))
		if err := kf.Readers[i].Close(); err != nil {
			logger.Error("Failed to close kafka reader", zap.Error(err))
		}
	}

}

func (kf *KafkaReaders) setupCloseReaderHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		logger.Info("\r- Ctrl+C pressed in Terminal")
		kf.closeReaders()
		os.Exit(0)
	}()
}

func (kf *KafkaWriter) setupCloseWriterHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		logger.Info("\r- Ctrl+C pressed in Terminal")
		kf.closeWriter()
		os.Exit(0)
	}()
}
