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

import io.stackdocker.moocsaas.webapp.entity.ServiceCatalog;
import io.stackdocker.moocsaas.webapp.util.k8s.BadParamException;
import io.stackdocker.moocsaas.webapp.util.k8s.KubeClient;

@Component
public class CatalogRepository {
    private  final static Logger logger = LoggerFactory.getLogger(CatalogRepository.class);

    private static AtomicLong counter = new AtomicLong();

    private final ConcurrentMap<Long, ServiceCatalog> catalogs = new ConcurrentHashMap<>();

    public Iterable<ServiceCatalog> findAll() {
        return this.catalogs.values();
    }

    public ServiceCatalog save(ServiceCatalog catalog) {
        Long id = catalog.getId();
        if (id == null) {
            id = counter.incrementAndGet();
            catalog.setId(id);
        }
        this.catalogs.put(id, catalog);
        return catalog;
    }

    public ServiceCatalog findServiceCatalog(Long id) {
        return this.catalogs.get(id);
    }

    public void deleteServiceCatalog(Long id) {
        this.catalogs.remove(id);
    }

}
