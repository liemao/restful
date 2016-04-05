package service

import (
	"github.com/Shopify/sarama"
    "github.com/wvanbergen/kafka/consumergroup"
    log "github.com/thinkboy/log4go"
    "encoding/json"
    "fmt"
    "time"
)

const (
    KafkaPushsTopic = "kafka:chat:topic"
    KAFKA_GROUP_NAME                   = "kafka_topic_push_group"
	OFFSETS_PROCESSING_TIMEOUT_SECONDS = 10 * time.Second
	OFFSETS_COMMIT_INTERVAL            = 10 * time.Second
)

var (
	producer sarama.AsyncProducer
)

type kafkaMsg struct {
    Type string
    Msg string
}

func InitKafka() (err error) {
    kafkaAddrs := []string{"172.16.9.157:9002"}
    config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.NoResponse
	config.Producer.Partitioner = sarama.NewHashPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	producer, err = sarama.NewAsyncProducer(kafkaAddrs, config)
    go HandleSuccess()
	go HandleError()
    //go pullMessagea()
    go pullMessageb()
    go pullMessage()
    go pushMessage()
    return 
}

func HandleSuccess() {
    var (
		pm *sarama.ProducerMessage
	)
	for {
		pm = <-producer.Successes()
		if pm != nil {
			log.Info("producer message success, partition:%d offset:%d key:%s valus:%s", pm.Partition, pm.Offset, pm.Key, pm.Value)
		}
	}
}

func HandleError() {
    var (
		err *sarama.ProducerError
	)
	for {
		err = <-producer.Errors()
		if err != nil {
			log.Error("producer message error, partition:%d offset:%d key:%s valus:%s error(%v)", err.Msg.Partition, err.Msg.Offset, err.Msg.Key, err.Msg.Value, err.Err)
		}
	}
}

func pushMessage() (err error){
    
    var (
		vBytes []byte
		v      = &kafkaMsg{Type: "msg", Msg: "Hello world"}
	)
	if vBytes, err = json.Marshal(v); err != nil {
		return
	}
    time.Sleep(time.Millisecond * 1000)
    for i := 0; i < 10; i++ {
        producer.Input() <- &sarama.ProducerMessage{Topic: "chat", Value: sarama.ByteEncoder(vBytes)}
        /*
        partition, offset, err := producer.SendMessage(&sarama.ProducerMessage{Topic: "test", Value: sarama.ByteEncoder(vBytes)})
        if err != nil {
            log.Info("Failed to produce message to kafka cluster.")
        }

        log.Info("Produced message to partition %d with offset %d", partition, offset)
        */
    }
	return
}

func pullMessage() {

    config := consumergroup.NewConfig()
	config.Offsets.Initial = sarama.OffsetNewest
	config.Offsets.ProcessingTimeout = OFFSETS_PROCESSING_TIMEOUT_SECONDS
	config.Offsets.CommitInterval = OFFSETS_COMMIT_INTERVAL
	config.Zookeeper.Chroot = ""
	kafkaTopics := []string{"chat"}
	cg, err := consumergroup.JoinConsumerGroup(KAFKA_GROUP_NAME, kafkaTopics, []string{"172.16.9.157:2185"}, config)
    if err != nil {
        log.Info("zookeeper error:", err)
    }

    for msg := range cg.Messages() {
        log.Info("deal with topic:%s, partitionId:%d, Offset:%d, Key:%s msg:%s", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
        cg.CommitUpto(msg)
    }
    
    

    /*
    consumer, err := sarama.NewConsumer([]string{"172.16.9.157:9092"}, nil)
    defer consumer.Close()
    if err != nil {
        panic(err)
    }

    partitionConsumer, err := consumer.ConsumePartition("test", 0, sarama.OffsetNewest)
    defer partitionConsumer.Close()
    if err != nil {
        panic(err)
    }
    consumed := 0
    for {
        select {
            case msg := <-partitionConsumer.Messages():
                log.Info("Consumed message offset %d,  partitionId:%d, Message: %s\n", msg.Offset, msg.Partition, string(msg.Value))
                consumed++
                log.Info("Consumed: %d\n", consumed)
            default:
                
        }
    }
    */
}

func pullMessageb() {

    config := consumergroup.NewConfig()
	config.Offsets.Initial = sarama.OffsetNewest
	config.Offsets.ProcessingTimeout = OFFSETS_PROCESSING_TIMEOUT_SECONDS
	config.Offsets.CommitInterval = OFFSETS_COMMIT_INTERVAL
	config.Zookeeper.Chroot = ""
	kafkaTopics := []string{"chat"}
	cg, err := consumergroup.JoinConsumerGroup(KAFKA_GROUP_NAME, kafkaTopics, []string{"172.16.9.157:2185"}, config)
    if err != nil {
        log.Info("zookeeper error:", err)
    }

    for msg := range cg.Messages() {
        log.Info("111deal with topic:%s, partitionId:%d, Offset:%d, Key:%s msg:%s", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
        cg.CommitUpto(msg)
    }
    
    

    /*
    consumer, err := sarama.NewConsumer([]string{"172.16.9.157:9092"}, nil)
    defer consumer.Close()
    if err != nil {
        panic(err)
    }

    partitionConsumer, err := consumer.ConsumePartition("test", 0, sarama.OffsetNewest)
    defer partitionConsumer.Close()
    if err != nil {
        panic(err)
    }
    consumed := 0
    for {
        select {
            case msg := <-partitionConsumer.Messages():
                log.Info("Consumed message offset %d,  partitionId:%d, Message: %s\n", msg.Offset, msg.Partition, string(msg.Value))
                consumed++
                log.Info("Consumed: %d\n", consumed)
            default:
                
        }
    }
    */
}

func pullMessagea() {
    consumer, err := sarama.NewConsumer([]string{"172.16.9.157:9001"}, nil)
    defer consumer.Close()
    if err != nil {
        panic(err)
    }
    
    topics, _ := consumer.Partitions("test")
    for _, topic := range topics {
        fmt.Println(topic)
    }
    partitionConsumer, err := consumer.ConsumePartition("test", 0, sarama.OffsetNewest)
    defer partitionConsumer.Close()
    if err != nil {
        panic(err)
    }
    //consumed := 0
    for {
        select {
            case msg := <-partitionConsumer.Messages():
                //log.Info("Consumed message offset %d,  partitionId:%d, Message: %s\n", msg.Offset, msg.Partition, string(msg.Value))  
                //consumed++
                log.Info("Consumed: %s\n", msg)
            default:
                
        }
    }
    
}