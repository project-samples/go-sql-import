# go-sql-import

## Import 1 million records into SQL from CSV Fixed-Length Format Files with high performance
In this demo, we import 1 Million Records into SQL from CSV Fixed-Length Format Files with high performance in GOLANG.

This article outlines strategies for high-performance batch processing in GOLANG, specifically for importing 1 million records into PostgreSQL from a CSV or fixed-length format file.

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
<td>RAM</td>
<td>Disk</td>
<td>Without data validation</td>
<td>With data validation</td>
</tr></thead><tbody>

<tr>
<td>Fix Length</td>
<td>1.15 GB</td>
<td>1,018,584</td>
<td>33 M</td>
<td>3.1 M/s</td>
<td>5 min 16 sec</td>
<td>6 min 10 sec</td>
</tr>

<tr>
<td>CSV</td>
<td>0.975 GB</td>
<td>1,018,584</td>
<td>34 M</td>
<td>2.8 M/s</td>
<td>5 min 12 sec</td>
<td>6 min</td>
</tr>

</tbody></table>

- For fix length format file, please unzip 'fixedlength_big.7z' file (1.15 GB)
- For CSV file, please unzip 'delimiter_big.7z' file (0.975 GB)

### Batch jobs
Differ from online processing:
- Long time running, often at night, after working hours.
- Non interactive, often include logic for handling errors
- Large volumes of data

### Challenges in Batch Processing
- Data Size: Large volumes can strain memory and I/O performance.
- Database Constraints: SQL (Postgres, My SQL, MS SQL...) may struggle with handling high transaction loads.
### High-Performance Strategy
#### Streaming and Buffered I/O:
- Instead of loading the entire file into memory, stream the file line by line using Go's buffered I/O. This reduces memory consumption.
- Example: Use bufio.Scanner to read records line by line.
#### Bulk Inserts:
Bulk inserts in SQL are typically faster because they reduce the overhead associated with individual inserts. Here’s why:
- Transaction Handling: In a bulk insert, multiple rows are inserted in a single transaction, reducing the need for multiple commits, which can be expensive in terms of I/O operations.
- Logging: Bulk inserts often minimize logging overhead, especially if the database is configured to use minimal logging for bulk operations (like in SQL Server with the "BULK_LOGGED" recovery model).
- Constraints: When inserting data in bulk, constraints such as foreign keys and unique indexes may be deferred or optimized by the database engine.
- Index Updates: Instead of updating indexes after each row insert, bulk inserts allow the database to update indexes in batches, improving performance.

However, it’s important to note that bulk inserts still need to ensure data integrity. Some databases provide options to temporarily disable constraints or logging to optimize performance further, but this can lead to data consistency issues if not handled properly.

#### Batch Inserts:
- Insert records into PostgreSQL in batches (e.g., 1000 records per transaction). This reduces transaction overhead and improves performance.
- In this sample, we use Batch Inserts. We still have a very good performance.
#### Error Handling and Logging:
- Implement robust error handling and logging. Track failed records to reprocess or fix them later.

### Advantages of This Approach
- Efficiency: Streaming and batching minimize memory usage and reduce database transaction overhead.
- Scalability: Parallel processing allows you to scale the import process across multiple cores.
- Flexibility: This approach can handle large datasets and can be adapted for other file formats or databases.
### Disadvantages
- Simple: Do not handle retries or parallel processing.
- Transaction Size: Large batches can still strain the database if not managed properly.

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
- SQL Writer: to insert or update data
- SQL Inserter: to insert data
- SQL Updater: to update data

- SQL Stream Writer: to insert or update data. When you write data, it keeps the data in the buffer, it does not write data. It just writes data when flush.
- SQL Inserter: to insert data. When you write data, it keeps the data in the buffer, it does not write data. It just writes data when flush. Especially, we build 1 single SQL statement to improve the performance.
- SQL Updater: to update data. When you write data, it keeps the data in the buffer, it does not write data. It just writes data when flush.

- Mongo Writer: to insert or update data
- Mongo Inserter: to insert data
- Mongo Updater: to update data

- Mongo Stream Writer: to insert or update data. When you write data, it keeps the data in the buffer, it does not write data. It just writes data when flush.
- Mongo Inserter: to insert data. When you write data, it keeps the data in the buffer, it does not write data. It just writes data when flush.
- Mongo Updater: to update data. When you write data, it keeps the data in the buffer, it does not write data. It just writes data when flush.