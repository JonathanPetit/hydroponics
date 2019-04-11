import yaml
import logging

from dht22.dht22 import DHT22


class Hydroponics:
    def __init__(self):
        self.config = self.__readConfig('config.yml')
        self.logger = self.__setupLogger()

        self.pins = self.config['pins']

        self.dht22 = DHT22(self.logger, self.pins['dht22'])


    def __readConfig(self, filename: str):
        with open(filename, 'r') as file:
            data = yaml.safe_load(file)
        return data

    
    def __setupLogger(self):
        formatter = logging.Formatter('%(asctime)s :: %(levelname)s :: %(message)s')
        handler = logging.FileHandler("log.log", mode="a", encoding="utf-8")

        handler.setFormatter(formatter)
        handler.setLevel(logging.DEBUG)

        logger = logging.getLogger("hydroponics")
        logger.setLevel(logging.INFO)
        logger.addHandler(handler)

        return logger