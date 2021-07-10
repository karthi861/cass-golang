# go-Products-restapi-Cassandra

..HI all , This is a basic golang Restful API using one of the NoSQL database called Cassandra.

**###How to install Apache Cassandra on Windows:**

Apache Cassandra requires Java 8 to run on a Windows system. Additionally, the Cassandra command-line shell (cqlsh) is dependent on Python 2.7 to work correctly.

To be able to install Cassandra on Windows, first you need to:

Download and Install Java 8 and set environment variables.( Java SE Development Kit 8u251)
Download and install Python 2.7 and set environment variables. (Go to This PC > Properties.and go to Advanced System Settings now Click the Environment Variables… button.>>Double click on the Path variable.Select New and then Browse.
                                                                Add the Python 2.7 path to the Path system variable)


**##Extract Cassandra tar.gz Folder**
1. Visit the official Apache Cassandra Download page "https://www.apache.org/dyn/closer.lua/cassandra/3.11.10/apache-cassandra-3.11.10-bin.tar.gz" and select the version you would prefer to download. Currently, the latest available version is 3.11.10.
2. Click the suggested Mirror download link to start the download process.
3. Unzip the compressed tar.gz folder using a compression tool such as 7-Zip or WinZip. In this example, the compressed folder was unzipped, and the content placed in the C:Cassandraapache-cassandra-3.11.10 folder.


**###Configure Environment Variables**
1. Go to This PC > Properties. and go to Advanced System Settings now Click the Environment Variables… button.
2. Select the New option.>>>CASSANDRA_HOME for Variable name>>>for theVariable value column select the location of the unzipped Apache Cassandra folder.
3.Double click on the Path variable.Select New and then Browse. In this instance, you need to add the full path to the bin folder located within the Apache Cassandra folder,( C:Cassandraapache-cassandra-3.11.6bin.)
**#Follow same steps for configuring JAVA_HOME**


**####How to start Cassandra server**
1.Open cmd and run Cassandra 
   ....the cassandra server starts....
2.Open another cmd and run cqlsh





