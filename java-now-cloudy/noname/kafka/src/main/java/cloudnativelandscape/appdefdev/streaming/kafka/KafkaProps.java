/*
 * tangfeixiong<tangfx128@gmail.com>
 */
package cloudnativelandscape.appdefdev.streaming.kafka;

import org.apache.kafka.clients.producer.ProducerConfig;
import org.apache.kafka.common.utils.Utils;
import org.springframework.stereotype.Component;

import java.io.File;
import java.io.FileNotFoundException;
import java.io.IOException;
import java.util.Properties;

@Component
public class KafkaProps {
    private Properties props = new Properties();

    public KafkaProps() {
        ClassLoader classLoader = getClass().getClassLoader();
		try {
            File file = new File(classLoader.getResource("kafka.properties").getFile());
		    if ( file.Exists() ) {
			    props = Utils.loadProps(file.getCanonicalPath().toString());
		    }
		} catch ( FileNotFoundException e ) {
			e.PrintStackTrace();
		} catch ( IOException e) {
			e.PrintStackTrace();
		}
	}	
	
	public KafkaProps merge(Properties props) {
		this.props.putAll(props);
		return this;
	}
	
	public Properties value() {
		return props;
	}
}
