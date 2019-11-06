# Direct FS

A simple webservice that serves contents from a single static file from the filesystem, when being contacted on the "/" route.

A buffer with the blocksize of the filesystem is allocated (default 4096 bytes), and filled with data from the file until it reaches capacity.

The OS filesystem cache is bypassed to ensure the contents are read again from disk on each request.

This webservice was created to track the downtime of a Xen VM during migration to another physical server.

