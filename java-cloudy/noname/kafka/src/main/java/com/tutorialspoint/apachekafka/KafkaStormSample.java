package com.tutorialspoint.apachekafka;
/*
    Source code:
	    https://www.tutorialspoint.com/apache_kafka/apache_kafka_integration_storm.htm
*/
import backtype.storm.Config;
import backtype.storm.LocalCluster;
import backtype.storm.spout.SchemeAsMultiScheme;
import backtype.storm.topology.TopologyBuilder;

import java.util.ArrayList;
import java.util.List;
import java.util.UUID;

import storm.kafka.ZkHosts;
import storm.kafka.Broker;
import storm.kafka.StaticHosts;
import storm.kafka.BrokerHosts;
import storm.kafka.SpoutConfig;
import storm.kafka.KafkaConfig;
import storm.kafka.KafkaSpout;
import storm.kafka.StringScheme;
import storm.kafka.trident.GlobalPartitionInformation;

public class KafkaStormSample {
   public static void main(String[] args) throws Exception{
      Config config = new Config();
      config.setDebug(true);
      config.put(Config.TOPOLOGY_MAX_SPOUT_PENDING, 1);
      String zkConnString = "localhost:2181";
      String topic = "my-first-topic";
      BrokerHosts hosts = new ZkHosts(zkConnString);
      
      SpoutConfig kafkaSpoutConfig = new SpoutConfig (hosts, topic, "/" + topic,    
         UUID.randomUUID().toString());
      kafkaSpoutConfig.bufferSizeBytes = 1024 * 1024 * 4;
      kafkaSpoutConfig.fetchSizeBytes = 1024 * 1024 * 4;
      kafkaSpoutConfig.forceFromStart = true;
      kafkaSpoutConfig.scheme = new SchemeAsMultiScheme(new StringScheme());

      TopologyBuilder builder = new TopologyBuilder();
      builder.setSpout("kafka-spout", new KafkaSpout(kafkaSpoutConfig));
      builder.setBolt("word-spitter", new SplitBolt()).shuffleGrouping("kafka-spout");
      builder.setBolt("word-counter", new CountBolt()).shuffleGrouping("word-spitter");
         
      LocalCluster cluster = new LocalCluster();
      cluster.submitTopology("KafkaStormSample", config, builder.createTopology());

      Thread.sleep(10000);
      
      cluster.shutdown();
   }
}