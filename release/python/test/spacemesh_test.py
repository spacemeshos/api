from concurrent import futures
import unittest

from spacemesh.v1 import node_pb2_grpc, node_types_pb2, types_pb2
import grpc


class NodeServicer(node_pb2_grpc.NodeServiceServicer):
    def Echo(self, request, context):
        return node_types_pb2.EchoResponse(msg=request.msg)


class TestGrpcApi(unittest.TestCase):
    def setUp(self):
        """
        Launch a stub server
        """
        server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        self.server = server
        node_pb2_grpc.add_NodeServiceServicer_to_server(
            NodeServicer(), server)
        server.add_insecure_port('[::]:50051')
        server.start()

    def tearDown(self):
        self.server.stop(None)

    def test_NodeService(self):
        message = "Hello, world!"
        channel = grpc.insecure_channel('[::]:50051')
        stub = node_pb2_grpc.NodeServiceStub(channel)
        response = stub.Echo(node_types_pb2.EchoRequest(msg=types_pb2.SimpleString(value=message)))
        self.assertEqual(response.msg.value, message)


if __name__ == '__main__':
    unittest.main()
