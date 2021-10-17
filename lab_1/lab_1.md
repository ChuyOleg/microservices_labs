1. Run minikube.  
   > minikube start --driver=hyperv --force
   
   (use --force flag in order to skip no-privileges warning, only for my Windows 10 home).


2. Give active user rights to file ".minikube/machines/minikube/id_rsa" manually (only for my Windows 10 home in order go fix --force flag).  
  

3. Use commands below that kubernetes will be able to use local images.  
   > minikube docker-env  
   
   > minikube -p minikube docker-env | Invoke-Expression
4. Build images from Dockerfiles of your services.  
   > docker build -t service1 -f your_path/Dockerfile ./  
   
   > docker build -t service2 -f your_path/Dockerfile ./
5. Apply deployments and services using kubectl.  
   > kubectl apply -f your_path/service1  
   
   > kubectl apply -f your_path/service2
6. Enable special addons for ingress.  
   > minikube addons enable ingress
7. Apply ingress using kubectl.  
   > kubectl apply -f your_path/ingress.yaml
8. Open application.  
   Urls:  
     ${minikube ip}://api/service1  
     ${minikube ip}://WebApp

   > curl $(minikube ip)://api/service1
   
   Or

   > curl $(minikube ip)://WebApp
   
   