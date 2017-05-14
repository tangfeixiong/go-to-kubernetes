
package cloudnativecomputing.kafkastreaming.connectormanagement;

@SpringBootApplication
public class App {
	
    public static void main(String[] args) {
        //SpringApplication.run(Application.class, args);
    	SpringApplicationBuilder builder = new SpringApplicationBuilder(app.class);
    	ConfigurableApplicationContext ctx;
    	
    	Runnable httpRunner = () -> {

    		ctx = builder.web(true).run(args);
    		
    	}

    	new Thread(httpRunner).start();
    	
    	context = ctx
 
        ctx.application().dispatch();
    			
        context.close();
    }

}