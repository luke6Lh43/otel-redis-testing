package com.example;

import redis.clients.jedis.Jedis;

public class RedisExample {
    public static void main(String[] args) {
        // Connect to Redis
        try (Jedis jedis = new Jedis("localhost", 6379)) {
            // Save a record
            jedis.set("mykey", "myvalue");

            // Retrieve the record
            String value = jedis.get("mykey");

            // Print the record
            System.out.println("Retrieved value: " + value);
        }
    }
}
