# eventsourcing-go
Simple bank app using Event Sourcing pattern with Kafka and Redis for studying CQRS pattern (Command Query Responsibility Segregation)

#### Step 1
Install Kafka 1.0.0 and run Zookeeper and Broker(s):

    $ bin/zookeeper-server-start.sh config/zookeeper.properties
    $ bin/kafka-server-start.sh config/server.properties

Then create a first topic with a single partition for testing:
    
    $ bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic example_topic
       
    
#### Step 2
Install govendor for dependency management 

    $ go get -u github.com/kardianos/govendor
    $ go init
    $ go add +external
    
Install ginkgo and gomega for BDD style testing

    $ go get github.com/onsi/ginkgo/ginkgo
    $ go get github.com/onsi/gomega
    $ ginkgo bootstrap
    

#### Step 3
Run test

    $ ginkgo
    
#### Step 4
Start a consumer

    $ ./eventsourcing-go --act=consumer
  
Start a producer in the separate terminal
   
    $ ./eventsourcing-go --act=producer
    # Create a new account in the prompt and deposit $1000
    -> create###Joe Smith 
    -> deposit###ba50259d-8a44-445a-90c7-433ce2e4417e#1000
    
 In a consumer console, the consumed events will be displayed.
 
    Welcome to Example Bank service: consumer
    
    Exiting mainConsumer
    Press [enter] to exit consumer
    
    Waiting a message from the channel
    Processing CreateEvent:
    {"AccID":"ba50259d-8a44-445a-90c7-433ce2e4417e","Type":"CreateEvent","AccName":"Andy Yoo"}
    {ID:ba50259d-8a44-445a-90c7-433ce2e4417e Name:Andy Yoo Balance:0}
    
    Processing DepositEvent:
    {"AccID":"ba50259d-8a44-445a-90c7-433ce2e4417e","Type":"DepositEvent","Amount":1000}
    {ID:ba50259d-8a44-445a-90c7-433ce2e4417e Name:Andy Yoo Balance:1000}
 
 
 #### In future version
    1. Use docker compose
    2. Implement REST API for microservie using event-sourcing
    3. Use Kafka consumer groups.
    4. Add React or Angular frontend 
    5. Add Websocket to make this real time application.