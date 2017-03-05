One name node and one data node
=================================

NameNode
------------

Format

```
ubuntu@datanode125:/opt/hadoop-3.0.0-alpha2$ cat etc/hadoop/core-site.xml
<configuration>
    <property>
        <name>fs.defaultFS</name>
        <value>hdfs://datanode125:9000</value>
    </property>
</configuration>

ubuntu@datanode125:/opt/hadoop-3.0.0-alpha2$ cat etc/hadoop/hdfs-site.xml
<configuration>
    <property>
        <name>dfs.replication</name>
        <value>1</value>
    </property>
</configuration>

ubuntu@datanode125:/opt/hadoop-3.0.0-alpha2$ ssh-keygen -t rsa -P '' -f ~/.ssh/id_rsa

ubuntu@datanode125:/opt/hadoop-3.0.0-alpha2$ ssh-copy-id ubuntu@10.64.33.126

ubuntu@datanode125:/opt/hadoop-3.0.0-alpha2$ bin/hdfs namenode -format

ubuntu@datanode125:/opt/hadoop-3.0.0-alpha2$ bin/hdfs --daemon start namenode

ubuntu@datanode125:/opt/hadoop-3.0.0-alpha2$ head logs/hadoop-ubuntu-namenode-datanode125.log 
2017-03-05 02:36:16,826 INFO org.apache.hadoop.hdfs.server.namenode.NameNode: STARTUP_MSG: 
/************************************************************
STARTUP_MSG: Starting NameNode
STARTUP_MSG:   user = ubuntu
STARTUP_MSG:   host = datanode125/10.64.33.125
STARTUP_MSG:   args = []
STARTUP_MSG:   version = 3.0.0-alpha2

```

DataNode
------------

Start

```
ubuntu@datanode126:/opt/hadoop-3.0.0-alpha2$ scp ubuntu@datanode125:/opt/hadoop-3.0.0-alpha2/etc/hadoop/core-site.xml etc/hadoop

ubuntu@datanode126:/opt/hadoop-3.0.0-alpha2$ scp ubuntu@datanode125:/opt/hadoop-3.0.0-alpha2/etc/hadoop/hdfs-site.xml etc/hadoop

ubuntu@datanode126:/opt/hadoop-3.0.0-alpha2$ bin/hdfs --daemon start datanode

ubuntu@datanode126:/opt/hadoop-3.0.0-alpha2$ ps -ef | grep java

ubuntu@datanode126:/opt/hadoop-3.0.0-alpha2$ head logs/hadoop-ubuntu-datanode-datanode126.log 
2017-03-05 03:30:18,554 INFO org.apache.hadoop.hdfs.server.datanode.DataNode: STARTUP_MSG: 
/************************************************************
STARTUP_MSG: Starting DataNode
STARTUP_MSG:   user = ubuntu
STARTUP_MSG:   host = datanode126/10.64.33.126
STARTUP_MSG:   args = []
STARTUP_MSG:   version = 3.0.0-alpha2

``` 

Word count

```
ubuntu@datanode125:/opt/hadoop-3.0.0-alpha2$ bin/hdfs dfs -ls /

ubuntu@datanode125:/opt/hadoop-3.0.0-alpha2$ bin/hdfs dfs -mkdir /input

ubuntu@datanode125:/opt/hadoop-3.0.0-alpha2$ bin/hdfs dfs -put etc/hadoop/*.xml /input

ubuntu@datanode125:/opt/hadoop-3.0.0-alpha2$ bin/hadoop jar share/hadoop/mapreduce/hadoop-mapreduce-examples-3.0.0-alpha2.jar grep /input /output 'dfs[a-z.]*'

ubuntu@datanode125:/opt/hadoop-3.0.0-alpha2$ bin/hdfs dfs -ls /output

ubuntu@datanode125:/opt/hadoop-3.0.0-alpha2$ bin/hdfs dfs -cat /output/*

ubuntu@datanode125:/opt/hadoop-3.0.0-alpha2$ tail logs/hadoop-ubuntu-namenode-datanode125.log 
2017-03-05 03:44:20,812 INFO org.apache.hadoop.hdfs.StateChange: BLOCK* allocate blk_1073741834_1010, replicas=10.64.33.126:9866 for /output/_temporary/0/_temporary/attempt_local1444179237_0002_r_000000_0/part-r-00000
2017-03-05 03:44:20,849 INFO org.apache.hadoop.hdfs.StateChange: DIR* completeFile: /output/_temporary/0/_temporary/attempt_local1444179237_0002_r_000000_0/part-r-00000 is closed by DFSClient_NONMAPREDUCE_1121327099_1
2017-03-05 03:44:20,918 INFO org.apache.hadoop.hdfs.StateChange: DIR* completeFile: /output/_SUCCESS is closed by DFSClient_NONMAPREDUCE_1121327099_1

ubuntu@datanode126:/opt/hadoop-3.0.0-alpha2$ tail logs/hadoop-ubuntu-datanode-datanode126.log 
2017-03-05 03:44:22,320 INFO org.apache.hadoop.hdfs.server.datanode.DataNode: Receiving BP-1444699866-10.64.33.125-1488652735000:blk_1073741834_1010 src: /10.64.33.125:58489 dest: /10.64.33.126:9866
2017-03-05 03:44:22,333 INFO org.apache.hadoop.hdfs.server.datanode.DataNode.clienttrace: src: /10.64.33.125:58489, dest: /10.64.33.126:9866, bytes: 35, op: HDFS_WRITE, cliID: DFSClient_NONMAPREDUCE_1121327099_1, offset: 0, srvID: 139ab23c-5247-463c-ada7-9477545c720b, blockid: BP-1444699866-10.64.33.125-1488652735000:blk_1073741834_1010, duration(ns): 7601332
2017-03-05 03:44:22,336 INFO org.apache.hadoop.hdfs.server.datanode.DataNode: PacketResponder: BP-1444699866-10.64.33.125-1488652735000:blk_1073741834_1010, type=LAST_IN_PIPELINE terminating
2017-03-05 03:44:25,212 INFO org.apache.hadoop.hdfs.server.datanode.fsdataset.impl.FsDatasetAsyncDiskService: Scheduling blk_1073741833_1009 replica FinalizedReplica, blk_1073741833_1009, FINALIZED
  getNumBytes()     = 163
  getBytesOnDisk()  = 163
  getVisibleLength()= 163
  getVolume()       = /tmp/hadoop-ubuntu/dfs/data
  getBlockURI()     = file:/tmp/hadoop-ubuntu/dfs/data/current/BP-1444699866-10.64.33.125-1488652735000/current/finalized/subdir0/subdir0/blk_1073741833 for deletion
2017-03-05 03:44:25,224 INFO org.apache.hadoop.hdfs.server.datanode.fsdataset.impl.FsDatasetAsyncDiskService: Deleted BP-1444699866-10.64.33.125-1488652735000 blk_1073741833_1009 URI file:/tmp/hadoop-ubuntu/dfs/data/current/BP-1444699866-10.64.33.125-1488652735000/current/finalized/subdir0/subdir0/blk_1073741833

ubuntu@datanode126:/opt/hadoop-3.0.0-alpha2$ bin/hdfs dfs -ls /output
Found 2 items
-rw-r--r--   1 ubuntu supergroup          0 2017-03-05 03:44 /output/_SUCCESS
-rw-r--r--   1 ubuntu supergroup         35 2017-03-05 03:44 /output/part-r-00000

ubuntu@datanode126:/opt/hadoop-3.0.0-alpha2$ bin/hdfs dfs -cat /output/*
1	dfsadmin
1	dfs.replication
1	dfs

```