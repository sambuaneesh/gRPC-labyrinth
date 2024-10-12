### Learning Materials

The `learning-go` and `learning-grpc` folders contain my learning materials for Go and gRPC respectively.

### Original Problem Statement

For the original problem statement, please refer to the [PROBLEM.md](Docs/PROBLEM.md) file.

### Progress

For the progress of the project, please refer to the [#1 issue](https://github.com/sambuaneesh/gRPC-labyrinth/issues/1) in this repo.

### How to run

To generate a labyrinth of size `width` and `height`, you run:

```bash
make gen width height
```

To build the project, you run:

```bash
make build
```

To run the server, you run:

```bash
cd bin
./server
```

To run the client first you need to install the dependencies:

```bash
pip install -r requirements.txt
```

Then you can run the client:

```bash
python client/client.py
```
