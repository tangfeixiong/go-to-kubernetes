package io.stackdocker.moocsaas.webapp.util.k8s;

import com.google.common.collect.ImmutableList;
import com.google.common.collect.ImmutableMap;
import io.kubernetes.client.custom.IntOrString;
import io.kubernetes.client.custom.Quantity;
import io.kubernetes.client.models.*;
import io.kubernetes.client.proto.IntStr;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class CtfTemplate {
    private final MariadbConfigure mariadb = new MariadbConfigure();
    private final ApiserverConfigure apiserver = new ApiserverConfigure();
    private final WebserverConfigure webserver = new WebserverConfigure();

    public V1Service mariadbSvc_SPOF(String ns, String name, Map<String, String> annotations,  Map<String, String> labels) {
        return mariadb.svc_spof(ns, name, annotations, labels);
    }

    public V1StatefulSet mariadbSts_SPOF(String ns, String name, Map<String, String> annotations,  Map<String, String> labels
            , int replicas) throws BadParamException {
        if ( 1 > replicas) {
            throw new BadParamException("Replicas show greater and equal than 1");
        }
        return mariadb.sts_spof(ns, name, annotations, labels, replicas);
    }

    public V1Service apiserverSvc(String ns, String name, Map<String, String> annotations,  Map<String, String> labels) {
        return apiserver.svcTemplate(ns, name, annotations, labels);
    }

    public V1Deployment apiserverDeploy(String ns, String name, Map<String, String> annotations,  Map<String, String> labels
            , int replicas) throws BadParamException {
        if ( 1 > replicas) {
            throw new BadParamException("Replicas show greater and equal than 1");
        }
        return apiserver.deployTemplate(ns, name, annotations, labels, replicas);
    }

    public V1Service webserverSvc(String ns, String name, Map<String, String> annotations,  Map<String, String> labels) {
        return webserver.svcTemplate(ns, name, annotations, labels);
    }

    public V1Deployment webserverDeploy(String ns, String name, Map<String, String> annotations,  Map<String, String> labels
            , int replicas) throws BadParamException {
        if ( 1 > replicas) {
            throw new BadParamException("Replicas show greater and equal than 1");
        }
        return webserver.deployTemplate(ns, name, annotations, labels, replicas);
    }
}


class MariadbConfigure {
    V1ObjectMeta metadata;
    List<V1PersistentVolumeClaim> pvclaims;
    V1StatefulSetSpec spec;
    V1StatefulSet sts;

    public V1Service svc_spof(String ns, String name, Map<String, String> annotations,  Map<String, String> labels) {
        return new V1Service()
                .apiVersion("v1")
                .kind("Service")
                .metadata(new V1ObjectMeta()
                        .namespace(ns)
                        .name(name)
                        .annotations(annotations)
                        .labels(labels))
                .spec(new V1ServiceSpec()
                        .ports(ImmutableList.of(new V1ServicePort()
                                .name("client")
                                .port(3306)
                                .protocol("TCP")
                                .targetPort(new IntOrString(3306))))
                        .selector(ImmutableMap.of("app", "ctf"
                                , "component", "mariadb"
                                , "stackdocker.io/go-to-kubernetes", "mooc-k8s"))
                        .type("ClusterIP"));
    }

    public V1StatefulSet sts_spof(String ns, String name, Map<String, String> annotations,  Map<String, String> labels
            , int replicas) {
        return new V1StatefulSet()
                .apiVersion("apps/v1")
                .kind("StatefulSet")
                .metadata(new V1ObjectMeta()
                        .namespace(ns)
                        .name(name)
                        .annotations(annotations)
                        .labels(labels))
                .spec(new V1StatefulSetSpec()
                        .podManagementPolicy("OrderedReady")
                        .replicas(replicas)
                        .selector(new V1LabelSelector()
                                .matchLabels(ImmutableMap.of("app", "ctf"
                                        , "component", "mariadb"
                                        , "stackdocker.io/go-to-kubernetes", "mooc-k8s")))
                        .serviceName("mariadb")
                        .template(new V1PodTemplateSpec()
                                .metadata(new V1ObjectMeta()
                                        .labels(ImmutableMap.of("app", "ctf"
                                                , "component", "mariadb"
                                                , "stackdocker.io/go-to-kubernetes", "mooc-k8s")))
                                .spec(new V1PodSpec()
                                        .containers(ImmutableList.of( new V1Container()
                                                .args(ImmutableList.of())
                                                .command(ImmutableList.of())
                                                .env(ImmutableList.of())
                                                .image("docker.io/mariadb")
                                                .imagePullPolicy("IfNotPresent")
                                                .name("mariadb")
                                                .ports(ImmutableList.of(new V1ContainerPort()
                                                        .containerPort(3306)
                                                        .name("client")
                                                        .protocol("TCP")))
                                                .resources(new V1ResourceRequirements())
                                                .securityContext(new V1SecurityContext()
                                                        .privileged(true))
                                                .terminationMessagePolicy("File")
                                                .terminationMessagePath("/dev/termination-log")
                                                .volumeMounts(ImmutableList.of(new V1VolumeMount()
                                                        .name("hostpath")
                                                        .mountPath("/var/lib/sql")))))
                                        .dnsPolicy("ClusterFirst")
                                        .initContainers(ImmutableList.of())
                                        .restartPolicy("Always")
                                        .schedulerName("default-scheduler")
                                        .securityContext(new V1PodSecurityContext())
                                        .terminationGracePeriodSeconds(30L)
                                        .volumes(ImmutableList.of(new V1Volume()
                                                .name("mariadb-data")
                                                .persistentVolumeClaim(new V1PersistentVolumeClaimVolumeSource()
                                                        .claimName("mariadb-hostpath-claim"))))))
                        .updateStrategy(new V1StatefulSetUpdateStrategy()
                                .type("OnDelete"))
                        .volumeClaimTemplates(ImmutableList.of(new V1PersistentVolumeClaim()
                                .metadata(new V1ObjectMeta()
                                        .name("hostpath"))
                                .spec(new V1PersistentVolumeClaimSpec()
                                        .accessModes(ImmutableList.of("ReadWriteOnce"))
                                        .storageClassName("mariadb-hostpath")
                                        .resources(new V1ResourceRequirements()
                                                .requests(ImmutableMap.of("storage", Quantity.fromString("100G"))))))));
    }
}

class ApiserverConfigure {

    public V1Service svcTemplate(String ns, String name, Map<String, String> annotations,  Map<String, String> labels) {
        return new V1Service()
                .apiVersion("v1")
                .kind("Service")
                .metadata(new V1ObjectMeta()
                        .namespace(ns)
                        .name(name)
                        .annotations(annotations)
                        .labels(labels))
                .spec(new V1ServiceSpec()
                        .ports(ImmutableList.of(new V1ServicePort()
                                .name("client")
                                .port(8090)
                                .protocol("TCP")
                                .targetPort(new IntOrString(8090))))
                        .selector(ImmutableMap.of("app", "ctf"
                                , "component", "api-server"
                                , "stackdocker.io/go-to-kubernetes", "mooc-k8s"))
                        .type("ClusterIP"));
    }

    public V1Deployment deployTemplate(String ns, String name, Map<String, String> annotations, Map<String, String> labels
            , int replicas) {
        return new V1Deployment()
                .apiVersion("apps/v1")
                .kind("Deployment")
                .metadata(new V1ObjectMeta()
                        .namespace(ns)
                        .name(name)
                        .annotations(annotations)
                        .labels(labels))
                .spec(new V1DeploymentSpec()
                        .replicas(replicas)
                        .selector(new V1LabelSelector()
                                .matchLabels(ImmutableMap.of("app", "ctf"
                                        , "component", "api-server"
                                        , "stackdocker.io/go-to-kubernetes", "mooc-k8s")))
                        .strategy(new V1DeploymentStrategy()
                                .type("RollingUpdate")
                                .rollingUpdate(new V1RollingUpdateDeployment()
                                        .maxSurge(new IntOrString(1)).maxUnavailable(new IntOrString(1))))
                        .template(new V1PodTemplateSpec()
                                .metadata(new V1ObjectMeta()
                                        .labels(ImmutableMap.of("app", "ctf"
                                                , "component", "api-server"
                                                , "stackdocker.io/go-to-kubernetes", "mooc-k8s")))
                                .spec(new V1PodSpec()
                                        .containers(ImmutableList.of( new V1Container()
                                                .args(ImmutableList.of())
                                                .command(ImmutableList.of())
                                                .env(ImmutableList.of())
                                                .image("docker.io/tangfeixiong/ctf-api-server")
                                                .imagePullPolicy("IfNotPresent")
                                                .name("api-server")
                                                .ports(ImmutableList.of(new V1ContainerPort()
                                                        .containerPort(8090)
                                                        .name("client")
                                                        .protocol("TCP")))
                                                .resources(new V1ResourceRequirements())
                                                .securityContext(new V1SecurityContext()
                                                        .privileged(true))
                                                .terminationMessagePolicy("File")
                                                .terminationMessagePath("/dev/termination-log")
                                                .volumeMounts(ImmutableList.of())))
                                        .dnsPolicy("ClusterFirst")
                                        .initContainers(ImmutableList.of())
                                        .restartPolicy("Always")
                                        .schedulerName("default-scheduler")
                                        .securityContext(new V1PodSecurityContext())
                                        .terminationGracePeriodSeconds(30L)
                                        .volumes(ImmutableList.of()))));
    }
}

class WebserverConfigure {

    public V1Service svcTemplate(String ns, String name, Map<String, String> annotations,  Map<String, String> labels) {
        return new V1Service()
                .apiVersion("v1")
                .kind("Service")
                .metadata(new V1ObjectMeta()
                        .namespace(ns)
                        .name(name)
                        .annotations(annotations)
                        .labels(labels))
                .spec(new V1ServiceSpec()
                        .ports(ImmutableList.of(new V1ServicePort()
                                .name("client")
                                .port(8080)
                                .protocol("TCP")
                                .targetPort(new IntOrString(8080))))
                        .selector(ImmutableMap.of("app", "ctf"
                                , "component", "web-server"
                                , "stackdocker.io/go-to-kubernetes", "mooc-k8s"))
                        .type("NodePort"));
    }

    public V1Deployment deployTemplate(String ns, String name, Map<String, String> annotations, Map<String, String> labels
            , int replicas) {
        return new V1Deployment()
                .apiVersion("apps/v1")
                .kind("Deployment")
                .metadata(new V1ObjectMeta()
                        .namespace(ns)
                        .name(name)
                        .annotations(annotations)
                        .labels(labels))
                .spec(new V1DeploymentSpec()
                        .replicas(replicas)
                        .selector(new V1LabelSelector()
                                .matchLabels(ImmutableMap.of("app", "ctf"
                                        , "component", "web-server"
                                        , "stackdocker.io/go-to-kubernetes", "mooc-k8s")))
                        .strategy(new V1DeploymentStrategy()
                                .type("RollingUpdate")
                                .rollingUpdate(new V1RollingUpdateDeployment()
                                        .maxSurge(new IntOrString(1)).maxUnavailable(new IntOrString(1))))
                        .template(new V1PodTemplateSpec()
                                .metadata(new V1ObjectMeta()
                                        .labels(ImmutableMap.of("app", "ctf"
                                                , "component", "web-server"
                                                , "stackdocker.io/go-to-kubernetes", "mooc-k8s")))
                                .spec(new V1PodSpec()
                                        .containers(ImmutableList.of( new V1Container()
                                                .args(ImmutableList.of())
                                                .command(ImmutableList.of())
                                                .env(ImmutableList.of())
                                                .image("docker.io/tangfeixiong/ctf-web-server")
                                                .imagePullPolicy("IfNotPresent")
                                                .name("api-server")
                                                .ports(ImmutableList.of(new V1ContainerPort()
                                                        .containerPort(8090)
                                                        .name("client")
                                                        .protocol("TCP")))
                                                .resources(new V1ResourceRequirements())
                                                .securityContext(new V1SecurityContext()
                                                        .privileged(true))
                                                .terminationMessagePolicy("File")
                                                .terminationMessagePath("/dev/termination-log")
                                                .volumeMounts(ImmutableList.of())))
                                        .dnsPolicy("ClusterFirst")
                                        .initContainers(ImmutableList.of())
                                        .restartPolicy("Always")
                                        .schedulerName("default-scheduler")
                                        .securityContext(new V1PodSecurityContext())
                                        .terminationGracePeriodSeconds(30L)
                                        .volumes(ImmutableList.of()))));
    }

}