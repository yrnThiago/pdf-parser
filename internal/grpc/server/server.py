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
user_personal_info = {
"Nome": "",
 "E-mail": "",
 "Telefone": "",
 "Localização": "",
 "LinkedIn": "",
 "GitHub": "",
}

user_experience = {
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

        extractAndSaveUserInfoByKeys(content)

        user = pdf_pb2.User(
            ID = request.ID,
            Name = user_personal_info["Nome"],
            Email = user_personal_info["E-mail"],
            CellNumber = user_personal_info["Telefone"],
            Address = user_personal_info["Localização"],
            LinkedIn = user_personal_info["LinkedIn"],
            Github = user_personal_info["GitHub"],
        )
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

def extractAndSaveUserInfoByKeys(content):
    personal_info_found = False

    rows = content.split("\n")
    qty_rows = len(rows)

    idx = 0
    while idx < qty_rows-1:
        row = rows[idx]
        words = row.split()
        if rowHasTitle(row):
            if not personal_info_found:
                personal_info_found = True

                for y in range(1, len(user_personal_info)+1):
                    # TODO: fix more than 2 values when "key:val | key:val"
                    key, val = rows[idx+y].split(":")
                    user_personal_info[key] = val.strip()

                idx += len(user_personal_info)+1

            key = words[0]
            next_row = rows[idx+1]
            while not rowHasTitle(next_row[idx+1]):
                user_experience[key] += next_row[idx+1]
                idx += 1
        idx += 1

def rowHasTitle(row):
    qty_words = len(row.split())
    if qty_words <= 2:
        return True

    return False


def getInsecureGrpcPort():
    return f"[::]:{grpc_port}"

def getPdfPath(id):
    return "internal/uploads/" + id + ".pdf"

if __name__ == '__main__':
    serve()
