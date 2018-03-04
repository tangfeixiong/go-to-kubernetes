package io.stackdocker.microservices;

import org.apache.hadoop.fs.FileStatus;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.data.hadoop.fs.FsShell;

import com.google.gson.reflect.TypeToken;

import io.kubernetes.client.ApiClient;
import io.kubernetes.client.ApiException;
import io.kubernetes.client.Configuration;
import io.kubernetes.client.apis.CoreV1Api;
import io.kubernetes.client.models.V1Namespace;
import io.kubernetes.client.models.V1Pod;
import io.kubernetes.client.models.V1PodList;
import io.kubernetes.client.util.Config;
import io.kubernetes.client.util.Watch;

import java.io.File;
import java.io.IOException;
import java.nio.charset.Charset;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.Arrays;
import java.util.Collection;
import java.util.List;

@SpringBootApplication
public class App implements CommandLineRunner {
    @Value("${demohdfsclassic.dfsuri}")
    private String dfsUri;    
    
	@Autowired
    private DiscoveryClient discoveryClient;
    
	@Autowired
	private FsShell shell;

	@Override
	public void run(String... args) {
        Sample.shell = shell;
		
		System.out.println(Sample.listFileStatusRecursivelyWithGlobPath("/") == 0 ? "Found nothing" : "Found!");
        
        System.out.println("Creating dir into HDFS");
        Sample.createDirectoriesIntoDFS("/foo", "/bar");

        try {
            Path local = Utils.createTempFile();
            System.out.println("Copy local files into HDFS");
            Sample.copyFromLocalToDFS(local.toString(), "/foo");
    
            System.out.println("Copy HDFS");
            Sample.copyDFS("/foo/*", "/bar");
    
            System.out.println("Copy HDFS files into local");
            local = Files.createTempDirectory("hdfs");
            Sample.copyFromDfsToLocal("/bar/*", local.toString());
            
            System.out.println("List HDFS again");
            Sample.listFileStatusRecursivelyWithGlobPath("/");
        } catch (Exception ex) {
            ex.printStackTrace();
        }
        
        System.out.println("Removing dir into HDFS");
        Sample.removeDirectoriesIntoDFS("/foo", "/bar");
	}

	public static void main(String[] args) {
		SpringApplication.run(App.class, args);
	}
}

/*
  https://docs.spring.io/spring-hadoop/docs/2.5.0.RELEASE/api/org/springframework/data/hadoop/fs/FsShell.html
*/
class Sample {
    static FsShell shell;
    
    // glob
    static int listFileStatusRecursivelyWithGlobPath(String match) {
        Collection<FileStatus> fss = shell.lsr(match);
		for (FileStatus s : fss) {
			System.out.println("> " + s.getPath());
		}
        return fss.size();
    }
    
    static void createDirectoriesIntoDFS(String... uris) {
        shell.mkdir(uris);
    }
    
    static void copyFromLocalToDFS(String local, String remote) {
        shell.copyFromLocal(local, remote);
    }
    
    static void copyDFS(String src, String dst) {
        shell.cp(src, dst);
    }
    
    static void copyFromDfsToLocal(String remote, String local) {
        shell.copyToLocal(remote, local);
    }
    
    static void removeDirectoriesIntoDFS(String... uris) {
        shell.rmr(uris);
    }
}

class Utils {
    
    static Path createTempFile() throws Exception {
        Path path = Files.createTempFile("hdfs", "txt");

        List<String> lines = Arrays.asList("hello", "world");

        Files.write(path, lines, Charset.forName("UTF-8"));
        return path;
    }

    static File createTempDirectory() throws IOException {
        final File temp;
    
        temp = File.createTempFile("temp", Long.toString(System.nanoTime()));
    
        if(!(temp.delete())) {
            throw new IOException("Could not delete temp file: " + temp.getAbsolutePath());
        }
    
        if(!(temp.mkdir())) {
            throw new IOException("Could not create temp directory: " + temp.getAbsolutePath());
        }
    
        return (temp);
    }
}

/*
  https://github.com/kubernetes-client/java
*/
class K8sUtils {
    public static void listAllPods() throws IOException, ApiException{
        ApiClient client = Config.defaultClient();
        Configuration.setDefaultApiClient(client);

        CoreV1Api api = new CoreV1Api();
        V1PodList list = api.listPodForAllNamespaces(null, null, null, null, null, null);
        for (V1Pod item : list.getItems()) {
            System.out.println(item.getMetadata().getName());
        }
    }

    public static void watchOnNamespace() throws IOException, ApiException{
        ApiClient client = Config.defaultClient();
        Configuration.setDefaultApiClient(client);

        CoreV1Api api = new CoreV1Api();

        Watch<V1Namespace> watch = Watch.createWatch(
                client,
                api.listNamespaceCall(null, null, null, null, null, 5, null, null, Boolean.TRUE, null, null),
                new TypeToken<Watch.Response<V1Namespace>>(){}.getType());

        for (Watch.Response<V1Namespace> item : watch) {
            System.out.printf("%s : %s%n", item.type, item.object.getMetadata().getName());
        }
    }
}