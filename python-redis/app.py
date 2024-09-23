import redis

# Connect to Redis server
client = redis.Redis(host='localhost', port=6379, db=0)

# Save a record
key = 'sample_key'
value = 'sample_value'
client.set(key, value)

# Retrieve the record
retrieved_value = client.get(key)

# Print the retrieved value
print(f'The value for "{key}" is "{retrieved_value.decode("utf-8")}"')
