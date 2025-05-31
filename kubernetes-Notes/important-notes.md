# Important notes
'kubectl expose' to create a service or expose deployment to service

'k get pods -l <field=value>' select specific pods

'kubectl delete all --all' Delete everything without deleting the namespace

cronjob to run something periodically in future

` kubectl top pods --all-namespaces --no-headers | sort -k3 -nr | head -n1 | awk '{print $2","$1}' > high_cpu_pod.txt` Get top consuming pods

`kubectl run sleep-pod --image=nginx --restart=Never --command -- sleep 3600
` Create pod and run some commands. 

`kubectl expose pod nginx-pod --name=nginx-service --port=80 --target-port=80 --type=ClusterIP` Expose pod

`kubectl port-forward service/nginx-service 8080:80` Port forwarding

`k expose deployment/nginx-app-cka --name=app-service-cka --port=80 --target-port=80 --type=NodePort ` node port service.

`kubectl run test-nslookup --rm -i --tty --image=busybox:1.28 --restart=Never -- /bin/sh` run busy box for commands

StorageClass is like a template or blueprint that tells Kubernetes how to create (provision) PersistentVolumes (PVs) — what kind of storage, parameters, performance, etc.

PersistentVolume (PV) is the actual storage resource created based on that StorageClass (or statically created).

PersistentVolumeClaim (PVC) is your request or claim to use a PV that matches your storage needs (size, access mode, StorageClass).
