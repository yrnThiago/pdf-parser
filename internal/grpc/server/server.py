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

load_dotenv()

grpc_port = os.getenv('GRPC_PORT', "50051")

# TODO: maybe move to service folder...
class PdfServiceServicer(pdf_pb2_grpc.PdfServiceServicer):

    def ExtractFromPdf(self, request, context):
        print(f"extracting text from pdf id: {request.ID}")

        reader = PdfReader(get_pdf_path(request.ID))
        page = reader.pages[0]
        content = page.extract_text()

        user_personal_info, user_experience = extract_user_info_and_save(content)

        user = pdf_pb2.User(
            ID = request.ID,
            Name = user_personal_info["NOME"],
            Email = user_personal_info["E-MAIL"],
            CellNumber = user_personal_info["TELEFONE"],
            Address = user_personal_info["ENDEREÇO"],
            LinkedIn = user_personal_info["LINKEDIN"],
            Github = user_personal_info["GITHUB"],
        )
        pdf_response = pdf_pb2.PdfResponse(User=user, Text=content)

        print(f"successfully pdf id: {request.ID}")
        return pdf_response

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    
    pdf_pb2_grpc.add_PdfServiceServicer_to_server(PdfServiceServicer(), server)
    
    print(f"gRPC server listening on port {grpc_port}")
    server.add_insecure_port(get_insecure_grpc_port())
    server.start()
    
    #TODO: I dont like these lines refactor later
    try:
        while True:
            time.sleep(60 * 60 * 24)
    except KeyboardInterrupt:
        server.stop(0)

def extract_user_info_and_save(content):
    rows = content.split("\n")

    user_personal_info, idx = get_personal_info(rows) 
    user_experience = get_user_experience_from_content(idx, rows)

    return user_personal_info, user_experience
   
def get_user_experience_from_content(idx, rows):
    user_experience = {}

    qty_rows = len(rows)
    while idx < qty_rows-1:
        row = rows[idx]
        words = row.split()
        if row_has_title(row):
            key = words[0]
            if not user_experience.get(key):
                user_experience[key] = ""

            while not row_has_title(rows[idx+1]):
                user_experience[key] += rows[idx+1]

                idx += 1
        idx += 1

    return user_experience

def get_formatted_key_and_value(row):
    values = row.split(":")
    if len(values) == 1:
        values.append("not found")

    key, val = values[0].strip().upper(), values[1].strip()
    values = []

    return key, val

def get_personal_info(rows):
    user_personal_info = {
        "NOME": "",
        "E-MAIL": "",
        "TELEFONE": "",
        "ENDEREÇO": "",
        "LINKEDIN": "",
        "GITHUB": "",
    }

    user_personal_info["NOME"] = rows[0]
    user_personal_info["RESUMO"] = rows[1]

    idx = 1
    while not row_has_title(rows[idx+1]):
        keys_same_row = rows[idx+1].split("|")
        if len(keys_same_row) > 1:
            for row in keys_same_row:
                key, val = get_formatted_key_and_value(row)
                user_personal_info[key] = val
                
        key, val = get_formatted_key_and_value(keys_same_row[0])
        user_personal_info[key] = val
        idx += 1

    return user_personal_info, idx

def row_has_title(row):
    qty_words = len(row.split())
    if qty_words <= 2:
        return True

    return False


def get_insecure_grpc_port():
    return f"[::]:{grpc_port}"

def get_pdf_path(id):
    return "internal/uploads/" + id + ".pdf"

if __name__ == '__main__':
    serve()
