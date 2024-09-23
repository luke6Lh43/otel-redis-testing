const express = require('express');
const redis = require('redis');
const { trace } = require('@opentelemetry/api');

// Initialize Redis client
const client = redis.createClient({
  url: 'redis://localhost:6379'
});

client.on('error', (err) => console.error('Redis Client Error', err));

const app = express();
const port = 3000;

app.use(express.json());

app.all('/', async (req, res) => {
  const key = 'myKey';

  try {
    await client.connect();

    if (req.method === 'POST') {
      const value = req.body.value || 'myValue';
      await client.set(key, value);
      res.send(`Set key ${key} with value ${value}`);
    } else if (req.method === 'GET') {
      const value = await client.get(key);
      res.send(`Got value ${value} for key ${key}`);
    } else {
      res.status(405).send('Method Not Allowed');
    }
  } catch (err) {
    console.error('Error interacting with Redis', err);
    res.status(500).send('Internal Server Error');
  } finally {
    await client.disconnect();
  }
});

app.listen(port, () => {
  console.log(`App listening at http://localhost:${port}`);
});
