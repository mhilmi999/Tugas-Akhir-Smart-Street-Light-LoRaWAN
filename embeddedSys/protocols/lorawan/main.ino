  #include <lorawan.h>
  #include <PZEM004Tv30.h>

  #if defined(ESP32)
  PZEM004Tv30 pzem(Serial2, 16, 17);
  #else
  PZEM004Tv30 pzem(Serial2);
  #endif

  //ABP Credentials
  const char *devAddr = "36009001";
  const char *nwkSKey = "01fe7c50a39803d00000000000000000";
  const char *appSKey = "000000000000000093a1cf61893c1605";


  const unsigned long interval = 10000;    // 10 s interval to send message
  unsigned long previousMillis = 0;  // will store last time message sent
  unsigned int counter = 0;     // message counter

  char myStr[100];
  char temp[50];
  Char outStr[20];
  int recvStatus = 0;
  int port, channel, freq;
  bool newmessage = false;
  String dataSend = "";

  const sRFM_pins RFM_pins = {
    .CS = 5,
    .RST = 17,
    .DIO0 = 2,
    .DIO1 = 4,
  };

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
  }

  void loop() {
    lorawanrxtx();
  }

  void lorawanrxtx(){
    // Check interval overflow
    if (millis() - previousMillis > interval) {
      previousMillis = millis();

      sprintf(myStr, "Lora Counter-%d\n", counter++);
      
      Serial.print("Custom Address:");
      Serial.println(pzem.readAddress(), HEX);

      // Read the data from the sensor
      float voltage = getVoltage();
      float current = getCurrent();
      float power = getPower();
      float energy = getEnergy();

      // Uplink Send Data form Sensor to LoRaWAN Gateway    
      Serial.print("Sending: ");
      dataSend = "{\"volt\": " + String(voltage,2) + ", \"cur\": " + String(current,3) +", \"pwr\": " + String(power,2) +", \"eng\": " + String(energy,3) "}";
      dataSend.toCharArray(myStr,100);
      Serial.println(myStr);
      lora.sendUplink(myStr, strlen(myStr), 0);
      port = lora.getFramePortTx();
      channel = lora.getChannel();
      freq = lora.getChannelFreq(channel);
      Serial.print(F("fport: "));    Serial.print(port);Serial.print(" ");
      Serial.print(F("Ch: "));    Serial.print(channel);Serial.print(" ");
      Serial.print(F("Freq: "));    Serial.print(freq);Serial.println(" ");
      Serial.print("Voltage: ");      Serial.print(voltage);      Serial.println("V");
      Serial.print("Current: ");      Serial.print(current);      Serial.println("A");
      Serial.print("Power: ");        Serial.print(power);        Serial.println("W");
      Serial.print("Energy: ");       Serial.print(energy,3);     Serial.println("kWh");

    }

    // Check Lora RX
    lora.update();

    // Downlink Receive Data from LoRaWAN Gateway
    recvStatus = lora.readdata(outStr);
    if (recvStatus) {
      Serial.println(outStr);
    }
  }

  float getVoltage(){
    float v = pzem.voltage();
    if (isnan(v)) {
      Serial.println("Error reading voltage");
      return 0;
    }
    return v;
  }
  float getCurrent(){
    float c = pzem.current();
    if (isnan(c)) {
      Serial.println("Error reading current");
      return 0;
    }
    return c;
  }
  float getPower(){
    float p = pzem.power();
    if (isnan(p)) {
      Serial.println("Error reading power");
      return 0;
    }
    return p;
  }
  float getEnergy(){
    float e = pzem.energy();
    if(isnan(e)){
      Serial.println("Error reading energy");
      return 0;
    }
  }