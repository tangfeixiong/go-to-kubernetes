/*
 * tangfeixiong<tangfx128@gmail.com>
 */
package cloudnativelandscape.appdefdev.streaming.kafka;

import kafka.common.MessageReader;
import kafka.producer.NewShinyProducer;

import kafka.tools.ConsoleProducer$;
import kafka.tools.ConsoleProducer.ProducerConfig;
import kafka.tools.ConsoleProducer.LineMessageReader;

import org.springframework.beans.factory.annotation.Autowired;

import org.springframework.stereotype.Component;

import java.util.Properties;

@Component
public class KafkaOps {

    @Autowired
	KafkaProps kafkaProps;
	
	public void main() {
		
		Properties props = ConsoleProducer$.Module$.getNewProducerProps(conf);
		props = kafkaProps.merge(props).value();
	
				NewShinyProducer producer = new ShinyProducer(props);
           if(config.useOldProducer) {
               new OldProducer(ConsoleProducer$.Module$.getOldProducerProps(config))
           } else {
               new NewShinyProducer(ConsoleProducer$.Module$.getNewProducerProps(config))
           }			


	    try {
	        ProducerConfig config = new ProducerConfig(args);
	        val reader = (MessageReader) Class.forName(config.readerClass).newInstance();
	        reader.init(System.in, ConsoleProducer$.Module$.getReaderProps(config))
	
	        val producer =
	          if(config.useOldProducer) {
	            new OldProducer(ConsoleProducer$.Module$.getOldProducerProps(config))
	          } else {
	            new NewShinyProducer(ConsoleProducer$.Module$.getNewProducerProps(config))
	          }
	
	        Runtime.getRuntime.addShutdownHook(new Thread() {
	          override def run() {
	            producer.close()
	          }
	        })
	
	        var message: ProducerRecord[Array[Byte], Array[Byte]] = null
	        do {
	          message = reader.readMessage()
	          if (message != null)
	            producer.send(message.topic, message.key, message.value)
	        } while (message != null)
	    } catch {
	      case e: joptsimple.OptionException =>
	        System.err.println(e.getMessage)
	        Exit.exit(1)
	      case e: Exception =>
	        e.printStackTrace
	        Exit.exit(1)
	    }
	    Exit.exit(0)
	}

	
}