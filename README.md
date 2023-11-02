# instantbus
A standalone service bus over network written in golang

## Service Bus component
- Allows services to subsribe to a topic .
- Receives messages in byte array over a topic (serialized), pushes it to all subscribers for that topic
- Subscribers need to be aware of de-serialization method.

## Producer/Publisher
Produces a message for a topic

## Consumer/Subsriber
Subscribes to a topic over the bus via n/w and get messages as they arrive.
