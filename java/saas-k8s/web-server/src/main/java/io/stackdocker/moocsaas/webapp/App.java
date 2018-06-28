package io.stackdocker.moocsaas.webapp;

import java.io.IOException;

import io.stackdocker.moocsaas.webapp.entity.ServiceCatalog;
import org.springframework.beans.factory.BeanCreationException;
import org.springframework.beans.factory.BeanCreationNotAllowedException;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;
import org.springframework.core.convert.converter.Converter;

import io.stackdocker.moocsaas.webapp.service.CatalogRepository;
import io.stackdocker.moocsaas.webapp.util.k8s.KubeClient;

import javax.xml.ws.Service;

@SpringBootApplication
public class App {


	@Bean
	public KubeClient kubernetesClient() {
		try {
			return new KubeClient();
		} catch (IOException e) {
			e.printStackTrace();
			throw new BeanCreationException("Failed to establish kubernetes client", e);
		}
	}

	@Bean
	public CatalogRepository catalogRepository() {
		return new CatalogRepository();
	}

	@Bean
	public Converter<String, ServiceCatalog> messageConverter() {
		return new Converter<String, ServiceCatalog>() {
			@Override
			public ServiceCatalog convert(String id) {
				return catalogRepository().findServiceCatalog(Long.valueOf(id));
			}
		};
	}

	public static void main(String[] args) {
		SpringApplication.run(App.class, args);
	}

}