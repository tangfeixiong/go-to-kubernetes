package io.stackdocker.moocsaas.webapp.service;

import java.io.IOException;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.ConcurrentMap;
import java.util.concurrent.atomic.AtomicLong;

import com.google.common.collect.ImmutableList;
import com.google.common.collect.ImmutableMap;
import io.kubernetes.client.ApiException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.BeanCreationException;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.stereotype.Component;

import io.stackdocker.moocsaas.webapp.entity.*;
import io.stackdocker.moocsaas.webapp.util.k8s.BadParamException;
import io.stackdocker.moocsaas.webapp.util.k8s.KubeClient;

@Component
public class InstanceRepository {
    private  final static Logger logger = LoggerFactory.getLogger(InstanceRepository.class);

    private static AtomicLong counter = new AtomicLong();

    @Autowired private KubeClient kubeclient;

    private final ConcurrentMap<Long, ServiceInstance> serviceInstances = new ConcurrentHashMap<>();

    public Iterable<ServiceInstance> findAll() {
        return this.serviceInstances.values();
    }

    public ServiceInstance saveIntoCreation(ServiceInstance si) {

        try {
            kubeclient.createCtf("demo1", "ctf1", ImmutableMap.of(), ImmutableMap.of());
        } catch (BadParamException e) {
            logger.warn("Invalid Kubernetes Resource", e);
            return null;
        } catch (ApiException e) {
            logger.error("Failed to execute Kubernetes API", e);
            return null;
        }

        Long id = si.getId();
        if (id == null) {
            id = counter.incrementAndGet();
            si.setId(id);
        }
        this.serviceInstances.put(id, si);
        return si;
    }

    public ServiceInstance findServiceInstance(Long id) {
        return this.serviceInstances.get(id);
    }

    public void deleteServiceInstance(Long id) {
        this.serviceInstances.remove(id);
    }

}
