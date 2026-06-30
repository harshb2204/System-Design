# Blob Storages and S3

Earlier when people uploaded any files, they uploaded it to "server" and were stored on the hard disk attached to it.

getting a file was simple, make an API call, the handler reads the file & return

* This is precisely how `/static` folders / routed worked

when user uploads the file -> accept it on HTTP POST
`'a.txt'` -> create absolute path using folder mapped to `'/static'`
`/home/arpit/www/static/a.txt`
-> store the file at that location

when user requests `'/static/a.txt'` -> get the path from URL
-> create absolute path
`/home/arpi/www/static/a.txt`
-> read the file and return

*(Early days of the internet)*

This worked well for quite some time, but won't work with multiple servers because each server will have its own disk.
![](/diagrams/diskstorage.png)

Hence, we need an infinitely scalable network attached storage / file system.
-> **This is S3 / BlobStorage**
![](/diagrams/s3.png)

## S3 Concepts

On S3 you have:
* **buckets**: (namespace) e.g., `insta-images`, `my-bucket` (must be any unique name)
* **keys**: path of the file within the bucket

Example: `s3://insta-images/user123/72896.png`
- `insta-images` is the **bucket**
- `user123/72896.png` is the **key**

You can seamlessly create the file, replace the file, delete the file, and read the entire file or a segment of it.
*Note: It is NOT a full-fledged file system.*

## Advantages & Disadvantages

### Advantages:
- cheap, durable storage
- can store literally any file (images, video, audio, text, DB backup, DB CSV exports, anything)
- scalable and available
- integration with a lot of AWS and BigData services

### Disadvantages:
- reads on S3 are slow
  - so if you want quick reads, you **should not** use S3
  - -> SSD and HDD attached to instances are better for this
- not a full-fledged file system

## When to use S3

You should use S3 when you want to store a 'blob' that is centrally accessible.
- Database backups
- Static website hosting
- Big Data storage
- logs archival
- infrequently accessed data dumping ground
