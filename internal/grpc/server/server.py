import grpc
from concurrent import futures
import time

import sys
sys.path.append('./api/pb') 

import pdf_pb2
import pdf_pb2_grpc
from pypdf import PdfReader
from dotenv import load_dotenv
import os

# TODO: storage user info extracted from pdf
user_info = {
    "Objetivo": "",
    "Conhecimento": "",
    "Experiência": "",
    "Educação": "",
}

load_dotenv()

grpc_port = os.getenv('GRPC_PORT', "50051")

# TODO: maybe move to service folder...
class PdfServiceServicer(pdf_pb2_grpc.PdfServiceServicer):

    def ExtractFromPdf(self, request, context):
        print(f"extracting text from pdf id: {request.ID}")

        reader = PdfReader(getPdfPath(request.ID))
        page = reader.pages[0]
        content = page.extract_text()

        for row in content.split("\n"):
            words = row.split()
            qty_words = len(words)
            if qty_words == 1:
                user_info[words[0]] = "_"

            print(row, len(row.split()))
        
        print(user_info)
        user = pdf_pb2.User()
        pdfResponse = pdf_pb2.PdfResponse(User=user, Text=content)
        
        print(f"successfully pdf id: {request.ID}")
        return pdfResponse

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    
    pdf_pb2_grpc.add_PdfServiceServicer_to_server(PdfServiceServicer(), server)
    
    print(f"gRPC server listening on port {grpc_port}")
    server.add_insecure_port(getInsecureGrpcPort())
    server.start()
    
    #TODO: I dont like these lines refactor later
    try:
        while True:
            time.sleep(60 * 60 * 24)
    except KeyboardInterrupt:
        server.stop(0)

def getInsecureGrpcPort():
    return f"[::]:{grpc_port}"

def getPdfPath(id):
    return "internal/uploads/" + id + ".pdf"

if __name__ == '__main__':
    serve()
