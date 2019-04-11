import requests
import logging

import Adafruit_DHT


class DHT22():
    def __init__(self, logger, pin: int):
        self.logger = logger

        self.pin = pin
        self.sensor = Adafruit_DHT.DHT22

    
    def readTemperature(self):
        _, temp = Adafruit_DHT.read_retry(self.sensor, self.pin)
        data = {'value': temp}

        try:
            requests.post(url = 'http://192.168.0.13:8080/temperature', json = data) 
            print(temp)
        except requests.exceptions.RequestException as e:
            self.logger.warning(e)


    def readHumidity(self):
        humidity, _ = Adafruit_DHT.read_retry(self.sensor, self.pin)
        data = {'value': humidity}

        try:
            requests.post(url = 'http://192.168.0.13:8080/humidity', json = data) 
            print(humidity)
        except requests.exceptions.RequestException as e:
            self.logger.warning(e)

