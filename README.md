# Milestone4
## Kubernetes

### k8s

Installation (all following sub-steps must be performed on each VMs - master, worker1, worker2 ... etc.):

- create new GCP instance with Debian 11

- install Docker

  ```bash
  ~ sudo apt update
  ~ sudo apt install \
      ca-certificates \
      curl \
      gnupg \
      lsb-release

  ~ curl -fsSL https://download.docker.com/linux/debian/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

  ~ echo \
    "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/debian \
    $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

  ~ sudo apt update
  ~ sudo apt install docker-ce docker-ce-cli containerd.io docker-compose-plugin

  # if you have .sock permission errors perform following steps ...
  ~ sudo groupadd docker
  ~ sudo usermod -aG docker ${USER}
  ~ sudo reboot now
  ```

- install k8s

  ```bash
  ~ sudo apt install apt-transport-https

  ~ curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add
  ~ cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
      deb https://apt.kubernetes.io/ kubernetes-xenial main
    EOF

  ~ sudo apt update
  ~ suo apt install -y kubelet kubeadm kubectl
  ```

- additional configuring for k8s

  ```bash
  ~ sudo swapoff -a
  ```

Initialize first k8s cluster (on master node)

  ```bash
  # '--pod-network-cidr' and it's argument is mandatory for this guide due to Flanel CNI !!!
  ~ sudo kubeadm init --pod-network-cidr=10.244.0.0/16
  ```

If everything is OK, in the end you can see string 'kubeadm join <ip:port> --token ...' for joining you worker-node to your master-node. Do not worry, you can retrieve this connect-string any time with next command

  ```bash
  ~ sudo kubeadm token create --print-join-command
  ```

Relocate k8s main config for default user

  ```bash
  # create home folder for cluster
  ~ mkdir -p $HOME/.kube
  
  # copy generated settings to this folder
  ~ sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config

  # change ownership from root to default/current user
  ~ sudo chown $(id -u):$(id -g) $HOME/.kube/config
  ```

Check base info about your cluster

  ```bash
  ~ sudo kubectl get nodes
  ~ sudo kubectl get namespaces
  ~ sudo kubectl get pods --all-namespaces
  ```

**TIP**: if you get error 'The connection to the server localhost:8080 was refused' after previous command(s) - there is:
- problem with *.kube* directory - check ownership, permissions, content of folder etc.  
**OR**
- problem with firewall - if you did some connection/firewall changes on instance.

At this step you must see after ***kubectl get nodes*** info with one node - master node with name of your machine. Node has status 'NotReady' because the cluster does not have pod-network configs / CNI. There are multiple CNI (container network interface):
- Flannel
- Calico
- Canal
- Weave  

... (first two are the most common)
  
To setup pod networking (Flannel in our case)

  ```bash
  ~ kubectl apply -f https://raw.githubusercontent.com/flannel-io/flannel/master/Documentation/kube-flannel.yml
  ```

Now status of master node must be 'Ready'.

**IMPORTANT**: Check if CNI is configured and work corectlly by runing ***kubectl get pods*** -> all pods have to have *Status* 'Runing' !!!  
Otherwise -> something went wrong with CNI !!!  
For Flannel CNI ***kubeadm init*** must be executed with *--pod-network-cidr=10.244.0.0/16* (this argument is default, but you can choose other one) !!!

Connect configured worker-nodes (on worker node)

  ```bash
  # sudo kubeadm token create --print-join-command - if you lost connection string :)
  ~ kubeadm join <ip>:<port> --token <token> --discovery-token-ca-cert-hash sha256:<hash>
  ```

If worker connected succesfully you have to see output like:

  ```md
  This node has joined the cluster:
  * Certificate signing request was sent to apiserver and a response was received.
  * The Kubelet was informed of the new secure connection details.

  Run 'kubectl get nodes' on the control-plane to see this node join the cluster.
  ```

So, run ***kubectl get nodes*** on master node to see new woker node.

**TIP**: new node has role with tag *none* by default. If you want to change this tag use command

  ```bash
  ~ kubectl label node <name> node-role.kubernetes.io/worker=<new-tag>
  ```

To prove correct configuration of k8s try a simple example -> [docker](https://docs.docker.com/get-started/orchestration/)

  - to connect to pods use
    
    ```bash
    ~ kubectl exec -it <name> -- sh
    ```

To create k8s cred to your private Docker registry

- register on Docker Hub

- create your public/private registry

- create access token to this registry with some rules

- on your k8s master-node instance try to login into the registry 

  ```bash
  # after this command you will be asked to prompt your pass/access token
  ~ docker login -u <username>

  # prompt your access token
  xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
  ```

- now you have to see that your creds will be stored in *${HOME}/.docker/config.json* unencrypted

- use this tutorial to create k8s creds with this registry cred -> [kubernetes](https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/)

  - in my case i used .yaml style of creating creds object, because you can specify aditional fields (like *namespace*) and manage this creds like a resource (this is cred_docker.yaml)

    ```yaml
    apiVersion: v1
    kind: Secret
    metadata:
      name: some-name
      namespace: some-space
    data:
      .dockerconfigjson: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx...
    type: kubernetes.io/dockerconfigjson
    ```

  - *.dockerconfigjson* is value generated by command

    ```bash
    ~ cat ~/.docker/config.json | base64
    ```

  - apply this yaml file to generate cred object

    ```bash
    ~ kuectl apply -f cred_docker.yaml
    ```

  - now you can use this cred in other yaml files to pull Docker images from your Docker registry without direct usage of Docker creds or routine passing this creds on each k8s worker-node

  - check this tutorial for more details also -> [blog](https://blog.cloudhelix.io/using-a-private-docker-registry-with-kubernetes-f8d5f6b8f646)
