1. Run minikube.  
   > minikube start --driver=hyperv --force
   
   (use --force flag in order to skip no-privileges warning, only for my Windows 10 home).


2. Give active user rights to file ".minikube/machines/minikube/id_rsa" manually (only for my Windows 10 home in order go fix --force flag).  
  

3. Use commands below that kubernetes will be able to use local images.  
   > minikube docker-env  
   
   > minikube -p minikube docker-env | Invoke-Expression
4. Build images from Dockerfiles of your services.  
   > make create docker images
5. Enable special addons for ingress.  
   > minikube addons enable ingress
6. Apply deployments and services using kubectl.  
   > kubectl apply -f k8s_v1
7. Test performance (average time for 100 requests) of service1 and services2.
   > go run test/main.go

   Broke one pod.
   > curl curl $(minikube ip):api/service2/untested-request
   
   Test services one more time.


8. Apply VirtualService with retry/timeout settings for services using kubectl.
   > kubectl apply -f k8s_v2

   Test services one more time.


9. Apply DestionationRule for disabling broken pod using circuit breaker pattern and kubectl.
   > kubectl apply -f k8s_v3

    Test services one more time.
   
   