package com.tutorialspoint.apachekafka;
/*
    Source code:
	    https://www.tutorialspoint.com/apache_kafka/apache_kafka_integration_spark.htm
*/
import java.util.HashMap

import org.apache.kafka.clients.producer.{KafkaProducer, ProducerConfig, Produc-erRecord}
import org.apache.spark.SparkConf
import org.apache.spark.streaming._
import org.apache.spark.streaming.kafka._

object KafkaWordCount {
   def main(args: Array[String]) {
      if (args.length < 4) {
         System.err.println("Usage: KafkaWordCount <zkQuorum><group> <topics> <numThreads>")
         System.exit(1)
      }

      val Array(zkQuorum, group, topics, numThreads) = args
      val sparkConf = new SparkConf().setAppName("KafkaWordCount")
      val ssc = new StreamingContext(sparkConf, Seconds(2))
      ssc.checkpoint("checkpoint")

      val topicMap = topics.split(",").map((_, numThreads.toInt)).toMap
      val lines = KafkaUtils.createStream(ssc, zkQuorum, group, topicMap).map(_._2)
      val words = lines.flatMap(_.split(" "))
      val wordCounts = words.map(x => (x, 1L))
         .reduceByKeyAndWindow(_ + _, _ - _, Minutes(10), Seconds(2), 2)
      wordCounts.print()

      ssc.start()
      ssc.awaitTermination()
   }
}