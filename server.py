import socket

# create chatroom class
class Chatroom:
    def __init__(self, room_name, max_users, key):
        self.room_name = room_name
        self.max_users = max_users
        self.key = key
        self.users = []

def main():
    # chatroom needs to be implemented as <str, ChatRoom>
    chatrooms = {}
    # server address and port
    server_address = '0.0.0.0'
    server_port = 9000

    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    sock.bind((server_address, server_port))
    sock.listen(5)
    while True:
        connection, client_address = sock.accept()
        data, _ = connection.recvfrom(4096)
        print(data)

if __name__ == '__main__':
    main() 
        