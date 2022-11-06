/**
   ABP device
   LoRaWAN class C
*/
#include <lorawan.h>
#include <HardwareSerial.h>
#include <PZEM004T.h>

//ABP Credentials
const char *devAddr = "368438d3";
const char *nwkSKey = "1fd8e4bfac90d41a0000000000000000";
const char *appSKey = "0000000000000000950213a0592cf0a8";



const unsigned long interval = 10000;    // 10 s interval to send message
unsigned long previousMillis = 0;  // will store last time message sent
unsigned int counter = 0;     // message counter

char myStr[50];
byte outStr[255];
byte recvStatus = 0;
int port, channel, freq;
bool newmessage = false;

const sRFM_pins RFM_pins = {
  .CS = 5,
  .RST = 17,
  .DIO0 = 2,
  .DIO1 = 4,
};

HardwareSerial Serial2(2);     // Use hwserial UART2 at pins IO-16 (RX2) and IO-17 (TX2)
PZEM004T pzem(&Serial2);
IPAddress ip(192,168,1,1);

int relayPin = 4;
int relayStat = 1; // 1 = relay is off, 0 = relay is on

void setup() {
  // Setup loraid access
  Serial.begin(115200);
  delay(2000);
  if (!lora.init()) {
    Serial.println("RFM95 not detected");
    delay(5000);
    return;
  }

  // Set LoRaWAN Class change CLASS_A or CLASS_C
  lora.setDeviceClass(CLASS_C);

  // Set Data Rate
  lora.setDataRate(SF10BW125);

  // Set FramePort Tx
  lora.setFramePortTx(5);

  // set channel to random
  lora.setChannel(MULTI);

  // Set TxPower to 15 dBi (max)
  lora.setTxPower(15);

  // Put ABP Key and DevAddress here
  lora.setNwkSKey(nwkSKey);
  lora.setAppSKey(appSKey);
  lora.setDevAddr(devAddr);
  
  pinMode(relayPin, OUTPUT);// define a pin as output for relay
  digitalWrite(relayPin, relayStat);//initial state either ON or OFF
  
  while (true) {
      Serial.println("Connecting to PZEM...");
      if(pzem.setAddress(ip))
        break;
      delay(1000);
   }
   
}

void loop() {
  lorawanProtocols();
  pzemSensor();
  lampControl();
}

void lorawanProtocols(){
    // Check interval overflow
  if (millis() - previousMillis > interval) {
    previousMillis = millis();

    sprintf(myStr, "Lora Counter-%d", counter++);

    
    Serial.print("Sending: ");
    Serial.println(myStr);
    lora.sendUplink(myStr, strlen(myStr), 0);
    port = lora.getFramePortTx();
    channel = lora.getChannel();
    freq = lora.getChannelFreq(channel);
    Serial.print(F("fport: "));    Serial.print(port);Serial.print(" ");
    Serial.print(F("Ch: "));    Serial.print(channel);Serial.print(" ");
    Serial.print(F("Freq: "));    Serial.print(freq);Serial.println(" ");

  }

  // Check Lora RX
  lora.update();

  recvStatus = lora.readDataByte(outStr);
  if (recvStatus) {
    newmessage = true;
    int counter = 0;
    port = lora.getFramePortRx();
    channel = lora.getChannelRx();
    freq = lora.getChannelRxFreq(channel);

    for (int i = 0; i < recvStatus; i++)
    {
      if (((outStr[i] >= 32) && (outStr[i] <= 126)) || (outStr[i] == 10) || (outStr[i] == 13))
        counter++;
    }
    if (port != 0)
    {
      if (counter == recvStatus)
      {
        Serial.print(F("Received String : "));
        for (int i = 0; i < recvStatus; i++)
        {
          Serial.print(char(outStr[i]));
        }
      }
      else
      {
        Serial.print(F("Received Hex: "));
        for (int i = 0; i < recvStatus; i++)
        {
          Serial.print(outStr[i], HEX); Serial.print(" ");
        }
      }
      Serial.println();
      Serial.print(F("fport: "));    Serial.print(port);Serial.print(" ");
      Serial.print(F("Ch: "));    Serial.print(channel);Serial.print(" ");
      Serial.print(F("Freq: "));    Serial.println(freq);Serial.println(" ");
    }
    else
    {
      Serial.print(F("Received Mac Cmd : "));
      for (int i = 0; i < recvStatus; i++)
      {
        Serial.print(outStr[i], HEX); Serial.print(" ");
      }
      Serial.println();
      Serial.print(F("fport: "));    Serial.print(port);Serial.print(" ");
      Serial.print(F("Ch: "));    Serial.print(channel);Serial.print(" ");
      Serial.print(F("Freq: "));    Serial.println(freq);Serial.println(" ");
    }
  }
}

void pzemSensor(){
    float v = pzem.voltage(ip);
  if (v < 0.0) v = 0.0;
   Serial.print(v);Serial.print("V; ");

  float i = pzem.current(ip);
   if(i >= 0.0){ Serial.print(i);Serial.print("A; "); }

  float p = pzem.power(ip);
   if(p >= 0.0){ Serial.print(p);Serial.print("W; "); }

  float e = pzem.energy(ip);
   if(e >= 0.0){ Serial.print(e);Serial.print("Wh; "); }

  Serial.println();

  delay(3000);
}

void lampControl(){
  if (relayStat) {
    digitalWrite(relayPin, HIGH);
    relayStat = 0;
  }
  Serial.println("Lampu menyala");
//   else {
//     digitalWrite(relayPin, LOW);
//     relayStat = 1;
//   }
}