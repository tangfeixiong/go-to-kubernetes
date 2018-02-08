ROM openjdk:8-jre
LABEL maintainer="tangfeixiong <tangfx128@gmail.com>" \
      project="https://github.com/tangfeixiong/go-to-kubernetes" \
      name="hadoop-operator" \
      annotation='{"example.com/hadoop-operator":"hdfs-springboot"}' \
      tag="centos java1.8 openjdk springboot"

ARG jarTgt
ARG javaOpt

COPY ${jarTgt:-hdfs-springboot.jar} /hdfs-springboot.jar

ENV JAVA_OPTIONS="${javaOpt:-'-Xms128m -Xmx512m -XX:PermSize=128m -XX:MaxPermSize=256m'}" \
    SPRING_HADOOP_FSURI="hdfs://demo-hdfs-classic-0.demo-hdfs-classic.default.svc.cluster.local:9000" \
    GO_TO_KUBERNETES="hadoop-operator"

ENTRYPOINT ["java","-Djava.security.egd=file:/dev/./urandom","-jar","/hdfs-springboot.jar"]
