A quick benchmark of persisting and getting JSON documents with Upspin and MongoDB.

Example execution:

```
$ go run *.go /path/to/upspin/config user@mail.com/appdata

Mongo accesser:
persist took 290.13472ms
get took 104.90109ms
persist took 24.954586ms
get took 25.122874ms
persist took 25.247414ms
get took 24.58505ms
persist took 24.331595ms
get took 24.579216ms
persist took 27.946452ms
get took 25.080292ms
persist took 27.224589ms
get took 24.769088ms
persist took 25.877404ms
get took 25.005334ms
persist took 24.826546ms
get took 25.63095ms
persist took 25.357371ms
get took 28.32153ms
persist took 31.020568ms
get took 28.744731ms

Upspin accesser:
persist took 1.471544477s
get took 53.47207ms
persist took 253.415334ms
get took 53.071953ms
persist took 200.974539ms
get took 52.971038ms
persist took 204.022989ms
get took 52.755499ms
persist took 188.978781ms
get took 66.460147ms
persist took 291.278249ms
get took 68.187328ms
persist took 224.832229ms
get took 57.860124ms
persist took 225.751839ms
get took 52.982705ms
persist took 219.969851ms
get took 59.638136ms
persist took 257.537105ms
get took 68.611404ms
```
