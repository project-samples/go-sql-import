# go-sql-import

## Import 1 million records into SQL from CSV Fixed-Length Format Files with good performance
In this demo, we import 1 Million Records into SQL from CSV Fixed-Length Format Files with high performance in GOLANG.

This article outlines strategies for high-performance batch processing in GOLANG, specifically for importing 1 million records into PostgreSQL from a CSV or fixed-length format file.

![Import into SQL from CSV or fixed-length format files](https://cdn-images-1.medium.com/max/800/1*rYaIdKGSd0HwZqZW7pMEiQ.png)

### Test Info
- RAM: 12 GB
- Disk: SSD KINGSTON SA400S37240G ATA Device
- Exec File Size: 12M
- Database: PosgreSQL 16
- Total of rows: 1.018.584 rows
- Total of columns: 76 columns

<table><thead><tr>
<td>Type</td>
<td>File Size</td>
<td>Rows</td>
<td>CPU</td>
<td>RAM</td>
<td>Disk</td>
<td>Power Usage</td>
<td>Without data validation</td>
<td>With data validation</td>
</tr></thead><tbody>

<tr>
<td>Fix Length</td>
<td>1.15 GB</td>
<td>1,018,584</td>
<td>6.6%</td>
<td>33 M</td>
<td>3.1 M/s</td>
<td>Moderate</td>
<td>5 min 16 sec</td>
<td>6 min 10 sec</td>
</tr>

<tr>
<td>CSV</td>
<td>0.975 GB</td>
<td>1,018,584</td>
<td>8%</td>
<td>34 M</td>
<td>2.8 M/s</td>
<td>Moderate</td>
<td>5 min 12 sec</td>
<td>6 min</td>
</tr>

</tbody></table>

- For fix length format file, please unzip 'fixedlength_big.7z' file (1.15 GB)
- For CSV file, please unzip 'delimiter_big.7z' file (0.975 GB)

### Batch jobs
Differ from online processing:
- Long time running, often at night, after working hours.
- Non-interactive, often include logic for handling errors
- Large volumes of data

### Challenges in Batch Processing
- <b>Data Size</b>: Large volumes can strain memory and I/O performance.
- <b>Database Constraints</b>: SQL (Postgres, My SQL, MS SQL...) may struggle with handling high transaction loads.
### High-Performance Strategy
#### Streaming and Buffered I/O:
- Instead of loading the entire file into memory, stream the file line by line using Go's buffered I/O. This reduces memory consumption.
- Example: Use bufio.Scanner to read records line by line.
#### Bulk Inserts:
Bulk inserts in SQL are typically faster because they reduce the overhead associated with individual inserts. Here’s why:
- <b>Transaction Handling</b>: In a bulk insert, multiple rows are inserted in a single transaction, reducing the need for multiple commits, which can be expensive in terms of I/O operations.
- <b>Logging</b>: Bulk inserts often minimize logging overhead, especially if the database is configured to use minimal logging for bulk operations (like in SQL Server with the "BULK_LOGGED" recovery model).
- <b>Constraints</b>: When inserting data in bulk, constraints such as foreign keys and unique indexes may be deferred or optimized by the database engine.
- <b>Index Updates</b>: Instead of updating indexes after each row insert, bulk inserts allow the database to update indexes in batches, improving performance.

However, it’s important to note that bulk inserts still need to ensure data integrity. Some databases provide options to temporarily disable constraints or logging to optimize performance further, but this can lead to data consistency issues if not handled properly.

#### Batch Inserts:
- Insert records into PostgreSQL in batches (e.g., 1000 records per transaction). This reduces transaction overhead and improves performance.
- In this sample, we use Batch Inserts. We still have a very good performance.
#### Error Handling and Logging:
- Implement robust error handling and logging. Track failed records to reprocess or fix them later.

### Advantages of This Approach
- <b>Efficiency</b>: Streaming and batching minimize memory usage and reduce database transaction overhead.
- <b>Scalability</b>: Parallel processing allows you to scale the import process across multiple cores.
- <b>Flexibility</b>: This approach can handle large datasets and can be adapted for other file formats or databases.
### Disadvantages
- <b>Simple</b>: Do not handle retries or parallel processing.
- <b>Transaction Size</b>: Large batches can still strain the database if not managed properly.

### Conclusion
By carefully handling file I/O, database interactions, and error management, you can ensure high performance when importing large datasets into SQL.

## Import flow
![Import flow with data validation](https://cdn-images-1.medium.com/max/800/1*Y4QUN6QnfmJgaKigcNHbQA.png)

## Common Architectures
### Layer Architecture
- Popular for web development

![Layer Architecture](https://cdn-images-1.medium.com/max/800/1*JDYTlK00yg0IlUjZ9-sp7Q.png)

### Hexagonal Architecture
- Suitable for Import Flow

![Hexagonal Architecture](https://cdn-images-1.medium.com/max/800/1*nMu5_jZJ1omzIB5VK5Lh-w.png)

#### Based on the flow, there are 4 main components (4 main ports):
- Reader, Validator, Transformer, Writer
##### Reader
Reader Adapter Sample: File Reader. We provide 2 file reader adapters:
- Delimiter (CSV format) File Reader
- Fix Length File Reader
##### Validator
- Validator Adapter Sample: Schema Validator
- We provide the Schema validator based on GOLANG Tags
##### Transformer
We provide 2 transformer adapters
- Delimiter Transformer (CSV)
- Fix Length Transformer
##### Writer
We provide many writer adapters:
- SQL:
    - [SQL Writer](https://github.com/core-go/sql/blob/main/writer/writer.go): to insert or update data
    - [SQL Inserter](https://github.com/core-go/sql/blob/main/writer/inserter.go): to insert data
    - [SQL Updater](https://github.com/core-go/sql/blob/main/writer/updater.go): to update data
    - [SQL Stream Writer](https://github.com/core-go/sql/blob/main/writer/stream_writer.go): to insert or update data. When you write data, it keeps the data in the buffer, it does not write data. It just writes data when flush.
    - [SQL Stream Inserter](https://github.com/core-go/sql/blob/main/writer/stream_inserter.go): to insert data. When you write data, it keeps the data in the buffer, it does not write data. It just writes data when flush. Especially, we build 1 single SQL statement to improve the performance.
    - [SQL Stream Updater](https://github.com/core-go/sql/blob/main/writer/stream_updater.go): to update data. When you write data, it keeps the data in the buffer, it does not write data. It just writes data when flush.

- Mongo:
    - [Mongo Writer](https://github.com/core-go/mongo/blob/main/writer/writer.go): to insert or update data
    - [Mongo Inserter](https://github.com/core-go/mongo/blob/main/writer/inserter.go): to insert data
    - [Mongo Updater](https://github.com/core-go/mongo/blob/main/writer/updater.go): to update data
    - [Mongo Stream Writer](https://github.com/core-go/mongo/blob/main/batch/stream_writer.go): to insert or update data. When you write data, it keeps the data in the buffer, it does not write data. It just writes data when flush.
    - [Mongo Stream Inserter](https://github.com/core-go/mongo/blob/main/batch/stream_inserter.go): to insert data. When you write data, it keeps the data in the buffer, it does not write data. It just writes data when flush.
    - [Mongo Stream Updater](https://github.com/core-go/mongo/blob/main/batch/stream_updater.go): to update data. When you write data, it keeps the data in the buffer, it does not write data. It just writes data when flush.

- Elastic Search
    - [Elastic Search Writer](https://github.com/core-go/elasticsearch/blob/main/writer/writer.go): to insert or update data
    - [Elastic Search Creator](https://github.com/core-go/elasticsearch/blob/main/writer/creator.go): to create data
    - [Elastic Search Updater](https://github.com/core-go/elasticsearch/blob/main/writer/updater.go): to update data
    - [Elastic Search Stream Writer](https://github.com/core-go/elasticsearch/blob/main/batch/stream_writer.go): to insert or update data. When you write data, it keeps the data in the buffer, it does not write data. It just writes data when flush.
    - [Elastic Search Stream Creator](https://github.com/core-go/elasticsearch/blob/main/batch/stream_creator.go): to create data. When you write data, it keeps the data in the buffer, it does not write data. It just writes data when flush.
    - [Elastic Search Stream Updater](https://github.com/core-go/elasticsearch/blob/main/batch/stream_updater.go): to update data. When you write data, it keeps the data in the buffer, it does not write data. It just writes data when flush.

- Firestore
    - [Firestore Writer](https://github.com/core-go/firestore/blob/main/writer/writer.go): to insert or update data
    - [Firestore Updater](https://github.com/core-go/firestore/blob/main/writer/updater.go): to update data

- Cassandra
    - [Cassandra Writer](https://github.com/core-go/cassandra/blob/main/writer/writer.go): to insert or update data
    - [Cassandra Inserter](https://github.com/core-go/cassandra/blob/main/writer/inserter.go): to insert data
    - [Cassandra Updater](https://github.com/core-go/cassandra/blob/main/writer/updater.go): to update data

- Hive
    - [Hive Writer](https://github.com/core-go/hive/blob/main/writer/writer.go): to insert or update data
    - [Hive Inserter](https://github.com/core-go/hive/blob/main/writer/inserter.go): to insert data
    - [Hive Updater](https://github.com/core-go/hive/blob/main/writer/updater.go): to update data
    - [Hive Stream Updater](https://github.com/core-go/hive/blob/main/batch/stream_writer.go): to update data. When you write data, it keeps the data in the buffer, it does not write data. It just writes data when flush.
