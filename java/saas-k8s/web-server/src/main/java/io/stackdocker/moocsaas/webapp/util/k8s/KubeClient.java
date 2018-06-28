package io.stackdocker.moocsaas.webapp.util.k8s;

import java.io.IOException;
import java.util.Map;

import com.google.common.collect.ImmutableList;
import io.kubernetes.client.ApiClient;
import io.kubernetes.client.ApiException;
import io.kubernetes.client.Configuration;
import io.kubernetes.client.apis.AppsV1Api;
import io.kubernetes.client.apis.CoreV1Api;
import io.kubernetes.client.models.*;
import io.kubernetes.client.util.Config;

public class KubeClient {
    private final ApiClient client;

    public KubeClient() throws IOException {
        ApiClient client = Config.defaultClient();
        Configuration.setDefaultApiClient(client);
        this.client = client;
    }

    public V1PodList listPods() throws ApiException {

        CoreV1Api api = new CoreV1Api();
        V1PodList list =
                api.listPodForAllNamespaces(null, null, null, null, null, null, null, null, null);
        for (V1Pod item : list.getItems()) {
            System.out.println(item.getMetadata().getName());
        }
        return list;
    }

    private V1Service createService(String ns, V1Service svc) throws  ApiException {
        CoreV1Api api = new CoreV1Api();
        V1Service created = api.createNamespacedService(ns, svc, null);
        return created;
    }

    private V1Deployment createDeployment(String ns, V1Deployment deploy) throws  ApiException {
        AppsV1Api api = new AppsV1Api();
        V1Deployment created = api.createNamespacedDeployment(ns, deploy, null);
        return created;
    }

    private V1StatefulSet createStatefulSet(String ns, V1StatefulSet sts) throws  ApiException {
        AppsV1Api api = new AppsV1Api();
        V1StatefulSet created = api.createNamespacedStatefulSet(ns, sts, null);
        return created;
    }

    public Object createCtf(String ns, String name, Map<String, String> annotations, Map<String, String> labels) throws  BadParamException, ApiException {
        CtfTemplate builder = new CtfTemplate();

        V1Service mariadbSvc = createService(ns, builder.mariadbSvc_SPOF(ns, name, annotations, labels));
        V1StatefulSet mariadbSts = createStatefulSet(ns, builder.mariadbSts_SPOF(ns, name, annotations, labels, 1));

        V1Deployment apiDeploy = createDeployment(ns, builder.apiserverDeploy(ns, name, annotations, labels, 1));
        V1Service apiSvc = createService(ns, builder.apiserverSvc(ns, name, annotations, labels));

        V1Deployment webDeploy = createDeployment(ns, builder.webserverDeploy(ns, name, annotations, labels, 1));
        V1Service webSvc = createService(ns, builder.webserverSvc(ns, name, annotations, labels));

        return ImmutableList.of(mariadbSvc, mariadbSts, apiSvc, apiDeploy, webSvc, webDeploy);
    }
}
