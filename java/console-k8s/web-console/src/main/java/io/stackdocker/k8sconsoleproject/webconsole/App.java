/*
 * Copyright 2002-2015 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package /* thymeleafsandbox.stsm */ /* thymeleafsandbox.biglistflux */ io.stackdocker.k8sconsoleproject.webconsole;

import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

import javax.sql.DataSource;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.jdbc.DataSourceBuilder;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.autoconfigure.data.mongo.MongoDataAutoConfiguration;
import org.springframework.boot.autoconfigure.mongo.MongoAutoConfiguration;
import org.springframework.context.ApplicationContext;
import org.springframework.context.ApplicationContextAware;
import org.springframework.context.ApplicationContextException;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;
import org.springframework.core.io.Resource;
import org.springframework.format.FormatterRegistry;
import org.springframework.web.reactive.HandlerMapping;
import org.springframework.web.reactive.config.CorsRegistry;
//import org.springframework.web.reactive.config.EnableWebReactive;
import org.springframework.web.reactive.config.EnableWebFlux;
//import org.springframework.web.reactive.config.WebReactiveConfigurer;
import org.springframework.web.reactive.config.WebFluxConfigurer;
import org.springframework.web.reactive.handler.SimpleUrlHandlerMapping;
import org.springframework.web.reactive.socket.WebSocketHandler;
import org.springframework.web.reactive.socket.server.WebSocketService;
import org.springframework.web.reactive.socket.server.support.HandshakeWebSocketService;
import org.springframework.web.reactive.socket.server.support.WebSocketHandlerAdapter;
import org.springframework.web.reactive.socket.server.upgrade.ReactorNettyRequestUpgradeStrategy;

import io.stackdocker.k8sconsoleproject.webconsole.business.entities.repositories.*;
import io.stackdocker.k8sconsoleproject.webconsole.business.services.*;
import io.stackdocker.k8sconsoleproject.webconsole.web.conversion.*;
import io.stackdocker.k8sconsoleproject.webconsole.xyz.*;

/**
 * @author FeiXiong Tang
 */
@SpringBootApplication
@Configuration
/* @EnableWebReactive */ //@EnableWebFlux
//@ComponentScan
@EnableAutoConfiguration(exclude={MongoAutoConfiguration.class, MongoDataAutoConfiguration.class})
public class App implements /* WebReactiveConfigurer */ WebFluxConfigurer, ApplicationContextAware {

    private static final Logger logger = LoggerFactory.getLogger(App.class);

    private ApplicationContext applicationContext;

    @Override
    public void setApplicationContext(final ApplicationContext applicationContext) {
        this.applicationContext = applicationContext;
    }

    @Bean
    public DataSource dataSource() {
        try {

            final Resource dbResource = this.applicationContext.getResource("classpath:data/chinook.sqlite");
            logger.debug("Database path: " + dbResource.getURL().getPath());

            final DataSourceBuilder dataSourceBuilder = DataSourceBuilder.create();
            dataSourceBuilder.driverClassName("org.sqlite.JDBC");
            dataSourceBuilder.url("jdbc:sqlite:" + dbResource.getURL().getPath());
            return dataSourceBuilder.build();

        } catch (final IOException e) {
            throw new ApplicationContextException("Error initializing database", e);
        }
    }

	public static void main(String[] args) throws Exception {
		SpringApplication.run(App.class, args);
	}

//    @Bean
//    public VarietyRepository varietyRepository() {
//        return new VarietyRepository();
//    }

//    @Bean
//    public VarietyService varietyService() {
//        return new VarietyService();
//    }

//    @Bean
//    public DateFormatter dateFormatter() {
//        return new DateFormatter();
//    }

//    @Bean
//    public VarietyFormatter varietyFormatter() {
//        return new VarietyFormatter();
//    }

//    @Override
//    public void addFormatters(FormatterRegistry registry) {
//        registry.addFormatter(dateFormatter());
//        registry.addFormatter(varietyFormatter());
//    }

//    @Bean
//    public SeedStarterRepository seedStarterRepository() {
//        return new SeedStarterRepository();
//    }

//    @Bean
//    public SeedStarterService seedStarterService() {
//        return new SeedStarterService();
//    }

//	@Override
//	public void addCorsMappings(CorsRegistry registry) {
//		registry.addMapping("/cors/config");
//	}

//	@Bean
//	public HandlerMapping handlerMapping() {

//		Map<String, WebSocketHandler> map = new HashMap<>();
//		map.put("/websocket/echo", new EchoWebSocketHandler());

//		SimpleUrlHandlerMapping mapping = new SimpleUrlHandlerMapping();
//		mapping.setUrlMap(map);
//		return mapping;
//	}

//	@Bean
//	public WebSocketHandlerAdapter handlerAdapter() {
//		return new WebSocketHandlerAdapter(webSocketService());
//	}

//	@Bean
//	public WebSocketService webSocketService() {
//		return new HandshakeWebSocketService(new ReactorNettyRequestUpgradeStrategy());
//	}

}
