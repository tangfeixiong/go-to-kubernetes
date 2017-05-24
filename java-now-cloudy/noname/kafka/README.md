

### develop

vagrant@vagrant-ubuntu-trusty-64:/work/src/github.com/tangfeixiong/go-to-kubernetes/java-devel/noname/kafka$ mvn clean compile exec:java -Dexec.mainClass="com.test.groups.ConsumerGroupExample" -Dexec.args="10.123.248.11:2181 group1 topic1" -e --offline

$ kubectl --namespace=kafka exec -ti kafka-0 -- /opt/kafka/bin/kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 --partitions 1 --topic topic1

## About development

### Use Case

任务：批次数据的处理周期

作业：

1. 解密所有用户的报价文件
2. 解密其它文件
3. 复制到内网

__Phase 1__

Topic: decrption-for-price

Partition: 

Key: 

Value: Database ORM

__Phase 2__

Topic: decrption-for-technology-and-business

Partition:

Key:

Value: Database ORM

__Phase 3__

Topic: sync-from-filesystem-into-database

Partition:

Connect: file-source-connector, database-sink-connector

__Phase 4__

Topic: sync-from-database-into-filesystem

Partition:

Connect: database-source-connector, file-sink-connector

