## How to build and deploy services to k8s

#### Before start

1. Install `goose` binary and make sure it's in the `$PATH` scope
2. Install docker
3. Authenticate to your docker account
4. Create k8s cluster for deployment
5. Configure `kubectl` to point to a deployment cluster

#### Account Service
1. Run `cd account-srv`
2. Update configuration params in `conf/config.yaml` if necessary
3. Run `make docker-image`
4. Run `docker push account-srv:latest`
5. Run `cd ..` (change directory to the project root)
6. Change port in deployment manifests if necessary in correspondence with `grpc_listen_addr` configuration value (manifests are stored in `.kube/writer-deployment.yml` and `.kube/writer-service.yml`)
7. Run `kubectl apply -f .kube/account-deployment.yml` to apply deployment manifest
8. Run `kubectl apply -f .kube/account-service.yml` to create service
9. Once the load balancer has rolled out, connect to grpc endpoint on port specified in `grpc_listen_addr` configuration value  

#### Writer Service
1. Run `cd writer-srv`
2. Update configuration params in `conf/config.yaml` if necessary
3. Run `make docker-image`
4. Run `docker push writer-srv`
5. Configure `kubectl` to point to a deployment cluster
6. Run `cd ..` (change directory to the project root)
7. Change port in deployment manifests if necessary in correspondence with `grpc_listen_addr` configuration value (manifests are stored in `.kube/writer-deployment.yml` and `.kube/writer-service.yml`)
8. Run `kubectl apply -f .kube/writer-deployment.yml` to apply deployment manifest
9. Run `kubectl apply -f .kube/writer-service.yml` to create service
10. Connect to grpc endpoint on port specified in `grpc_listen_addr` configuration value

#### Coinsrate Service
1. Run `cd coinsrate-srv`
2. Put postgres connection string to `migrationdsn.txt` (create database if necessary)
3. Run `./migrate.sh up` to apply migration
4. Update configuration params in `conf/config.yaml` if necessary
5. Run `make docker-image`
6. Run `docker push coinsrate-srv`
7. Configure `kubectl` to point to a deployment cluster
8. Run `cd ..` (change directory to the project root)
9. Run `kubectl apply -f .kube/coinsrate-deployment.yml` to apply deployment manifest
10. Run `kubectl apply -f .kube/coinsrate-service.yml` to create service
11. Connect to grpc endpoint on port specified in `grpc_listen_addr` configuration value

#### Solana Stake Service
1. Run `cd solana-stake-srv`
2. Put postgres connection string to `migrationdsn.txt` (create database if necessary)
3. Run `./migrate.sh up` to apply migration
4. Update configuration params in `conf/config.yaml` if necessary
5. Run `make docker-image`
6. Run `docker push solana-stake-srv`
7. Configure `kubectl` to point to a deployment cluster
8. Run `cd ..` (change directory to the project root)
9. Run `kubectl apply -f .kube/solana-stake-deployment.yml` to apply deployment manifest
10. Run `kubectl apply -f .kube/solana-stake-service.yml` to create service

#### ETH Listener Service
1. Run `cd listener-srv`
2. Put postgres connection string to `migrationdsn_eth.txt` (create database if necessary)
3. Run `./migrate_eth.sh up` to apply migration
4. Update configuration params in `conf/config-eth.yaml` if necessary
5. Run `make docker-image-eth`
6. Run `docker push eth-listener-srv`
7. Configure `kubectl` to point to a deployment cluster
8. Run `cd ..` (change directory to the project root)
9. Run `kubectl apply -f .kube/eth-listener-deployment.yml` to apply deployment manifest
10. Run `kubectl apply -f .kube/eth-listener-service.yml` to create service

#### SOL Listener Service
1. Run `cd listener-srv`
2. Put postgres connection string to `migrationdsn_sol.txt` (create database if necessary)
3. Database must be different from one used for ETH listener service
4. Run `./migrate_sol.sh up` to apply migration
5. Update configuration params in `conf/config-sol.yaml` if necessary
6. Run `make docker-image-sol`
7. Run `docker push sol-listener-srv`
8. Configure `kubectl` to point to a deployment cluster
9. Run `cd ..` (change directory to the project root)
10. Run `kubectl apply -f .kube/sol-listener-deployment.yml` to apply deployment manifest
11. Run `kubectl apply -f .kube/sol-listener-service.yml` to create service
