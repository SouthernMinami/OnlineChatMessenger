import socket

class ChatClient:
    def __init__(self, address, port):
        self.address = address
        self.port = port
        # When creating a new chat room, TCP socket is used to avoid loosing data
        self.sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

    def create_room(self, room_name, max_users, key, server_address, server_port):
        self.sock.connect((server_address, server_port))
        message = 'new room name: {}, max number of users: {}, key: {}'.format(room_name, max_users, key).encode()
        self.sock.send(message)
        data, _ = self.sock.recvfrom(4096)
        return data
    
def main():
    
    # server address and port
    server_address = '0.0.0.0'
    server_port = 9000

    client1 = ChatClient('', '9900')
    client2 = ChatClient('', '9901')
    client3 = ChatClient('', '9902')
    client4 = ChatClient('', '9903')

    while True:
        room_name = input('Enter a new room name: ')
        max_users = input('Maximum number of users?:')
        key = input('Key number?: ')
        print(client1.create_room(room_name, max_users, key, server_address, server_port))

if __name__ == '__main__':
    main()