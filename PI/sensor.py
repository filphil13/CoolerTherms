import dht
import urequests
import network
import socket
from time import sleep
import machine

SENSOR_NAME = "" #Sensor Name
SLEEP_TIME = 15
SENSOR = dht.DHT22(machine.Pin(16))
LED = machine.Pin("LED", machine.Pin.OUT)

ssid = "" #WIFI HostName
password = '' #WIFI Password

def blinkLED(numOfBlinks):
    for i in range(numOfBlinks):
        LED.on()
        sleep(0.2)
        LED.off()
        sleep(0.2)

def connect():
    #Connect to WLAN
    wlan = network.WLAN(network.STA_IF)
    wlan.active(True)
    wlan.connect(ssid, password)
    while wlan.isconnected() == False:
        print('Waiting for connection...')
        sleep(1)
    ip = wlan.ifconfig()[0]
    print(f'Connected on {ip}')
    return ip

def initSensor():
    
    rqst = urequests.post('http://192.168.2.11:8080/InitSensor', json=
        {
            "name": SENSOR_NAME,
            "address": None
        })
    rqst.close()
    blinkLED(4)

LED.on()

try:
    ip = connect()
except KeyboardInterrupt:
    machine.reset()

initSensor()
    
LED.off()
blinkLED(5)
while True:
    SENSOR.measure()
    humidity = SENSOR.humidity() 
    temperature = SENSOR.temperature()
    
    if humidity is not None and temperature is not None:
        print("Temp={0:0.1f}*C  Humidity={1:0.1f}%".format(temperature, humidity))
        try:
            rqst = urequests.post('http://192.168.2.11:8080/TempScans', json=
                {
                    "name": SENSOR_NAME,
                    "temperature": temperature,
                    "humidity": humidity,
                    "time": None
                }
            )
            rqst.close()
            blinkLED(2)
            
            
        except Exception as e:
            print(e)
            print("Error sending the post request")
            LED.on()
    else:
        print("Failed to retrieve data from humidity sensor")
    
    
    sleep(SLEEP_TIME)
