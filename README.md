# random

random json doc generator in go

<h5>random.exe: a go program to generate random JSON data and put it into a local(127.0.0.1)  mongodb database called test using a collection called test_data. You can specify on the console how many records to create.
</h5>
For each JSON record there is a Value entry which is randomly set to either a 1 or 0 with a 10%  probability that the value will be a 1. This is achieved using the rand.Float64() function.

```go
var num int
if rand.Float64() < 0.1 {
	num = 1
} else {
	num = 0
}
```


<h5>publisher_test.py: a python script to read all records form the mongodb test_data collection. It then loops through the returned collection and every second it will put one of these onto a rabbitMQ queue called test_queue.
https://github.com/mjdilworth/publisher.git
</h5>

<h5>consumer.exe: a go program which connects to the rabbitMQ test_queue and consumes messages. When it receives a message that has a "Value" attribute which is a "1" it writes out to a text file (/tmp/stratagem_golang_output.txt) the timestamp (placed on the queue) Thu Jan 29 08:48:11 +0000 UTC 2015, the JSON object message ID 54c9f00c14339b3c3ea5052c and the string "Got a 1!". these are delimitated by the pipe symbol.
https://github.com/mjdilworth/consumer.git
</h5>



<h5>Overall notes</h5>

<ul>Installed Go 1.4.1
https://storage.googleapis.com/golang/go1.4.1.windows-amd64.msi
</ul>
<ul>Installed Python 2.7.9
Python 2.7.9 is used as 3.4 does not work with pika
https://www.python.org/ftp/python/2.7.9/python-2.7.9.msi
</ul>
<ul>Installed Erlang
http://www.erlang.org/download/otp_win64_17.4.exe
</ul>
<ul>Installed RabbitMQ
http://www.rabbitmq.com/releases/rabbitmq-server/v3.4.3/rabbitmq-server-3.4.3.exe
</ul>
<ul>Installed MongoDB
https://fastdl.mongodb.org/win32/mongodb-win32-x86_64-2008plus-2.6.7-signed.msi?_ga=1.216185685.235954760.1422360900
</ul>

All installations used default parameters
<p>
For Python we use pika for rabbitMQ  and pymongo for mongodb

these are installed using pip 
```
pip install pika
pip install pymongo
```
<br>
for go you need to install the Go RabbtMQ client
```
go get github.com/streadway/amqp
```

and the mgo driver
```
go get gopkg.in/mgo.v2
```


To export test data from mongodb
```
C:\Program Files\MongoDB 2.6 Standard\bin>mongoexport --db test --collection test_data --csv  -f _id,Value,DT --out /tmp/test_data.csv
connected to: 127.0.0.1
exported 100 records
```
