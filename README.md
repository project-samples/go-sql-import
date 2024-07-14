# go-import

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