

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

