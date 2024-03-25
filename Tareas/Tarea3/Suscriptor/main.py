import redis

def subscribe_and_receive(r, channel):
    p = r.pubsub()
    p.subscribe(channel)
    for message in p.listen():
        yield message

def process_message(message):
    if message["type"] == "message":
        data = message["data"].decode("utf-8")
        print(data)

r = redis.Redis(
    host='10.12.128.3',
    port=6379,
    decode_responses=True,
)

messages = subscribe_and_receive(r, "test")

for message in messages:
    process_message(message)
