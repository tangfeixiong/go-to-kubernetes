package operator

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/go-redis/redis"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	//appsv1beta1 "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	//appsv1beta2 "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
)

type RedisRole string

const (
	RedisHighAvailabilityRedis    RedisRole = "redis"
	RedisHighAvailabilitySentinel RedisRole = "sentinel"
)

type RedisBootstrapConfig struct {
	Kubeconfig string
	Kind       string
	Name       string
	Namespace  string
	Role       RedisRole
	Port       int
	Dir        string
	LogLevel   int
}

type RedisReplicationInfo struct {
	Role       string
	MasterHost string
	MasterPort int
}

type RedisInfo struct {
	RedisReplicationInfo
}

func ConfigConf(bc *RedisBootstrapConfig) {
	fmt.Printf("with bootstrap config: %v", bc)
	clientset, err := K8sClientset(bc.Kubeconfig)
	if err != nil {
		fmt.Println("Get kubernetes clientset failed:", err.Error())
		os.Exit(1)
	}
	switch bc.Role {
	case RedisHighAvailabilityRedis:
		bc.generateHighAvailabilityRedisConf(clientset)
		break
	case RedisHighAvailabilitySentinel:
		bc.generateHighAvailabilitySentinelConf(clientset)
		break
	default:
		fmt.Println("Unexpected")
	}
}

func (bc *RedisBootstrapConfig) generateHighAvailabilityRedisConf(clientset kubernetes.Interface) {
	var ns string = apiv1.NamespaceAll
	if bc.Namespace == "" {
		ns = apiv1.NamespaceDefault
	} else {
		ns = bc.Namespace
	}
	stsClient := clientset.AppsV1beta1().StatefulSets(ns)

	name := "my-redis"
	if bc.Name != "" {
		name = bc.Name
	}
	result, err := stsClient.Get(name, metav1.GetOptions{})
	if err != nil {
		fmt.Printf("Get failed: %+v\n", err)
		os.Exit(10)
	}
	if result == nil {
		fmt.Println("Get nothing")
		os.Exit(11)
	}

	selector := result.Spec.Selector
	if selector == nil {
		// If empty, defaulted to labels on the pod template.
		selector = &metav1.LabelSelector{
			MatchLabels: result.Spec.Template.Labels,
		}
	}

	podClient := clientset.CoreV1().Pods(ns)

	list, err := podClient.List(metav1.ListOptions{
		LabelSelector: metav1.FormatLabelSelector(selector),
	})
	if err != nil {
		fmt.Println("Failed to list PODs:", err.Error())
		os.Exit(12)
	}

	path := filepath.Join(bc.Dir, "redis.conf")
	// appendonly yes
	// protected-mode no
	// slaveof 127.0.0.1 6379
	mconf := "appendonly yes\nprotected-mode no"
	sf := "%s\nslaveof %s %d\n"
	if len(list.Items) == 0 {
		f, err := os.Create(path)
		if err != nil {
			fmt.Println("Failed to create redis.conf:", err.Error())
			os.Exit(13)
		}
		defer f.Close()

		if _, err := f.WriteString(mconf); err != nil {
			fmt.Println("Failed to wirte redis.conf:", err.Error())
			os.Exit(14)
		}
		fmt.Println("Master of redis.conf is written")
		return
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Failed to read hostname:", err.Error())
		os.Exit(15)
	}
	for _, pod := range list.Items {
		if pod.Name != hostname {
			redisClient := redis.NewClient(&redis.Options{
				Addr:     fmt.Sprintf("%s:%d", pod.Status.PodIP, 6379),
				Password: "", // no password set
				DB:       0,  // use default DB
			})
			defer redisClient.Close()
			pong, err := redisClient.Ping().Result()
			fmt.Println(pong, err)
			if err != nil {
				fmt.Println("Redis is not existed")
				continue
			}

			info := redisClient.Info("replication")
			if info.Err() != nil {
				fmt.Println("Redis client could not work:", info.Err().Error())
				continue
			}
			content := strings.Split(info.Val(), "\r\n")
			var redisinfo RedisInfo
			for _, line := range content {
				if line == "role:master" {
					redisinfo.Role = "master"
					continue
				}
				if line == "role:slave" {
					redisinfo.Role = "slave"
					continue
				}
				if line == "role:sentinel" {
					redisinfo.Role = "sentinel"
					continue
				}
				if strings.HasPrefix(line, "master_host:") {
					redisinfo.MasterHost = line[12:]
					continue
				}
				if strings.HasPrefix(line, "master_port:") {
					redisinfo.MasterPort, _ = strconv.Atoi(line[12:])
					if redisinfo.MasterPort == 0 {
						redisinfo.MasterPort = 6379
					}
					continue
				}
			}
			switch redisinfo.Role {
			case "master":
				f, err := os.Create(path)
				if err != nil {
					fmt.Println("Failed to create redis.conf:", err.Error())
					os.Exit(13)
				}
				defer f.Close()

				conf := fmt.Sprintf(sf, mconf, pod.Status.PodIP, 6379)
				if _, err := f.WriteString(conf); err != nil {
					fmt.Println("Failed to wirte redis.conf for slave:", err.Error())
					os.Exit(14)
				}
				fmt.Println("Slave of redis.conf is written")
				return
			case "slave":
				if redisinfo.MasterHost == "" || redisinfo.MasterPort == 0 {
					fmt.Println("Unexpected information from master host or port")
					continue
				}
				f, err := os.Create(path)
				if err != nil {
					fmt.Println("Failed to create redis.conf:", err.Error())
					os.Exit(13)
				}
				defer f.Close()

				conf := fmt.Sprintf(sf, mconf, redisinfo.MasterHost, redisinfo.MasterPort)
				if _, err := f.WriteString(conf); err != nil {
					fmt.Println("Failed to wirte redis.conf for slave:", err.Error())
					os.Exit(14)
				}
				fmt.Println("Slave of redis.conf is written")
				return
			default:
				fmt.Println("Skip to write redis.conf for role", redisinfo.Role)
				continue
			}
		}
	}
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("Failed to create redis.conf:", err.Error())
		os.Exit(13)
	}
	defer f.Close()

	if _, err := f.WriteString(mconf); err != nil {
		fmt.Println("Failed to wirte redis.conf:", err.Error())
		os.Exit(14)
	}
	fmt.Println("Master of redis.conf is written")
}

func (bc *RedisBootstrapConfig) generateHighAvailabilitySentinelConf(clientset kubernetes.Interface) {
	var ns string = apiv1.NamespaceAll
	if bc.Namespace == "" {
		ns = apiv1.NamespaceDefault
	} else {
		ns = bc.Namespace
	}
	stsClient := clientset.AppsV1beta1().StatefulSets(ns)

	name := "my-redis"
	if bc.Name != "" {
		name = bc.Name
	}
	result, err := stsClient.Get(name, metav1.GetOptions{})
	if err != nil {
		fmt.Printf("Get failed: %+v\n", err)
		os.Exit(10)
	}
	if result == nil {
		fmt.Println("Get nothing")
		os.Exit(11)
	}

	selector := result.Spec.Selector
	if selector == nil {
		// If empty, defaulted to labels on the pod template.
		selector = &metav1.LabelSelector{
			MatchLabels: result.Spec.Template.Labels,
		}
	}

	podClient := clientset.CoreV1().Pods(ns)

	list, err := podClient.List(metav1.ListOptions{
		LabelSelector: metav1.FormatLabelSelector(selector),
	})
	if err != nil {
		fmt.Println("Failed to list PODs:", err.Error())
		os.Exit(12)
	}

	path := filepath.Join(bc.Dir, "sentinel.conf")
	//sentinel monitor mymaster 127.0.0.1 6379 2
	//sentinel down-after-milliseconds mymaster 60000
	//sentinel failover-timeout mymaster 180000
	//sentinel parallel-syncs mymaster 1
	conf := "sentinel monitor %s %s %d %d\nsentinel down-after-milliseconds %s 60000\nsentinel failover-timeout %s 180000\nsentinel parallel-syncs %s 1"

	for _, pod := range list.Items {
		redisClient := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", pod.Status.PodIP, 6379),
			Password: "", // no password set
			DB:       0,  // use default DB
		})
		defer redisClient.Close()
		pong, err := redisClient.Ping().Result()
		fmt.Println(pong, err)
		if err != nil {
			fmt.Println("Redis is not existed")
			continue
		}

		info := redisClient.Info("replication")
		if info.Err() != nil {
			fmt.Println("Redis client could not work:", info.Err().Error())
			continue
		}
		content := strings.Split(info.Val(), "\r\n")
		var redisinfo RedisInfo
		for _, line := range content {
			if line == "role:master" {
				redisinfo.Role = "master"
				continue
			}
			if line == "role:slave" {
				redisinfo.Role = "slave"
				continue
			}
			if line == "role:sentinel" {
				redisinfo.Role = "sentinel"
				continue
			}
			if strings.HasPrefix(line, "master_host:") {
				redisinfo.MasterHost = line[12:]
				continue
			}
			if strings.HasPrefix(line, "master_port:") {
				redisinfo.MasterPort, _ = strconv.Atoi(line[12:])
				if redisinfo.MasterPort == 0 {
					redisinfo.MasterPort = 6379
				}
				continue
			}
		}
		switch redisinfo.Role {
		case "master":
			f, err := os.Create(path)
			if err != nil {
				fmt.Println("Failed to create redis.conf:", err.Error())
				os.Exit(13)
			}
			defer f.Close()

			conf1 := fmt.Sprintf(conf, name, pod.Status.PodIP, 6379, 2, name, name, name)
			if _, err := f.WriteString(conf1); err != nil {
				fmt.Println("Failed to wirte sentinel.conf:", err.Error())
				os.Exit(14)
			}
			fmt.Println("sentinel.conf is written")
			return
		case "slave":
			if redisinfo.MasterHost == "" || redisinfo.MasterPort == 0 {
				fmt.Println("Unexpected information from master host or port")
				continue
			}
			f, err := os.Create(path)
			if err != nil {
				fmt.Println("Failed to create redis.conf:", err.Error())
				os.Exit(13)
			}
			defer f.Close()

			conf1 := fmt.Sprintf(conf, name, redisinfo.MasterHost, redisinfo.MasterPort, 2, name, name, name)
			if _, err := f.WriteString(conf1); err != nil {
				fmt.Println("Failed to wirte redis.conf:", err.Error())
				os.Exit(14)
			}
			fmt.Println("sentinel.conf is written")
			return
		default:
			fmt.Println("Skip to write sentinel.conf for role", redisinfo.Role)
			continue
		}
	}

	fmt.Println("Master not found")
	os.Exit(15)
}
