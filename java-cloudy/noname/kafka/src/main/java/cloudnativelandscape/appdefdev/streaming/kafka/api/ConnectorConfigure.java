
package cloudnativecomputing.kafkastreaming.connectormanagement.api;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

@JsonIgnoreProperties(ignoreUnknown = true)
public class ConnectorConfigure {

    private String name;
    private Map<String, String> config;
    private ConnectorTask[] tasks; 

    public ConnectorConfigure() {
    	cfg = new HashMap<String, String>();
    }

    public String getName() {
        return name;
    }

    public void setName(String Name) {
        this.name = name;
    }

    public Map<String, String> getConfig() {
        return cfg;
    }

    public void setConfig(Map<String, String> config) {
        this.config = config;
    }
    
    public ConnectorTask[] getTasks() {
    	return this.tasks;
    }
    
    public void setTasks(ConnectorTasks tasks) {
    	this.tasks = tasks;
    }

    @Override
    public String toString() {
        return "ConnectorConfigure{" +
                "name='" + name + '\'' +
                ", config=" + config +
                '}';
    }
}