#### Basic MQTT example using Golang

> MQTT stands for `"Message Queuing Telemetry Transport."` It's a lightweight and efficient messaging protocol designed 
> for use in situations where bandwidth is limited or connectivity is unreliable, such as in remote or constrained 
> environments. MQTT is commonly used for Internet of Things (IoT) applications and other scenarios where devices need 
> to communicate with each other or with a central server.

Here are some key features of the `MQTT` protocol:

- `Efficiency:`
MQTT is designed to be lightweight and efficient. It minimizes the overhead of the communication protocol, making it 
suitable for low-bandwidth or high-latency networks.

- `Low Bandwidth and Low Power:`
MQTT's design makes it suitable for IoT (Internet of Things) devices that might have limited resources, battery 
power, or bandwidth.

- `Publish/Subscribe Model:` 
MQTT follows a `publish/subscribe` messaging pattern. Clients can publish messages to specific 
topics, and other clients can subscribe to those topics to receive the messages.

- `Quality of Service (QoS):` 
MQTT supports different levels of message delivery guarantees, known as QoS levels. There are three levels: QoS 0 (at most once), 
QoS 1 (at least once), and QoS 2 (exactly once).

- `Retained Messages:` 
Clients can send retained messages, which are stored by the broker and sent to new subscribers as soon as they 
subscribe to the corresponding topic. This is useful for providing the most recent status information to new clients.

- `Bi-directional Communication:` 
MQTT supports both sending and receiving messages, allowing devices to communicate in both directions.

- `Security:` 
While MQTT itself does not define specific security mechanisms, it can be used over secure connections, such as TLS/SSL, 
to ensure data confidentiality and integrity.

- `Last Will and Testament (LWT):` 
Clients can specify a "last will" message that will be sent by the broker if the client unexpectedly disconnects. 
This can be used to notify others of the client's status.

- `Broker-Based:` 
MQTT communication typically involves a central message broker that acts as an intermediary, facilitating communication between clients.

- `Small Footprint:` 
MQTT is designed to be lightweight and efficient, making it suitable for resource-constrained devices and low-bandwidth networks.


Overall, MQTT is widely adopted in various industries for its simplicity, efficiency, and suitability for IoT applications.
It's important to note that my information is accurate up until September 2021, and there may have been developments or changes 
in the MQTT protocol since then. MQTT is favored for applications where efficient, reliable, and lightweight communication is necessary. 
It's commonly used in IoT, home automation, remote monitoring, and other scenarios where devices need to communicate 
with each other or with central systems.


![mqtt-protocol3](https://github.com/faisal-porag/basic_mqtt_example_with_golang/blob/master/photo_files/MQTT-architecture.jpg)

---

##### Some uses of MQTT

- `Internet of Things (IoT) Applications:`
MQTT is widely used in IoT applications to enable communication between devices, sensors, and central servers. 
It's ideal for scenarios where devices need to transmit data to a central system or receive commands from that system.

- `Home Automation:`
MQTT can be used to build smart home systems, where devices like smart lights, thermostats, and security 
cameras communicate with a central server or each other.

- `Industrial Automation:`
MQTT is employed in industrial environments for machine-to-machine communication, remote monitoring, and control. 
It's used to gather data from sensors and control actuators in industrial processes.

- `Telemetry and Remote Monitoring:`
MQTT is used to collect telemetry data, such as temperature, pressure, and other measurements, from remote locations 
and transmit it to a central monitoring system.

- `Vehicle Tracking and Fleet Management:`
MQTT can be used to track the location of vehicles in real-time and monitor their status, allowing fleet managers to 
optimize routes and maintenance.

- `Financial Services:`
MQTT can be used in financial trading systems to disseminate real-time market data to subscribers.

- `Notification Systems:`
MQTT can power instant notification systems, where messages are delivered in real-time to subscribers, ensuring prompt updates.

- `Healthcare Monitoring:`
MQTT can be applied in healthcare systems for patient monitoring. Medical devices can publish data to MQTT topics, and 
healthcare providers can subscribe to receive real-time patient information.

- `Mobile Applications:`
MQTT can be used in mobile apps to enable real-time updates and messaging between users or devices.


Remember that while MQTT is `versatile`, its suitability depends on factors such as network reliability, message delivery requirements, 
and the nature of the application. Additionally, security considerations are crucial, especially when dealing with sensitive data.
Before implementing `MQTT` in any specific scenario, it's essential to evaluate your application's requirements and MQTT's capabilities 
to ensure it's the right fit for your needs.


---


> To set up an MQTT broker (such as Mosquitto) within a Docker container, you can follow these general steps:

`Eclipse Mosquitto:` 
Eclipse Mosquitto is one of the most well-known and widely used MQTT brokers. It's lightweight and open-source, 
making it a popular choice for various IoT and messaging applications.

```shell
 make docker_up 
```

> Use `broker.emqx.io` for free mqtt broker.

---

> Is clientID is different for client and server ?

Yes, the `clientID` should be unique for each MQTT client that connects to the broker, including both clients and servers.
The `clientID` is used by the broker to identify individual clients and manage their connections. Each client connecting 
to the broker should have a distinct `clientID`.
The client and server should have separate `clientID` values to ensure proper identification and handling of their 
connections by the MQTT broker.


***NOTE*** Both the client and the server should agree on the topic they are using. If you want to change the topic name, 
you need to make sure that the client publishes to the new topic name, and the server subscribes to the same new topic name.

For that, I am using different `clientID` values for my `client` & `server` in the example service.

FOR CLIENT SIDE:
```shell
ClientID      = "client_rider_tracker"
```

FOR SERVER SIDE:
```shell
ClientID      = "server_rider_tracker"
```

---


[Resource](https://aws.amazon.com/what-is/mqtt/#:~:text=MQTT%20is%20a%20standards%2Dbased,constrained%20network%20with%20limited%20bandwidth.)



---
THANKS




























