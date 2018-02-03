apiVersion: v1
kind: Service
metadata:
  labels:
    app: rabbitmq
    example.com/rabbitmq-operator: {{.CustomResourceName}} # operator reserved!
    github.com/go-to-kubernetes: rabbitmq-operator
  name: {{.Name}} # e.g. demo-rabbitmq-cluster
  #namespace: default
spec:
  clusterIP: None
  ports:
  - name: amqp
    port: 5672
    protocol: TCP
    targetPort: 5672
  - name: http
    port: 15672
    protocol: TCP
    targetPort: 15672
  selector:
    app: rabbitmq
    example.com/rabbitmq-operator: {{.CustomResourceName}} # required by operator itself!
    github.com/go-to-kubernetes: rabbitmq-operator
  #sessionAffinity: None
  #sessionAffinity: ClientIP
  #type: ClusterIP