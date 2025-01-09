### KubeNotifyPlus
- Acts as a webhook for Alertmanager to handle and process alerts.
- Sends detailed notifications to specific users or groups responsible for the affected services.
- In case of a pod crash, provide the exact reason for the crash, including pod state details.
- Support notify for ingress and other notify of k8s, also support for common alert, this is capability of using this tools for all webhook in alert manager.
- Demo data and other resources provided in resources folder

### Usage
- Create new `config.json` with template of `config.json.dist`
- After that:
```
go mod tidy
```

- Init database and tables, then update to `config.json`
```
create database kubenotifyplus_db;
use kubenotifyplus_db;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    telegram_id BIGINT NOT NULL UNIQUE
);

CREATE TABLE service_mappings (
    id INT AUTO_INCREMENT PRIMARY KEY,
    service_name VARCHAR(255) NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
```
- Apply manifest in `resources/k8s/pod_inspector_sa.yaml` for creeate ServiceAccount for k8s which have permission to get pods, events and logs...
- Copy rules to prometheus in `resources/prometheus_rules/`
- Update alert manager config with example in `resources/alertmanager/alertmanager.yaml`

### Run:
```
go run main.go
```
Example output:
```
2025/01/09 15:29:32 MySQL connection established
2025/01/09 15:29:32 Server started on port 6789
```

- Test: you can test with Post request via `PostMan` i guess.

### Module
- This steps is for create new module, module already created at this time.
- Need to init the fucking module for the first time, it will creates `god.mod` file
```
go mod init KubeNotifyPlus
```

### K8S Deployment rule:
- Support send notify to users with deployment and ingress.
- The deployment should be matched with container name for notify alert.
- Right now, this only support 1 container per 1 deployment in 1 cluster if using feature pod investigation.
- Pod investigation: get pod logs and pod events.
- Support for send log is not k8s also