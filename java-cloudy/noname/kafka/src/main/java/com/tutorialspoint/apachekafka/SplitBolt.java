package com.tutorialspoint.apachekafka;
/*
    Source hosted:
	    https://www.tutorialspoint.com/apache_kafka/apache_kafka_integration_storm.htm
*/
import java.util.Map;

import backtype.storm.tuple.Tuple;
import backtype.storm.tuple.Fields;
import backtype.storm.tuple.Values;

import backtype.storm.task.OutputCollector;
import backtype.storm.topology.OutputFieldsDeclarer;
import backtype.storm.topology.IRichBolt;
import backtype.storm.task.TopologyContext;

public class SplitBolt implements IRichBolt {
   private OutputCollector collector;
   
   @Override
   public void prepare(Map stormConf, TopologyContext context,
      OutputCollector collector) {
      this.collector = collector;
   }
   
   @Override
   public void execute(Tuple input) {
      String sentence = input.getString(0);
      String[] words = sentence.split(" ");
      
      for(String word: words) {
         word = word.trim();
         
         if(!word.isEmpty()) {
            word = word.toLowerCase();
            collector.emit(new Values(word));
         }
         
      }

      collector.ack(input);
   }
   
   @Override
   public void declareOutputFields(OutputFieldsDeclarer declarer) {
      declarer.declare(new Fields("word"));
   }

   @Override
   public void cleanup() {}
   
   @Override
   public Map<String, Object> getComponentConfiguration() {
      return null;
   }
   
}