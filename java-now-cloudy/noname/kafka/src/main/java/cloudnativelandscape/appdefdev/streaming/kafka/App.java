/*
 * tangfeixiong<tangfx128@gmail.com>
 */
package cloudnativelandscape.appdefdev.streaming.kafka;

import org.springframework.boot.Banner;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.builder.SpringApplicationBuilder;
import org.springframework.context.annotation.Bean;

import static java.lang.System.exit;

@SpringBootApplication
public class App {
	    
	@Bean
	CommandLineRunner init(KafkaOps kafkaOps) {
		/*
		 * https://docs.spring.io/spring-boot/docs/current/reference/htmlsingle/#boot-features-command-line-runner
		 * ... offer a single run method which will be called just before SpringApplication.run(…​) completes.
		 */
		return (args) -> {
            kafkaOps.main(args);
		};
	}
	
    public static void main(String[] args) {
        //SpringApplication.run(Application.class, args);
    	SpringApplicationBuilder builder = new SpringApplicationBuilder(App.class);
    	ConfigurableApplicationContext ctx;
    	
    	Runnable httpRunner = () -> {
			// app = builder.web(false);
            // app.setBannerMode(Banner.Mode.ON);
			// ctx = app.run(args);

    		ctx = builder.web(true).run(args);
    		
    	}

    	new Thread(httpRunner).start();
 
        ctx.application().dispatch();
    			
        ctx.close();

		// System.exit is common for Batch applications since the exit code can be used to drive a workflow
		System.exit(SpringApplication.exit(ctx));
		// System.exit(SpringApplication.exit(SpringApplication.run(SampleBatchApplication.class, args)));
    }

    /*
	 * In Spring Boot, to create a non-web application, implements CommandLineRunner and override the run method, for example :
     * https://www.mkyong.com/spring-boot/spring-boot-non-web-application-example/
     */
    public void run(String... args) throws Exception {

        if (args.length > 0) {
            System.out.println(args);
        } else {
            System.out.println("no command flags");
        }

        exit(0);
    }
	
}