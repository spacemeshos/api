import unittest

from spacemesh.v1 import node_pb2_grpc
import grpc


class MyTestCase(unittest.TestCase):
    def test_something(self):
        channel = grpc.insecure_channel('localhost:5901')
        stub = node_pb2_grpc.NodeServiceStub(channel)
        version = stub.Version
        self.assertEqual(1, version)


if __name__ == '__main__':
    unittest.main()
