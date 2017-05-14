
package cloudnativecomputing.kafkastreaming.connectormanagement.rest;

import org.springframework.boot.web.client.RestTemplateBuilder;
import org.springframework.context.annotation.Bean;
import org.springframework.web.client.RestTemplate;

import cloudnativecomputing.kafkastreaming.connectormanagement.api.ConnectorConfigure;

@RestController
public class ConnectorManager {
    private static final Logger log = LoggerFactory.getLogger(ConnectorManager.class);

    @Bean
	public RestTemplate restTemplate(RestTemplateBuilder builder) {
		return builder.build();
	}
    
    @Autowired
    private String kafkaRestUrl;
    
    public String getKafkaRestUrl() {
        return this.kafkaRestUrl;
    }
    
    @HttpMethod()
    public void StartSync() {
    
    }
    
    public void setKafkaRestUrl(String kafkaRestUrl) {
        if ( null == kafkaRestUrl || 0 == kafkaRestUrl.trimLeft().length() )
            this.kafkaRestUrl = "http://broker.kafka.svc.cluster.local:8083/connectors" 
        else
            this.kafkaRestUrl = kafkaRestUrl;
    }    
    
    public void getConnectors() {
        RestTemplate restTemplate = new RestTemplate();
        ArrayList<String> connectors = restTemplate.getForObject(kafkaRestUrl, ArrayList<String>.class);
        log.info(connectors.toString());
    }
    
    public void postConnector(String name, Map<String, String> cfg) {
        ConnectorConfigure req = new ConnectorConfigure();
        req.setName(name);
        req.setCfg(cfg);
        postConnector(req);
    }
    
    public void postConnector(ConnectorConfigure req) {
        RestTemplate restTemplate = new RestTemplate();
        ConnectorConfigure resp = restTemplate.postForObject(kafkaRestUrl, req, ConnectorConfigure.class);
        log.info(resp.toString());
        /*HttpHeaders requestHeaders = new HttpHeaders();
        requestHeaders.setContentType(MediaType.MULTIPART_FORM_DATA);
        HttpEntity<MultiValueMap<String, String>> requestEntity = new HttpEntity<MultiValueMap<String, String>>(mvm, requestHeaders);
        ResponseEntity<InstagramResult> response = restTemplate.exchange(url, HttpMethod.POST, requestEntity, InstagramResult.class);
        InstagramResult result = response.getBody();*/
    }
    
    
    public void getConnector(String name) {
        RestTemplate restTemplate = new RestTemplate();
        ConnectorConfigure resp = restTemplate.getForObject(kafkaRestUrl + "/" + name, 
                ConnectorConfigure.class);
        log.info(resp.toString());
    }
    
    public void getConnectorConfig(String name) {
        RestTemplate restTemplate = new RestTemplate();
        HashMap<String, String> cfg = restTemplate.getForObject(kafkaRestUrl + "/" + name + "/config, 
                HashMap<String, String>.class);
        log.info(cfg.toString());
    }
    
    public void putConnectorConfig(String name, Map<String, String> config) {
        RestTemplate restTemplate = new RestTemplate();
        ConnectorConfigure resp = restTemplate.putForObject(kafkaRestUrl + "/" + name + "/config, 
                config, HashMap<String, String>.class);
        log.info(resp.toString());
    }
    
    public void getConnectorStatus(String name) {
        RestTemplate restTemplate = new RestTemplate();
        HashMap<String, String> cfg = restTemplate.getForObject(kafkaRestUrl + "/" + name + "/status, 
                HashMap<String, String>.class);
        log.info(cfg.toString());
    }
    
    public void postConnectorRestart(String name) {
       RestTemplate restTemplate = new RestTemplate();
        HashMap<String, String> cfg = restTemplate.postForObject(kafkaRestUrl + "/" + name + "/restart);
    }
    
    
    public void putConnectorPause(String name) {
       RestTemplate restTemplate = new RestTemplate();
        HashMap<String, String> cfg = restTemplate.putForObject(kafkaRestUrl + "/" + name + "/pause);
    }
    
    public void postConnectorResume(String name) {
       RestTemplate restTemplate = new RestTemplate();
        HashMap<String, String> cfg = restTemplate.postForObject(kafkaRestUrl + "/" + name + "/resume);
    }
    
    public void deleteConnector(String name) {
       RestTemplate restTemplate = new RestTemplate();
        HashMap<String, String> cfg = restTemplate.deleteForObject(kafkaRestUrl + "/" + name);
    }
    
    public void getConnectorTasks(String name) {
       RestTemplate restTemplate = new RestTemplate();
        HashMap<String, String> cfg = restTemplate.getForObject(kafkaRestUrl + "/" + name + "/tasks");
    }
    
    public void getConnectorTaskStatus(String name, int index) {
       RestTemplate restTemplate = new RestTemplate();
        HashMap<String, String> cfg = restTemplate.getForObject(kafkaRestUrl + "/" + name + "/tasks/" + index + "/status");
    }
    
    public void getConnectorTaskRestart(String name, int index) {
       RestTemplate restTemplate = new RestTemplate();
        HashMap<String, String> cfg = restTemplate.getForObject(kafkaRestUrl + "/" + name + "/tasks/" + index + "/restart");
    }
    
    public void getConnectorPlugins() {
       RestTemplate restTemplate = new RestTemplate();
        HashMap<String, String> cfg = restTemplate.getForObject(kafkaRestUrl + "/connector-plugins");
    }
    
    public void putConnectorPluginConfigValidate(String name) {
       RestTemplate restTemplate = new RestTemplate();
        HashMap<String, String> cfg = restTemplate.putForObject(kafkaRestUrl + "/connector-plugins" + name + "/config/validate");
    }

    
}