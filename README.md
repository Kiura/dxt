# dxt
###### Template to files using bash commands
---
**WARNING** 
The package evaluates shell commands so be careful with custom inputs, only use if you trust the input
The tool is not tested, beta, and needs to be polished
it can be extended to have config files or templates or can be extended to use flags etc (PRs are welcome, or open an issue if you want some functionality to be added)

examples of usage:

```bash
cat yourfile.dxt | dxt
```
```bash
echo "currend date: $(date)" | dxt
```
image worth 1000 words:
![image](https://imagebin.ca/3pjDkLBAXmgN/dxt.png)

###### Real world example:

1) lets check the template file (optional)

```bash
$ cat test-deployment.dxt
```

```yaml
apiVersion: apps/v1beta1
kind: Deployment
metadata:
  name: test-$(echo $DEP_NAME)-deployment
  namespace: default
  labels:
    app: test-$(echo $DEP_NAME)-deployment
spec:
  replicas: 3
  selector:
    matchLabels:
      app: test-$(echo $DEP_NAME)-deployment
  template:
    metadata:
      labels:
        app: test-$(echo $DEP_NAME)-deployment
    spec:
      containers:
        - name: test-$(echo $DEP_NAME)
          image: test-$(echo $DEP_NAME):latest
          ports:
            - name: $(echo $DEP_NAME)
              containerPort: 3000
      imagePullSecrets:
        - name: testsecret
```
2) loop over an array of strings, set var for each, pipe through dxt, save the result into a file
```bash
deps=( users files hashes ); for dep in "${deps[@]}";
do
export DEP_NAME=$dep; cat test-deployment.dxt | dxt > test-$dep-deployment.yml
done
```

3) output result of one fail to check (optional)
```bash
$ cat test-users-deployment.yml
```
```yaml
apiVersion: apps/v1beta1
kind: Deployment
metadata:
  name: test-users-deployment
  namespace: default
  labels:
    app: test-users-deployment
spec:
  replicas: 3
  selector:
    matchLabels:
      app: test-users-deployment
  template:
    metadata:
      labels:
        app: test-users-deployment
    spec:
      containers:
        - name: test-users
          image: test-users:latest
          ports:
            - name: users
              containerPort: 3000
      imagePullSecrets:
        - name: testsecret
```
