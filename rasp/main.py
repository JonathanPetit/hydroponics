import yaml 
import logging
import schedule

from hydroponics import Hydroponics

def main():
    hydroponics = Hydroponics()

    schedule.every(1).minutes.do(hydroponics.dht22.readTemperature)
    schedule.every(1).minutes.do(hydroponics.dht22.readHumidity)

    while True:
        schedule.run_pending()


if __name__ == "__main__":
    main()
