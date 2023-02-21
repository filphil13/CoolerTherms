import Adafruit_DHT
import requests
import time

SLEEP_TIME = 20
DHT_SENSOR = Adafruit_DHT.DHT22

DHT_PIN = 4

while True:
    humidity, temperature = Adafruit_DHT.read_retry(DHT_SENSOR, DHT_PIN)
    
    if humidity is not None and temperature is not None:
        print("Temp={0:0.1f}*C  Humidity={1:0.1f}%".format(temperature, humidity))
        requests.post('http://192.168.2.11:8080/TempScans', json=
            {
                "sensorid": 1,
                "temperature": temperature,
                "humidity": humidity,
                "time": time.ctime(time.time())
            }
        )
    else:
        print("Failed to retrieve data from humidity sensor")
    
    time.sleep(SLEEP_TIME)