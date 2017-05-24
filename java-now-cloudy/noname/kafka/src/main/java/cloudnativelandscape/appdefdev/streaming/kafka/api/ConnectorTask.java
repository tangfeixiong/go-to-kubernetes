
package cloudnativecomputing.kafkastreaming.connectormanagement.api;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

@JsonIgnoreProperties(ignoreUnknown = true)
public class ConnectorTask {

    private String name;
    private int task;

    public ConnectorTask() { }

    public String getName() {
        return name;
    }

    public void setName(String Name) {
        this.name = name;
    }

    public int getTask() {
        return task;
    }

    public void setTask(int task) {
        this.task = task;
    }

    @Override
    public String toString() {
        return "ConnectorTask{" +
                "name='" + name + '\'' +
                ", task=" + task +
                '}';
    }
}