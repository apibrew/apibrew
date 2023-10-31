from apibrew.client import Client

client = Client.new_client()

print(client.list_resources())
